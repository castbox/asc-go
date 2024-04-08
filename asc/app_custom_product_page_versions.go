package asc

import (
	"context"
	"fmt"
)

// GetAppCustomProductPageVersionsRequest defines model for GetAppCustomProductPageVersionsRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpageversions_id
type GetAppCustomProductPageVersionsRequest struct {
	FieldsAppCustomProductPageLocalizations []string `url:"fields[appCustomProductPageLocalizations],omitempty"`
	FieldsAppCustomProductPageVersions      []string `url:"fields[appCustomProductPageVersions],omitempty"`
	Include                                 []string `url:"include,omitempty"`
	LimitAppCustomProductPageLocalizations  int      `url:"limit[appCustomProductPageLocalizations],omitempty"`
}

// AppCustomProductPageVersionResponse defines model for AppCustomProductPageVersionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversionresponse
type AppCustomProductPageVersionResponse struct {
	Data     *AppCustomProductPageVersion       `json:"data"`
	Included []AppCustomProductPageLocalization `json:"included"`
	Links    DocumentLinks                      `json:"links"`
}

// AppCustomProductPageVersionsResponse defines model for AppCustomProductPageVersionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversionsresponse
type AppCustomProductPageVersionsResponse struct {
	Data     []AppCustomProductPageVersion      `json:"data"`
	Included []AppCustomProductPageLocalization `json:"included,omitempty"`
	Links    PagedDocumentLinks                 `json:"links"`
	Meta     *PagingInformation                 `json:"meta,omitempty"`
}

// AppCustomProductPageVersion defines model for AppCustomProductPageVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversion
type AppCustomProductPageVersion struct {
	Attributes    *AppCustomProductPageVersionAttributes    `json:"attributes,omitempty"`
	ID            string                                    `json:"id"`
	Links         *ResourceLinks                            `json:"links,omitempty"`
	Relationships *AppCustomProductPageVersionRelationships `json:"relationships,omitempty"`
	Type          string                                    `json:"type"`
}

// AppCustomProductPageVersionAttributes defines model for AppCustomProductPageVersionAttributes.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversion/attributes
type AppCustomProductPageVersionAttributes struct {
	State   string `json:"state,omitempty"`
	Version string `json:"version,omitempty"`
}

// AppCustomProductPageVersionRelationships defines model for AppCustomProductPageVersionRelationships.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversion/relationships
type AppCustomProductPageVersionRelationships struct {
	AppCustomProductPage              *RelationshipsAppCustomProductPage `json:"appCustomProductPage,omitempty"`
	AppCustomProductPageLocalizations *AppCustomProductPageLocalizations `json:"appCustomProductPageLocalizations,omitempty"`
}

// RelationshipsAppCustomProductPage defines model for RelationshipsAppCustomProductPage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversion/relationships/appcustomproductpage
type RelationshipsAppCustomProductPage struct {
	Data  *AppCustomProductPageVersionData  `json:"data,omitempty"`
	Links *AppCustomProductPageVersionLinks `json:"links,omitempty"`
}

// AppCustomProductPageLocalizations defines model for AppCustomProductPageLocalizations.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversion/relationships/appcustomproductpagelocalizations
type AppCustomProductPageLocalizations struct {
	Data  []AppCustomProductPageLocalizationsData `json:"data,omitempty"`
	Links *AppCustomProductPageVersionLinks       `json:"links,omitempty"`
	Meta  *PagingInformation                      `json:"meta,omitempty"`
}

// AppCustomProductPageLocalizationsData defines model for AppCustomProductPageLocalizationsData.
// https://developer.apple.com/documentation/appstoreconnectapi/appcustomproductpageversion/relationships/appcustomproductpagelocalizations/data
type AppCustomProductPageLocalizationsData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// GetAppCustomProductPageVersion get AppCustomProductPageVersions
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_v1_appcustomproductpageversions_id
func (s *AppCustomProductPageService) GetAppCustomProductPageVersion(ctx context.Context, id string, req *GetAppCustomProductPageVersionsRequest) (*AppCustomProductPageVersionResponse, *Response, error) {
	url := fmt.Sprintf("/v1/appCustomProductPageVersions/%s", id)
	res := new(AppCustomProductPageVersionResponse)
	resp, err := s.client.get(ctx, url, req, res)
	if err != nil {
		return nil, nil, err
	}
	return res, resp, nil
}
