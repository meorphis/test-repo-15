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

func TestHTTPTopBrowserFamiliesWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.BrowserFamilies(context.TODO(), radar.HTTPTopBrowserFamiliesParams{
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		BotClass:      cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsBotClass{radar.HTTPTopBrowserFamiliesParamsBotClassLikelyAutomated, radar.HTTPTopBrowserFamiliesParamsBotClassLikelyHuman}),
		BrowserFamily: cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsBrowserFamily{radar.HTTPTopBrowserFamiliesParamsBrowserFamilyChrome, radar.HTTPTopBrowserFamiliesParamsBrowserFamilyEdge, radar.HTTPTopBrowserFamiliesParamsBrowserFamilyFirefox}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsDeviceType{radar.HTTPTopBrowserFamiliesParamsDeviceTypeDesktop, radar.HTTPTopBrowserFamiliesParamsDeviceTypeMobile, radar.HTTPTopBrowserFamiliesParamsDeviceTypeOther}),
		Format:        cloudflare.F(radar.HTTPTopBrowserFamiliesParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsHTTPProtocol{radar.HTTPTopBrowserFamiliesParamsHTTPProtocolHTTP, radar.HTTPTopBrowserFamiliesParamsHTTPProtocolHTTPS}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsHTTPVersion{radar.HTTPTopBrowserFamiliesParamsHTTPVersionHttPv1, radar.HTTPTopBrowserFamiliesParamsHTTPVersionHttPv2, radar.HTTPTopBrowserFamiliesParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsIPVersion{radar.HTTPTopBrowserFamiliesParamsIPVersionIPv4, radar.HTTPTopBrowserFamiliesParamsIPVersionIPv6}),
		Limit:         cloudflare.F(int64(5)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		OS:            cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsOS{radar.HTTPTopBrowserFamiliesParamsOSWindows, radar.HTTPTopBrowserFamiliesParamsOSMacosx, radar.HTTPTopBrowserFamiliesParamsOSIos}),
		TLSVersion:    cloudflare.F([]radar.HTTPTopBrowserFamiliesParamsTLSVersion{radar.HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_0, radar.HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_1, radar.HTTPTopBrowserFamiliesParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPTopBrowsersWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Top.Browsers(context.TODO(), radar.HTTPTopBrowsersParams{
		ASN:           cloudflare.F([]string{"string", "string", "string"}),
		BotClass:      cloudflare.F([]radar.HTTPTopBrowsersParamsBotClass{radar.HTTPTopBrowsersParamsBotClassLikelyAutomated, radar.HTTPTopBrowsersParamsBotClassLikelyHuman}),
		BrowserFamily: cloudflare.F([]radar.HTTPTopBrowsersParamsBrowserFamily{radar.HTTPTopBrowsersParamsBrowserFamilyChrome, radar.HTTPTopBrowsersParamsBrowserFamilyEdge, radar.HTTPTopBrowsersParamsBrowserFamilyFirefox}),
		Continent:     cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:       cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:     cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:    cloudflare.F([]radar.HTTPTopBrowsersParamsDeviceType{radar.HTTPTopBrowsersParamsDeviceTypeDesktop, radar.HTTPTopBrowsersParamsDeviceTypeMobile, radar.HTTPTopBrowsersParamsDeviceTypeOther}),
		Format:        cloudflare.F(radar.HTTPTopBrowsersParamsFormatJson),
		HTTPProtocol:  cloudflare.F([]radar.HTTPTopBrowsersParamsHTTPProtocol{radar.HTTPTopBrowsersParamsHTTPProtocolHTTP, radar.HTTPTopBrowsersParamsHTTPProtocolHTTPS}),
		HTTPVersion:   cloudflare.F([]radar.HTTPTopBrowsersParamsHTTPVersion{radar.HTTPTopBrowsersParamsHTTPVersionHttPv1, radar.HTTPTopBrowsersParamsHTTPVersionHttPv2, radar.HTTPTopBrowsersParamsHTTPVersionHttPv3}),
		IPVersion:     cloudflare.F([]radar.HTTPTopBrowsersParamsIPVersion{radar.HTTPTopBrowsersParamsIPVersionIPv4, radar.HTTPTopBrowsersParamsIPVersionIPv6}),
		Limit:         cloudflare.F(int64(5)),
		Location:      cloudflare.F([]string{"string", "string", "string"}),
		Name:          cloudflare.F([]string{"string", "string", "string"}),
		OS:            cloudflare.F([]radar.HTTPTopBrowsersParamsOS{radar.HTTPTopBrowsersParamsOSWindows, radar.HTTPTopBrowsersParamsOSMacosx, radar.HTTPTopBrowsersParamsOSIos}),
		TLSVersion:    cloudflare.F([]radar.HTTPTopBrowsersParamsTLSVersion{radar.HTTPTopBrowsersParamsTLSVersionTlSv1_0, radar.HTTPTopBrowsersParamsTLSVersionTlSv1_1, radar.HTTPTopBrowsersParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
