//go:build ignore

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/castbox/asc-go/asc"
	"github.com/castbox/asc-go/examples/util"
)

var (
	bundleID = flag.String("bundleid", "", "Bundle ID to verify (optional)")
)

func main() {
	flag.Parse()

	fmt.Println("========================================")
	fmt.Println("Verifying App Store Connect Credentials")
	fmt.Println("========================================")

	// Create client using util.TokenConfig (reads -kid, -iss, -privatekeypath flags)
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("❌ Failed to create auth config: %s", err)
	}
	fmt.Println("✅ Auth config created successfully")

	client := asc.NewClient(auth.Client())
	fmt.Println("✅ Client created successfully")

	ctx := context.Background()

	// Test 1: List apps (read-only operation)
	fmt.Println("\nTesting API connection...")
	query := &asc.ListAppsQuery{
		Limit: 1,
	}
	if *bundleID != "" {
		query.FilterBundleID = []string{*bundleID}
	}

	apps, _, err := client.Apps.ListApps(ctx, query)
	if err != nil {
		log.Fatalf("❌ API call failed: %s", err)
	}

	fmt.Println("✅ API connection successful!")
	fmt.Printf("   Found %d app(s)\n", len(apps.Data))

	if len(apps.Data) > 0 {
		app := apps.Data[0]
		fmt.Printf("   First app: %s (Bundle ID: %s)\n",
			*app.Attributes.Name,
			*app.Attributes.BundleID)
	}

	fmt.Println("\n========================================")
	fmt.Println("✅ All credentials are valid!")
	fmt.Println("========================================")
}
