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

// GameCenterLeaderboardVersion defines model for GameCenterLeaderboardVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2
type GameCenterLeaderboardVersion struct {
	Attributes    *GameCenterLeaderboardVersionAttributes    `json:"attributes,omitempty"`
	ID            string                                     `json:"id"`
	Links         ResourceLinks                              `json:"links"`
	Relationships *GameCenterLeaderboardVersionRelationships `json:"relationships,omitempty"`
	Type          string                                     `json:"type"`
}

// GameCenterLeaderboardVersionAttributes defines model for GameCenterLeaderboardVersion.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2/attributes
type GameCenterLeaderboardVersionAttributes struct {
	State   *string `json:"state,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

// GameCenterLeaderboardVersionRelationships defines model for GameCenterLeaderboardVersion.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2/relationships
type GameCenterLeaderboardVersionRelationships struct {
	Leaderboard   *Relationship      `json:"leaderboard,omitempty"`
	Localizations *PagedRelationship `json:"localizations,omitempty"`
}

// gameCenterLeaderboardVersionCreateRequest defines model for GameCenterLeaderboardVersionV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2createrequest
type gameCenterLeaderboardVersionCreateRequest struct {
	Relationships gameCenterLeaderboardVersionCreateRequestRelationships `json:"relationships"`
	Type          string                                                 `json:"type"`
}

// gameCenterLeaderboardVersionCreateRequestRelationships are relationships for GameCenterLeaderboardVersionV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2/relationships
type gameCenterLeaderboardVersionCreateRequestRelationships struct {
	Leaderboard relationshipDeclaration `json:"leaderboard"`
}

// GameCenterLeaderboardVersionResponse defines model for GameCenterLeaderboardVersionV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionv2response
type GameCenterLeaderboardVersionResponse struct {
	Data     GameCenterLeaderboardVersion                   `json:"data"`
	Included []GameCenterLeaderboardVersionResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                  `json:"links"`
}

// GameCenterLeaderboardVersionsResponse defines model for GameCenterLeaderboardVersionsV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardversionsv2response
type GameCenterLeaderboardVersionsResponse struct {
	Data     []GameCenterLeaderboardVersion                 `json:"data"`
	Included []GameCenterLeaderboardVersionResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                             `json:"links"`
	Meta     *PagingInformation                             `json:"meta,omitempty"`
}

// GameCenterLeaderboardVersionResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterLeaderboardVersionResponse or GameCenterLeaderboardVersionsResponse.
type GameCenterLeaderboardVersionResponseIncluded included

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in GameCenterLeaderboardVersionResponseIncluded.
func (i *GameCenterLeaderboardVersionResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// GameCenterLeaderboardLocalization returns the GameCenterLeaderboardLocalization stored within, if one is present.
func (i *GameCenterLeaderboardVersionResponseIncluded) GameCenterLeaderboardLocalization() *GameCenterLeaderboardLocalization {
	return extractIncludedGameCenterLeaderboardLocalization(i.inner)
}

// ListGameCenterLeaderboardVersionsQuery defines model for ListGameCenterLeaderboardVersions.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboard-versions
type ListGameCenterLeaderboardVersionsQuery struct {
	FieldsGameCenterLeaderboardLocalizations []string `url:"fields[gameCenterLeaderboardLocalizations],omitempty"`
	FieldsGameCenterLeaderboardVersions      []string `url:"fields[gameCenterLeaderboardVersions],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	Limit                                    int      `url:"limit,omitempty"`
	LimitLocalizations                       int      `url:"limit[localizations],omitempty"`
	Cursor                                   string   `url:"cursor,omitempty"`
}

// GetGameCenterLeaderboardVersionQuery defines model for GetGameCenterLeaderboardVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardversions-_id_
type GetGameCenterLeaderboardVersionQuery struct {
	FieldsGameCenterLeaderboardLocalizations []string `url:"fields[gameCenterLeaderboardLocalizations],omitempty"`
	FieldsGameCenterLeaderboardVersions      []string `url:"fields[gameCenterLeaderboardVersions],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	LimitLocalizations                       int      `url:"limit[localizations],omitempty"`
}

// CreateGameCenterLeaderboardVersion creates a new version for a leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post-v2-gamecenterleaderboardversions
func (s *GameCenterService) CreateGameCenterLeaderboardVersion(ctx context.Context, gameCenterLeaderboardID string) (*GameCenterLeaderboardVersionResponse, *Response, error) {
	req := gameCenterLeaderboardVersionCreateRequest{
		Relationships: gameCenterLeaderboardVersionCreateRequestRelationships{
			Leaderboard: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterLeaderboardID,
					Type: "gameCenterLeaderboards",
				},
			},
		},
		Type: "gameCenterLeaderboardVersions",
	}
	res := new(GameCenterLeaderboardVersionResponse)
	resp, err := s.client.post(ctx, "../v2/gameCenterLeaderboardVersions", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterLeaderboardVersion gets information about a specific leaderboard version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardversions-_id_
func (s *GameCenterService) GetGameCenterLeaderboardVersion(ctx context.Context, id string, params *GetGameCenterLeaderboardVersionQuery) (*GameCenterLeaderboardVersionResponse, *Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboardVersions/%s", id)
	res := new(GameCenterLeaderboardVersionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListGameCenterLeaderboardVersionsForLeaderboard lists all versions for a leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboards-_id_-versions
func (s *GameCenterService) ListGameCenterLeaderboardVersionsForLeaderboard(ctx context.Context, gameCenterLeaderboardID string, params *ListGameCenterLeaderboardVersionsQuery) (*GameCenterLeaderboardVersionsResponse, *Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboards/%s/versions", gameCenterLeaderboardID)
	res := new(GameCenterLeaderboardVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
