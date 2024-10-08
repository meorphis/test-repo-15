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

func TestHTTPAseHTTPProtocolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Ases.HTTPProtocol.Get(
		context.TODO(),
		radar.HTTPAseHTTPProtocolGetParamsHTTPProtocolHTTP,
		radar.HTTPAseHTTPProtocolGetParams{
			ASN:           cloudflare.F([]string{"string", "string", "string"}),
			BotClass:      cloudflare.F([]radar.HTTPAseHTTPProtocolGetParamsBotClass{radar.HTTPAseHTTPProtocolGetParamsBotClassLikelyAutomated, radar.HTTPAseHTTPProtocolGetParamsBotClassLikelyHuman}),
			BrowserFamily: cloudflare.F([]radar.HTTPAseHTTPProtocolGetParamsBrowserFamily{radar.HTTPAseHTTPProtocolGetParamsBrowserFamilyChrome, radar.HTTPAseHTTPProtocolGetParamsBrowserFamilyEdge, radar.HTTPAseHTTPProtocolGetParamsBrowserFamilyFirefox}),
			Continent:     cloudflare.F([]string{"string", "string", "string"}),
			DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
			DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:    cloudflare.F([]radar.HTTPAseHTTPProtocolGetParamsDeviceType{radar.HTTPAseHTTPProtocolGetParamsDeviceTypeDesktop, radar.HTTPAseHTTPProtocolGetParamsDeviceTypeMobile, radar.HTTPAseHTTPProtocolGetParamsDeviceTypeOther}),
			Format:        cloudflare.F(radar.HTTPAseHTTPProtocolGetParamsFormatJson),
			IPVersion:     cloudflare.F([]radar.HTTPAseHTTPProtocolGetParamsIPVersion{radar.HTTPAseHTTPProtocolGetParamsIPVersionIPv4, radar.HTTPAseHTTPProtocolGetParamsIPVersionIPv6}),
			Limit:         cloudflare.F(int64(5)),
			Location:      cloudflare.F([]string{"string", "string", "string"}),
			Name:          cloudflare.F([]string{"string", "string", "string"}),
			OS:            cloudflare.F([]radar.HTTPAseHTTPProtocolGetParamsOS{radar.HTTPAseHTTPProtocolGetParamsOSWindows, radar.HTTPAseHTTPProtocolGetParamsOSMacosx, radar.HTTPAseHTTPProtocolGetParamsOSIos}),
			TLSVersion:    cloudflare.F([]radar.HTTPAseHTTPProtocolGetParamsTLSVersion{radar.HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_0, radar.HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_1, radar.HTTPAseHTTPProtocolGetParamsTLSVersionTlSv1_2}),
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
