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

// GameCenterAchievementImage defines model for GameCenterAchievementImage.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimage
type GameCenterAchievementImage struct {
	Attributes    *GameCenterAchievementImageAttributes    `json:"attributes,omitempty"`
	ID            string                                   `json:"id"`
	Links         ResourceLinks                            `json:"links"`
	Relationships *GameCenterAchievementImageRelationships `json:"relationships,omitempty"`
	Type          string                                   `json:"type"`
}

// GameCenterAchievementImageAttributes defines model for GameCenterAchievementImage.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimage/attributes
type GameCenterAchievementImageAttributes struct {
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	FileName           *string             `json:"fileName,omitempty"`
	FileSize           *int                `json:"fileSize,omitempty"`
	ImageAsset         *ImageAsset         `json:"imageAsset,omitempty"`
	UploadOperations   []UploadOperation   `json:"uploadOperations,omitempty"`
}

// GameCenterAchievementImageRelationships defines model for GameCenterAchievementImage.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimage/relationships
type GameCenterAchievementImageRelationships struct {
	GameCenterAchievementLocalization *Relationship `json:"gameCenterAchievementLocalization,omitempty"`
}

// gameCenterAchievementImageCreateRequest defines model for GameCenterAchievementImageCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimagecreaterequest/data
type gameCenterAchievementImageCreateRequest struct {
	Attributes    GameCenterAchievementImageCreateRequestAttributes    `json:"attributes"`
	Relationships gameCenterAchievementImageCreateRequestRelationships `json:"relationships"`
	Type          string                                               `json:"type"`
}

// GameCenterAchievementImageCreateRequestAttributes are attributes for GameCenterAchievementImageCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimagecreaterequest/data/attributes
type GameCenterAchievementImageCreateRequestAttributes struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
}

// gameCenterAchievementImageCreateRequestRelationships are relationships for GameCenterAchievementImageCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimagecreaterequest/data/relationships
type gameCenterAchievementImageCreateRequestRelationships struct {
	GameCenterAchievementLocalization relationshipDeclaration `json:"gameCenterAchievementLocalization"`
}

// gameCenterAchievementImageUpdateRequest defines model for GameCenterAchievementImageUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimageupdaterequest/data
type gameCenterAchievementImageUpdateRequest struct {
	Attributes *GameCenterAchievementImageUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                             `json:"id"`
	Type       string                                             `json:"type"`
}

// GameCenterAchievementImageUpdateRequestAttributes are attributes for GameCenterAchievementImageUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimageupdaterequest/data/attributes
type GameCenterAchievementImageUpdateRequestAttributes struct {
	Uploaded *bool `json:"uploaded,omitempty"`
}

// GameCenterAchievementImageResponse defines model for GameCenterAchievementImageResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/gamecenterachievementimageresponse
type GameCenterAchievementImageResponse struct {
	Data     GameCenterAchievementImage                   `json:"data"`
	Included []GameCenterAchievementImageResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                                `json:"links"`
}

// GameCenterAchievementImageResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a GameCenterAchievementImageResponse.
type GameCenterAchievementImageResponseIncluded included

// GetGameCenterAchievementImageQuery defines model for GetGameCenterAchievementImage
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_image_information
type GetGameCenterAchievementImageQuery struct {
	FieldsGameCenterAchievementImages        []string `url:"fields[gameCenterAchievementImages],omitempty"`
	FieldsGameCenterAchievementLocalizations []string `url:"fields[gameCenterAchievementLocalizations],omitempty"`
	Include                                  []string `url:"include,omitempty"`
}

// CreateGameCenterAchievementImage creates a new image for an achievement localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_achievement_image
func (s *GameCenterService) CreateGameCenterAchievementImage(ctx context.Context, attributes GameCenterAchievementImageCreateRequestAttributes, gameCenterAchievementLocalizationID string) (*GameCenterAchievementImageResponse, *Response, error) {
	req := gameCenterAchievementImageCreateRequest{
		Attributes: attributes,
		Relationships: gameCenterAchievementImageCreateRequestRelationships{
			GameCenterAchievementLocalization: relationshipDeclaration{
				Data: RelationshipData{
					ID:   gameCenterAchievementLocalizationID,
					Type: "gameCenterAchievementLocalizations",
				},
			},
		},
		Type: "gameCenterAchievementImages",
	}
	res := new(GameCenterAchievementImageResponse)
	resp, err := s.client.post(ctx, "gameCenterAchievementImages", newRequestBody(req), res)

	return res, resp, err
}

// GetGameCenterAchievementImage gets information about a specific achievement image.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_achievement_image_information
func (s *GameCenterService) GetGameCenterAchievementImage(ctx context.Context, id string, params *GetGameCenterAchievementImageQuery) (*GameCenterAchievementImageResponse, *Response, error) {
	url := fmt.Sprintf("gameCenterAchievementImages/%s", id)
	res := new(GameCenterAchievementImageResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// UpdateGameCenterAchievementImage commits an achievement image after uploading.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_achievement_image
func (s *GameCenterService) UpdateGameCenterAchievementImage(ctx context.Context, id string, attributes *GameCenterAchievementImageUpdateRequestAttributes) (*GameCenterAchievementImageResponse, *Response, error) {
	req := gameCenterAchievementImageUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "gameCenterAchievementImages",
	}
	url := fmt.Sprintf("gameCenterAchievementImages/%s", id)
	res := new(GameCenterAchievementImageResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteGameCenterAchievementImage deletes an achievement image.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_achievement_image
func (s *GameCenterService) DeleteGameCenterAchievementImage(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("gameCenterAchievementImages/%s", id)

	return s.client.delete(ctx, url, nil)
}
