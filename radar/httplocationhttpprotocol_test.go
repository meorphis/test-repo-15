// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/radar"
)

func TestHTTPLocationHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Locations.HTTPProtocol.Get(
		context.TODO(),
		radar.HTTPLocationHTTPProtocolGetParamsHTTPProtocolHTTP,
		radar.HTTPLocationHTTPProtocolGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPLocationHTTPProtocolGetParamsBotClass{radar.HTTPLocationHTTPProtocolGetParamsBotClassLikelyAutomated, radar.HTTPLocationHTTPProtocolGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPLocationHTTPProtocolGetParamsBrowserFamily{radar.HTTPLocationHTTPProtocolGetParamsBrowserFamilyChrome, radar.HTTPLocationHTTPProtocolGetParamsBrowserFamilyEdge, radar.HTTPLocationHTTPProtocolGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPLocationHTTPProtocolGetParamsDeviceType{radar.HTTPLocationHTTPProtocolGetParamsDeviceTypeDesktop, radar.HTTPLocationHTTPProtocolGetParamsDeviceTypeMobile, radar.HTTPLocationHTTPProtocolGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPLocationHTTPProtocolGetParamsFormatJson),
			IPVersion:     cloudflare.F([]radar.HTTPLocationHTTPProtocolGetParamsIPVersion{radar.HTTPLocationHTTPProtocolGetParamsIPVersionIPv4, radar.HTTPLocationHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPLocationHTTPProtocolGetParamsOS{radar.HTTPLocationHTTPProtocolGetParamsOSWindows, radar.HTTPLocationHTTPProtocolGetParamsOSMacosx, radar.HTTPLocationHTTPProtocolGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPLocationHTTPProtocolGetParamsTLSVersion{radar.HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_0, radar.HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_1, radar.HTTPLocationHTTPProtocolGetParamsTLSVersionTlSv1_2}),
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
