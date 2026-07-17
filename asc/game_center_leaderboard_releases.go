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

// GameCenterLeaderboardRelease defines model for GameCenterLeaderboardRelease.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardrelease
type GameCenterLeaderboardRelease struct {
	Attributes    *GameCenterLeaderboardReleaseAttributes    `json:"attributes,omitempty"`
	ID            string                                     `json:"id"`
	Links         ResourceLinks                              `json:"links"`
	Relationships *GameCenterLeaderboardReleaseRelationships `json:"relationships,omitempty"`
	Type          string                                     `json:"type"`
}

// GameCenterLeaderboardReleaseAttributes defines model for GameCenterLeaderboardRelease.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardrelease
type GameCenterLeaderboardReleaseAttributes struct {
	Live *bool `json:"live,omitempty"`
}

// GameCenterLeaderboardReleaseRelationships defines model for GameCenterLeaderboardRelease.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardrelease
type GameCenterLeaderboardReleaseRelationships struct {
	GameCenterDetail      *Relationship `json:"gameCenterDetail,omitempty"`
	GameCenterLeaderboard *Relationship `json:"gameCenterLeaderboard,omitempty"`
}

// gameCenterLeaderboardReleaseCreateRequest defines model for GameCenterLeaderboardReleaseCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardreleasecreaterequest
type gameCenterLeaderboardReleaseCreateRequest struct {
	Relationships gameCenterLeaderboardReleaseCreateRequestRelationships `json:"relationships"`
	Type          string                                                 `json:"type"`
}

// gameCenterLeaderboardReleaseCreateRequestRelationships are relationships for GameCenterLeaderboardReleaseCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardreleasecreaterequest
type gameCenterLeaderboardReleaseCreateRequestRelationships struct {
	GameCenterDetail      relationshipDeclaration `json:"gameCenterDetail"`
	GameCenterLeaderboard relationshipDeclaration `json:"gameCenterLeaderboard"`
}

// GameCenterLeaderboardReleaseResponse defines model for GameCenterLeaderboardReleaseResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardreleaseresponse
type GameCenterLeaderboardReleaseResponse struct {
	Data     GameCenterLeaderboardRelease                   `json:"data"`
	Included []GameCenterLeaderboardReleaseResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                  `json:"links"`
}

// GameCenterLeaderboardReleasesResponse defines model for GameCenterLeaderboardReleasesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardreleasesresponse
type GameCenterLeaderboardReleasesResponse struct {
	Data     []GameCenterLeaderboardRelease                 `json:"data"`
	Included []GameCenterLeaderboardReleaseResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                             `json:"links"`
	Meta     *PagingInformation                             `json:"meta,omitempty"`
}

// GameCenterLeaderboardReleaseResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterLeaderboardReleaseResponse or GameCenterLeaderboardReleasesResponse.
type GameCenterLeaderboardReleaseResponseIncluded included

// ListGameCenterLeaderboardReleasesQuery defines model for ListGameCenterLeaderboardReleases.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboard-releases
type ListGameCenterLeaderboardReleasesQuery struct {
	FieldsGameCenterDetails             []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterLeaderboardReleases []string `url:"fields[gameCenterLeaderboardReleases],omitempty"`
	FieldsGameCenterLeaderboards        []string `url:"fields[gameCenterLeaderboards],omitempty"`
	FilterGameCenterLeaderboard         []string `url:"filter[gameCenterLeaderboard],omitempty"`
	FilterLive                          []string `url:"filter[live],omitempty"`
	Include                             []string `url:"include,omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	Cursor                              string   `url:"cursor,omitempty"`
}

// GetGameCenterLeaderboardReleaseQuery defines model for GetGameCenterLeaderboardRelease.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-gamecenterleaderboardreleases-_id_
type GetGameCenterLeaderboardReleaseQuery struct {
	FieldsGameCenterDetails             []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterLeaderboardReleases []string `url:"fields[gameCenterLeaderboardReleases],omitempty"`
	FieldsGameCenterLeaderboards        []string `url:"fields[gameCenterLeaderboards],omitempty"`
	Include                             []string `url:"include,omitempty"`
}

// CreateGameCenterLeaderboardRelease creates a new release for a leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post-v1-gamecenterleaderboardreleases
func (s *GameCenterService) CreateGameCenterLeaderboardRelease(ctx context.Context, gameCenterLeaderboardID string, gameCenterDetailID string) (*GameCenterLeaderboardReleaseResponse, *Response, error) {
	req := gameCenterLeaderboardReleaseCreateRequest{
		Relationships: gameCenterLeaderboardReleaseCreateRequestRelationships{
			GameCenterDetail: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterDetailID,
					Type: "gameCenterDetails",
				},
			},
			GameCenterLeaderboard: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterLeaderboardID,
					Type: "gameCenterLeaderboards",
				},
			},
		},
		Type: "gameCenterLeaderboardReleases",
	}
	res := new(GameCenterLeaderboardReleaseResponse)
	resp, err := s.client.post(ctx, "gameCenterLeaderboardReleases", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterLeaderboardRelease gets information about a specific leaderboard release.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-gamecenterleaderboardreleases-_id_
func (s *GameCenterService) GetGameCenterLeaderboardRelease(ctx context.Context, id string, params *GetGameCenterLeaderboardReleaseQuery) (*GameCenterLeaderboardReleaseResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterLeaderboardReleases/%s", id)
	res := new(GameCenterLeaderboardReleaseResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// DeleteGameCenterLeaderboardRelease deletes a leaderboard release.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete-v1-gamecenterleaderboardreleases-_id_
func (s *GameCenterService) DeleteGameCenterLeaderboardRelease(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("gameCenterLeaderboardReleases/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListGameCenterLeaderboardReleasesForDetail lists all leaderboard releases for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-gamecenterleaderboards-_id_-releases
func (s *GameCenterService) ListGameCenterLeaderboardReleasesForDetail(ctx context.Context, gameCenterDetailID string, params *ListGameCenterLeaderboardReleasesQuery) (*GameCenterLeaderboardReleasesResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterDetails/%s/leaderboardReleases", gameCenterDetailID)
	res := new(GameCenterLeaderboardReleasesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListGameCenterLeaderboardReleasesForLeaderboard lists all releases for a leaderboard.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v1-gamecenterleaderboards-_id_-releases
func (s *GameCenterService) ListGameCenterLeaderboardReleasesForLeaderboard(ctx context.Context, gameCenterLeaderboardID string, params *ListGameCenterLeaderboardReleasesQuery) (*GameCenterLeaderboardReleasesResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterLeaderboards/%s/releases", gameCenterLeaderboardID)
	res := new(GameCenterLeaderboardReleasesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ReplaceGameCenterLeaderboardReleasesForDetail replaces all leaderboard releases for a Game Center detail.
// This is used to reorder leaderboards by providing the release IDs in the desired order.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboard-releases
func (s *GameCenterService) ReplaceGameCenterLeaderboardReleasesForDetail(ctx context.Context, gameCenterDetailID string, gameCenterLeaderboardReleaseIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(gameCenterLeaderboardReleaseIDs, "gameCenterLeaderboardReleases")
	url := fmt.Sprintf("gameCenterDetails/%s/relationships/leaderboardReleases", gameCenterDetailID)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}
