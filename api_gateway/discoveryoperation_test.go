// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/api_gateway"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestDiscoveryOperationListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Discovery.Operations.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		api_gateway.DiscoveryOperationListParams{
			Diff:      cloudflare.F(true),
			Direction: cloudflare.F(api_gateway.DiscoveryOperationListParamsDirectionAsc),
			Endpoint:  cloudflare.F("/api/v1"),
			Host:      cloudflare.F([]string{"api.cloudflare.com"}),
			Method:    cloudflare.F([]string{"GET"}),
			Order:     cloudflare.F(api_gateway.DiscoveryOperationListParamsOrderHost),
			Origin:    cloudflare.F(api_gateway.DiscoveryOperationListParamsOriginMl),
			Page:      cloudflare.F(int64(1)),
			PerPage:   cloudflare.F(int64(5)),
			State:     cloudflare.F(api_gateway.DiscoveryOperationListParamsStateReview),
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

func TestDiscoveryOperationEditWithOptionalParams(t *testing.T) {
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
	_, err := client.APIGateway.Discovery.Operations.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		api_gateway.DiscoveryOperationEditParams{
			State: cloudflare.F(api_gateway.DiscoveryOperationEditParamsStateReview),
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
