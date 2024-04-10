package asc

import (
	"context"
	"fmt"
)

// AppCustomProductPageLocalization defines model for AppCustomProductPageLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalization
type AppCustomProductPageLocalization struct {
	Attributes    *AppCustomProductPageLocalizationAttributes    `json:"attributes,omitempty"`
	ID            string                                         `json:"id"`
	Links         *ResourceLinks                                 `json:"links"`
	Relationships *AppCustomProductPageLocalizationRelationships `json:"relationships,omitempty"`
	Type          string                                         `json:"type"`
}

// AppCustomProductPageLocalizationAttributes defines model for AppCustomProductPageLocalizationAttributes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalization/attributes
type AppCustomProductPageLocalizationAttributes struct {
	Locale          string `json:"locale"`
	PromotionalText string `json:"promotionalText,omitempty"`
}

// AppCustomProductPageLocalizationRelationships defines model for AppCustomProductPageLocalizationRelationships.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalization/relationships
type AppCustomProductPageLocalizationRelationships struct {
	AppCustomProductPageVersion *RelationshipsAppCustomProductPageVersion `json:"appCustomProductPageVersion,omitempty"`
	AppPreviewSets              *PagedRelationship                        `json:"appPreviewSets,omitempty"`
	AppScreenshotSets           *PagedRelationship                        `json:"appScreenshotSets,omitempty"`
}

// RelationshipsAppCustomProductPageVersion defines model for RelationshipsAppCustomProductPageVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalization/relationships/appcustomproductpageversion
type RelationshipsAppCustomProductPageVersion struct {
	Data  *AppCustomProductPageVersionData  `json:"data,omitempty"`
	Links *AppCustomProductPageVersionLinks `json:"links,omitempty"`
}

// AppCustomProductPageVersionData defines model for AppCustomProductPageVersionData.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalization/relationships/appcustomproductpageversion/data
type AppCustomProductPageVersionData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// AppCustomProductPageVersionLinks defines model for AppCustomProductPageVersionLinks.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalization/relationships/appcustomproductpageversion/links
type AppCustomProductPageVersionLinks struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
}

// AppCustomProductPageLocalizationCreateRequest defines model for AppCustomProductPageLocalizationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationcreaterequest
type AppCustomProductPageLocalizationCreateRequest struct {
	Data *AppCustomProductPageLocalizationCreateRequestData `json:"data"`
}

// AppCustomProductPageLocalizationCreateRequestData defines model for AppCustomProductPageLocalizationCreateRequestData.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationcreaterequest/data
type AppCustomProductPageLocalizationCreateRequestData struct {
	Attributes    *AppCustomProductPageLocalizationAttributes    `json:"attributes"`
	Relationships *AppCustomProductPageLocalizationRelationships `json:"relationships"`
	Type          string                                         `json:"type"`
}

// AppCustomProductPageLocalizationResponse defines model for AppCustomProductPageLocalizationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationresponse
type AppCustomProductPageLocalizationResponse struct {
	Data     *AppCustomProductPageLocalization `json:"data"`
	Included []AppScreenshotSet                `json:"included"`
	Links    DocumentLinks                     `json:"links"`
}

// AppCustomProductPageLocalizationUpdateRequest defines model for AppCustomProductPageLocalizationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationupdaterequest
type AppCustomProductPageLocalizationUpdateRequest struct {
	Data *AppCustomProductPageLocalizationUpdateRequestData `json:"data"`
}

type AppCustomProductPageLocalizationUpdateRequestData struct {
	Attributes *AppCustomProductPageLocalizationAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

// AppCustomProductPageLocalizationsResponse defines model for AppCustomProductPageLocalizationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationsresponse
type AppCustomProductPageLocalizationsResponse struct {
	Data     []AppCustomProductPageLocalization `json:"data"`
	Included []AppScreenshotSet                 `json:"included,omitempty"`
	Links    DocumentLinks                      `json:"links"`
	Meta     *PagingInformation                 `json:"meta,omitempty"`
}

// GetCustomProductPageLocalizationAppScreenshotSetsRequest defines model for GetCustomProductPageLocalizationAppScreenshotSetsRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpagelocalizations_id_appscreenshotsets
type GetCustomProductPageLocalizationAppScreenshotSetsRequest struct {
	FieldsAppScreenshotSets                               []string `url:"fields[appScreenshotSets],omitempty"`
	FieldsAppScreenshots                                  []string `url:"fields[appScreenshots],omitempty"`
	FilterAppStoreVersionExperimentTreatmentLocalization  []string `url:"filter[appStoreVersionExperimentTreatmentLocalization],omitempty"`
	FilterAppStoreVersionLocalization                     []string `url:"filter[appStoreVersionLocalization],omitempty"`
	FilterScreenshotDisplayType                           []string `url:"filter[screenshotDisplayType],omitempty"`
	Include                                               []string `url:"include,omitempty"`
	Limit                                                 int      `url:"limit,omitempty"`
	LimitAppScreenshots                                   int      `url:"limit[appScreenshots],omitempty"`
	FieldsAppCustomProductPageLocalizations               []string `url:"fields[appCustomProductPageLocalizations],omitempty"`
	FieldsAppStoreVersionExperimentTreatmentLocalizations []string `url:"fields[appStoreVersionExperimentTreatmentLocalizations],omitempty"`
	FieldsAppStoreVersionLocalizations                    []string `url:"fields[appStoreVersionLocalizations],omitempty"`
}

// AppCustomProductPageLocalizationInlineCreate defines model for AppCustomProductPageLocalizationInlineCreate
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationinlinecreate
type AppCustomProductPageLocalizationInlineCreate struct {
	Attributes    *AppCustomProductPageLocalizationInlineCreateAttributes    `json:"attributes"`
	Id            string                                                     `json:"id,omitempty"`
	Relationships *AppCustomProductPageLocalizationInlineCreateRelationships `json:"relationships,omitempty"`
	Type          string                                                     `json:"type"`
}

// AppCustomProductPageLocalizationInlineCreateAttributes defines model for AppCustomProductPageLocalizationInlineCreate.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationinlinecreate/attributes
type AppCustomProductPageLocalizationInlineCreateAttributes struct {
	Locale          string `json:"locale"`
	PromotionalText string `json:"promotionalText"`
}

// AppCustomProductPageLocalizationInlineCreateRelationships defines model for AppCustomProductPageLocalizationInlineCreate.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagelocalizationinlinecreate/relationships
type AppCustomProductPageLocalizationInlineCreateRelationships struct {
	AppCustomProductPageVersion *RelationShipAppCustomProductPageVersion `json:"appCustomProductPageVersion"`
}

type RelationShipAppCustomProductPageVersion struct {
	Data *RelationshipData `json:"data"`
}

// GetCustomProductPageLocalizationAppScreenshotSets
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpagelocalizations_id_appscreenshotsets
func (s *AppCustomProductPageService) GetCustomProductPageLocalizationAppScreenshotSets(ctx context.Context, id string, params *GetCustomProductPageLocalizationAppScreenshotSetsRequest) (*AppScreenshotSetsResponse, *Response, error) {
	url := fmt.Sprintf("/v1/appCustomProductPageLocalizations/%s/appScreenshotSets", id)
	res := new(AppScreenshotSetsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	if err != nil {
		return nil, nil, err
	}
	return res, resp, nil
}

// CreatAppCustomProductPageLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/post_v1_appcustomproductpagelocalizations
func (s *AppCustomProductPageService) CreatAppCustomProductPageLocalization(ctx context.Context, local, promotionalText, appCustomProductPageVersionId string, appPreviewSet, screenshotSet *PagedRelationship) (*AppCustomProductPageLocalizationResponse, *Response, error) {
	req := &AppCustomProductPageLocalizationCreateRequest{
		Data: &AppCustomProductPageLocalizationCreateRequestData{
			Type: "appCustomProductPageLocalizations",
			Attributes: &AppCustomProductPageLocalizationAttributes{
				Locale:          local,
				PromotionalText: promotionalText,
			},
			Relationships: &AppCustomProductPageLocalizationRelationships{
				AppCustomProductPageVersion: &RelationshipsAppCustomProductPageVersion{
					Data: &AppCustomProductPageVersionData{
						Id:   appCustomProductPageVersionId,
						Type: "appCustomProductPageVersions",
					},
				},
				AppPreviewSets:    appPreviewSet,
				AppScreenshotSets: screenshotSet,
			},
		},
	}

	url := fmt.Sprintf("/v1/appCustomProductPageLocalizations")
	res := new(AppCustomProductPageLocalizationResponse)
	resp, err := s.client.post(ctx, url, newRequestBody(req.Data), res)
	if err != nil {
		return nil, nil, err
	}
	return res, resp, nil
}
