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
	"testing"
)

func TestCreateGameCenterAchievement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterAchievement(ctx, GameCenterAchievementCreateRequestAttributes{
			ReferenceName:    "Test Achievement",
			VendorIdentifier: "com.example.achievement1",
			Points:           10,
			ShowBeforeEarned: true,
			Repeatable:       false,
		}, "gameCenterDetailID")
	})
}

func TestGetGameCenterAchievement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterAchievement(ctx, "10", &GetGameCenterAchievementQuery{})
	})
}

func TestUpdateGameCenterAchievement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterAchievement(ctx, "10", &GameCenterAchievementUpdateRequestAttributes{
			ReferenceName: String("Updated Achievement"),
		})
	})
}

func TestDeleteGameCenterAchievement(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterAchievement(ctx, "10")
	})
}

func TestListGameCenterAchievementsForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterAchievementsForDetail(ctx, "gameCenterDetailID", &ListGameCenterAchievementsQuery{})
	})
}

func TestCreateGameCenterAchievementLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterAchievementLocalization(ctx, GameCenterAchievementLocalizationCreateRequestAttributes{
			Locale:                  "en-US",
			Name:                    "Test Achievement",
			BeforeEarnedDescription: "Earn this achievement",
			AfterEarnedDescription:  "You earned this achievement",
		}, "achievementID")
	})
}

func TestGetGameCenterAchievementLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterAchievementLocalization(ctx, "10", &GetGameCenterAchievementLocalizationQuery{})
	})
}

func TestUpdateGameCenterAchievementLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterAchievementLocalization(ctx, "10", &GameCenterAchievementLocalizationUpdateRequestAttributes{
			Name: String("Updated Achievement Name"),
		})
	})
}

func TestDeleteGameCenterAchievementLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterAchievementLocalization(ctx, "10")
	})
}

func TestListGameCenterAchievementLocalizationsForAchievement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterAchievementLocalizationsForAchievement(ctx, "achievementID", &ListGameCenterAchievementLocalizationsQuery{})
	})
}

func TestCreateGameCenterAchievementImage(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementImageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterAchievementImage(ctx, GameCenterAchievementImageCreateRequestAttributes{
			FileName: "achievement.png",
			FileSize: 1024,
		}, "localizationID")
	})
}

func TestGetGameCenterAchievementImage(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementImageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterAchievementImage(ctx, "10", &GetGameCenterAchievementImageQuery{})
	})
}

func TestUpdateGameCenterAchievementImage(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementImageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterAchievementImage(ctx, "10", &GameCenterAchievementImageUpdateRequestAttributes{
			Uploaded: Bool(true),
		})
	})
}

func TestDeleteGameCenterAchievementImage(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterAchievementImage(ctx, "10")
	})
}

func TestCreateGameCenterAchievementRelease(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterAchievementRelease(ctx, "achievementID", "gameCenterDetailID")
	})
}

func TestGetGameCenterAchievementRelease(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterAchievementRelease(ctx, "10", &GetGameCenterAchievementReleaseQuery{})
	})
}

func TestDeleteGameCenterAchievementRelease(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterAchievementRelease(ctx, "10")
	})
}

func TestListGameCenterAchievementReleasesForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementReleasesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterAchievementReleasesForDetail(ctx, "gameCenterDetailID", &ListGameCenterAchievementReleasesQuery{})
	})
}

func TestListGameCenterAchievementReleasesForAchievement(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterAchievementReleasesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterAchievementReleasesForAchievement(ctx, "achievementID", &ListGameCenterAchievementReleasesQuery{})
	})
}

func TestReplaceGameCenterAchievementReleasesForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.ReplaceGameCenterAchievementReleasesForDetail(ctx, "gameCenterDetailID", []string{"release1", "release2", "release3"})
	})
}

func TestCreateGameCenterDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterDetail(ctx, "appID")
	})
}

func TestGetGameCenterDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterDetail(ctx, "10", &GetGameCenterDetailQuery{})
	})
}

func TestGetGameCenterDetailForApp(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterDetailForApp(ctx, "appID", &GetGameCenterDetailForAppQuery{})
	})
}

func TestUpdateGameCenterDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterDetailResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterDetail(ctx, "10", &GameCenterDetailUpdateRequestAttributes{
			ChallengeEnabled: Bool(true),
		}, nil, nil, nil)
	})
}
