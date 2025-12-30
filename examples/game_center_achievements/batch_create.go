//go:build ignore

/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/castbox/asc-go/asc"
	"github.com/castbox/asc-go/examples/util"
	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

var (
	bundleID    = flag.String("bundleid", "", "Bundle ID for an app (required)")
	configFile  = flag.String("config", "", "Path to JSON config file with achievements (required)")
	resume      = flag.Bool("resume", false, "Resume mode: skip existing achievements and localizations, only upload missing images")
	concurrency = flag.Int("concurrency", 10, "Number of concurrent localization/image uploads (default: 10)")
)

// rateLimiter limits API requests to 4 per second to avoid hitting the undocumented per-minute limit (~300/min)
var rateLimiter = rate.NewLimiter(rate.Limit(4), 10)

// AchievementConfig represents a single achievement configuration
type AchievementConfig struct {
	ReferenceName    string               `json:"referenceName"`
	VendorIdentifier string               `json:"vendorIdentifier"`
	Points           int                  `json:"points"`
	ShowBeforeEarned bool                 `json:"showBeforeEarned"`
	Repeatable       bool                 `json:"repeatable"`
	Position         int                  `json:"position"` // 1-based position in the final order. 0 means append at end.
	Localizations    []LocalizationConfig `json:"localizations"`
}

// LocalizationConfig represents localization for an achievement
type LocalizationConfig struct {
	Locale                  string `json:"locale"`
	Name                    string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription  string `json:"afterEarnedDescription"`
	ImageFile               string `json:"imageFile,omitempty"`
}

// BatchConfig represents the batch configuration file
type BatchConfig struct {
	Achievements []AchievementConfig `json:"achievements"`
}

func main() {
	flag.Parse()

	if *bundleID == "" {
		log.Fatal("bundleid is required")
	}
	if *configFile == "" {
		log.Fatal("config is required")
	}

	// Load configuration
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
	fmt.Printf("Loaded %d achievements from config\n", len(config.Achievements))

	ctx := context.Background()

	// Create client
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}
	client := asc.NewClient(auth.Client())

	// Get app
	fmt.Printf("Looking up app with bundle ID: %s\n", *bundleID)
	app, err := util.GetApp(ctx, client, &asc.ListAppsQuery{
		FilterBundleID: []string{*bundleID},
	})
	if err != nil {
		log.Fatalf("Failed to find app: %s", err)
	}
	fmt.Printf("Found app: %s (ID: %s)\n", *app.Attributes.Name, app.ID)

	// Initialize Game Center
	gcInfo, err := initializeGameCenter(ctx, client, app)
	if err != nil {
		log.Fatalf("Failed to initialize Game Center: %s", err)
	}

	// Build existing achievements map for resume mode
	var existingAchievements map[string]*asc.GameCenterAchievement
	if *resume {
		existingAchievements, err = fetchExistingAchievements(ctx, client, gcInfo.detailID)
		if err != nil {
			log.Fatalf("Failed to fetch existing achievements: %s", err)
		}
	}

	// Create each achievement and collect new release IDs
	var createdAchievements []string
	var skippedAchievements []string
	var failedAchievements []achievementFailure
	var newReleases []achievementRelease
	for i, achConfig := range config.Achievements {
		fmt.Printf("========================================\n")
		fmt.Printf("Processing achievement %d/%d: %s\n", i+1, len(config.Achievements), achConfig.ReferenceName)
		fmt.Printf("========================================\n")

		var achievementID, releaseID string
		var err error

		// Check if achievement already exists (resume mode)
		if *resume {
			if existingAch, ok := existingAchievements[achConfig.VendorIdentifier]; ok {
				achievementID = existingAch.ID
				fmt.Printf("  Achievement already exists (ID: %s), checking localizations...\n", achievementID)

				// Process localizations for existing achievement
				failedLocs := processLocalizationsForExistingAchievement(ctx, client, achievementID, achConfig.Localizations)
				if len(failedLocs) > 0 {
					failedAchievements = append(failedAchievements, achievementFailure{
						name:                achConfig.ReferenceName,
						vendorIdentifier:    achConfig.VendorIdentifier,
						failedLocalizations: failedLocs,
					})
					log.Printf("Warning: %d localization(s) failed for %s", len(failedLocs), achConfig.ReferenceName)
				}

				skippedAchievements = append(skippedAchievements, achievementID)
				fmt.Printf("Achievement processed (resume mode): %s\n\n", achievementID)
				continue
			}
		}

		// Create new achievement
		achievementID, releaseID, err = createAchievementWithRelease(ctx, client, gcInfo.detailID, gcInfo.groupID, achConfig)
		if err != nil {
			log.Printf("Failed to create achievement %s: %s", achConfig.ReferenceName, err)
			continue
		}
		createdAchievements = append(createdAchievements, achievementID)
		if releaseID != "" {
			newReleases = append(newReleases, achievementRelease{
				achievementID: achievementID,
				releaseID:     releaseID,
				position:      achConfig.Position,
				name:          achConfig.ReferenceName,
			})
		}
		fmt.Printf("Achievement created successfully: %s\n\n", achievementID)
	}

	// Reorder achievements based on position field
	reorderAchievements(ctx, client, gcInfo.detailID, gcInfo.existingReleases, newReleases)

	// Summary
	printSummary(len(config.Achievements), len(createdAchievements), len(skippedAchievements), len(newReleases), failedAchievements, *resume)
}

func loadConfig(path string) (*BatchConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config BatchConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// gameCenterInfo holds Game Center initialization information
type gameCenterInfo struct {
	detailID         string
	groupID          string
	existingReleases []string
}

// initializeGameCenter initializes Game Center and returns relevant IDs
func initializeGameCenter(ctx context.Context, client *asc.Client, app *asc.App) (*gameCenterInfo, error) {
	fmt.Println("Getting Game Center detail...")
	gameCenterDetail, _, err := client.GameCenter.GetGameCenterDetailForApp(ctx, app.ID, &asc.GetGameCenterDetailForAppQuery{
		Include: []string{"gameCenterGroup"},
	})
	if err != nil {
		fmt.Println("Game Center not enabled, enabling...")
		gameCenterDetail, _, err = client.GameCenter.CreateGameCenterDetail(ctx, app.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to enable Game Center: %w", err)
		}
	}
	fmt.Printf("Game Center Detail ID: %s\n\n", gameCenterDetail.Data.ID)

	// Check if the app belongs to a Game Center Group
	var gameCenterGroupID string
	gameCenterGroup, _, err := client.GameCenter.GetGameCenterGroupForDetail(ctx, gameCenterDetail.Data.ID, nil)
	if err == nil && gameCenterGroup != nil && gameCenterGroup.Data.ID != "" {
		gameCenterGroupID = gameCenterGroup.Data.ID
		fmt.Printf("App belongs to Game Center Group: %s\n", gameCenterGroupID)
		fmt.Println("Achievements will be created at the GROUP level.")
		fmt.Println()
	} else {
		// Try to get group ID from gameCenterDetail relationships
		if gameCenterDetail.Data.Relationships != nil && gameCenterDetail.Data.Relationships.GameCenterGroup != nil {
			if gameCenterDetail.Data.Relationships.GameCenterGroup.Data != nil {
				gameCenterGroupID = gameCenterDetail.Data.Relationships.GameCenterGroup.Data.ID
				fmt.Printf("App belongs to Game Center Group (from relationship): %s\n", gameCenterGroupID)
				fmt.Println("Achievements will be created at the GROUP level.")
				fmt.Println()
			}
		}
		if gameCenterGroupID == "" {
			fmt.Println("App does not belong to a Game Center Group.")
			fmt.Println("Achievements will be created at the APP level.")
			fmt.Println()
		}
	}

	// Get existing achievement releases for ordering
	fmt.Println("Getting existing achievement releases...")
	var existingReleaseIDs []string
	existingReleases, _, err := client.GameCenter.ListGameCenterAchievementReleasesForDetail(ctx, gameCenterDetail.Data.ID, &asc.ListGameCenterAchievementReleasesQuery{
		Limit: 200,
	})
	if err != nil {
		fmt.Printf("Note: Could not get existing releases (this is normal if no releases exist yet): %v\n", err)
	} else if existingReleases != nil {
		for _, r := range existingReleases.Data {
			existingReleaseIDs = append(existingReleaseIDs, r.ID)
		}
	}
	fmt.Printf("Found %d existing achievement releases\n\n", len(existingReleaseIDs))

	return &gameCenterInfo{
		detailID:         gameCenterDetail.Data.ID,
		groupID:          gameCenterGroupID,
		existingReleases: existingReleaseIDs,
	}, nil
}

// fetchExistingAchievements fetches existing achievements for resume mode
func fetchExistingAchievements(ctx context.Context, client *asc.Client, gameCenterDetailID string) (map[string]*asc.GameCenterAchievement, error) {
	existingAchievements := make(map[string]*asc.GameCenterAchievement)
	fmt.Println("Resume mode enabled, fetching existing achievements...")
	existingAchList, _, err := client.GameCenter.ListGameCenterAchievementsForDetail(ctx, gameCenterDetailID, &asc.ListGameCenterAchievementsQuery{
		Limit: 200,
	})
	if err != nil {
		fmt.Printf("Warning: Could not fetch existing achievements: %v\n", err)
		return existingAchievements, nil
	}
	if existingAchList != nil {
		for i := range existingAchList.Data {
			ach := &existingAchList.Data[i]
			if ach.Attributes != nil && ach.Attributes.VendorIdentifier != nil {
				existingAchievements[*ach.Attributes.VendorIdentifier] = ach
			}
		}
	}
	fmt.Printf("Found %d existing achievements\n\n", len(existingAchievements))
	return existingAchievements, nil
}

// achievementRelease represents a new achievement release for ordering
type achievementRelease struct {
	achievementID string
	releaseID     string
	position      int
	name          string
}

// reorderAchievements reorders achievements based on position field
func reorderAchievements(ctx context.Context, client *asc.Client, gameCenterDetailID string, existingReleaseIDs []string, newReleases []achievementRelease) error {
	if len(newReleases) == 0 || len(existingReleaseIDs)+len(newReleases) == 0 {
		return nil
	}

	fmt.Println("========================================")
	fmt.Println("Reordering achievements...")
	fmt.Println("========================================")

	// Start with existing releases
	finalOrder := make([]string, len(existingReleaseIDs))
	copy(finalOrder, existingReleaseIDs)

	// Sort new releases by position (higher positions first to avoid index shifting issues)
	for i := 0; i < len(newReleases); i++ {
		for j := i + 1; j < len(newReleases); j++ {
			if newReleases[i].position < newReleases[j].position {
				newReleases[i], newReleases[j] = newReleases[j], newReleases[i]
			}
		}
	}

	// Insert each new release at its position (process from highest to lowest position)
	for _, nr := range newReleases {
		if nr.position <= 0 {
			// Position 0 or negative means append at end
			finalOrder = append(finalOrder, nr.releaseID)
			fmt.Printf("  %s -> appended at END\n", nr.name)
		} else if nr.position > len(finalOrder) {
			// Position beyond current length, append at end
			finalOrder = append(finalOrder, nr.releaseID)
			fmt.Printf("  %s -> position %d (appended at END, beyond current length)\n", nr.name, nr.position)
		} else {
			// Insert at specific position (1-based, so position 1 = index 0)
			idx := nr.position - 1
			finalOrder = append(finalOrder[:idx], append([]string{nr.releaseID}, finalOrder[idx:]...)...)
			fmt.Printf("  %s -> position %d\n", nr.name, nr.position)
		}
	}

	fmt.Printf("Final order: %d achievements total\n", len(finalOrder))
	_, err := client.GameCenter.ReplaceGameCenterAchievementReleasesForDetail(ctx, gameCenterDetailID, finalOrder)
	if err != nil {
		fmt.Printf("Note: Could not reorder achievements: %v\n", err)
		fmt.Println("(Reordering requires an editable Game Center enabled app version)")
	} else {
		fmt.Println("Achievements reordered successfully!")
	}
	fmt.Println()

	return nil
}

// achievementFailure represents an achievement with failed localizations
type achievementFailure struct {
	name                string
	vendorIdentifier    string
	failedLocalizations []localizationError
}

// printSummary prints the final summary of the batch operation
func printSummary(totalCount, createdCount, skippedCount, releasesCount int, failedAchievements []achievementFailure, resumeMode bool) {
	fmt.Println("\n========================================")
	fmt.Println("Batch Creation Complete!")
	fmt.Println("========================================")
	fmt.Printf("Total achievements in config: %d\n", totalCount)
	fmt.Printf("Successfully created: %d\n", createdCount)
	if resumeMode {
		fmt.Printf("Skipped (already existed): %d\n", skippedCount)
	}
	fmt.Printf("Releases created: %d\n", releasesCount)
	fmt.Printf("Failed: %d\n", totalCount-createdCount-skippedCount)

	if len(failedAchievements) > 0 {
		fmt.Printf("\nAchievements with localization failures: %d\n", len(failedAchievements))
		for _, fa := range failedAchievements {
			fmt.Printf("  - %s (%s): %d failed localization(s)\n", fa.name, fa.vendorIdentifier, len(fa.failedLocalizations))
			for _, fl := range fa.failedLocalizations {
				fmt.Printf("    * [%s] %v\n", fl.locale, fl.err)
			}
		}
	}
	fmt.Println("========================================")
}

// localizationError represents a localization processing error
type localizationError struct {
	locale string
	err    error
}

// isImageUploadComplete checks if an image upload is complete
func isImageUploadComplete(imageInfo *asc.GameCenterAchievementImageResponse, err error) bool {
	return err == nil &&
		imageInfo != nil &&
		imageInfo.Data.Attributes != nil &&
		imageInfo.Data.Attributes.AssetDeliveryState != nil &&
		imageInfo.Data.Attributes.AssetDeliveryState.State != nil &&
		*imageInfo.Data.Attributes.AssetDeliveryState.State == "COMPLETE"
}

// verifyAndHandleImage verifies if an image was successfully uploaded and re-uploads if necessary
func verifyAndHandleImage(ctx context.Context, client *asc.Client, imageID string, localizationID string, imagePath string, locale string, mu *sync.Mutex) error {
	if err := rateLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limiter error for getting image: %w", err)
	}

	imageInfo, _, err := client.GameCenter.GetGameCenterAchievementImage(ctx, imageID, nil)

	if !isImageUploadComplete(imageInfo, err) {
		// Image upload incomplete or failed, delete and re-upload
		mu.Lock()
		if err != nil {
			fmt.Printf("    [%s] Failed to get image info (error: %v), deleting and re-uploading...\n", locale, err)
		} else if imageInfo != nil && imageInfo.Data.Attributes != nil && imageInfo.Data.Attributes.AssetDeliveryState != nil && imageInfo.Data.Attributes.AssetDeliveryState.State != nil {
			fmt.Printf("    [%s] Image upload incomplete (state: %s), deleting and re-uploading...\n", locale, *imageInfo.Data.Attributes.AssetDeliveryState.State)
		} else {
			fmt.Printf("    [%s] Image upload incomplete/failed, deleting and re-uploading...\n", locale)
		}
		mu.Unlock()

		// Delete the incomplete image
		if err := rateLimiter.Wait(ctx); err != nil {
			return fmt.Errorf("rate limiter error for deleting image: %w", err)
		}

		_, deleteErr := client.GameCenter.DeleteGameCenterAchievementImage(ctx, imageID)
		if deleteErr != nil {
			return fmt.Errorf("failed to delete incomplete image: %w", deleteErr)
		}

		// Re-upload the image
		if err := rateLimiter.Wait(ctx); err != nil {
			return fmt.Errorf("rate limiter error for re-uploading image: %w", err)
		}

		if imgErr := uploadImage(ctx, client, localizationID, imagePath); imgErr != nil {
			return fmt.Errorf("failed to re-upload image: %w", imgErr)
		}

		mu.Lock()
		fmt.Printf("    [%s] Image re-uploaded successfully\n", locale)
		mu.Unlock()
	} else {
		// Image uploaded successfully, skip
		mu.Lock()
		fmt.Printf("    [%s] OK (image uploaded successfully)\n", locale)
		mu.Unlock()
	}

	return nil
}

// uploadImageIfProvided uploads an image if the imagePath is not empty
func uploadImageIfProvided(ctx context.Context, client *asc.Client, localizationID string, imagePath string, locale string, mu *sync.Mutex) error {
	if imagePath == "" {
		return nil
	}

	if err := rateLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limiter error for image: %w", err)
	}

	if imgErr := uploadImage(ctx, client, localizationID, imagePath); imgErr != nil {
		return fmt.Errorf("failed to upload image: %w", imgErr)
	}

	mu.Lock()
	fmt.Printf("    [%s] Image uploaded successfully\n", locale)
	mu.Unlock()

	return nil
}

// createNewLocalization creates a new localization and uploads its image if provided
func createNewLocalization(ctx context.Context, client *asc.Client, achievementID string, locConfig LocalizationConfig, mu *sync.Mutex) error {
	if err := rateLimiter.Wait(ctx); err != nil {
		return fmt.Errorf("rate limiter error: %w", err)
	}

	newLoc, _, locErr := client.GameCenter.CreateGameCenterAchievementLocalization(ctx, asc.GameCenterAchievementLocalizationCreateRequestAttributes{
		Locale:                  locConfig.Locale,
		Name:                    locConfig.Name,
		BeforeEarnedDescription: locConfig.BeforeEarnedDescription,
		AfterEarnedDescription:  locConfig.AfterEarnedDescription,
	}, achievementID)
	if locErr != nil {
		return fmt.Errorf("failed to create localization: %w", locErr)
	}

	mu.Lock()
	fmt.Printf("    [%s] Created localization ID: %s\n", locConfig.Locale, newLoc.Data.ID)
	mu.Unlock()

	// Upload image if provided
	if err := uploadImageIfProvided(ctx, client, newLoc.Data.ID, locConfig.ImageFile, locConfig.Locale, mu); err != nil {
		return err
	}

	return nil
}

// handleExistingLocalization handles an existing localization, checking and uploading images as needed
func handleExistingLocalization(ctx context.Context, client *asc.Client, existingLoc *asc.GameCenterAchievementLocalization, locConfig LocalizationConfig, mu *sync.Mutex) error {
	hasImage := existingLoc.Relationships != nil &&
		existingLoc.Relationships.GameCenterAchievementImage != nil &&
		existingLoc.Relationships.GameCenterAchievementImage.Data != nil &&
		existingLoc.Relationships.GameCenterAchievementImage.Data.ID != ""

	if locConfig.ImageFile != "" {
		if hasImage {
			// Verify if the image was successfully uploaded
			imageID := existingLoc.Relationships.GameCenterAchievementImage.Data.ID
			if err := verifyAndHandleImage(ctx, client, imageID, existingLoc.ID, locConfig.ImageFile, locConfig.Locale, mu); err != nil {
				return err
			}
		} else {
			// No image exists, upload new image
			if err := rateLimiter.Wait(ctx); err != nil {
				return fmt.Errorf("rate limiter error for image: %w", err)
			}

			if imgErr := uploadImage(ctx, client, existingLoc.ID, locConfig.ImageFile); imgErr != nil {
				return fmt.Errorf("failed to upload image: %w", imgErr)
			}

			mu.Lock()
			fmt.Printf("    [%s] Missing image uploaded successfully\n", locConfig.Locale)
			mu.Unlock()
		}
	} else {
		mu.Lock()
		fmt.Printf("    [%s] OK (no image configured)\n", locConfig.Locale)
		mu.Unlock()
	}

	return nil
}

// processLocalizationsForExistingAchievement handles localizations for an achievement that already exists
// It checks which localizations exist and which need images, then only uploads missing images
// Returns a slice of failed localizations
func processLocalizationsForExistingAchievement(ctx context.Context, client *asc.Client, achievementID string, localizations []LocalizationConfig) []localizationError {
	// Get existing localizations for this achievement
	existingLocs, _, err := client.GameCenter.ListGameCenterAchievementLocalizationsForAchievement(ctx, achievementID, &asc.ListGameCenterAchievementLocalizationsQuery{
		Limit:   200,
		Include: []string{"gameCenterAchievementImage"},
	})
	if err != nil {
		return []localizationError{{locale: "unknown", err: fmt.Errorf("failed to list existing localizations: %w", err)}}
	}

	// Build map of existing localizations: locale -> localization
	existingLocMap := make(map[string]*asc.GameCenterAchievementLocalization)
	for i := range existingLocs.Data {
		loc := &existingLocs.Data[i]
		if loc.Attributes != nil && loc.Attributes.Locale != nil {
			existingLocMap[*loc.Attributes.Locale] = loc
		}
	}
	fmt.Printf("  Found %d existing localizations\n", len(existingLocMap))

	// Process each localization from config concurrently
	fmt.Printf("  Processing %d localizations concurrently (concurrency=%d)...\n", len(localizations), *concurrency)

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(*concurrency)

	var mu sync.Mutex
	var failedLocalizations []localizationError

	for _, locConfig := range localizations {
		locConfig := locConfig // capture loop variable
		existingLoc, exists := existingLocMap[locConfig.Locale]

		g.Go(func() error {
			var err error
			if !exists {
				err = createNewLocalization(gCtx, client, achievementID, locConfig, &mu)
			} else {
				err = handleExistingLocalization(gCtx, client, existingLoc, locConfig, &mu)
			}

			if err != nil {
				mu.Lock()
				failedLocalizations = append(failedLocalizations, localizationError{locale: locConfig.Locale, err: err})
				mu.Unlock()
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return []localizationError{{locale: "unknown", err: err}}
	}

	if len(failedLocalizations) > 0 {
		fmt.Printf("  Localizations processed with %d failures\n", len(failedLocalizations))
	} else {
		fmt.Println("  All localizations processed successfully")
	}

	return failedLocalizations
}

func createAchievementWithRelease(ctx context.Context, client *asc.Client, gameCenterDetailID string, gameCenterGroupID string, config AchievementConfig) (achievementID string, releaseID string, err error) {
	// 1. Create the achievement (at group level if app belongs to a group, otherwise at app level)
	var achievement *asc.GameCenterAchievementResponse
	attrs := asc.GameCenterAchievementCreateRequestAttributes{
		ReferenceName:    config.ReferenceName,
		VendorIdentifier: config.VendorIdentifier,
		Points:           config.Points,
		ShowBeforeEarned: config.ShowBeforeEarned,
		Repeatable:       config.Repeatable,
	}

	if gameCenterGroupID != "" {
		// App belongs to a group, create achievement at group level
		achievement, _, err = client.GameCenter.CreateGameCenterAchievementForGroup(ctx, attrs, gameCenterGroupID)
	} else {
		// App does not belong to a group, create achievement at app level
		achievement, _, err = client.GameCenter.CreateGameCenterAchievement(ctx, attrs, gameCenterDetailID)
	}
	if err != nil {
		return "", "", fmt.Errorf("failed to create achievement: %w", err)
	}
	fmt.Printf("  Achievement ID: %s\n", achievement.Data.ID)

	// 2. Create release for the achievement (for ordering - may fail if no editable version)
	release, _, releaseErr := client.GameCenter.CreateGameCenterAchievementRelease(ctx, achievement.Data.ID, gameCenterDetailID)
	if releaseErr != nil {
		fmt.Printf("  Note: Could not create release: %v\n", releaseErr)
		fmt.Println("  (Release creation requires an editable Game Center enabled app version)")
	} else {
		releaseID = release.Data.ID
		fmt.Printf("  Release ID: %s\n", releaseID)
	}

	// 3. Create localizations concurrently
	fmt.Printf("  Creating %d localizations concurrently (concurrency=%d)...\n", len(config.Localizations), *concurrency)

	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(*concurrency)

	var mu sync.Mutex
	var successCount int

	for _, locConfig := range config.Localizations {
		locConfig := locConfig // capture loop variable
		g.Go(func() error {
			// Rate limit to avoid API throttling
			if err := rateLimiter.Wait(gCtx); err != nil {
				return fmt.Errorf("rate limiter error for %s: %w", locConfig.Locale, err)
			}

			localization, _, locErr := client.GameCenter.CreateGameCenterAchievementLocalization(gCtx, asc.GameCenterAchievementLocalizationCreateRequestAttributes{
				Locale:                  locConfig.Locale,
				Name:                    locConfig.Name,
				BeforeEarnedDescription: locConfig.BeforeEarnedDescription,
				AfterEarnedDescription:  locConfig.AfterEarnedDescription,
			}, achievement.Data.ID)
			if locErr != nil {
				return fmt.Errorf("failed to create localization for %s: %w", locConfig.Locale, locErr)
			}

			mu.Lock()
			fmt.Printf("    [%s] Localization ID: %s\n", locConfig.Locale, localization.Data.ID)
			mu.Unlock()

			// Upload image if provided
			if locConfig.ImageFile != "" {
				// Rate limit for image upload API calls
				if err := rateLimiter.Wait(gCtx); err != nil {
					return fmt.Errorf("rate limiter error for image %s: %w", locConfig.Locale, err)
				}

				if imgErr := uploadImage(gCtx, client, localization.Data.ID, locConfig.ImageFile); imgErr != nil {
					return fmt.Errorf("failed to upload image for %s: %w", locConfig.Locale, imgErr)
				}

				mu.Lock()
				fmt.Printf("    [%s] Image uploaded successfully\n", locConfig.Locale)
				mu.Unlock()
			}

			mu.Lock()
			successCount++
			mu.Unlock()

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return "", "", err
	}

	fmt.Printf("  All %d localizations created successfully\n", successCount)
	return achievement.Data.ID, releaseID, nil
}

func uploadImage(ctx context.Context, client *asc.Client, localizationID string, imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	fmt.Printf("      File: %s, Size: %d bytes\n", stat.Name(), stat.Size())

	// Reserve
	imageReservation, _, err := client.GameCenter.CreateGameCenterAchievementImage(ctx, asc.GameCenterAchievementImageCreateRequestAttributes{
		FileName: stat.Name(),
		FileSize: int(stat.Size()),
	}, localizationID)
	if err != nil {
		return fmt.Errorf("failed to reserve image: %w", err)
	}
	fmt.Printf("      Image reservation ID: %s\n", imageReservation.Data.ID)

	// Upload
	uploadOperations := imageReservation.Data.Attributes.UploadOperations
	fmt.Printf("      Upload operations count: %d\n", len(uploadOperations))
	for i, op := range uploadOperations {
		if op.URL != nil {
			fmt.Printf("      Operation %d: method=%s, offset=%d, length=%d\n", i, *op.Method, *op.Offset, *op.Length)
			fmt.Printf("      Operation %d URL: %s\n", i, *op.URL)
			fmt.Printf("      Operation %d Headers:\n", i)
			for _, h := range op.RequestHeaders {
				if h.Name != nil && h.Value != nil {
					fmt.Printf("        %s: %s\n", *h.Name, *h.Value)
				}
			}
		}
	}

	if err := client.Upload(ctx, uploadOperations, file); err != nil {
		return fmt.Errorf("failed to upload: %w", err)
	}

	// Commit
	_, _, err = client.GameCenter.UpdateGameCenterAchievementImage(ctx, imageReservation.Data.ID, &asc.GameCenterAchievementImageUpdateRequestAttributes{
		Uploaded: asc.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	return nil
}

func md5Checksum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
