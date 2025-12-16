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

// GameCenterAchievementLocalization defines model for GameCenterAchievementLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization
type GameCenterAchievementLocalization struct {
	Attributes    *GameCenterAchievementLocalizationAttributes    `json:"attributes,omitempty"`
	ID            string                                          `json:"id"`
	Links         ResourceLinks                                   `json:"links"`
	Relationships *GameCenterAchievementLocalizationRelationships `json:"relationships,omitempty"`
	Type          string                                          `json:"type"`
}

// GameCenterAchievementLocalizationAttributes defines model for GameCenterAchievementLocalization.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization/attributes
type GameCenterAchievementLocalizationAttributes struct {
	Locale                  *string `json:"locale,omitempty"`
	Name                    *string `json:"name,omitempty"`
	BeforeEarnedDescription *string `json:"beforeEarnedDescription,omitempty"`
	AfterEarnedDescription  *string `json:"afterEarnedDescription,omitempty"`
}

// GameCenterAchievementLocalizationRelationships defines model for GameCenterAchievementLocalization.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalization/relationships
type GameCenterAchievementLocalizationRelationships struct {
	GameCenterAchievement      *Relationship `json:"gameCenterAchievement,omitempty"`
	GameCenterAchievementImage *Relationship `json:"gameCenterAchievementImage,omitempty"`
}

// gameCenterAchievementLocalizationCreateRequest defines model for GameCenterAchievementLocalizationCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationcreaterequest/data
type gameCenterAchievementLocalizationCreateRequest struct {
	Attributes    GameCenterAchievementLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships gameCenterAchievementLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                                      `json:"type"`
}

// GameCenterAchievementLocalizationCreateRequestAttributes are attributes for GameCenterAchievementLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationcreaterequest/data/attributes
type GameCenterAchievementLocalizationCreateRequestAttributes struct {
	Locale                  string `json:"locale"`
	Name                    string `json:"name"`
	BeforeEarnedDescription string `json:"beforeEarnedDescription"`
	AfterEarnedDescription  string `json:"afterEarnedDescription"`
}

// gameCenterAchievementLocalizationCreateRequestRelationships are relationships for GameCenterAchievementLocalizationCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationcreaterequest/data/relationships
type gameCenterAchievementLocalizationCreateRequestRelationships struct {
	GameCenterAchievement relationshipDeclaration `json:"gameCenterAchievement"`
}

// gameCenterAchievementLocalizationUpdateRequest defines model for GameCenterAchievementLocalizationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationupdaterequest/data
type gameCenterAchievementLocalizationUpdateRequest struct {
	Attributes *GameCenterAchievementLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                                    `json:"id"`
	Type       string                                                    `json:"type"`
}

// GameCenterAchievementLocalizationUpdateRequestAttributes are attributes for GameCenterAchievementLocalizationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationupdaterequest/data/attributes
type GameCenterAchievementLocalizationUpdateRequestAttributes struct {
	Name                    *string `json:"name,omitempty"`
	BeforeEarnedDescription *string `json:"beforeEarnedDescription,omitempty"`
	AfterEarnedDescription  *string `json:"afterEarnedDescription,omitempty"`
}

// GameCenterAchievementLocalizationResponse defines model for GameCenterAchievementLocalizationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationresponse
type GameCenterAchievementLocalizationResponse struct {
	Data     GameCenterAchievementLocalization                   `json:"data"`
	Included []GameCenterAchievementLocalizationResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                       `json:"links"`
}

// GameCenterAchievementLocalizationsResponse defines model for GameCenterAchievementLocalizationsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementlocalizationsresponse
type GameCenterAchievementLocalizationsResponse struct {
	Data     []GameCenterAchievementLocalization                 `json:"data"`
	Included []GameCenterAchievementLocalizationResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                                  `json:"links"`
	Meta     *PagingInformation                                  `json:"meta,omitempty"`
}

// GameCenterAchievementLocalizationResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterAchievementLocalizationResponse or GameCenterAchievementLocalizationsResponse.
type GameCenterAchievementLocalizationResponseIncluded included

// ListGameCenterAchievementLocalizationsQuery defines model for ListGameCenterAchievementLocalizations
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_localizations_for_an_achievement
type ListGameCenterAchievementLocalizationsQuery struct {
	FieldsGameCenterAchievementImages        []string `url:"fields[gameCenterAchievementImages],omitempty"`
	FieldsGameCenterAchievementLocalizations []string `url:"fields[gameCenterAchievementLocalizations],omitempty"`
	FieldsGameCenterAchievements             []string `url:"fields[gameCenterAchievements],omitempty"`
	FilterLocale                             []string `url:"filter[locale],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	Limit                                    int      `url:"limit,omitempty"`
	Cursor                                   string   `url:"cursor,omitempty"`
}

// GetGameCenterAchievementLocalizationQuery defines model for GetGameCenterAchievementLocalization
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_localization_information
type GetGameCenterAchievementLocalizationQuery struct {
	FieldsGameCenterAchievementImages        []string `url:"fields[gameCenterAchievementImages],omitempty"`
	FieldsGameCenterAchievementLocalizations []string `url:"fields[gameCenterAchievementLocalizations],omitempty"`
	FieldsGameCenterAchievements             []string `url:"fields[gameCenterAchievements],omitempty"`
	Include                                  []string `url:"include,omitempty"`
}

// CreateGameCenterAchievementLocalization creates a new localization for an achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_localization
func (s *GameCenterService) CreateGameCenterAchievementLocalization(ctx context.Context, attributes GameCenterAchievementLocalizationCreateRequestAttributes, gameCenterAchievementID string) (*GameCenterAchievementLocalizationResponse, *Response, error) {
	req := gameCenterAchievementLocalizationCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterAchievementLocalizationCreateRequestRelationships{
			GameCenterAchievement: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterAchievementID,
					Type: "gameCenterAchievements",
				},
			},
		},
		Type: "gameCenterAchievementLocalizations",
	}
	res := new(GameCenterAchievementLocalizationResponse)
	resp, err := s.client.post(ctx, "gameCenterAchievementLocalizations", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterAchievementLocalization gets information about a specific achievement localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_localization_information
func (s *GameCenterService) GetGameCenterAchievementLocalization(ctx context.Context, id string, params *GetGameCenterAchievementLocalizationQuery) (*GameCenterAchievementLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterAchievementLocalizations/%s", id)
	res := new(GameCenterAchievementLocalizationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterAchievementLocalization updates an existing achievement localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_achievement_localization
func (s *GameCenterService) UpdateGameCenterAchievementLocalization(ctx context.Context, id string, attributes *GameCenterAchievementLocalizationUpdateRequestAttributes) (*GameCenterAchievementLocalizationResponse, *Response, error) {
	req := gameCenterAchievementLocalizationUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterAchievementLocalizations",
	}
	url := fmt.Sprintf("gameCenterAchievementLocalizations/%s", id)
	res := new(GameCenterAchievementLocalizationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterAchievementLocalization deletes an achievement localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_localization
func (s *GameCenterService) DeleteGameCenterAchievementLocalization(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("gameCenterAchievementLocalizations/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListGameCenterAchievementLocalizationsForAchievement lists all localizations for an achievement.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_localizations_for_an_achievement
func (s *GameCenterService) ListGameCenterAchievementLocalizationsForAchievement(ctx context.Context, gameCenterAchievementID string, params *ListGameCenterAchievementLocalizationsQuery) (*GameCenterAchievementLocalizationsResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterAchievements/%s/localizations", gameCenterAchievementID)
	res := new(GameCenterAchievementLocalizationsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
