package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/castbox/asc-go/asc"
	"github.com/castbox/asc-go/examples/util"
	"github.com/gogf/gf/v2/encoding/gjson"
)

var (
	appid                              = flag.String("appid", "", "ios appid")
	cppName                            = flag.String("cppName", "", "custom product page name")
	ppid                               = flag.String("ppid", "", "customProductPageId, can be find from the url query parameter ppid, for example: https://apps.apple.com/us/app/gurutest/id{appid}?ppid={ppid}")
	customProductPageVersionId         = flag.String("customProductPageVersionId", "", "customProductPageVersionId, can be find from this url: https://appstoreconnect.apple.com/apps/{appid}/distribution/productpages/{customProductPageVersionId}")
	appCustomProductPageLocalizationId = flag.String("appCustomProductPageLocalizationId", "", "appCustomProductPageLocalizationId")
	locale                             = flag.String("locale", "", "locale, for example: en-US, en-GB, ja")
	deleteAppScreenshotsSetId          = flag.String("deleteAppScreenshotsSetId", "", "deleteAppScreenshotsSetId")
	deleteAppScreenshotId              = flag.String("deleteAppScreenshotId", "", "deleteAppScreenshotId")
	screenshotDisplayType              = flag.String("screenshotDisplayType", "", "screenshotDisplayType, https://developer.apple.com/documentation/appstoreconnectapi/screenshotdisplaytype")
	screenshotFile                     = flag.String("screenshotfile", "", "Path to a file to upload as a screenshot")
	sortScreenshotSetId                = flag.String("sortScreenshotSetId", "", "sortScreenshotSetId, the id of the screenshot set to sort")
	sortScreenShotIds                  = flag.String("sortScreenShotIds", "", "sortScreenShotIds, the ids of the screenshots to sort")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	// Create an Authorization header value with bearer token (JWT).
	//    The token is set to expire in 20 minutes, and is used for all App Store
	//    Connect API calls.
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())
	// Get AppCustomProductPages： 获取指定App下的所有CPP页面列表
	customProductPagesRes, res, err := client.CustomProductPage.GetAllAppCustomProductPagesForAnApp(ctx, *appid, &asc.GetAppCustomProductPagesForAnAppQuery{
		Include: []string{"appCustomProductPageVersions"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("customProductPagesRes: %v\nres: %v\n", gjson.MustEncodeString(customProductPagesRes), res)

	// Create AppCustomProductPage： 创建新的CPP页面
	createCustomProductPageRes, res, err := client.CustomProductPage.CreateAppCustomProductPage(ctx, *cppName, *appid, "", nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("customProductPageRes:%v \nres: %v\n", gjson.MustEncodeString(createCustomProductPageRes), res)

	// Get AppCustomProductPage： 根据customProductPageId获取指定CPP页面的信息
	customProductPageRes, res, err := client.CustomProductPage.GetAppCustomProductPage(ctx, *ppid, &asc.GetAppCustomProductPageQuery{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("customProductPageRes: %v\nres: %v\n", gjson.MustEncodeString(customProductPageRes), res)

	// Get AppCustomProductPageVersions： 根据customProductPageId获取指定CPP页面的CustomProductPageVersions信息
	customProductPageVersionsRes, res, err := client.CustomProductPage.GetAppCustomProductPageVersionsByAppCustomProductPageId(ctx, *ppid, &asc.GetAppCustomProductPageVersionsByAppCustomProductPagesIdQuery{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("customProductPageVersionsRes: %v\nres: %v\n", gjson.MustEncodeString(customProductPageVersionsRes), res)

	// 1.Get AppCustomProductPageVersion： 根据customProductPageVersionId获取指定版本的所有语言的信息
	getAppCustomProductPageVersionsRequest := &asc.GetAppCustomProductPageVersionsRequest{
		FieldsAppCustomProductPageLocalizations: nil,
		FieldsAppCustomProductPageVersions:      nil,
		Include:                                 []string{"appCustomProductPage,appCustomProductPageLocalizations"},
		LimitAppCustomProductPageLocalizations:  40,
	}
	appCustomProductPageVersionsRes, res, err := client.CustomProductPage.GetAppCustomProductPageVersion(ctx, *customProductPageVersionId, getAppCustomProductPageVersionsRequest)
	if err != nil {
		fmt.Printf("GetAppCustomProductPageVersion err: %v\n", err)
		return
	}
	fmt.Printf("appCustomProductPageVersionsRes: %v\nres: %v\n", gjson.MustEncodeString(appCustomProductPageVersionsRes), res)

	// 2.根据appCustomProductPageLocalizationId获取指定语言的appScreenshotsSets
	getCustomProductPageLocalizationAppScreenshotSetsRequest := &asc.GetCustomProductPageLocalizationAppScreenshotSetsRequest{
		Include: []string{"appScreenshots"},
	}
	appScreenshotsSetsRes, res, err := client.CustomProductPage.GetCustomProductPageLocalizationAppScreenshotSets(ctx, *appCustomProductPageLocalizationId, getCustomProductPageLocalizationAppScreenshotSetsRequest)
	if err != nil {
		fmt.Printf("GetCustomProductPageLocalizationAppScreenshotSets err: %v\n", err)
		return
	}
	fmt.Printf("appScreenshotsSetsRes: %v\nres: %v\n", gjson.MustEncodeString(appScreenshotsSetsRes), res)

	// 3.Create an App Custom Product Page Localization，获取到appCustomProductPageLocalizationId
	appCustomProductPageLocalizationRes, res, err := client.CustomProductPage.CreateAppCustomProductPageLocalization(ctx, *locale, "", *customProductPageVersionId, nil, nil)
	if err != nil {
		fmt.Printf("CreateAppCustomProductPageLocalization err: %v\n", err)
		return
	}
	fmt.Printf("appCustomProductPageLocalizationRes: %v\nres: %v\n", gjson.MustEncodeString(appCustomProductPageLocalizationRes), res)

	// 4.删除一个已有的 AppScreenShotSet
	res, err = client.Apps.DeleteAppScreenshotSet(ctx, *deleteAppScreenshotsSetId)
	if err != nil {
		fmt.Printf("DeleteAppScreenshotSet err: %v\n", err)
		return
	}
	fmt.Printf("res: %v\n", gjson.MustEncodeString(res.Rate))

	// 5.删除单个截图
	res, err = client.Apps.DeleteAppScreenshot(ctx, *deleteAppScreenshotId)
	if err != nil {
		fmt.Printf("DeleteAppScreenshot err: %v\n", err)
		return
	}
	fmt.Printf("res: %v\n", gjson.MustEncodeString(res.Rate))

	// 6. 创建AppScreenShotSet
	displayType := asc.ScreenshotDisplayType(*screenshotDisplayType)
	appScreenshotsSetRes, res, err := client.Apps.CreateAppScreenshotSet(ctx, displayType, "", *appCustomProductPageLocalizationId, "")
	if err != nil {
		fmt.Printf("CreateAppScreenshotSet err: %v\n", err)
		return
	}
	fmt.Printf("appScreenshotsSetRes: %v\nres: %v\n", gjson.MustEncodeString(appScreenshotsSetRes), res)

	// 7. 上传截图
	appScreenshotSetId := appScreenshotsSetRes.Data.ID
	file, err := os.Open(*screenshotFile)
	if err != nil {
		log.Fatalf("file could not be read: %s", err)
	}
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("file could not be read: %s", err)
	}

	baseFileName := filepath.Base(file.Name())
	fmt.Printf("Reserving space for a new app screenshot. fileName:%s fileSize:%d\n", baseFileName, stat.Size())
	reserveScreenshot, _, err := client.Apps.CreateAppScreenshot(ctx, baseFileName, stat.Size(), appScreenshotSetId)
	if err != nil {
		fmt.Println(err)
	}
	screenshot := reserveScreenshot.Data
	fmt.Printf("ReserveScreenshot: %v\n", gjson.MustEncodeString(screenshot))

	// 8. Upload each part according to the returned upload operations.
	//     The reservation returned uploadOperations, which instructs us how
	//     to split the asset into parts. Upload each part individually.
	//     Note: To speed up the process, upload multiple parts asynchronously
	//     if you have the bandwidth.
	uploadOperations := screenshot.Attributes.UploadOperations
	fmt.Printf("Uploading %d screenshot components\n", len(uploadOperations))

	currentCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	err = client.Upload(currentCtx, uploadOperations, file)
	if err != nil {
		log.Fatalf("file could not be read: %s", err)
	}

	// 9. Commit the reservation and provide a checksum.
	//     Committing tells App Store Connect the script is finished uploading parts.
	//     App Store Connect uses the checksum to ensure the parts were uploaded
	//     successfully.
	fmt.Println("Commit the reservation")
	screenshotURL := screenshot.Links.Self
	checksum, err := md5Checksum(*screenshotFile)
	if err != nil {
		log.Fatalf("file checksum could not be calculated: %s", err)
	}

	appScreenshotRes, res, err := client.Apps.CommitAppScreenshot(ctx, screenshot.ID, asc.Bool(true), &checksum)
	if err != nil {
		fmt.Printf("CommitAppScreenshot err: %v\n", err)
		return
	}

	fmt.Printf("appScreenshotRes: %v\nres: %v\n", gjson.MustEncodeString(appScreenshotRes), res)
	// Report success to the caller.
	fmt.Printf("\nApp Screenshot successfully uploaded to:\n%s\nYou can verify success in App Store Connect or using the API.\n\n", screenshotURL.String())

	// 10.截图排序
	screenshotIds := strings.Split(*sortScreenShotIds, ",")
	res, err = client.Apps.ReplaceAppScreenshotsForSet(ctx, *sortScreenshotSetId, screenshotIds)
	if err != nil {
		fmt.Printf("ReplaceAppScreenshotsForSet err: %v\n", err)
		return
	}
	fmt.Printf("res: %v\n", gjson.MustEncodeString(res.Rate))
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
