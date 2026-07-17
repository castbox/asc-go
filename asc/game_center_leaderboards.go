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

package asc

import (
	"context"
	"fmt"
)

// StringToStringMap defines model for StringToStringMap.
//
// https://developer.apple.com/documentation/appstoreconnectapi/stringtostringmap
type StringToStringMap map[string]string

// GameCenterLeaderboard defines model for GameCenterLeaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2
type GameCenterLeaderboard struct {
	Attributes    *GameCenterLeaderboardAttributes    `json:"attributes,omitempty"`
	ID            string                              `json:"id"`
	Links         ResourceLinks                       `json:"links"`
	Relationships *GameCenterLeaderboardRelationships `json:"relationships,omitempty"`
	Type          string                              `json:"type"`
}

// GameCenterLeaderboardAttributes defines model for GameCenterLeaderboard.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2/attributes
type GameCenterLeaderboardAttributes struct {
	ActivityProperties  *StringToStringMap `json:"activityProperties,omitempty"`
	Archived            *bool              `json:"archived,omitempty"`
	DefaultFormatter    *string            `json:"defaultFormatter,omitempty"`
	RecurrenceDuration  *string            `json:"recurrenceDuration,omitempty"`
	RecurrenceRule      *string            `json:"recurrenceRule,omitempty"`
	RecurrenceStartDate *string            `json:"recurrenceStartDate,omitempty"`
	ReferenceName       *string            `json:"referenceName,omitempty"`
	ScoreRangeEnd       *int64             `json:"scoreRangeEnd,omitempty"`
	ScoreRangeStart     *int64             `json:"scoreRangeStart,omitempty"`
	ScoreSortType       *string            `json:"scoreSortType,omitempty"`
	SubmissionType      *string            `json:"submissionType,omitempty"`
	VendorIdentifier    *string            `json:"vendorIdentifier,omitempty"`
	Visibility          *string            `json:"visibility,omitempty"`
}

// GameCenterLeaderboardRelationships defines model for GameCenterLeaderboard.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2/relationships
type GameCenterLeaderboardRelationships struct {
	Activity                  *Relationship      `json:"activity,omitempty"`
	Challenge                 *Relationship      `json:"challenge,omitempty"`
	GameCenterDetail          *Relationship      `json:"gameCenterDetail,omitempty"`
	GameCenterGroup           *Relationship      `json:"gameCenterGroup,omitempty"`
	GameCenterLeaderboardSets *PagedRelationship `json:"gameCenterLeaderboardSets,omitempty"`
	Versions                  *PagedRelationship `json:"versions,omitempty"`
}

// gameCenterLeaderboardCreateRequest defines model for GameCenterLeaderboardV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2createrequest/data-data.dictionary
type gameCenterLeaderboardCreateRequest struct {
	Attributes    GameCenterLeaderboardCreateRequestAttributes    `json:"attributes"`
	Relationships gameCenterLeaderboardCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// GameCenterLeaderboardCreateRequestAttributes are attributes for GameCenterLeaderboardV2CreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2/attributes
type GameCenterLeaderboardCreateRequestAttributes struct {
	ActivityProperties  *StringToStringMap `json:"activityProperties,omitempty"`
	DefaultFormatter    string             `json:"defaultFormatter"`
	RecurrenceDuration  *string            `json:"recurrenceDuration,omitempty"`
	RecurrenceRule      *string            `json:"recurrenceRule,omitempty"`
	RecurrenceStartDate *string            `json:"recurrenceStartDate,omitempty"`
	ReferenceName       string             `json:"referenceName"`
	ScoreRangeEnd       *int64             `json:"scoreRangeEnd,omitempty"`
	ScoreRangeStart     *int64             `json:"scoreRangeStart,omitempty"`
	ScoreSortType       string             `json:"scoreSortType"`
	SubmissionType      string             `json:"submissionType"`
	VendorIdentifier    string             `json:"vendorIdentifier"`
	Visibility          *string            `json:"visibility,omitempty"`
}

// gameCenterLeaderboardCreateRequestRelationships are relationships for GameCenterLeaderboardV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2/relationships
type gameCenterLeaderboardCreateRequestRelationships struct {
	GameCenterDetail *relationshipDeclaration     `json:"gameCenterDetail,omitempty"`
	GameCenterGroup  *relationshipDeclaration     `json:"gameCenterGroup,omitempty"`
	Versions         pagedRelationshipDeclaration `json:"versions"`
}

// GameCenterLeaderboardVersionInlineCreate defines model for GameCenterLeaderboardVersionV2InlineCreate.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2inlinecreate
type GameCenterLeaderboardVersionInlineCreate struct {
	ID            string                                                 `json:"id,omitempty"`
	Relationships *GameCenterLeaderboardVersionInlineCreateRelationships `json:"relationships,omitempty"`
	Type          string                                                 `json:"type"`
}

// GameCenterLeaderboardVersionInlineCreateRelationships defines model for GameCenterLeaderboardVersionV2InlineCreate.Relationships.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2inlinecreate/relationships
type GameCenterLeaderboardVersionInlineCreateRelationships struct {
	Leaderboard *relationshipDeclaration `json:"leaderboard,omitempty"`
}

// gameCenterLeaderboardUpdateRequest defines model for GameCenterLeaderboardV2UpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2updaterequest
type gameCenterLeaderboardUpdateRequest struct {
	Attributes *GameCenterLeaderboardUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                        `json:"id"`
	Type       string                                        `json:"type"`
}

// GameCenterLeaderboardUpdateRequestAttributes are attributes for GameCenterLeaderboardV2UpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2/attributes
type GameCenterLeaderboardUpdateRequestAttributes struct {
	ActivityProperties  *StringToStringMap `json:"activityProperties,omitempty"`
	Archived            *bool              `json:"archived,omitempty"`
	DefaultFormatter    *string            `json:"defaultFormatter,omitempty"`
	RecurrenceDuration  *string            `json:"recurrenceDuration,omitempty"`
	RecurrenceRule      *string            `json:"recurrenceRule,omitempty"`
	RecurrenceStartDate *string            `json:"recurrenceStartDate,omitempty"`
	ReferenceName       *string            `json:"referenceName,omitempty"`
	ScoreRangeEnd       *int64             `json:"scoreRangeEnd,omitempty"`
	ScoreRangeStart     *int64             `json:"scoreRangeStart,omitempty"`
	ScoreSortType       *string            `json:"scoreSortType,omitempty"`
	SubmissionType      *string            `json:"submissionType,omitempty"`
	Visibility          *string            `json:"visibility,omitempty"`
}

// GameCenterLeaderboardResponse defines model for GameCenterLeaderboardV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardv2response
type GameCenterLeaderboardResponse struct {
	Data     GameCenterLeaderboard                   `json:"data"`
	Included []GameCenterLeaderboardResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                           `json:"links"`
}

// GameCenterLeaderboardsResponse defines model for GameCenterLeaderboardsV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardsv2response
type GameCenterLeaderboardsResponse struct {
	Data     []GameCenterLeaderboard                 `json:"data"`
	Included []GameCenterLeaderboardResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                      `json:"links"`
	Meta     *PagingInformation                      `json:"meta,omitempty"`
}

// GameCenterLeaderboardResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterLeaderboardResponse or GameCenterLeaderboardsResponse.
type GameCenterLeaderboardResponseIncluded included

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in GameCenterLeaderboardResponseIncluded.
func (i *GameCenterLeaderboardResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// GameCenterLeaderboardVersion returns the GameCenterLeaderboardVersion stored within, if one is present.
func (i *GameCenterLeaderboardResponseIncluded) GameCenterLeaderboardVersion() *GameCenterLeaderboardVersion {
	return extractIncludedGameCenterLeaderboardVersion(i.inner)
}

// ListGameCenterLeaderboardsQuery defines model for ListGameCenterLeaderboards.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboards
type ListGameCenterLeaderboardsQuery struct {
	FieldsGameCenterDetails      []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups       []string `url:"fields[gameCenterGroups],omitempty"`
	FieldsGameCenterLeaderboards []string `url:"fields[gameCenterLeaderboards],omitempty"`
	FilterArchived               []string `url:"filter[archived],omitempty"`
	FilterID                     []string `url:"filter[id],omitempty"`
	FilterReferenceName          []string `url:"filter[referenceName],omitempty"`
	FilterVendorIdentifier       []string `url:"filter[vendorIdentifier],omitempty"`
	Include                      []string `url:"include,omitempty"`
	Limit                        int      `url:"limit,omitempty"`
	LimitVersions                int      `url:"limit[versions],omitempty"`
	Sort                         []string `url:"sort,omitempty"`
	Cursor                       string   `url:"cursor,omitempty"`
}

// GetGameCenterLeaderboardQuery defines model for GetGameCenterLeaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboards-_id_
type GetGameCenterLeaderboardQuery struct {
	FieldsGameCenterDetails      []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups       []string `url:"fields[gameCenterGroups],omitempty"`
	FieldsGameCenterLeaderboards []string `url:"fields[gameCenterLeaderboards],omitempty"`
	Include                      []string `url:"include,omitempty"`
	LimitVersions                int      `url:"limit[versions],omitempty"`
}

// CreateGameCenterLeaderboard creates a new leaderboard for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post-v2-gamecenterleaderboards
func (s *GameCenterService) CreateGameCenterLeaderboard(ctx context.Context, attributes GameCenterLeaderboardCreateRequestAttributes, gameCenterDetailID string) (*GameCenterLeaderboardResponse, *Response, error) {
	const inlineVersionID = "${new-gameCenterLeaderboardVersion-id}"

	req := gameCenterLeaderboardCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterLeaderboardCreateRequestRelationships{
			GameCenterDetail: &relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterDetailID,
					Type: "gameCenterDetails",
				},
			},
			Versions: newPagedRelationshipDeclaration([]string{inlineVersionID}, "gameCenterLeaderboardVersions"),
		},
		Type: "gameCenterLeaderboards",
	}
	included := []GameCenterLeaderboardVersionInlineCreate{
		{
			ID:   inlineVersionID,
			Type: "gameCenterLeaderboardVersions",
		},
	}
	res := new(GameCenterLeaderboardResponse)
	resp, err := s.client.post(ctx, "../v2/gameCenterLeaderboards", newRequestBodyWithIncluded(req, included), res)

	return res, resp, err
}

// CreateGameCenterLeaderboardForGroup creates a new leaderboard for a Game Center group.
// Use this method when the app belongs to a Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post-v2-gamecenterleaderboards
func (s *GameCenterService) CreateGameCenterLeaderboardForGroup(ctx context.Context, attributes GameCenterLeaderboardCreateRequestAttributes, gameCenterGroupID string) (*GameCenterLeaderboardResponse, *Response, error) {
	const inlineVersionID = "${new-gameCenterLeaderboardVersion-id}"

	req := gameCenterLeaderboardCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterLeaderboardCreateRequestRelationships{
			GameCenterGroup: &relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterGroupID,
					Type: "gameCenterGroups",
				},
			},
			Versions: newPagedRelationshipDeclaration([]string{inlineVersionID}, "gameCenterLeaderboardVersions"),
		},
		Type: "gameCenterLeaderboards",
	}
	included := []GameCenterLeaderboardVersionInlineCreate{
		{
			ID:   inlineVersionID,
			Type: "gameCenterLeaderboardVersions",
		},
	}
	res := new(GameCenterLeaderboardResponse)
	resp, err := s.client.post(ctx, "../v2/gameCenterLeaderboards", newRequestBodyWithIncluded(req, included), res)

	return res, resp, err
}

// GetGameCenterLeaderboard gets information about a specific leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboards-_id_
func (s *GameCenterService) GetGameCenterLeaderboard(ctx context.Context, id string, params *GetGameCenterLeaderboardQuery) (*GameCenterLeaderboardResponse, *Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboards/%s", id)
	res := new(GameCenterLeaderboardResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterLeaderboard updates an existing leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/patch-v2-gamecenterleaderboards-_id_
func (s *GameCenterService) UpdateGameCenterLeaderboard(ctx context.Context, id string, attributes *GameCenterLeaderboardUpdateRequestAttributes) (*GameCenterLeaderboardResponse, *Response, error) {
	req := gameCenterLeaderboardUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterLeaderboards",
	}
	url := fmt.Sprintf("../v2/gameCenterLeaderboards/%s", id)
	res := new(GameCenterLeaderboardResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterLeaderboard deletes a leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete-v2-gamecenterleaderboards-_id_
func (s *GameCenterService) DeleteGameCenterLeaderboard(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboards/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListGameCenterLeaderboardsForDetail lists all leaderboards for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboards
func (s *GameCenterService) ListGameCenterLeaderboardsForDetail(ctx context.Context, gameCenterDetailID string, params *ListGameCenterLeaderboardsQuery) (*GameCenterLeaderboardsResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterDetails/%s/gameCenterLeaderboardsV2", gameCenterDetailID)
	res := new(GameCenterLeaderboardsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListGameCenterLeaderboardsForGroup lists all leaderboards for a Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboards
func (s *GameCenterService) ListGameCenterLeaderboardsForGroup(ctx context.Context, gameCenterGroupID string, params *ListGameCenterLeaderboardsQuery) (*GameCenterLeaderboardsResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterGroups/%s/gameCenterLeaderboardsV2", gameCenterGroupID)
	res := new(GameCenterLeaderboardsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ReplaceGameCenterLeaderboardsForDetail replaces all leaderboards for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/patch-v1-gamecenterdetails-_id_-relationships-gamecenterleaderboardsv2
func (s *GameCenterService) ReplaceGameCenterLeaderboardsForDetail(ctx context.Context, gameCenterDetailID string, gameCenterLeaderboardIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(gameCenterLeaderboardIDs, "gameCenterLeaderboards")
	url := fmt.Sprintf("gameCenterDetails/%s/relationships/gameCenterLeaderboardsV2", gameCenterDetailID)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}
