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

type AppCustomProductPageCreateRequest struct {
	Data *AppCustomProductPage `json:"data"`
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
	Data     AppCustomProductPage `json:"data"`
	Included []App                `json:"included,omitempty"`
	Links    DocumentLinks        `json:"links"`
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
