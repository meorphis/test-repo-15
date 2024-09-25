// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/ai_gateway"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestLogListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIGateway.Logs.List(
		context.TODO(),
		"0d37909e38d3e99c29fa2cd343ac421a",
		"my-gateway",
		ai_gateway.LogListParams{
			Cached:    cloudflare.F(true),
			Direction: cloudflare.F(ai_gateway.LogListParamsDirectionAsc),
			EndDate:   cloudflare.F(time.Now()),
			OrderBy:   cloudflare.F(ai_gateway.LogListParamsOrderByCreatedAt),
			Page:      cloudflare.F(int64(1)),
			PerPage:   cloudflare.F(int64(5)),
			Search:    cloudflare.F("search"),
			StartDate: cloudflare.F(time.Now()),
			Success:   cloudflare.F(true),
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
