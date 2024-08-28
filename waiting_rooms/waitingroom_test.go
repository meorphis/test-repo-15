// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/waiting_rooms"
)

func TestWaitingRoomNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.WaitingRooms.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		waiting_rooms.WaitingRoomNewParams{
			Query: waiting_rooms.QueryParam{
				Host:              cloudflare.F("shop.example.com"),
				Name:              cloudflare.F("production_webinar"),
				NewUsersPerMinute: cloudflare.F(int64(200)),
				TotalActiveUsers:  cloudflare.F(int64(200)),
				AdditionalRoutes: cloudflare.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}, {
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}, {
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}}),
				CookieAttributes: cloudflare.F(waiting_rooms.CookieAttributesParam{
					Samesite: cloudflare.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   cloudflare.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cloudflare.F("abcd"),
				CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cloudflare.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             cloudflare.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cloudflare.F(false),
				JsonResponseEnabled:     cloudflare.F(false),
				Path:                    cloudflare.F("/shop/checkout"),
				QueueAll:                cloudflare.F(true),
				QueueingMethod:          cloudflare.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      cloudflare.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         cloudflare.F(int64(1)),
				Suspended:               cloudflare.F(true),
			},
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.WaitingRooms.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomUpdateParams{
			Query: waiting_rooms.QueryParam{
				Host:              cloudflare.F("shop.example.com"),
				Name:              cloudflare.F("production_webinar"),
				NewUsersPerMinute: cloudflare.F(int64(200)),
				TotalActiveUsers:  cloudflare.F(int64(200)),
				AdditionalRoutes: cloudflare.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}, {
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}, {
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}}),
				CookieAttributes: cloudflare.F(waiting_rooms.CookieAttributesParam{
					Samesite: cloudflare.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   cloudflare.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cloudflare.F("abcd"),
				CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cloudflare.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             cloudflare.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cloudflare.F(false),
				JsonResponseEnabled:     cloudflare.F(false),
				Path:                    cloudflare.F("/shop/checkout"),
				QueueAll:                cloudflare.F(true),
				QueueingMethod:          cloudflare.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      cloudflare.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         cloudflare.F(int64(1)),
				Suspended:               cloudflare.F(true),
			},
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.WaitingRooms.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		waiting_rooms.WaitingRoomListParams{
			Page:    cloudflare.F[any](map[string]interface{}{}),
			PerPage: cloudflare.F[any](map[string]interface{}{}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.WaitingRooms.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomEditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.WaitingRooms.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		waiting_rooms.WaitingRoomEditParams{
			Query: waiting_rooms.QueryParam{
				Host:              cloudflare.F("shop.example.com"),
				Name:              cloudflare.F("production_webinar"),
				NewUsersPerMinute: cloudflare.F(int64(200)),
				TotalActiveUsers:  cloudflare.F(int64(200)),
				AdditionalRoutes: cloudflare.F([]waiting_rooms.AdditionalRoutesParam{{
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}, {
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}, {
					Host: cloudflare.F("shop2.example.com"),
					Path: cloudflare.F("/shop2/checkout"),
				}}),
				CookieAttributes: cloudflare.F(waiting_rooms.CookieAttributesParam{
					Samesite: cloudflare.F(waiting_rooms.CookieAttributesSamesiteAuto),
					Secure:   cloudflare.F(waiting_rooms.CookieAttributesSecureAuto),
				}),
				CookieSuffix:            cloudflare.F("abcd"),
				CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
				DefaultTemplateLanguage: cloudflare.F(waiting_rooms.QueryDefaultTemplateLanguageEnUs),
				Description:             cloudflare.F("Production - DO NOT MODIFY"),
				DisableSessionRenewal:   cloudflare.F(false),
				JsonResponseEnabled:     cloudflare.F(false),
				Path:                    cloudflare.F("/shop/checkout"),
				QueueAll:                cloudflare.F(true),
				QueueingMethod:          cloudflare.F(waiting_rooms.QueryQueueingMethodFifo),
				QueueingStatusCode:      cloudflare.F(waiting_rooms.QueryQueueingStatusCode200),
				SessionDuration:         cloudflare.F(int64(1)),
				Suspended:               cloudflare.F(true),
			},
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWaitingRoomGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.WaitingRooms.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
