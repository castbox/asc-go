package asc

import (
	"context"
	"fmt"
)

// AppCustomProductPageService defines model for AppCustomProductPageService.
// https://developer.apple.com/documentation/appstoreconnectapi/app_store/custom_product_pages_and_localizations
// https://developer.apple.com/documentation/appstoreconnectapi/app_store/custom_product_pages_and_localizations/app_custom_product_page_localizations
type AppCustomProductPageService service

// AppCustomProductPage defines model for AppCustomProductPage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/app_store/custom_product_pages_and_localizations/app_custom_product_pages
type AppCustomProductPage struct {
	Attributes    *AppCustomProductPageAttributes    `json:"attributes,omitempty"`
	ID            string                             `json:"id"`
	Links         ResourceLinks                      `json:"links"`
	Relationships *AppCustomProductPageRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

type AppCustomProductPageAttributes struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Visible bool   `json:"visible"`
}

type AppCustomProductPageRelationships struct {
	App      *AppCustomProductPageRelationshipsApp      `json:"app,omitempty"`
	Versions *AppCustomProductPageRelationshipsVersions `json:"versions,omitempty"`
}

type AppCustomProductPageRelationshipsApp struct {
	Data  *AppCustomProductPageRelationshipsAppData  `json:"data,omitempty"`
	Links *AppCustomProductPageRelationshipsAppLinks `json:"links,omitempty"`
}

type AppCustomProductPageRelationshipsAppData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AppCustomProductPageRelationshipsAppLinks struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
}

type AppCustomProductPageRelationshipsVersions struct {
	Data  []*AppCustomProductPageRelationshipsVersionsData `json:"data,omitempty"`
	Links *AppCustomProductPageRelationshipsVersionsLinks  `json:"links,omitempty"`
	Meta  *PagingInformation                               `json:"meta,omitempty"`
}

type AppCustomProductPageRelationshipsVersionsData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AppCustomProductPageRelationshipsVersionsLinks struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
}

// appCustomProductPageCreateRequest defines model for appCustomProductPageCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest
type appCustomProductPageCreateRequest struct {
	Data     *AppCustomProductPageCreateRequestData      `json:"data"`
	Included []AppCustomProductPageCreateRequestIncluded `json:"included"`
}

type AppCustomProductPageCreateRequestIncluded interface{}

// AppCustomProductPageCreateRequestData defines model for appCustomProductPageCreateRequest.Data
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data
type AppCustomProductPageCreateRequestData struct {
	Attributes    *AppCustomProductPageCreateRequestDataAttributes    `json:"attributes"`
	Relationships *AppCustomProductPageCreateRequestDataRelationships `json:"relationships"`
	Type          string                                              `json:"type"`
}

// AppCustomProductPageCreateRequestDataAttributes defines model for appCustomProductPageCreateRequest.Data.Attributes
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data/attributes
type AppCustomProductPageCreateRequestDataAttributes struct {
	Name string `json:"name"`
}

// AppCustomProductPageCreateRequestDataRelationships defines model for appCustomProductPageCreateRequest.Data.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data/relationships
type AppCustomProductPageCreateRequestDataRelationships struct {
	App                          *AppCustomProductPageCreateRequestDataRelationshipsApp                          `json:"app"`
	AppCustomProductPageVersions *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions `json:"appCustomProductPageVersions,omitempty"`
	AppStoreVersionTemplate      *AppCustomProductPageCreateRequestDataRelationshipsAppStoreVersionTemplate      `json:"appStoreVersionTemplate,omitempty"`
	CustomProductPageTemplate    *AppCustomProductPageCreateRequestDataRelationshipsCustomProductPageTemplate    `json:"customProductPageTemplate,omitempty"`
}

// AppCustomProductPageCreateRequestDataRelationshipsApp defines model for appCustomProductPageCreateRequest.Data.Relationships.App
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data/relationships/app
type AppCustomProductPageCreateRequestDataRelationshipsApp struct {
	Data *RelationshipData `json:"data"`
}

// AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions defines model for appCustomProductPageCreateRequest.Data.Relationships.AppCustomProductPageVersions
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data/relationships/appcustomproductpageversions
type AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions struct {
	Data []*RelationshipData `json:"data,omitempty"`
}

// AppCustomProductPageCreateRequestDataRelationshipsAppStoreVersionTemplate defines model for appCustomProductPageCreateRequest.Data.Relationships.AppStoreVersionTemplate
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data/relationships/appstoreversiontemplate
type AppCustomProductPageCreateRequestDataRelationshipsAppStoreVersionTemplate struct {
	Data *RelationshipData `json:"data,omitempty"`
}

// AppCustomProductPageCreateRequestDataRelationshipsCustomProductPageTemplate defines model for appCustomProductPageCreateRequest.Data.Relationships.AppCustomProductPageVersions
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagecreaterequest/data/relationships/appcustomproductpageversions
type AppCustomProductPageCreateRequestDataRelationshipsCustomProductPageTemplate struct {
	Data *RelationshipData `json:"data,omitempty"`
}

type GetAppCustomProductPageQuery struct {
	FieldsAppCustomProductPageVersions []string `url:"fields[appCustomProductPageVersions],omitempty"`
	FieldsAppCustomProductPages        []string `url:"fields[appCustomProductPages],omitempty"`
	Include                            []string `url:"include,omitempty"`
	LimitAppCustomProductPageVersions  int      `url:"limit[appCustomProductPageVersions],omitempty"`
}

// AppCustomProductPageResponse defines model for AppCustomProductPageResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagesresponse
type AppCustomProductPageResponse struct {
	Data     AppCustomProductPage          `json:"data"`
	Included []AppCustomProductPageVersion `json:"included,omitempty"`
	Links    DocumentLinks                 `json:"links"`
}

// AppCustomProductPagesResponse defines model for AppCustomProductPagesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpagesresponse
type AppCustomProductPagesResponse struct {
	Data     []AppCustomProductPage        `json:"data"`
	Included []AppCustomProductPageVersion `json:"included,omitempty"`
	Links    PagedDocumentLinks            `json:"links"`
	Meta     *PagingInformation            `json:"meta,omitempty"`
}

// GetAppCustomProductPagesForAnAppQuery defines model for GetAppCustomProductPagesForAnAppQuery.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_custom_product_pages_for_an_app/
type GetAppCustomProductPagesForAnAppQuery struct {
	FieldsAppCustomProductPageVersions []string `url:"fields[appCustomProductPageVersions],omitempty"`
	FieldsAppCustomProductPages        []string `url:"fields[appCustomProductPages],omitempty"`
	filterVisible                      []string `url:"filter[visible],omitempty"`
	Include                            []string `url:"include,omitempty"`
	Limit                              int      `url:"limit,omitempty"`
	LimitAppCustomProductPageVersions  int      `url:"limit[appCustomProductPageVersions],omitempty"`
	FieldsApps                         []string `url:"fields[apps],omitempty"`
}

// GetAppCustomProductPageVersionsByAppCustomProductPagesIdQuery defines model for GetAppCustomProductPageVersionsByAppCustomProductPagesIdQuery.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpages_id_appcustomproductpageversions/
type GetAppCustomProductPageVersionsByAppCustomProductPagesIdQuery struct {
	FieldsAppCustomProductPageLocalizations []string `url:"fields[appCustomProductPageLocalizations],omitempty"`
	FieldsAppCustomProductPageVersions      []string `url:"fields[appCustomProductPageVersions],omitempty"`
	FilterState                             []string `url:"filter[state],omitempty"`
	Include                                 []string `url:"include,omitempty"`
	Limit                                   int      `url:"limit,omitempty"`
	LimitAppCustomProductPageLocalizations  int      `url:"limit[appCustomProductPageLocalizations],omitempty"`
	FieldsAppCustomProductPages             []string `url:"fields[appCustomProductPages],omitempty"`
}

// GetAppCustomProductPage  get app custom product page by id
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpages_id
func (s *AppCustomProductPageService) GetAppCustomProductPage(ctx context.Context, id string, params *GetAppCustomProductPageQuery) (*AppCustomProductPageResponse, *Response, error) {
	url := fmt.Sprintf("/v1/appCustomProductPages/%s", id)
	res := new(AppCustomProductPageResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAllAppCustomProductPagesForAnApp get all app custom product pages for an app
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_custom_product_pages_for_an_app/
func (s *AppCustomProductPageService) GetAllAppCustomProductPagesForAnApp(ctx context.Context, appId string, params *GetAppCustomProductPagesForAnAppQuery) (*AppCustomProductPagesResponse, *Response, error) {
	url := fmt.Sprintf("/v1/apps/%s/appCustomProductPages", appId)
	res := new(AppCustomProductPagesResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// GetAppCustomProductPageVersionsByAppCustomProductPageId get app custom product page versions by app custom product pages id
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpages_id_appcustomproductpageversions/
func (s *AppCustomProductPageService) GetAppCustomProductPageVersionsByAppCustomProductPageId(ctx context.Context, customProductPageId string,
	params *GetAppCustomProductPageVersionsByAppCustomProductPagesIdQuery) (*AppCustomProductPageVersionsResponse, *Response, error) {
	url := fmt.Sprintf("/v1/appCustomProductPages/%s/appCustomProductPageVersions", customProductPageId)
	res := new(AppCustomProductPageVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreateAppCustomProductPage create an app custom product page
//
// https://developer.apple.com/documentation/appstoreconnectapi/post_v1_appcustomproductpages
func (s *AppCustomProductPageService) CreateAppCustomProductPage(ctx context.Context, pageName, appid, enPromotionalText string,
	appStoreVersionTemplateData, customProductPageTemplateData *RelationshipData) (*AppCustomProductPageResponse, *Response, error) {
	req := &appCustomProductPageCreateRequest{
		Data: &AppCustomProductPageCreateRequestData{
			Attributes: &AppCustomProductPageCreateRequestDataAttributes{
				Name: pageName,
			},
			Relationships: &AppCustomProductPageCreateRequestDataRelationships{
				App: &AppCustomProductPageCreateRequestDataRelationshipsApp{
					Data: &RelationshipData{
						ID:   appid,
						Type: "apps",
					},
				},
				AppCustomProductPageVersions: &AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions{
					Data: []*RelationshipData{
						{
							ID:   "${new-appCustomProductPageVersion-id}",
							Type: "appCustomProductPageVersions",
						},
					},
				},
				AppStoreVersionTemplate: &AppCustomProductPageCreateRequestDataRelationshipsAppStoreVersionTemplate{
					Data: appStoreVersionTemplateData,
				},
				CustomProductPageTemplate: &AppCustomProductPageCreateRequestDataRelationshipsCustomProductPageTemplate{
					Data: customProductPageTemplateData,
				},
			},
			Type: "appCustomProductPages",
		},
		Included: []AppCustomProductPageCreateRequestIncluded{
			AppCustomProductPageVersionInlineCreate{
				Type: "appCustomProductPageVersions",
				Id:   "${new-appCustomProductPageVersion-id}",
				Relationships: &AppCustomProductPageVersionInlineCreateRelationships{
					AppCustomProductPage: nil,
					AppCustomProductPageLocalizations: &RelationShipAppCustomProductPageLocalizations{
						Data: []*RelationshipData{
							{
								ID:   "${new-appCustomProductPageLocalization-id}",
								Type: "appCustomProductPageLocalizations",
							},
						},
					},
				},
			},
			AppCustomProductPageLocalizationInlineCreate{
				Id: "${new-appCustomProductPageLocalization-id}",
				Attributes: &AppCustomProductPageLocalizationInlineCreateAttributes{
					Locale:          "en-US",
					PromotionalText: enPromotionalText,
				},
				Type: "appCustomProductPageLocalizations",
			},
		},
	}

	url := "/v1/appCustomProductPages"
	res := new(AppCustomProductPageResponse)
	resp, err := s.client.post(ctx, url, newRequestBodyWithIncluded(req.Data, req.Included), res)
	return res, resp, err
}

// DeleteAnAppCustomProductPage delete an app custom product page
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_custom_product_page
func (s *AppCustomProductPageService) DeleteAnAppCustomProductPage(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("/v1/appCustomProductPages/%s", id)

	return s.client.delete(ctx, url, nil)
}
