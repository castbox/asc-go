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

// GameCenterDetail defines model for GameCenterDetail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail
type GameCenterDetail struct {
	Attributes    *GameCenterDetailAttributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Links         ResourceLinks                  `json:"links"`
	Relationships *GameCenterDetailRelationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}

// GameCenterDetailAttributes defines model for GameCenterDetail.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail/attributes
type GameCenterDetailAttributes struct {
	ArcadeEnabled           *bool `json:"arcadeEnabled,omitempty"`
	ChallengeEnabled        *bool `json:"challengeEnabled,omitempty"`
	DefaultGroupLeaderboard *bool `json:"defaultGroupLeaderboard,omitempty"`
	DefaultLeaderboard      *bool `json:"defaultLeaderboard,omitempty"`
	GameCenterEnabled       *bool `json:"gameCenterEnabled,omitempty"`
}

// GameCenterDetailRelationships defines model for GameCenterDetail.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetail/relationships
type GameCenterDetailRelationships struct {
	AchievementReleases       *PagedRelationship `json:"achievementReleases,omitempty"`
	App                       *Relationship      `json:"app,omitempty"`
	DefaultGroupLeaderboard   *Relationship      `json:"defaultGroupLeaderboard,omitempty"`
	DefaultLeaderboard        *Relationship      `json:"defaultLeaderboard,omitempty"`
	GameCenterAchievements    *PagedRelationship `json:"gameCenterAchievements,omitempty"`
	GameCenterAppVersions     *PagedRelationship `json:"gameCenterAppVersions,omitempty"`
	GameCenterGroup           *Relationship      `json:"gameCenterGroup,omitempty"`
	GameCenterLeaderboardSets *PagedRelationship `json:"gameCenterLeaderboardSets,omitempty"`
	GameCenterLeaderboards    *PagedRelationship `json:"gameCenterLeaderboards,omitempty"`
	LeaderboardReleases       *PagedRelationship `json:"leaderboardReleases,omitempty"`
	LeaderboardSetReleases    *PagedRelationship `json:"leaderboardSetReleases,omitempty"`
}

// gameCenterDetailCreateRequest defines model for GameCenterDetailCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailcreaterequest/data
type gameCenterDetailCreateRequest struct {
	Relationships gameCenterDetailCreateRequestRelationships `json:"relationships"`
	Type          string                                     `json:"type"`
}

// gameCenterDetailCreateRequestRelationships are relationships for GameCenterDetailCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailcreaterequest/data/relationships
type gameCenterDetailCreateRequestRelationships struct {
	App relationshipDeclaration `json:"app"`
}

// gameCenterDetailUpdateRequest defines model for GameCenterDetailUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailupdaterequest/data
type gameCenterDetailUpdateRequest struct {
	Attributes    *GameCenterDetailUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                                      `json:"id"`
	Relationships *gameCenterDetailUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                      `json:"type"`
}

// GameCenterDetailUpdateRequestAttributes are attributes for GameCenterDetailUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailupdaterequest/data/attributes
type GameCenterDetailUpdateRequestAttributes struct {
	ChallengeEnabled *bool `json:"challengeEnabled,omitempty"`
}

// gameCenterDetailUpdateRequestRelationships are relationships for GameCenterDetailUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailupdaterequest/data/relationships
type gameCenterDetailUpdateRequestRelationships struct {
	DefaultGroupLeaderboard *relationshipDeclaration `json:"defaultGroupLeaderboard,omitempty"`
	DefaultLeaderboard      *relationshipDeclaration `json:"defaultLeaderboard,omitempty"`
	GameCenterGroup         *relationshipDeclaration `json:"gameCenterGroup,omitempty"`
}

// GameCenterDetailResponse defines model for GameCenterDetailResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterdetailresponse
type GameCenterDetailResponse struct {
	Data     GameCenterDetail                   `json:"data"`
	Included []GameCenterDetailResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                      `json:"links"`
}

// GameCenterDetailResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterDetailResponse.
type GameCenterDetailResponseIncluded included

// GetGameCenterDetailQuery defines model for GetGameCenterDetail
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_game_center_details
type GetGameCenterDetailQuery struct {
	FieldsGameCenterAchievementReleases    []string `url:"fields[gameCenterAchievementReleases],omitempty"`
	FieldsGameCenterAchievements           []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterAppVersions            []string `url:"fields[gameCenterAppVersions],omitempty"`
	FieldsGameCenterDetails                []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups                 []string `url:"fields[gameCenterGroups],omitempty"`
	FieldsGameCenterLeaderboardReleases    []string `url:"fields[gameCenterLeaderboardReleases],omitempty"`
	FieldsGameCenterLeaderboardSetReleases []string `url:"fields[gameCenterLeaderboardSetReleases],omitempty"`
	FieldsGameCenterLeaderboardSets        []string `url:"fields[gameCenterLeaderboardSets],omitempty"`
	FieldsGameCenterLeaderboards           []string `url:"fields[gameCenterLeaderboards],omitempty"`
	Include                                []string `url:"include,omitempty"`
	LimitAchievementReleases               int      `url:"limit[achievementReleases],omitempty"`
	LimitGameCenterAchievements            int      `url:"limit[gameCenterAchievements],omitempty"`
	LimitGameCenterAppVersions             int      `url:"limit[gameCenterAppVersions],omitempty"`
	LimitGameCenterLeaderboardSets         int      `url:"limit[gameCenterLeaderboardSets],omitempty"`
	LimitGameCenterLeaderboards            int      `url:"limit[gameCenterLeaderboards],omitempty"`
	LimitLeaderboardReleases               int      `url:"limit[leaderboardReleases],omitempty"`
	LimitLeaderboardSetReleases            int      `url:"limit[leaderboardSetReleases],omitempty"`
}

// GetGameCenterDetailForAppQuery defines model for GetGameCenterDetailForApp
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_game_center_detail_for_an_app
type GetGameCenterDetailForAppQuery struct {
	FieldsGameCenterAchievementReleases    []string `url:"fields[gameCenterAchievementReleases],omitempty"`
	FieldsGameCenterAchievements           []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterAppVersions            []string `url:"fields[gameCenterAppVersions],omitempty"`
	FieldsGameCenterDetails                []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups                 []string `url:"fields[gameCenterGroups],omitempty"`
	FieldsGameCenterLeaderboardReleases    []string `url:"fields[gameCenterLeaderboardReleases],omitempty"`
	FieldsGameCenterLeaderboardSetReleases []string `url:"fields[gameCenterLeaderboardSetReleases],omitempty"`
	FieldsGameCenterLeaderboardSets        []string `url:"fields[gameCenterLeaderboardSets],omitempty"`
	FieldsGameCenterLeaderboards           []string `url:"fields[gameCenterLeaderboards],omitempty"`
	Include                                []string `url:"include,omitempty"`
	LimitAchievementReleases               int      `url:"limit[achievementReleases],omitempty"`
	LimitGameCenterAchievements            int      `url:"limit[gameCenterAchievements],omitempty"`
	LimitGameCenterAppVersions             int      `url:"limit[gameCenterAppVersions],omitempty"`
	LimitGameCenterLeaderboardSets         int      `url:"limit[gameCenterLeaderboardSets],omitempty"`
	LimitGameCenterLeaderboards            int      `url:"limit[gameCenterLeaderboards],omitempty"`
	LimitLeaderboardReleases               int      `url:"limit[leaderboardReleases],omitempty"`
	LimitLeaderboardSetReleases            int      `url:"limit[leaderboardSetReleases],omitempty"`
}

// CreateGameCenterDetail creates a Game Center detail for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enable_game_center_for_an_app
func (s *GameCenterService) CreateGameCenterDetail(ctx context.Context, appID string) (*GameCenterDetailResponse, *Response, error) {
	req := gameCenterDetailCreateRequest{
		Relationships: gameCenterDetailCreateRequestRelationships{
			App: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appID,
					Type: "apps",
				},
			},
		},
		Type: "gameCenterDetails",
	}
	res := new(GameCenterDetailResponse)
	resp, err := s.client.post(ctx, "gameCenterDetails", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterDetail gets information about a specific Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_game_center_details
func (s *GameCenterService) GetGameCenterDetail(ctx context.Context, id string, params *GetGameCenterDetailQuery) (*GameCenterDetailResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterDetails/%s", id)
	res := new(GameCenterDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetGameCenterDetailForApp gets the Game Center detail for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_game_center_detail_for_an_app
func (s *GameCenterService) GetGameCenterDetailForApp(ctx context.Context, appID string, params *GetGameCenterDetailForAppQuery) (*GameCenterDetailResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/gameCenterDetail", appID)
	res := new(GameCenterDetailResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterDetail updates a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_game_center_detail
func (s *GameCenterService) UpdateGameCenterDetail(ctx context.Context, id string, attributes *GameCenterDetailUpdateRequestAttributes, defaultLeaderboardID *string, defaultGroupLeaderboardID *string, gameCenterGroupID *string) (*GameCenterDetailResponse, *Response, error) {
	req := gameCenterDetailUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterDetails",
	}

	if defaultLeaderboardID != nil || defaultGroupLeaderboardID != nil || gameCenterGroupID != nil {
		req.Relationships = &gameCenterDetailUpdateRequestRelationships{
			DefaultLeaderboard:      newRelationshipDeclaration(defaultLeaderboardID, "gameCenterLeaderboards"),
			DefaultGroupLeaderboard: newRelationshipDeclaration(defaultGroupLeaderboardID, "gameCenterLeaderboards"),
			GameCenterGroup:         newRelationshipDeclaration(gameCenterGroupID, "gameCenterGroups"),
		}
	}

	url := fmt.Sprintf("gameCenterDetails/%s", id)
	res := new(GameCenterDetailResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}
