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

	"github.com/castbox/asc-go/asc"
	"github.com/castbox/asc-go/examples/util"
)

var (
	bundleID   = flag.String("bundleid", "", "Bundle ID for an app (required)")
	configFile = flag.String("config", "", "Path to JSON config file with achievements (required)")
)

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

	// Get or create Game Center detail
	fmt.Println("Getting Game Center detail...")
	gameCenterDetail, _, err := client.GameCenter.GetGameCenterDetailForApp(ctx, app.ID, &asc.GetGameCenterDetailForAppQuery{
		Include: []string{"gameCenterGroup"},
	})
	if err != nil {
		fmt.Println("Game Center not enabled, enabling...")
		gameCenterDetail, _, err = client.GameCenter.CreateGameCenterDetail(ctx, app.ID)
		if err != nil {
			log.Fatalf("Failed to enable Game Center: %s", err)
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

	// Create each achievement and collect new release IDs
	var createdAchievements []string
	var newReleases []struct {
		achievementID string
		releaseID     string
		position      int
		name          string
	}
	for i, achConfig := range config.Achievements {
		fmt.Printf("========================================\n")
		fmt.Printf("Creating achievement %d/%d: %s\n", i+1, len(config.Achievements), achConfig.ReferenceName)
		fmt.Printf("========================================\n")

		achievementID, releaseID, err := createAchievementWithRelease(ctx, client, gameCenterDetail.Data.ID, gameCenterGroupID, achConfig)
		if err != nil {
			log.Printf("Failed to create achievement %s: %s", achConfig.ReferenceName, err)
			continue
		}
		createdAchievements = append(createdAchievements, achievementID)
		if releaseID != "" {
			newReleases = append(newReleases, struct {
				achievementID string
				releaseID     string
				position      int
				name          string
			}{achievementID, releaseID, achConfig.Position, achConfig.ReferenceName})
		}
		fmt.Printf("Achievement created successfully: %s\n\n", achievementID)
	}

	// Reorder achievements based on position field
	if len(newReleases) > 0 && len(existingReleaseIDs)+len(newReleases) > 0 {
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
		_, err := client.GameCenter.ReplaceGameCenterAchievementReleasesForDetail(ctx, gameCenterDetail.Data.ID, finalOrder)
		if err != nil {
			fmt.Printf("Note: Could not reorder achievements: %v\n", err)
			fmt.Println("(Reordering requires an editable Game Center enabled app version)")
		} else {
			fmt.Println("Achievements reordered successfully!")
		}
		fmt.Println()
	}

	// Summary
	fmt.Println("\n========================================")
	fmt.Println("Batch Creation Complete!")
	fmt.Println("========================================")
	fmt.Printf("Total achievements in config: %d\n", len(config.Achievements))
	fmt.Printf("Successfully created: %d\n", len(createdAchievements))
	fmt.Printf("Releases created: %d\n", len(newReleases))
	fmt.Printf("Failed: %d\n", len(config.Achievements)-len(createdAchievements))
	fmt.Println("========================================")
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

	// 3. Create localizations
	for _, locConfig := range config.Localizations {
		fmt.Printf("  Creating localization: %s\n", locConfig.Locale)
		localization, _, locErr := client.GameCenter.CreateGameCenterAchievementLocalization(ctx, asc.GameCenterAchievementLocalizationCreateRequestAttributes{
			Locale:                  locConfig.Locale,
			Name:                    locConfig.Name,
			BeforeEarnedDescription: locConfig.BeforeEarnedDescription,
			AfterEarnedDescription:  locConfig.AfterEarnedDescription,
		}, achievement.Data.ID)
		if locErr != nil {
			return "", "", fmt.Errorf("failed to create localization for %s: %w", locConfig.Locale, locErr)
		}
		fmt.Printf("    Localization ID: %s\n", localization.Data.ID)

		// 4. Upload image if provided
		if locConfig.ImageFile != "" {
			fmt.Printf("    Uploading image: %s\n", locConfig.ImageFile)
			if imgErr := uploadImage(ctx, client, localization.Data.ID, locConfig.ImageFile); imgErr != nil {
				return "", "", fmt.Errorf("failed to upload image for %s: %w", locConfig.Locale, imgErr)
			}
			fmt.Println("    Image uploaded successfully")
		}
	}

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
