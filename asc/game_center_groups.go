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

// GameCenterGroup defines model for GameCenterGroup.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroup
type GameCenterGroup struct {
	Attributes    *GameCenterGroupAttributes    `json:"attributes,omitempty"`
	ID            string                        `json:"id"`
	Links         ResourceLinks                 `json:"links"`
	Relationships *GameCenterGroupRelationships `json:"relationships,omitempty"`
	Type          string                        `json:"type"`
}

// GameCenterGroupAttributes defines model for GameCenterGroup.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroup/attributes
type GameCenterGroupAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
}

// GameCenterGroupRelationships defines model for GameCenterGroup.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroup/relationships
type GameCenterGroupRelationships struct {
	GameCenterAchievements    *PagedRelationship `json:"gameCenterAchievements,omitempty"`
	GameCenterDetails         *PagedRelationship `json:"gameCenterDetails,omitempty"`
	GameCenterLeaderboardSets *PagedRelationship `json:"gameCenterLeaderboardSets,omitempty"`
	GameCenterLeaderboards    *PagedRelationship `json:"gameCenterLeaderboards,omitempty"`
}

// gameCenterGroupCreateRequest defines model for GameCenterGroupCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroupcreaterequest/data
type gameCenterGroupCreateRequest struct {
	Attributes gameCenterGroupCreateRequestAttributes `json:"attributes"`
	Type       string                                 `json:"type"`
}

// gameCenterGroupCreateRequestAttributes are attributes for GameCenterGroupCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroupcreaterequest/data/attributes
type gameCenterGroupCreateRequestAttributes struct {
	ReferenceName string `json:"referenceName"`
}

// gameCenterGroupUpdateRequest defines model for GameCenterGroupUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroupupdaterequest/data
type gameCenterGroupUpdateRequest struct {
	Attributes *GameCenterGroupUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                  `json:"id"`
	Type       string                                  `json:"type"`
}

// GameCenterGroupUpdateRequestAttributes are attributes for GameCenterGroupUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroupupdaterequest/data/attributes
type GameCenterGroupUpdateRequestAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
}

// GameCenterGroupResponse defines model for GameCenterGroupResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroupresponse
type GameCenterGroupResponse struct {
	Data     GameCenterGroup                   `json:"data"`
	Included []GameCenterGroupResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                     `json:"links"`
}

// GameCenterGroupResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterGroupResponse.
type GameCenterGroupResponseIncluded included

// GameCenterGroupsResponse defines model for GameCenterGroupsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecentergroupsresponse
type GameCenterGroupsResponse struct {
	Data     []GameCenterGroup                 `json:"data"`
	Included []GameCenterGroupResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                `json:"links"`
	Meta     *PagingInformation                `json:"meta,omitempty"`
}

// ListGameCenterGroupsQuery defines model for ListGameCenterGroups
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_game_center_groups
type ListGameCenterGroupsQuery struct {
	FieldsGameCenterAchievements    []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterDetails         []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups          []string `url:"fields[gameCenterGroups],omitempty"`
	FieldsGameCenterLeaderboardSets []string `url:"fields[gameCenterLeaderboardSets],omitempty"`
	FieldsGameCenterLeaderboards    []string `url:"fields[gameCenterLeaderboards],omitempty"`
	FilterGameCenterDetails         []string `url:"filter[gameCenterDetails],omitempty"`
	Include                         []string `url:"include,omitempty"`
	Limit                           int      `url:"limit,omitempty"`
	LimitGameCenterAchievements     int      `url:"limit[gameCenterAchievements],omitempty"`
	LimitGameCenterDetails          int      `url:"limit[gameCenterDetails],omitempty"`
	LimitGameCenterLeaderboardSets  int      `url:"limit[gameCenterLeaderboardSets],omitempty"`
	LimitGameCenterLeaderboards     int      `url:"limit[gameCenterLeaderboards],omitempty"`
	Cursor                          string   `url:"cursor,omitempty"`
}

// GetGameCenterGroupQuery defines model for GetGameCenterGroup
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_game_center_group_information
type GetGameCenterGroupQuery struct {
	FieldsGameCenterAchievements    []string `url:"fields[gameCenterAchievements],omitempty"`
	FieldsGameCenterDetails         []string `url:"fields[gameCenterDetails],omitempty"`
	FieldsGameCenterGroups          []string `url:"fields[gameCenterGroups],omitempty"`
	FieldsGameCenterLeaderboardSets []string `url:"fields[gameCenterLeaderboardSets],omitempty"`
	FieldsGameCenterLeaderboards    []string `url:"fields[gameCenterLeaderboards],omitempty"`
	Include                         []string `url:"include,omitempty"`
	LimitGameCenterAchievements     int      `url:"limit[gameCenterAchievements],omitempty"`
	LimitGameCenterDetails          int      `url:"limit[gameCenterDetails],omitempty"`
	LimitGameCenterLeaderboardSets  int      `url:"limit[gameCenterLeaderboardSets],omitempty"`
	LimitGameCenterLeaderboards     int      `url:"limit[gameCenterLeaderboards],omitempty"`
}

// ListGameCenterGroups lists all Game Center groups.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_game_center_groups
func (s *GameCenterService) ListGameCenterGroups(ctx context.Context, params *ListGameCenterGroupsQuery) (*GameCenterGroupsResponse, *Response, error) {
	res := new(GameCenterGroupsResponse)
	resp, err := s.client.get(ctx, "gameCenterGroups", params, res)

	return res, resp, err
}

// GetGameCenterGroup gets information about a specific Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_game_center_group_information
func (s *GameCenterService) GetGameCenterGroup(ctx context.Context, id string, params *GetGameCenterGroupQuery) (*GameCenterGroupResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterGroups/%s", id)
	res := new(GameCenterGroupResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateGameCenterGroup creates a new Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_game_center_group
func (s *GameCenterService) CreateGameCenterGroup(ctx context.Context, referenceName string) (*GameCenterGroupResponse, *Response, error) {
	req := gameCenterGroupCreateRequest{
		Attributes: gameCenterGroupCreateRequestAttributes{
			ReferenceName: referenceName,
		},
		Type: "gameCenterGroups",
	}
	res := new(GameCenterGroupResponse)
	resp, err := s.client.post(ctx, "gameCenterGroups", newRequestBody(req), res)

	return res, resp, err
}

// UpdateGameCenterGroup updates a Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_game_center_group
func (s *GameCenterService) UpdateGameCenterGroup(ctx context.Context, id string, attributes *GameCenterGroupUpdateRequestAttributes) (*GameCenterGroupResponse, *Response, error) {
	req := gameCenterGroupUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterGroups",
	}
	url := fmt.Sprintf("gameCenterGroups/%s", id)
	res := new(GameCenterGroupResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterGroup deletes a Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_game_center_group
func (s *GameCenterService) DeleteGameCenterGroup(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("gameCenterGroups/%s", id)

	return s.client.delete(ctx, url, nil)
}

// GetGameCenterGroupForDetail gets the Game Center group for a Game Center detail.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_the_game_center_group_for_a_game_center_detail
func (s *GameCenterService) GetGameCenterGroupForDetail(ctx context.Context, gameCenterDetailID string, params *GetGameCenterGroupQuery) (*GameCenterGroupResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterDetails/%s/gameCenterGroup", gameCenterDetailID)
	res := new(GameCenterGroupResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// ListGameCenterAchievementsForGroup lists all achievements for a Game Center group.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_achievements_for_a_game_center_group
func (s *GameCenterService) ListGameCenterAchievementsForGroup(ctx context.Context, groupID string, params *ListGameCenterAchievementsQuery) (*GameCenterAchievementsResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterGroups/%s/gameCenterAchievements", groupID)
	res := new(GameCenterAchievementsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
