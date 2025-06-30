package asc

import (
	"context"
	"fmt"
)

// AppCustomerReviewsService defines model for AppCustomerReviewsService.
type AppCustomerReviewsService service

// CustomerReview defines model for CustomerReview.
type CustomerReview struct {
	ID            string                    `json:"id"`
	Type          string                    `json:"type"`
	Attributes    *CustomerReviewAttributes `json:"attributes,omitempty"`
	Links         *ResourceLinks            `json:"links,omitempty"`
	Relationships *Relationships            `json:"relationships,omitempty"`
}

// CustomerReviewAttributes defines attributes for CustomerReview.
type CustomerReviewAttributes struct {
	Rating           int    `json:"rating"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	CreatedDate      string `json:"createdDate"`
	ReviewerNickname string `json:"reviewerNickname"`
	Territory        string `json:"territory"`
}

type Relationships struct {
	Response *Relationship `json:"response,omitempty"`
}

// CustomerReviewsResponse defines model for CustomerReviewsResponse.
type CustomerReviewsResponse struct {
	Data     []CustomerReview           `json:"data"`
	Links    PagedDocumentLinks         `json:"links"`
	Meta     *PagingInformation         `json:"meta,omitempty"`
	Included []CustomerReviewResponseV1 `json:"included,omitempty"`
}

// CustomerReviewResponse defines model for CustomerReviewResponse.
type CustomerReviewResponse struct {
	Data  CustomerReview `json:"data"`
	Links DocumentLinks  `json:"links"`
}

type CustomerReviewResponseV1 struct {
	Id            string                                 `json:"id"`
	Links         ResourceLinks                          `json:"links"`
	Type          string                                 `json:"type"`
	Relationships *CustomerReviewResponseV1Relationships `json:"relationships"`
	Attributes    *CustomerReviewResponseAttributes      `json:"attributes,omitempty"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/customerreviewresponsev1/attributes
type CustomerReviewResponseAttributes struct {
	ResponseBody     string `json:"responseBody"`
	LastModifiedDate string `json:"lastModifiedDate"`
	State            string `json:"state"`
}

type CustomerReviewResponseV1Relationships struct {
	Review *Relationship `json:"review"`
}

// GetCustomerReviewsQuery defines query parameters for getting customer reviews.
type GetCustomerReviewsQuery struct {
	FieldsCustomerReviews         string `url:"fields[customerReviews],omitempty"`         // 要返回的客户评论字段
	FieldsCustomerReviewResponses string `url:"fields[customerReviewResponses],omitempty"` // 要返回的客户评论回复字段
	FilterRating                  string `url:"filter[rating],omitempty"`                  // 评级过滤（如：1, 2, 5）
	FilterTerritory               string `url:"filter[territory],omitempty"`               // 国家或地区过滤
	Include                       string `url:"include,omitempty"`                         // 包含的相关数据，如评论回复
	Limit                         int    `url:"limit,omitempty"`                           // 返回的记录数量，最大值为200
	Sort                          string `url:"sort,omitempty"`                            // 排序方式，如：createdDate, -createdDate, rating, -rating
	ExistsPublishedResponse       bool   `url:"exists[publishedResponse],omitempty"`       // 过滤是否有已发布回复的评论
	Cursor                        string `url:"cursor,omitempty"`
}

// GetCustomerReviewQuery defines query parameters for getting a specific customer review.
type GetCustomerReviewQuery struct {
	FieldsCustomerReviews         string `url:"fields[customerReviews],omitempty"`         // 要返回的客户评论字段
	FieldsCustomerReviewResponses string `url:"fields[customerReviewResponses],omitempty"` // 要返回的客户评论回复字段
	Include                       string `url:"include,omitempty"`                         // 包含的相关数据，如评论回复
}

type CustomerReviewResponseV1CreateRequestDataAttributes struct {
	ResponseBody string `json:"responseBody"`
}

type CustomerReviewResponseV1CreateRequestRelationships struct {
	Review CustomerReviewResponseV1CreateRequestRelationshipsReview `json:"review"`
}

type CustomerReviewResponseV1CreateRequestRelationshipsReview struct {
	Data struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

//	type CustomerReviewResponseV1CreateRequestData struct {
//		Attributes    CustomerReviewResponseV1CreateRequestDataAttributes `json:"attributes"`
//		Relationships CustomerReviewResponseV1CreateRequestRelationships  `json:"relationships"`
//		Type          string                                              `json:"type"`
//	}
type CustomerReviewResponseV1CreateRequest struct {
	Attributes    CustomerReviewResponseV1CreateRequestDataAttributes `json:"attributes"`
	Relationships CustomerReviewResponseV1CreateRequestRelationships  `json:"relationships"`
	Type          string                                              `json:"type"`
}

type CustomerReviewResponseV1Response struct {
	Data     CustomerReviewResponseV1
	Links    DocumentLinks
	Included []CustomerReview
}

// GetCustomerReviewsForApp gets all customer reviews for a specific app.
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-apps-_id_-customerreviews
// GET https://api.appstoreconnect.apple.com/v1/apps/{id}/customerReviews
func (s *AppCustomerReviewsService) GetCustomerReviewsForApp(ctx context.Context, appId string, params *GetCustomerReviewsQuery) (*CustomerReviewsResponse, *Response, error) {
	url := fmt.Sprintf("/v1/apps/%s/customerReviews", appId)
	res := new(CustomerReviewsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}

// CreateOrUpdateCustomerReviewResponse Create a response or replace an existing response you wrote to a customer review.
// https://developer.apple.com/documentation/appstoreconnectapi/post-v1-customerreviewresponses
// POST GET POST https://api.appstoreconnect.apple.com/v1/customerReviewResponses
func (s *AppCustomerReviewsService) CreateOrUpdateCustomerReviewResponse(ctx context.Context, appId string, params *CustomerReviewResponseV1CreateRequest) (*CustomerReviewResponseV1Response, *Response, error) {
	url := "/v1/customerReviewResponses"
	res := new(CustomerReviewResponseV1Response)
	resp, err := s.client.post(ctx, url, newRequestBody(params), res)
	return res, resp, err
}

// GetCustomerReview gets information about a specific customer review, including the review content.
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-customerreviews-_id_
// GET https://api.appstoreconnect.apple.com/v1/customerReviews/{id}
func (s *AppCustomerReviewsService) GetCustomerReview(ctx context.Context, id string, params *GetCustomerReviewQuery) (*CustomerReviewResponse, *Response, error) {
	url := fmt.Sprintf("/v1/customerReviews/%s", id)
	res := new(CustomerReviewResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
