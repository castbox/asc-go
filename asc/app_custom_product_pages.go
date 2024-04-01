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
	Meta  *PaginationMeta                                  `json:"meta,omitempty"`
}

type AppCustomProductPageRelationshipsVersionsData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AppCustomProductPageRelationshipsVersionsLinks struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
}

type PaginationMeta struct {
	Paging *Paging `json:"paging,omitempty"`
}

type Paging struct {
	Total int `json:"total,omitempty"`
	Limit int `json:"limit,omitempty"`
}

type AppCustomProductPageCreateRequest struct {
	Data *AppCustomProductPage `json:"data"`
}

type GetAppCustomProductPagesQuery struct {
	FieldsAppCustomProductPageVersions []string `url:"fields[appCustomProductPageVersions],omitempty"`
	FieldsAppCustomProductPages        []string `url:"fields[appCustomProductPages],omitempty"`
	Include                            []string `url:"include,omitempty"`
	LimitAppCustomProductPageVersions  int      `url:"limit[appCustomProductPageVersions],omitempty"`
}

type AppCustomProductPageResponse struct {
	Data     AppCustomProductPage `json:"data"`
	Included []App                `json:"included,omitempty"`
	Links    DocumentLinks        `json:"links"`
}

func (s *AppCustomProductPageService) GetAppCustomProductPages(ctx context.Context, id string, params *GetAppCustomProductPagesQuery) (*AppCustomProductPageResponse, *Response, error) {
	url := fmt.Sprintf("appCustomProductPages/%s", id)
	res := new(AppCustomProductPageResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
