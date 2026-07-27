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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGameCenterLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterLeaderboard(ctx, GameCenterLeaderboardCreateRequestAttributes{
			DefaultFormatter: "INTEGER",
			ReferenceName:    "Test Leaderboard",
			VendorIdentifier: "com.example.leaderboard1",
			ScoreSortType:    "DESC",
			SubmissionType:   "BEST_SCORE",
		}, "gameCenterDetailID")
	})
}

func TestGameCenterLeaderboardV2EndpointPaths(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		wantPath string
		call     func(context.Context, *Client) (*Response, error)
	}{
		{
			name:     "create leaderboard",
			wantPath: "/v2/gameCenterLeaderboards",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				_, resp, err := client.GameCenter.CreateGameCenterLeaderboard(ctx, GameCenterLeaderboardCreateRequestAttributes{
					DefaultFormatter: "INTEGER",
					ReferenceName:    "Test Leaderboard",
					VendorIdentifier: "com.example.leaderboard1",
					ScoreSortType:    "DESC",
					SubmissionType:   "BEST_SCORE",
				}, "gameCenterDetailID")

				return resp, err
			},
		},
		{
			name:     "get leaderboard version",
			wantPath: "/v2/gameCenterLeaderboardVersions/10",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				_, resp, err := client.GameCenter.GetGameCenterLeaderboardVersion(ctx, "10", nil)

				return resp, err
			},
		},
		{
			name:     "create localization",
			wantPath: "/v2/gameCenterLeaderboardLocalizations",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				_, resp, err := client.GameCenter.CreateGameCenterLeaderboardLocalization(ctx, GameCenterLeaderboardLocalizationCreateRequestAttributes{
					Locale: "en-US",
					Name:   "Test Leaderboard",
				}, "versionID")

				return resp, err
			},
		},
		{
			name:     "create image",
			wantPath: "/v2/gameCenterLeaderboardImages",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				_, resp, err := client.GameCenter.CreateGameCenterLeaderboardImage(ctx, GameCenterLeaderboardImageCreateRequestAttributes{
					FileName: "leaderboard.png",
					FileSize: 1024,
				}, "localizationID")

				return resp, err
			},
		},
		{
			name:     "list detail leaderboards v2",
			wantPath: "/v1/gameCenterDetails/gameCenterDetailID/gameCenterLeaderboardsV2",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				_, resp, err := client.GameCenter.ListGameCenterLeaderboardsForDetail(ctx, "gameCenterDetailID", nil)

				return resp, err
			},
		},
		{
			name:     "list detail leaderboard releases",
			wantPath: "/v1/gameCenterDetails/gameCenterDetailID/leaderboardReleases",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				_, resp, err := client.GameCenter.ListGameCenterLeaderboardReleasesForDetail(ctx, "gameCenterDetailID", nil)

				return resp, err
			},
		},
		{
			name:     "replace detail leaderboards v2",
			wantPath: "/v1/gameCenterDetails/gameCenterDetailID/relationships/gameCenterLeaderboardsV2",
			call: func(ctx context.Context, client *Client) (*Response, error) {
				return client.GameCenter.ReplaceGameCenterLeaderboardsForDetail(ctx, "gameCenterDetailID", []string{"leaderboard1"})
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var gotPath string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.Path
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "{}")
			}))
			defer server.Close()

			client := NewClient(server.Client())
			base, err := url.Parse(server.URL + "/v1/")
			assert.NoError(t, err)
			client.baseURL = base

			resp, err := tc.call(context.Background(), client)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, tc.wantPath, gotPath)
		})
	}
}

func TestCreateGameCenterLeaderboardRequestBody(t *testing.T) {
	t.Parallel()
	scoreRangeEnd := int64(1000)
	scoreRangeStart := int64(-1000)

	var gotBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		gotBody = string(body)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "{}")
	}))
	defer server.Close()

	client := NewClient(server.Client())
	base, err := url.Parse(server.URL + "/v1/")
	assert.NoError(t, err)
	client.baseURL = base

	_, _, err = client.GameCenter.CreateGameCenterLeaderboard(context.Background(), GameCenterLeaderboardCreateRequestAttributes{
		DefaultFormatter: "INTEGER",
		ReferenceName:    "Test Leaderboard",
		VendorIdentifier: "com.example.leaderboard1",
		ScoreSortType:    "DESC",
		SubmissionType:   "BEST_SCORE",
		ScoreRangeEnd:    &scoreRangeEnd,
		ScoreRangeStart:  &scoreRangeStart,
	}, "gameCenterDetailID")
	assert.NoError(t, err)

	assert.JSONEq(t, `{
		"data": {
			"attributes": {
				"defaultFormatter": "INTEGER",
				"referenceName": "Test Leaderboard",
				"vendorIdentifier": "com.example.leaderboard1",
				"scoreSortType": "DESC",
				"submissionType": "BEST_SCORE",
				"scoreRangeEnd": 1000,
				"scoreRangeStart": -1000
			},
			"relationships": {
				"gameCenterDetail": {
					"data": {
						"id": "gameCenterDetailID",
						"type": "gameCenterDetails"
					}
				},
				"versions": {
					"data": [
						{
							"id": "${new-gameCenterLeaderboardVersion-id}",
							"type": "gameCenterLeaderboardVersions"
						}
					]
				}
			},
			"type": "gameCenterLeaderboards"
		},
		"included": [
			{
				"id": "${new-gameCenterLeaderboardVersion-id}",
				"type": "gameCenterLeaderboardVersions"
			}
		]
	}`, gotBody)
}

func TestGetGameCenterDetailIncludesLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointCustomBehavior(`{
		"data": {
			"id": "detailID",
			"type": "gameCenterDetails",
			"links": {
				"self": "https://api.appstoreconnect.apple.com/v1/gameCenterDetails/detailID"
			}
		},
		"included": [
			{
				"id": "leaderboardID",
				"type": "gameCenterLeaderboards",
				"links": {
					"self": "https://api.appstoreconnect.apple.com/v2/gameCenterLeaderboards/leaderboardID"
				}
			}
		]
	}`, func(ctx context.Context, client *Client) {
		detail, _, err := client.GameCenter.GetGameCenterDetail(ctx, "detailID", &GetGameCenterDetailQuery{
			Include: []string{"gameCenterLeaderboardsV2"},
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, detail.Included)
		assert.NotNil(t, detail.Included[0].GameCenterLeaderboard())
		assert.Equal(t, "leaderboardID", detail.Included[0].GameCenterLeaderboard().ID)
	})
}

func TestCreateGameCenterLeaderboardForGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterLeaderboardForGroup(ctx, GameCenterLeaderboardCreateRequestAttributes{
			DefaultFormatter: "INTEGER",
			ReferenceName:    "Test Group Leaderboard",
			VendorIdentifier: "com.example.group.leaderboard1",
			ScoreSortType:    "ASC",
			SubmissionType:   "MOST_RECENT_SCORE",
		}, "gameCenterGroupID")
	})
}

func TestGetGameCenterLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterLeaderboard(ctx, "10", &GetGameCenterLeaderboardQuery{})
	})
}

func TestListGameCenterLeaderboardsForDetailDecodesStringScoreRanges(t *testing.T) {
	t.Parallel()

	scoreRangeEnd := int64(9223372036854775807)
	scoreRangeStart := int64(-9223372036854775808)
	vendorIdentifier := "com.example.leaderboard1"
	testEndpointWithResponse(t, `{
		"data": [{
			"type": "gameCenterLeaderboards",
			"id": "leaderboardID",
			"attributes": {
				"scoreRangeEnd": "9223372036854775807",
				"scoreRangeStart": "-9223372036854775808",
				"vendorIdentifier": "com.example.leaderboard1"
			}
		}]
	}`, &GameCenterLeaderboardsResponse{
		Data: []GameCenterLeaderboard{{
			Attributes: &GameCenterLeaderboardAttributes{
				ScoreRangeEnd:    &scoreRangeEnd,
				ScoreRangeStart:  &scoreRangeStart,
				VendorIdentifier: &vendorIdentifier,
			},
			ID:   "leaderboardID",
			Type: "gameCenterLeaderboards",
		}},
	}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardsForDetail(ctx, "gameCenterDetailID", nil)
	})
}

func TestGameCenterLeaderboardAttributesDecodesScoreRangeVariants(t *testing.T) {
	t.Parallel()
	scoreRangeEnd := int64(100)
	scoreRangeStart := int64(-10)

	tests := []struct {
		name      string
		input     string
		wantStart *int64
		wantEnd   *int64
		wantError bool
	}{
		{
			name:      "numbers",
			input:     `{"scoreRangeStart":-10,"scoreRangeEnd":100}`,
			wantStart: &scoreRangeStart,
			wantEnd:   &scoreRangeEnd,
		},
		{
			name:      "numeric strings",
			input:     `{"scoreRangeStart":"-10","scoreRangeEnd":"100"}`,
			wantStart: &scoreRangeStart,
			wantEnd:   &scoreRangeEnd,
		},
		{
			name:  "nulls",
			input: `{"scoreRangeStart":null,"scoreRangeEnd":null}`,
		},
		{
			name:      "invalid numeric string",
			input:     `{"scoreRangeStart":"invalid"}`,
			wantError: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var attributes GameCenterLeaderboardAttributes
			err := json.Unmarshal([]byte(tc.input), &attributes)
			if tc.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.wantStart, attributes.ScoreRangeStart)
			assert.Equal(t, tc.wantEnd, attributes.ScoreRangeEnd)
		})
	}
}

func TestUpdateGameCenterLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterLeaderboard(ctx, "10", &GameCenterLeaderboardUpdateRequestAttributes{
			ReferenceName: String("Updated Leaderboard"),
		})
	})
}

func TestDeleteGameCenterLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterLeaderboard(ctx, "10")
	})
}

func TestListGameCenterLeaderboardsForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardsForDetail(ctx, "gameCenterDetailID", &ListGameCenterLeaderboardsQuery{})
	})
}

func TestListGameCenterLeaderboardsForGroup(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardsForGroup(ctx, "gameCenterGroupID", &ListGameCenterLeaderboardsQuery{})
	})
}

func TestReplaceGameCenterLeaderboardsForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.ReplaceGameCenterLeaderboardsForDetail(ctx, "gameCenterDetailID", []string{"leaderboard1", "leaderboard2"})
	})
}

func TestCreateGameCenterLeaderboardVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterLeaderboardVersion(ctx, "leaderboardID")
	})
}

func TestGetGameCenterLeaderboardVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardVersionResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterLeaderboardVersion(ctx, "10", &GetGameCenterLeaderboardVersionQuery{})
	})
}

func TestGetGameCenterLeaderboardVersionDecodesNumericVersion(t *testing.T) {
	t.Parallel()

	version := int64(1)
	testEndpointWithResponse(t, `{
		"data": {
			"type": "gameCenterLeaderboardVersions",
			"id": "10",
			"attributes": {
				"state": "PREPARE_FOR_SUBMISSION",
				"version": 1
			}
		}
	}`, &GameCenterLeaderboardVersionResponse{
		Data: GameCenterLeaderboardVersion{
			Attributes: &GameCenterLeaderboardVersionAttributes{
				State:   String("PREPARE_FOR_SUBMISSION"),
				Version: &version,
			},
			ID:   "10",
			Type: "gameCenterLeaderboardVersions",
		},
	}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterLeaderboardVersion(ctx, "10", nil)
	})
}

func TestListGameCenterLeaderboardVersionsForLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardVersionsForLeaderboard(ctx, "leaderboardID", &ListGameCenterLeaderboardVersionsQuery{})
	})
}

func TestCreateGameCenterLeaderboardLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterLeaderboardLocalization(ctx, GameCenterLeaderboardLocalizationCreateRequestAttributes{
			Locale: "en-US",
			Name:   "Test Leaderboard",
		}, "versionID")
	})
}

func TestGetGameCenterLeaderboardLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterLeaderboardLocalization(ctx, "10", &GetGameCenterLeaderboardLocalizationQuery{})
	})
}

func TestUpdateGameCenterLeaderboardLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardLocalizationResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterLeaderboardLocalization(ctx, "10", &GameCenterLeaderboardLocalizationUpdateRequestAttributes{
			Name: String("Updated Leaderboard Name"),
		})
	})
}

func TestDeleteGameCenterLeaderboardLocalization(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterLeaderboardLocalization(ctx, "10")
	})
}

func TestListGameCenterLeaderboardLocalizationsForVersion(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardLocalizationsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardLocalizationsForVersion(ctx, "versionID", &ListGameCenterLeaderboardLocalizationsQuery{})
	})
}

func TestCreateGameCenterLeaderboardImage(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardImageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterLeaderboardImage(ctx, GameCenterLeaderboardImageCreateRequestAttributes{
			FileName: "leaderboard.png",
			FileSize: 1024,
		}, "localizationID")
	})
}

func TestGetGameCenterLeaderboardImage(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardImageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterLeaderboardImage(ctx, "10", &GetGameCenterLeaderboardImageQuery{})
	})
}

func TestUpdateGameCenterLeaderboardImage(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardImageResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.UpdateGameCenterLeaderboardImage(ctx, "10", &GameCenterLeaderboardImageUpdateRequestAttributes{
			Uploaded: Bool(true),
		})
	})
}

func TestDeleteGameCenterLeaderboardImage(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterLeaderboardImage(ctx, "10")
	})
}

func TestCreateGameCenterLeaderboardRelease(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.CreateGameCenterLeaderboardRelease(ctx, "leaderboardID", "gameCenterDetailID")
	})
}

func TestGetGameCenterLeaderboardRelease(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardReleaseResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.GetGameCenterLeaderboardRelease(ctx, "10", &GetGameCenterLeaderboardReleaseQuery{})
	})
}

func TestDeleteGameCenterLeaderboardRelease(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.DeleteGameCenterLeaderboardRelease(ctx, "10")
	})
}

func TestListGameCenterLeaderboardReleasesForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardReleasesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardReleasesForDetail(ctx, "gameCenterDetailID", &ListGameCenterLeaderboardReleasesQuery{})
	})
}

func TestListGameCenterLeaderboardReleasesForLeaderboard(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &GameCenterLeaderboardReleasesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.GameCenter.ListGameCenterLeaderboardReleasesForLeaderboard(ctx, "leaderboardID", &ListGameCenterLeaderboardReleasesQuery{})
	})
}

func TestReplaceGameCenterLeaderboardReleasesForDetail(t *testing.T) {
	t.Parallel()

	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.GameCenter.ReplaceGameCenterLeaderboardReleasesForDetail(ctx, "gameCenterDetailID", []string{"release1", "release2", "release3"})
	})
}
