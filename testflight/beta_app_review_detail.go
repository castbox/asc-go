package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/apps"
	"github.com/aaronsky/asc-go/internal/services"
	"github.com/aaronsky/asc-go/internal/types"
)

// BetaAppReviewDetail defines model for BetaAppReviewDetail.
type BetaAppReviewDetail struct {
	Attributes *struct {
		ContactEmail        *string `json:"contactEmail,omitempty"`
		ContactFirstName    *string `json:"contactFirstName,omitempty"`
		ContactLastName     *string `json:"contactLastName,omitempty"`
		ContactPhone        *string `json:"contactPhone,omitempty"`
		DemoAccountName     *string `json:"demoAccountName,omitempty"`
		DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
		DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
		Notes               *string `json:"notes,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string              `json:"id"`
	Links         types.ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *types.RelationshipsData  `json:"data,omitempty"`
			Links *types.RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppReviewDetailUpdateRequest defines model for BetaAppReviewDetailUpdateRequest.
type BetaAppReviewDetailUpdateRequest struct {
	Attributes *BetaAppReviewDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

type BetaAppReviewDetailUpdateRequestAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// BetaAppReviewDetailResponse defines model for BetaAppReviewDetailResponse.
type BetaAppReviewDetailResponse struct {
	Data     BetaAppReviewDetail `json:"data"`
	Included *[]apps.App         `json:"included,omitempty"`
	Links    types.DocumentLinks `json:"links"`
}

// BetaAppReviewDetailsResponse defines model for BetaAppReviewDetailsResponse.
type BetaAppReviewDetailsResponse struct {
	Data     []BetaAppReviewDetail    `json:"data"`
	Included *[]apps.App              `json:"included,omitempty"`
	Links    types.PagedDocumentLinks `json:"links"`
	Meta     *types.PagingInformation `json:"meta,omitempty"`
}

// ListBetaAppReviewDetailsQuery defines model for ListBetaAppReviewDetails
type ListBetaAppReviewDetailsQuery struct {
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsBetaAppReviewDetails *[]string `url:"fields[betaAppReviewDetails],omitempty"`
	FilterApp                  *[]string `url:"filter[app],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// GetBetaAppReviewDetailQuery defines model for GetBetaAppReviewDetail
type GetBetaAppReviewDetailQuery struct {
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsBetaAppReviewDetails *[]string `url:"fields[betaAppReviewDetails],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
}

// GetAppForBetaAppReviewDetailQuery defines model for GetAppForBetaAppReviewDetail
type GetAppForBetaAppReviewDetailQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// GetBetaAppReviewDetailsForAppQuery defines model for GetBetaAppReviewDetailsForApp
type GetBetaAppReviewDetailsForAppQuery struct {
	FieldsBetaAppReviewDetails *[]string `url:"fields[betaAppReviewDetails],omitempty"`
}

// ListBetaAppReviewDetails finds and lists beta app review details for all apps.
func (s *Service) ListBetaAppReviewDetails(params *ListBetaAppReviewDetailsQuery) (*BetaAppReviewDetailsResponse, *services.Response, error) {
	res := new(BetaAppReviewDetailsResponse)
	resp, err := s.GetWithQuery("betaAppReviewDetails", params, res)
	return res, resp, err
}

// GetBetaAppReviewDetail gets beta app review details for a specific app.
func (s *Service) GetBetaAppReviewDetail(id string, params *GetBetaAppReviewDetailQuery) (*BetaAppReviewDetailResponse, *services.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBetaAppReviewDetail gets the app information for a specific beta app review details resource.
func (s *Service) GetAppForBetaAppReviewDetail(id string, params *GetAppForBetaAppReviewDetailQuery) (*apps.AppResponse, *services.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s/app", id)
	res := new(apps.AppResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewDetailsForApp gets the beta app review details for a specific app.
func (s *Service) GetBetaAppReviewDetailsForApp(id string, params *GetBetaAppReviewDetailsForAppQuery) (*BetaAppReviewDetailResponse, *services.Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppReviewDetail", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBetaAppReviewDetail updates the details for a specific app's beta app review.
func (s *Service) UpdateBetaAppReviewDetail(id string, body *BetaAppReviewDetailUpdateRequest) (*BetaAppReviewDetailResponse, *services.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
