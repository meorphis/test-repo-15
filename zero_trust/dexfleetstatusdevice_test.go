// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/zero_trust"
)

func TestDEXFleetStatusDeviceListWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.ZeroTrust.DEX.FleetStatus.Devices.List(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		zero_trust.DEXFleetStatusDeviceListParams{
			From:     cloudflare.F("2023-10-11T00:00:00Z"),
			Page:     cloudflare.F(1.000000),
			PerPage:  cloudflare.F(10.000000),
			Source:   cloudflare.F(zero_trust.DEXFleetStatusDeviceListParamsSourceLastSeen),
			To:       cloudflare.F("2023-10-11T00:00:00Z"),
			Colo:     cloudflare.F("SJC"),
			DeviceID: cloudflare.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
			Mode:     cloudflare.F("proxy"),
			Platform: cloudflare.F("windows"),
			SortBy:   cloudflare.F(zero_trust.DEXFleetStatusDeviceListParamsSortByColo),
			Status:   cloudflare.F("connected"),
			Version:  cloudflare.F("1.0.0"),
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
