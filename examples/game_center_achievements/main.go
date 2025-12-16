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
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/castbox/asc-go/asc"
	"github.com/castbox/asc-go/examples/util"
)

var (
	bundleID = flag.String("bundleid", "", "Bundle ID for an app (required)")

	// Achievement attributes
	referenceName    = flag.String("name", "", "Reference name for the achievement (required)")
	vendorIdentifier = flag.String("vendor", "", "Vendor identifier for the achievement (required)")
	points           = flag.Int("points", 10, "Points for the achievement (1-100)")
	showBeforeEarned = flag.Bool("showbefore", true, "Show achievement before earned")
	repeatable       = flag.Bool("repeatable", false, "Achievement is repeatable")

	// Localization attributes
	locale                  = flag.String("locale", "en-US", "Locale for the achievement")
	localizedName           = flag.String("localizedname", "", "Localized name for the achievement (required)")
	beforeEarnedDescription = flag.String("beforedesc", "", "Description before earning (required)")
	afterEarnedDescription  = flag.String("afterdesc", "", "Description after earning (required)")

	// Image
	imageFile = flag.String("imagefile", "", "Path to achievement image file (512x512 PNG, optional)")
)

func main() {
	flag.Parse()

	// Validate required flags
	if *bundleID == "" {
		log.Fatal("bundleid is required")
	}
	if *referenceName == "" {
		log.Fatal("name is required")
	}
	if *vendorIdentifier == "" {
		log.Fatal("vendor is required")
	}
	if *localizedName == "" {
		log.Fatal("localizedname is required")
	}
	if *beforeEarnedDescription == "" {
		log.Fatal("beforedesc is required")
	}
	if *afterEarnedDescription == "" {
		log.Fatal("afterdesc is required")
	}

	ctx := context.Background()

	// 1. Create an Authorization header value with bearer token (JWT).
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	// 2. Look up the app by bundle ID.
	fmt.Printf("Looking up app with bundle ID: %s\n", *bundleID)
	app, err := util.GetApp(ctx, client, &asc.ListAppsQuery{
		FilterBundleID: []string{*bundleID},
	})
	if err != nil {
		log.Fatalf("Failed to find app: %s", err)
	}
	fmt.Printf("Found app: %s (ID: %s)\n", *app.Attributes.Name, app.ID)

	// 3. Get or create Game Center detail for the app.
	fmt.Println("Getting Game Center detail for app...")
	gameCenterDetail, _, err := client.GameCenter.GetGameCenterDetailForApp(ctx, app.ID, nil)
	if err != nil {
		// If Game Center is not enabled, create it
		fmt.Println("Game Center not enabled, enabling...")
		gameCenterDetail, _, err = client.GameCenter.CreateGameCenterDetail(ctx, app.ID)
		if err != nil {
			log.Fatalf("Failed to enable Game Center: %s", err)
		}
	}
	fmt.Printf("Game Center Detail ID: %s\n", gameCenterDetail.Data.ID)

	// 4. Create the achievement.
	fmt.Printf("Creating achievement: %s\n", *referenceName)
	achievement, _, err := client.GameCenter.CreateGameCenterAchievement(ctx, asc.GameCenterAchievementCreateRequestAttributes{
		ReferenceName:    *referenceName,
		VendorIdentifier: *vendorIdentifier,
		Points:           *points,
		ShowBeforeEarned: *showBeforeEarned,
		Repeatable:       *repeatable,
	}, gameCenterDetail.Data.ID)
	if err != nil {
		log.Fatalf("Failed to create achievement: %s", err)
	}
	fmt.Printf("Achievement created with ID: %s\n", achievement.Data.ID)

	// 5. Create localization for the achievement.
	fmt.Printf("Creating localization for locale: %s\n", *locale)
	localization, _, err := client.GameCenter.CreateGameCenterAchievementLocalization(ctx, asc.GameCenterAchievementLocalizationCreateRequestAttributes{
		Locale:                  *locale,
		Name:                    *localizedName,
		BeforeEarnedDescription: *beforeEarnedDescription,
		AfterEarnedDescription:  *afterEarnedDescription,
	}, achievement.Data.ID)
	if err != nil {
		log.Fatalf("Failed to create localization: %s", err)
	}
	fmt.Printf("Localization created with ID: %s\n", localization.Data.ID)

	// 6. Upload achievement image if provided.
	if *imageFile != "" {
		fmt.Printf("Uploading achievement image: %s\n", *imageFile)
		err = uploadAchievementImage(ctx, client, localization.Data.ID, *imageFile)
		if err != nil {
			log.Fatalf("Failed to upload image: %s", err)
		}
		fmt.Println("Achievement image uploaded successfully!")
	}

	// 7. Summary
	fmt.Println("\n========================================")
	fmt.Println("Achievement created successfully!")
	fmt.Println("========================================")
	fmt.Printf("Achievement ID:     %s\n", achievement.Data.ID)
	fmt.Printf("Reference Name:     %s\n", *referenceName)
	fmt.Printf("Vendor Identifier:  %s\n", *vendorIdentifier)
	fmt.Printf("Points:             %d\n", *points)
	fmt.Printf("Show Before Earned: %t\n", *showBeforeEarned)
	fmt.Printf("Repeatable:         %t\n", *repeatable)
	fmt.Printf("Locale:             %s\n", *locale)
	fmt.Printf("Localized Name:     %s\n", *localizedName)
	fmt.Println("========================================")
}

func uploadAchievementImage(ctx context.Context, client *asc.Client, localizationID string, imagePath string) error {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %w", err)
	}
	defer util.Close(file)

	// Get file info
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	// 1. Reserve the image upload
	fmt.Println("  Reserving space for achievement image...")
	imageReservation, _, err := client.GameCenter.CreateGameCenterAchievementImage(ctx, asc.GameCenterAchievementImageCreateRequestAttributes{
		FileName: stat.Name(),
		FileSize: int(stat.Size()),
	}, localizationID)
	if err != nil {
		return fmt.Errorf("failed to reserve image: %w", err)
	}
	fmt.Printf("  Image reservation ID: %s\n", imageReservation.Data.ID)

	// 2. Upload the image parts
	uploadOperations := imageReservation.Data.Attributes.UploadOperations
	fmt.Printf("  Uploading %d image components...\n", len(uploadOperations))
	err = client.Upload(ctx, uploadOperations, file)
	if err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	// 3. Calculate checksum
	checksum, err := md5Checksum(imagePath)
	if err != nil {
		return fmt.Errorf("failed to calculate checksum: %w", err)
	}
	fmt.Printf("  Image checksum: %s\n", checksum)

	// 4. Commit the upload
	fmt.Println("  Committing image upload...")
	_, _, err = client.GameCenter.UpdateGameCenterAchievementImage(ctx, imageReservation.Data.ID, &asc.GameCenterAchievementImageUpdateRequestAttributes{
		Uploaded: asc.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to commit image: %w", err)
	}

	return nil
}

func md5Checksum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer util.Close(f)

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
