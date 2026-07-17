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

// GameCenterLeaderboardImage defines model for GameCenterLeaderboardImage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2
type GameCenterLeaderboardImage struct {
	Attributes    *GameCenterLeaderboardImageAttributes    `json:"attributes,omitempty"`
	ID            string                                   `json:"id"`
	Links         ResourceLinks                            `json:"links"`
	Relationships *GameCenterLeaderboardImageRelationships `json:"relationships,omitempty"`
	Type          string                                   `json:"type"`
}

// GameCenterLeaderboardImageAttributes defines model for GameCenterLeaderboardImage.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2/attributes
type GameCenterLeaderboardImageAttributes struct {
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	FileName           *string             `json:"fileName,omitempty"`
	FileSize           *int                `json:"fileSize,omitempty"`
	ImageAsset         *ImageAsset         `json:"imageAsset,omitempty"`
	UploadOperations   []UploadOperation   `json:"uploadOperations,omitempty"`
}

// GameCenterLeaderboardImageRelationships defines model for GameCenterLeaderboardImage.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2/relationships
type GameCenterLeaderboardImageRelationships struct {
	Localization *Relationship `json:"localization,omitempty"`
}

// gameCenterLeaderboardImageCreateRequest defines model for GameCenterLeaderboardImageV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2createrequest
type gameCenterLeaderboardImageCreateRequest struct {
	Attributes    GameCenterLeaderboardImageCreateRequestAttributes    `json:"attributes"`
	Relationships gameCenterLeaderboardImageCreateRequestRelationships `json:"relationships"`
	Type          string                                               `json:"type"`
}

// GameCenterLeaderboardImageCreateRequestAttributes are attributes for GameCenterLeaderboardImageV2CreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2/attributes
type GameCenterLeaderboardImageCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
}

// gameCenterLeaderboardImageCreateRequestRelationships are relationships for GameCenterLeaderboardImageV2CreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2/relationships
type gameCenterLeaderboardImageCreateRequestRelationships struct {
	Localization relationshipDeclaration `json:"localization"`
}

// gameCenterLeaderboardImageUpdateRequest defines model for GameCenterLeaderboardImageV2UpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2updaterequest
type gameCenterLeaderboardImageUpdateRequest struct {
	Attributes *GameCenterLeaderboardImageUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                             `json:"id"`
	Type       string                                             `json:"type"`
}

// GameCenterLeaderboardImageUpdateRequestAttributes are attributes for GameCenterLeaderboardImageV2UpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2/attributes
type GameCenterLeaderboardImageUpdateRequestAttributes struct {
	Uploaded *bool `json:"uploaded,omitempty"`
}

// GameCenterLeaderboardImageResponse defines model for GameCenterLeaderboardImageV2Response.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterleaderboardimagev2response
type GameCenterLeaderboardImageResponse struct {
	Data     GameCenterLeaderboardImage                   `json:"data"`
	Included []GameCenterLeaderboardImageResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                `json:"links"`
}

// GameCenterLeaderboardImageResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterLeaderboardImageResponse.
type GameCenterLeaderboardImageResponseIncluded included

// GetGameCenterLeaderboardImageQuery defines model for GetGameCenterLeaderboardImage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardimages-_id_
type GetGameCenterLeaderboardImageQuery struct {
	FieldsGameCenterLeaderboardImages        []string `url:"fields[gameCenterLeaderboardImages],omitempty"`
	FieldsGameCenterLeaderboardLocalizations []string `url:"fields[gameCenterLeaderboardLocalizations],omitempty"`
	Include                                  []string `url:"include,omitempty"`
}

// CreateGameCenterLeaderboardImage creates a new image for a leaderboard localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post-v2-gamecenterleaderboardimages
func (s *GameCenterService) CreateGameCenterLeaderboardImage(ctx context.Context, attributes GameCenterLeaderboardImageCreateRequestAttributes, gameCenterLeaderboardLocalizationID string) (*GameCenterLeaderboardImageResponse, *Response, error) {
	req := gameCenterLeaderboardImageCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterLeaderboardImageCreateRequestRelationships{
			Localization: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterLeaderboardLocalizationID,
					Type: "gameCenterLeaderboardLocalizations",
				},
			},
		},
		Type: "gameCenterLeaderboardImages",
	}
	res := new(GameCenterLeaderboardImageResponse)
	resp, err := s.client.post(ctx, "../v2/gameCenterLeaderboardImages", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterLeaderboardImage gets information about a specific leaderboard image.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get-v2-gamecenterleaderboardimages-_id_
func (s *GameCenterService) GetGameCenterLeaderboardImage(ctx context.Context, id string, params *GetGameCenterLeaderboardImageQuery) (*GameCenterLeaderboardImageResponse, *Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboardImages/%s", id)
	res := new(GameCenterLeaderboardImageResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterLeaderboardImage commits a leaderboard image after uploading.
//
// https://developer.apple.com/documentation/appstoreconnectapi/patch-v2-gamecenterleaderboardimages-_id_
func (s *GameCenterService) UpdateGameCenterLeaderboardImage(ctx context.Context, id string, attributes *GameCenterLeaderboardImageUpdateRequestAttributes) (*GameCenterLeaderboardImageResponse, *Response, error) {
	req := gameCenterLeaderboardImageUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterLeaderboardImages",
	}
	url := fmt.Sprintf("../v2/gameCenterLeaderboardImages/%s", id)
	res := new(GameCenterLeaderboardImageResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterLeaderboardImage deletes a leaderboard image.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete-v2-gamecenterleaderboardimages-_id_
func (s *GameCenterService) DeleteGameCenterLeaderboardImage(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("../v2/gameCenterLeaderboardImages/%s", id)

	return s.client.delete(ctx, url, nil)
}
