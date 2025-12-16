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

// GameCenterService handles communication with Game Center related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/game_center
type GameCenterService service

// GameCenterAchievement defines model for GameCenterAchievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement
type GameCenterAchievement struct {
	Attributes    *GameCenterAchievementAttributes    `json:"attributes,omitempty"`
	ID            string                              `json:"id"`
	Links         ResourceLinks                       `json:"links"`
	Relationships *GameCenterAchievementRelationships `json:"relationships,omitempty"`
	Type          string                              `json:"type"`
}

// GameCenterAchievementAttributes defines model for GameCenterAchievement.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement/attributes
type GameCenterAchievementAttributes struct {
	Archived         *bool   `json:"archived,omitempty"`
	ReferenceName    *string `json:"referenceName,omitempty"`
	VendorIdentifier *string `json:"vendorIdentifier,omitempty"`
	Points           *int    `json:"points,omitempty"`
	ShowBeforeEarned *bool   `json:"showBeforeEarned,omitempty"`
	Repeatable       *bool   `json:"repeatable,omitempty"`
}

// GameCenterAchievementRelationships defines model for GameCenterAchievement.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievement/relationships
type GameCenterAchievementRelationships struct {
	GameCenterDetail *Relationship      `json:"gameCenterDetail,omitempty"`
	GameCenterGroup  *Relationship      `json:"gameCenterGroup,omitempty"`
	GroupAchievement *Relationship      `json:"groupAchievement,omitempty"`
	Localizations    *PagedRelationship `json:"localizations,omitempty"`
	Releases         *PagedRelationship `json:"releases,omitempty"`
}

// gameCenterAchievementCreateRequest defines model for GameCenterAchievementCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementcreaterequest/data
type gameCenterAchievementCreateRequest struct {
	Attributes    GameCenterAchievementCreateRequestAttributes    `json:"attributes"`
	Relationships gameCenterAchievementCreateRequestRelationships `json:"relationships"`
	Type          string                                          `json:"type"`
}

// GameCenterAchievementCreateRequestAttributes are attributes for GameCenterAchievementCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementcreaterequest/data/attributes
type GameCenterAchievementCreateRequestAttributes struct {
	ReferenceName    string `json:"referenceName"`
	VendorIdentifier string `json:"vendorIdentifier"`
	Points           int    `json:"points"`
	ShowBeforeEarned bool   `json:"showBeforeEarned"`
	Repeatable       bool   `json:"repeatable"`
}

// gameCenterAchievementCreateRequestRelationships are relationships for GameCenterAchievementCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementcreaterequest/data/relationships
type gameCenterAchievementCreateRequestRelationships struct {
	GameCenterDetail *relationshipDeclaration `json:"gameCenterDetail,omitempty"`
	GameCenterGroup  *relationshipDeclaration `json:"gameCenterGroup,omitempty"`
}

// gameCenterAchievementUpdateRequest defines model for GameCenterAchievementUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementupdaterequest/data
type gameCenterAchievementUpdateRequest struct {
	Attributes *GameCenterAchievementUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                        `json:"id"`
	Type       string                                        `json:"type"`
}

// GameCenterAchievementUpdateRequestAttributes are attributes for GameCenterAchievementUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementupdaterequest/data/attributes
type GameCenterAchievementUpdateRequestAttributes struct {
	ReferenceName    *string `json:"referenceName,omitempty"`
	Points           *int    `json:"points,omitempty"`
	ShowBeforeEarned *bool   `json:"showBeforeEarned,omitempty"`
	Repeatable       *bool   `json:"repeatable,omitempty"`
	Archived         *bool   `json:"archived,omitempty"`
}

// GameCenterAchievementResponse defines model for GameCenterAchievementResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementresponse
type GameCenterAchievementResponse struct {
	Data     GameCenterAchievement                   `json:"data"`
	Included []GameCenterAchievementResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                           `json:"links"`
}

// GameCenterAchievementsResponse defines model for GameCenterAchievementsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementsresponse
type GameCenterAchievementsResponse struct {
	Data     []GameCenterAchievement                 `json:"data"`
	Included []GameCenterAchievementResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                      `json:"links"`
	Meta     *PagingInformation                      `json:"meta,omitempty"`
}

// GameCenterAchievementResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterAchievementResponse or GameCenterAchievementsResponse.
type GameCenterAchievementResponseIncluded included

// ListGameCenterAchievementsQuery defines model for ListGameCenterAchievements
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_achievements
type ListGameCenterAchievementsQuery struct {
	FieldsGameCenterAchievementLocalizations []string `url:"fields[gameCenterAchievementLocalizations],omitempty"`
	FieldsGameCenterAchievementReleases      []string `url:"fields[gameCenterAchievementReleases],omitempty"`
	FieldsGameCenterAchievements             []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterDetails                  []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups                   []string `url:"fields[gameCenterGroups],omitempty"`
	FilterArchived                           []string `url:"filter[archived],omitempty"`
	FilterID                                 []string `url:"filter[id],omitempty"`
	FilterReferenceName                      []string `url:"filter[referenceName],omitempty"`
	FilterVendorIdentifier                   []string `url:"filter[vendorIdentifier],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	Limit                                    int      `url:"limit,omitempty"`
	LimitLocalizations                       int      `url:"limit[localizations],omitempty"`
	LimitReleases                            int      `url:"limit[releases],omitempty"`
	Sort                                     []string `url:"sort,omitempty"`
	Cursor                                   string   `url:"cursor,omitempty"`
}

// GetGameCenterAchievementQuery defines model for GetGameCenterAchievement
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_information
type GetGameCenterAchievementQuery struct {
	FieldsGameCenterAchievementLocalizations []string `url:"fields[gameCenterAchievementLocalizations],omitempty"`
	FieldsGameCenterAchievementReleases      []string `url:"fields[gameCenterAchievementReleases],omitempty"`
	FieldsGameCenterAchievements             []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterDetails                  []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups                   []string `url:"fields[gameCenterGroups],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	LimitLocalizations                       int      `url:"limit[localizations],omitempty"`
	LimitReleases                            int      `url:"limit[releases],omitempty"`
}

// CreateGameCenterAchievement creates a new achievement for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement
func (s *GameCenterService) CreateGameCenterAchievement(ctx context.Context, attributes GameCenterAchievementCreateRequestAttributes, gameCenterDetailID string) (*GameCenterAchievementResponse, *Response, error) {
	req := gameCenterAchievementCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterAchievementCreateRequestRelationships{
			GameCenterDetail: &relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterDetailID,
					Type: "gameCenterDetails",
				},
			},
		},
		Type: "gameCenterAchievements",
	}
	res := new(GameCenterAchievementResponse)
	resp, err := s.client.post(ctx, "gameCenterAchievements", newRequestBody(req), res)

	return res, resp, err
}

// CreateGameCenterAchievementForGroup creates a new achievement for a Game Center group.
// Use this method when the app belongs to a Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement
func (s *GameCenterService) CreateGameCenterAchievementForGroup(ctx context.Context, attributes GameCenterAchievementCreateRequestAttributes, gameCenterGroupID string) (*GameCenterAchievementResponse, *Response, error) {
	req := gameCenterAchievementCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterAchievementCreateRequestRelationships{
			GameCenterGroup: &relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterGroupID,
					Type: "gameCenterGroups",
				},
			},
		},
		Type: "gameCenterAchievements",
	}
	res := new(GameCenterAchievementResponse)
	resp, err := s.client.post(ctx, "gameCenterAchievements", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterAchievement gets information about a specific achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_information
func (s *GameCenterService) GetGameCenterAchievement(ctx context.Context, id string, params *GetGameCenterAchievementQuery) (*GameCenterAchievementResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterAchievements/%s", id)
	res := new(GameCenterAchievementResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterAchievement updates an existing achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_achievement
func (s *GameCenterService) UpdateGameCenterAchievement(ctx context.Context, id string, attributes *GameCenterAchievementUpdateRequestAttributes) (*GameCenterAchievementResponse, *Response, error) {
	req := gameCenterAchievementUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterAchievements",
	}
	url := fmt.Sprintf("gameCenterAchievements/%s", id)
	res := new(GameCenterAchievementResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterAchievement deletes an achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement
func (s *GameCenterService) DeleteGameCenterAchievement(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("gameCenterAchievements/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListGameCenterAchievementsForDetail lists all achievements for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_achievements
func (s *GameCenterService) ListGameCenterAchievementsForDetail(ctx context.Context, gameCenterDetailID string, params *ListGameCenterAchievementsQuery) (*GameCenterAchievementsResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterDetails/%s/gameCenterAchievements", gameCenterDetailID)
	res := new(GameCenterAchievementsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
