// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_headers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/managed_headers"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestManagedHeaderList(t *testing.T) {
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
	_, err := client.ManagedHeaders.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestManagedHeaderEdit(t *testing.T) {
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
	_, err := client.ManagedHeaders.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		managed_headers.ManagedHeaderEditParams{
			ManagedRequestHeaders: cloudflare.F([]managed_headers.RequestModelParam{{
				ID:      cloudflare.F("add_cf-bot-score_header"),
				Enabled: cloudflare.F(true),
			}, {
				ID:      cloudflare.F("add_cf-bot-score_header"),
				Enabled: cloudflare.F(true),
			}, {
				ID:      cloudflare.F("add_cf-bot-score_header"),
				Enabled: cloudflare.F(true),
			}}),
			ManagedResponseHeaders: cloudflare.F([]managed_headers.RequestModelParam{{
				ID:      cloudflare.F("add_cf-bot-score_header"),
				Enabled: cloudflare.F(true),
			}, {
				ID:      cloudflare.F("add_cf-bot-score_header"),
				Enabled: cloudflare.F(true),
			}, {
				ID:      cloudflare.F("add_cf-bot-score_header"),
				Enabled: cloudflare.F(true),
			}}),
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
