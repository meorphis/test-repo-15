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

func TestHTTPAseTLSVersionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.TLSVersion.Get(
		context.TODO(),
		radar.HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0,
		radar.HTTPAseTLSVersionGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPAseTLSVersionGetParamsBotClass{radar.HTTPAseTLSVersionGetParamsBotClassLikelyAutomated, radar.HTTPAseTLSVersionGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseTLSVersionGetParamsBrowserFamily{radar.HTTPAseTLSVersionGetParamsBrowserFamilyChrome, radar.HTTPAseTLSVersionGetParamsBrowserFamilyEdge, radar.HTTPAseTLSVersionGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseTLSVersionGetParamsDeviceType{radar.HTTPAseTLSVersionGetParamsDeviceTypeDesktop, radar.HTTPAseTLSVersionGetParamsDeviceTypeMobile, radar.HTTPAseTLSVersionGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPAseTLSVersionGetParamsFormatJson),
			HTTPProtocol:  cloudflare.F([]radar.HTTPAseTLSVersionGetParamsHTTPProtocol{radar.HTTPAseTLSVersionGetParamsHTTPProtocolHTTP, radar.HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:   cloudflare.F([]radar.HTTPAseTLSVersionGetParamsHTTPVersion{radar.HTTPAseTLSVersionGetParamsHTTPVersionHttPv1, radar.HTTPAseTLSVersionGetParamsHTTPVersionHttPv2, radar.HTTPAseTLSVersionGetParamsHTTPVersionHttPv3}),
			IPVersion:     cloudflare.F([]radar.HTTPAseTLSVersionGetParamsIPVersion{radar.HTTPAseTLSVersionGetParamsIPVersionIPv4, radar.HTTPAseTLSVersionGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPAseTLSVersionGetParamsOS{radar.HTTPAseTLSVersionGetParamsOSWindows, radar.HTTPAseTLSVersionGetParamsOSMacosx, radar.HTTPAseTLSVersionGetParamsOSIos}),
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
