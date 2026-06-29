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

// GameCenterLeaderboardLocalization defines model for GameCenterLeaderboardLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2
type GameCenterLeaderboardLocalization struct {
	Attributes    *GameCenterLeaderboardLocalizationAttributes    `json:"attributes,omitempty"`
	ID            string                                          `json:"id"`
	Links         ResourceLinks                                   `json:"links"`
	Relationships *GameCenterLeaderboardLocalizationRelationships `json:"relationships,omitempty"`
	Type          string                                          `json:"type"`
}

// GameCenterLeaderboardLocalizationAttributes defines model for GameCenterLeaderboardLocalization.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2/attributes
type GameCenterLeaderboardLocalizationAttributes struct {
	Description             *string `json:"description,omitempty"`
	FormatterOverride       *string `json:"formatterOverride,omitempty"`
	FormatterSuffix         *string `json:"formatterSuffix,omitempty"`
	FormatterSuffixSingular *string `json:"formatterSuffixSingular,omitempty"`
	Locale                  *string `json:"locale,omitempty"`
	Name                    *string `json:"name,omitempty"`
}

// GameCenterLeaderboardLocalizationRelationships defines model for GameCenterLeaderboardLocalization.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2/relationships
type GameCenterLeaderboardLocalizationRelationships struct {
	Image   *Relationship `json:"image,omitempty"`
	Version *Relationship `json:"version,omitempty"`
}

// gameCenterLeaderboardLocalizationCreateRequest defines model for GameCenterLeaderboardLocalizationV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2createrequest
type gameCenterLeaderboardLocalizationCreateRequest struct {
	Attributes    GameCenterLeaderboardLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships gameCenterLeaderboardLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                                      `json:"type"`
}

// GameCenterLeaderboardLocalizationCreateRequestAttributes are attributes for GameCenterLeaderboardLocalizationV2CreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2/attributes
type GameCenterLeaderboardLocalizationCreateRequestAttributes struct {
	Description             *string `json:"description,omitempty"`
	FormatterOverride       *string `json:"formatterOverride,omitempty"`
	FormatterSuffix         *string `json:"formatterSuffix,omitempty"`
	FormatterSuffixSingular *string `json:"formatterSuffixSingular,omitempty"`
	Locale                  string  `json:"locale"`
	Name                    string  `json:"name"`
}

// gameCenterLeaderboardLocalizationCreateRequestRelationships are relationships for GameCenterLeaderboardLocalizationV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2/relationships
type gameCenterLeaderboardLocalizationCreateRequestRelationships struct {
	Version relationshipDeclaration `json:"version"`
}

// gameCenterLeaderboardLocalizationUpdateRequest defines model for GameCenterLeaderboardLocalizationV2UpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2updaterequest
type gameCenterLeaderboardLocalizationUpdateRequest struct {
	Attributes *GameCenterLeaderboardLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                                    `json:"id"`
	Type       string                                                    `json:"type"`
}

// GameCenterLeaderboardLocalizationUpdateRequestAttributes are attributes for GameCenterLeaderboardLocalizationV2UpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2/attributes
type GameCenterLeaderboardLocalizationUpdateRequestAttributes struct {
	Description             *string `json:"description,omitempty"`
	FormatterOverride       *string `json:"formatterOverride,omitempty"`
	FormatterSuffix         *string `json:"formatterSuffix,omitempty"`
	FormatterSuffixSingular *string `json:"formatterSuffixSingular,omitempty"`
	Name                    *string `json:"name,omitempty"`
}

// GameCenterLeaderboardLocalizationResponse defines model for GameCenterLeaderboardLocalizationV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationv2response
type GameCenterLeaderboardLocalizationResponse struct {
	Data     GameCenterLeaderboardLocalization                   `json:"data"`
	Included []GameCenterLeaderboardLocalizationResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                       `json:"links"`
}

// GameCenterLeaderboardLocalizationsResponse defines model for GameCenterLeaderboardLocalizationsV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardlocalizationsv2response
type GameCenterLeaderboardLocalizationsResponse struct {
	Data     []GameCenterLeaderboardLocalization                 `json:"data"`
	Included []GameCenterLeaderboardLocalizationResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                                  `json:"links"`
	Meta     *PagingInformation                                  `json:"meta,omitempty"`
}

// GameCenterLeaderboardLocalizationResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterLeaderboardLocalizationResponse or GameCenterLeaderboardLocalizationsResponse.
type GameCenterLeaderboardLocalizationResponseIncluded included

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in GameCenterLeaderboardLocalizationResponseIncluded.
func (i *GameCenterLeaderboardLocalizationResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// GameCenterLeaderboardImage returns the GameCenterLeaderboardImage stored within, if one is present.
func (i *GameCenterLeaderboardLocalizationResponseIncluded) GameCenterLeaderboardImage() *GameCenterLeaderboardImage {
	return extractIncludedGameCenterLeaderboardImage(i.inner)
}

// ListGameCenterLeaderboardLocalizationsQuery defines model for ListGameCenterLeaderboardLocalizations.
//
// https://developer.apple.com/documentation/appstoreconnectapi/game-center-leaderboard-localizations
type ListGameCenterLeaderboardLocalizationsQuery struct {
	FieldsGameCenterLeaderboardImages        []string `url:"fields[gameCenterLeaderboardImages],omitempty"`
	FieldsGameCenterLeaderboardLocalizations []string `url:"fields[gameCenterLeaderboardLocalizations],omitempty"`
	FieldsGameCenterLeaderboardVersions      []string `url:"fields[gameCenterLeaderboardVersions],omitempty"`
	FilterLocale                             []string `url:"filter[locale],omitempty"`
	Include                                  []string `url:"include,omitempty"`
	Limit                                    int      `url:"limit,omitempty"`
	Cursor                                   string   `url:"cursor,omitempty"`
}

// GetGameCenterLeaderboardLocalizationQuery defines model for GetGameCenterLeaderboardLocalization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardlocalizations-_id_
type GetGameCenterLeaderboardLocalizationQuery struct {
	FieldsGameCenterLeaderboardImages        []string `url:"fields[gameCenterLeaderboardImages],omitempty"`
	FieldsGameCenterLeaderboardLocalizations []string `url:"fields[gameCenterLeaderboardLocalizations],omitempty"`
	FieldsGameCenterLeaderboardVersions      []string `url:"fields[gameCenterLeaderboardVersions],omitempty"`
	Include                                  []string `url:"include,omitempty"`
}

// CreateGameCenterLeaderboardLocalization creates a new localization for a leaderboard version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post-v2-gamecenterleaderboardlocalizations
func (s *GameCenterService) CreateGameCenterLeaderboardLocalization(ctx context.Context, attributes GameCenterLeaderboardLocalizationCreateRequestAttributes, gameCenterLeaderboardVersionID string) (*GameCenterLeaderboardLocalizationResponse, *Response, error) {
	req := gameCenterLeaderboardLocalizationCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterLeaderboardLocalizationCreateRequestRelationships{
			Version: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterLeaderboardVersionID,
					Type: "gameCenterLeaderboardVersions",
				},
			},
		},
		Type: "gameCenterLeaderboardLocalizations",
	}
	res := new(GameCenterLeaderboardLocalizationResponse)
	resp, err := s.client.post(ctx, "../v2/gameCenterLeaderboardLocalizations", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterLeaderboardLocalization gets information about a specific leaderboard localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardlocalizations-_id_
func (s *GameCenterService) GetGameCenterLeaderboardLocalization(ctx context.Context, id string, params *GetGameCenterLeaderboardLocalizationQuery) (*GameCenterLeaderboardLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboardLocalizations/%s", id)
	res := new(GameCenterLeaderboardLocalizationResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterLeaderboardLocalization updates an existing leaderboard localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/patch-v2-gamecenterleaderboardlocalizations-_id_
func (s *GameCenterService) UpdateGameCenterLeaderboardLocalization(ctx context.Context, id string, attributes *GameCenterLeaderboardLocalizationUpdateRequestAttributes) (*GameCenterLeaderboardLocalizationResponse, *Response, error) {
	req := gameCenterLeaderboardLocalizationUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterLeaderboardLocalizations",
	}
	url := fmt.Sprintf("../v2/gameCenterLeaderboardLocalizations/%s", id)
	res := new(GameCenterLeaderboardLocalizationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterLeaderboardLocalization deletes a leaderboard localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete-v2-gamecenterleaderboardlocalizations-_id_
func (s *GameCenterService) DeleteGameCenterLeaderboardLocalization(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboardLocalizations/%s", id)

	return s.client.delete(ctx, url, nil)
}

// ListGameCenterLeaderboardLocalizationsForVersion lists all localizations for a leaderboard version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardversions-_id_-localizations
func (s *GameCenterService) ListGameCenterLeaderboardLocalizationsForVersion(ctx context.Context, gameCenterLeaderboardVersionID string, params *ListGameCenterLeaderboardLocalizationsQuery) (*GameCenterLeaderboardLocalizationsResponse, *Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboardVersions/%s/localizations", gameCenterLeaderboardVersionID)
	res := new(GameCenterLeaderboardLocalizationsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}
