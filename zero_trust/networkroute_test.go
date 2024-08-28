// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/zero_trust"
)

func TestNetworkRouteNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		zero_trust.NetworkRouteNewParams{
			Network:          cloudflare.F("172.16.0.0/16"),
			TunnelID:         cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
			Comment:          cloudflare.F("Example comment for this route."),
			VirtualNetworkID: cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
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

func TestNetworkRouteListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.List(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		zero_trust.NetworkRouteListParams{
			Comment:          cloudflare.F("Example comment for this route."),
			ExistedAt:        cloudflare.F(time.Now()),
			IsDeleted:        cloudflare.F(true),
			NetworkSubset:    cloudflare.F("172.16.0.0/16"),
			NetworkSuperset:  cloudflare.F("172.16.0.0/16"),
			Page:             cloudflare.F(1.000000),
			PerPage:          cloudflare.F(1.000000),
			RouteID:          cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
			TunTypes:         cloudflare.F("cfd_tunnel,warp_connector"),
			TunnelID:         cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
			VirtualNetworkID: cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
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

func TestNetworkRouteDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouteEditWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Networks.Routes.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.NetworkRouteEditParams{
			Comment:          cloudflare.F("Example comment for this route."),
			Network:          cloudflare.F("172.16.0.0/16"),
			TunnelID:         cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
			VirtualNetworkID: cloudflare.F("f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"),
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
