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

// GameCenterAchievementRelease defines model for GameCenterAchievementRelease.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementrelease
type GameCenterAchievementRelease struct {
	Attributes    *GameCenterAchievementReleaseAttributes    `json:"attributes,omitempty"`
	ID            string                                     `json:"id"`
	Links         ResourceLinks                              `json:"links"`
	Relationships *GameCenterAchievementReleaseRelationships `json:"relationships,omitempty"`
	Type          string                                     `json:"type"`
}

// GameCenterAchievementReleaseAttributes defines model for GameCenterAchievementRelease.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementrelease/attributes
type GameCenterAchievementReleaseAttributes struct {
	Live *bool `json:"live,omitempty"`
}

// GameCenterAchievementReleaseRelationships defines model for GameCenterAchievementRelease.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementrelease/relationships
type GameCenterAchievementReleaseRelationships struct {
	GameCenterAchievement *Relationship `json:"gameCenterAchievement,omitempty"`
	GameCenterDetail      *Relationship `json:"gameCenterDetail,omitempty"`
}

// gameCenterAchievementReleaseCreateRequest defines model for GameCenterAchievementReleaseCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementreleasecreaterequest/data
type gameCenterAchievementReleaseCreateRequest struct {
	Relationships gameCenterAchievementReleaseCreateRequestRelationships `json:"relationships"`
	Type          string                                                 `json:"type"`
}

// gameCenterAchievementReleaseCreateRequestRelationships are relationships for GameCenterAchievementReleaseCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementreleasecreaterequest/data/relationships
type gameCenterAchievementReleaseCreateRequestRelationships struct {
	GameCenterAchievement relationshipDeclaration `json:"gameCenterAchievement"`
	GameCenterDetail      relationshipDeclaration `json:"gameCenterDetail"`
}

// GameCenterAchievementReleaseResponse defines model for GameCenterAchievementReleaseResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementreleaseresponse
type GameCenterAchievementReleaseResponse struct {
	Data     GameCenterAchievementRelease                   `json:"data"`
	Included []GameCenterAchievementReleaseResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                  `json:"links"`
}

// GameCenterAchievementReleasesResponse defines model for GameCenterAchievementReleasesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementreleasesresponse
type GameCenterAchievementReleasesResponse struct {
	Data     []GameCenterAchievementRelease                 `json:"data"`
	Included []GameCenterAchievementReleaseResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                             `json:"links"`
	Meta     *PagingInformation                             `json:"meta,omitempty"`
}

// GameCenterAchievementReleaseResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterAchievementReleaseResponse or GameCenterAchievementReleasesResponse.
type GameCenterAchievementReleaseResponseIncluded included

// ListGameCenterAchievementReleasesQuery defines model for ListGameCenterAchievementReleases
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_achievement_releases
type ListGameCenterAchievementReleasesQuery struct {
	FieldsGameCenterAchievementReleases []string `url:"fields[gameCenterAchievementReleases],omitempty"`
	FieldsGameCenterAchievements        []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterDetails             []string `url:"fields[gameCenterDetails],omitempty"`
	FilterGameCenterAchievement         []string `url:"filter[gameCenterAchievement],omitempty"`
	FilterLive                          []string `url:"filter[live],omitempty"`
	Include                             []string `url:"include,omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	Cursor                              string   `url:"cursor,omitempty"`
}

// GetGameCenterAchievementReleaseQuery defines model for GetGameCenterAchievementRelease
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_release_information
type GetGameCenterAchievementReleaseQuery struct {
	FieldsGameCenterAchievementReleases []string `url:"fields[gameCenterAchievementReleases],omitempty"`
	FieldsGameCenterAchievements        []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterDetails             []string `url:"fields[gameCenterDetails],omitempty"`
	Include                             []string `url:"include,omitempty"`
}

// CreateGameCenterAchievementRelease creates a new release for an achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_release
func (s *GameCenterService) CreateGameCenterAchievementRelease(ctx context.Context, gameCenterAchievementID string, gameCenterDetailID string) (*GameCenterAchievementReleaseResponse, *Response, error) {
	req := gameCenterAchievementReleaseCreateRequest{
		Relationships: gameCenterAchievementReleaseCreateRequestRelationships{
			GameCenterAchievement: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterAchievementID,
					Type: "gameCenterAchievements",
				},
			},
			GameCenterDetail: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterDetailID,
					Type: "gameCenterDetails",
				},
			},
		},
		Type: "gameCenterAchievementReleases",
	}
	res := new(GameCenterAchievementReleaseResponse)
	resp, err := s.client.post(ctx, "gameCenterAchievementReleases", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterAchievementRelease gets information about a specific achievement release.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_release_information
func (s *GameCenterService) GetGameCenterAchievementRelease(ctx context.Context, id string, params *GetGameCenterAchievementReleaseQuery) (*GameCenterAchievementReleaseResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterAchievementReleases/%s", id)
	res := new(GameCenterAchievementReleaseResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// DeleteGameCenterAchievementRelease deletes an achievement release.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_release
func (s *GameCenterService) DeleteGameCenterAchievementRelease(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("gameCenterAchievementReleases/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListGameCenterAchievementReleasesForDetail lists all achievement releases for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_achievement_releases
func (s *GameCenterService) ListGameCenterAchievementReleasesForDetail(ctx context.Context, gameCenterDetailID string, params *ListGameCenterAchievementReleasesQuery) (*GameCenterAchievementReleasesResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterDetails/%s/gameCenterAchievementReleases", gameCenterDetailID)
	res := new(GameCenterAchievementReleasesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListGameCenterAchievementReleasesForAchievement lists all releases for an achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_releases_for_an_achievement
func (s *GameCenterService) ListGameCenterAchievementReleasesForAchievement(ctx context.Context, gameCenterAchievementID string, params *ListGameCenterAchievementReleasesQuery) (*GameCenterAchievementReleasesResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterAchievements/%s/releases", gameCenterAchievementID)
	res := new(GameCenterAchievementReleasesResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ReplaceGameCenterAchievementReleasesForDetail replaces all achievement releases for a Game Center detail.
// This is used to reorder achievements by providing the release IDs in the desired order.
//
// https://developer.apple.com/documentation/appstoreconnectapi/replace_all_game_center_achievement_releases
func (s *GameCenterService) ReplaceGameCenterAchievementReleasesForDetail(ctx context.Context, gameCenterDetailID string, gameCenterAchievementReleaseIDs []string) (*Response, error) {
	linkages := newPagedRelationshipDeclaration(gameCenterAchievementReleaseIDs, "gameCenterAchievementReleases")
	url := fmt.Sprintf("gameCenterDetails/%s/relationships/gameCenterAchievementReleases", gameCenterDetailID)

	return s.client.patch(ctx, url, newRequestBody(linkages.Data), nil)
}
