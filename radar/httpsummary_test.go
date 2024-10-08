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

func TestHTTPSummaryBotClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.BotClass(context.TODO(), radar.HTTPSummaryBotClassParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryBotClassParamsDeviceType{radar.HTTPSummaryBotClassParamsDeviceTypeDesktop, radar.HTTPSummaryBotClassParamsDeviceTypeMobile, radar.HTTPSummaryBotClassParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPSummaryBotClassParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryBotClassParamsHTTPProtocol{radar.HTTPSummaryBotClassParamsHTTPProtocolHTTP, radar.HTTPSummaryBotClassParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryBotClassParamsHTTPVersion{radar.HTTPSummaryBotClassParamsHTTPVersionHttPv1, radar.HTTPSummaryBotClassParamsHTTPVersionHttPv2, radar.HTTPSummaryBotClassParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryBotClassParamsIPVersion{radar.HTTPSummaryBotClassParamsIPVersionIPv4, radar.HTTPSummaryBotClassParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryBotClassParamsOS{radar.HTTPSummaryBotClassParamsOSWindows, radar.HTTPSummaryBotClassParamsOSMacosx, radar.HTTPSummaryBotClassParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryBotClassParamsTLSVersion{radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_0, radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_1, radar.HTTPSummaryBotClassParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryDeviceTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.DeviceType(context.TODO(), radar.HTTPSummaryDeviceTypeParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsBotClass{radar.HTTPSummaryDeviceTypeParamsBotClassLikelyAutomated, radar.HTTPSummaryDeviceTypeParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Format:       cloudflare.F(radar.HTTPSummaryDeviceTypeParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsHTTPProtocol{radar.HTTPSummaryDeviceTypeParamsHTTPProtocolHTTP, radar.HTTPSummaryDeviceTypeParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsHTTPVersion{radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv1, radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv2, radar.HTTPSummaryDeviceTypeParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsIPVersion{radar.HTTPSummaryDeviceTypeParamsIPVersionIPv4, radar.HTTPSummaryDeviceTypeParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsOS{radar.HTTPSummaryDeviceTypeParamsOSWindows, radar.HTTPSummaryDeviceTypeParamsOSMacosx, radar.HTTPSummaryDeviceTypeParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryDeviceTypeParamsTLSVersion{radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_0, radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_1, radar.HTTPSummaryDeviceTypeParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryHTTPProtocolWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPProtocol(context.TODO(), radar.HTTPSummaryHTTPProtocolParams{
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		BotClass:    cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsBotClass{radar.HTTPSummaryHTTPProtocolParamsBotClassLikelyAutomated, radar.HTTPSummaryHTTPProtocolParamsBotClassLikelyHuman}),
		Continent:   cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:  cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsDeviceType{radar.HTTPSummaryHTTPProtocolParamsDeviceTypeDesktop, radar.HTTPSummaryHTTPProtocolParamsDeviceTypeMobile, radar.HTTPSummaryHTTPProtocolParamsDeviceTypeOther}),
		Format:      cloudflare.F(radar.HTTPSummaryHTTPProtocolParamsFormatJson),
		HTTPVersion: cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsHTTPVersion{radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv1, radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv2, radar.HTTPSummaryHTTPProtocolParamsHTTPVersionHttPv3}),
		IPVersion:   cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsIPVersion{radar.HTTPSummaryHTTPProtocolParamsIPVersionIPv4, radar.HTTPSummaryHTTPProtocolParamsIPVersionIPv6}),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
		OS:          cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsOS{radar.HTTPSummaryHTTPProtocolParamsOSWindows, radar.HTTPSummaryHTTPProtocolParamsOSMacosx, radar.HTTPSummaryHTTPProtocolParamsOSIos}),
		TLSVersion:  cloudflare.F([]radar.HTTPSummaryHTTPProtocolParamsTLSVersion{radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_0, radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_1, radar.HTTPSummaryHTTPProtocolParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryHTTPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.HTTPVersion(context.TODO(), radar.HTTPSummaryHTTPVersionParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsBotClass{radar.HTTPSummaryHTTPVersionParamsBotClassLikelyAutomated, radar.HTTPSummaryHTTPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsDeviceType{radar.HTTPSummaryHTTPVersionParamsDeviceTypeDesktop, radar.HTTPSummaryHTTPVersionParamsDeviceTypeMobile, radar.HTTPSummaryHTTPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPSummaryHTTPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsHTTPProtocol{radar.HTTPSummaryHTTPVersionParamsHTTPProtocolHTTP, radar.HTTPSummaryHTTPVersionParamsHTTPProtocolHTTPS}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsIPVersion{radar.HTTPSummaryHTTPVersionParamsIPVersionIPv4, radar.HTTPSummaryHTTPVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsOS{radar.HTTPSummaryHTTPVersionParamsOSWindows, radar.HTTPSummaryHTTPVersionParamsOSMacosx, radar.HTTPSummaryHTTPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryHTTPVersionParamsTLSVersion{radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_0, radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_1, radar.HTTPSummaryHTTPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryIPVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.IPVersion(context.TODO(), radar.HTTPSummaryIPVersionParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryIPVersionParamsBotClass{radar.HTTPSummaryIPVersionParamsBotClassLikelyAutomated, radar.HTTPSummaryIPVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryIPVersionParamsDeviceType{radar.HTTPSummaryIPVersionParamsDeviceTypeDesktop, radar.HTTPSummaryIPVersionParamsDeviceTypeMobile, radar.HTTPSummaryIPVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPSummaryIPVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryIPVersionParamsHTTPProtocol{radar.HTTPSummaryIPVersionParamsHTTPProtocolHTTP, radar.HTTPSummaryIPVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryIPVersionParamsHTTPVersion{radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv1, radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv2, radar.HTTPSummaryIPVersionParamsHTTPVersionHttPv3}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryIPVersionParamsOS{radar.HTTPSummaryIPVersionParamsOSWindows, radar.HTTPSummaryIPVersionParamsOSMacosx, radar.HTTPSummaryIPVersionParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryIPVersionParamsTLSVersion{radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_0, radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_1, radar.HTTPSummaryIPVersionParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryOSWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.OS(context.TODO(), radar.HTTPSummaryOSParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryOSParamsBotClass{radar.HTTPSummaryOSParamsBotClassLikelyAutomated, radar.HTTPSummaryOSParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryOSParamsDeviceType{radar.HTTPSummaryOSParamsDeviceTypeDesktop, radar.HTTPSummaryOSParamsDeviceTypeMobile, radar.HTTPSummaryOSParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPSummaryOSParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryOSParamsHTTPProtocol{radar.HTTPSummaryOSParamsHTTPProtocolHTTP, radar.HTTPSummaryOSParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryOSParamsHTTPVersion{radar.HTTPSummaryOSParamsHTTPVersionHttPv1, radar.HTTPSummaryOSParamsHTTPVersionHttPv2, radar.HTTPSummaryOSParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryOSParamsIPVersion{radar.HTTPSummaryOSParamsIPVersionIPv4, radar.HTTPSummaryOSParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryOSParamsTLSVersion{radar.HTTPSummaryOSParamsTLSVersionTlSv1_0, radar.HTTPSummaryOSParamsTLSVersionTlSv1_1, radar.HTTPSummaryOSParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryPostQuantumWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.PostQuantum(context.TODO(), radar.HTTPSummaryPostQuantumParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryPostQuantumParamsBotClass{radar.HTTPSummaryPostQuantumParamsBotClassLikelyAutomated, radar.HTTPSummaryPostQuantumParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryPostQuantumParamsDeviceType{radar.HTTPSummaryPostQuantumParamsDeviceTypeDesktop, radar.HTTPSummaryPostQuantumParamsDeviceTypeMobile, radar.HTTPSummaryPostQuantumParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPSummaryPostQuantumParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryPostQuantumParamsHTTPProtocol{radar.HTTPSummaryPostQuantumParamsHTTPProtocolHTTP, radar.HTTPSummaryPostQuantumParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryPostQuantumParamsHTTPVersion{radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv1, radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv2, radar.HTTPSummaryPostQuantumParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryPostQuantumParamsIPVersion{radar.HTTPSummaryPostQuantumParamsIPVersionIPv4, radar.HTTPSummaryPostQuantumParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryPostQuantumParamsOS{radar.HTTPSummaryPostQuantumParamsOSWindows, radar.HTTPSummaryPostQuantumParamsOSMacosx, radar.HTTPSummaryPostQuantumParamsOSIos}),
		TLSVersion:   cloudflare.F([]radar.HTTPSummaryPostQuantumParamsTLSVersion{radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_0, radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_1, radar.HTTPSummaryPostQuantumParamsTLSVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHTTPSummaryTLSVersionWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.Summary.TLSVersion(context.TODO(), radar.HTTPSummaryTLSVersionParams{
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]radar.HTTPSummaryTLSVersionParamsBotClass{radar.HTTPSummaryTLSVersionParamsBotClassLikelyAutomated, radar.HTTPSummaryTLSVersionParamsBotClassLikelyHuman}),
		Continent:    cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]string{"7d", "7d", "7d"}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]radar.HTTPSummaryTLSVersionParamsDeviceType{radar.HTTPSummaryTLSVersionParamsDeviceTypeDesktop, radar.HTTPSummaryTLSVersionParamsDeviceTypeMobile, radar.HTTPSummaryTLSVersionParamsDeviceTypeOther}),
		Format:       cloudflare.F(radar.HTTPSummaryTLSVersionParamsFormatJson),
		HTTPProtocol: cloudflare.F([]radar.HTTPSummaryTLSVersionParamsHTTPProtocol{radar.HTTPSummaryTLSVersionParamsHTTPProtocolHTTP, radar.HTTPSummaryTLSVersionParamsHTTPProtocolHTTPS}),
		HTTPVersion:  cloudflare.F([]radar.HTTPSummaryTLSVersionParamsHTTPVersion{radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv1, radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv2, radar.HTTPSummaryTLSVersionParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]radar.HTTPSummaryTLSVersionParamsIPVersion{radar.HTTPSummaryTLSVersionParamsIPVersionIPv4, radar.HTTPSummaryTLSVersionParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		OS:           cloudflare.F([]radar.HTTPSummaryTLSVersionParamsOS{radar.HTTPSummaryTLSVersionParamsOSWindows, radar.HTTPSummaryTLSVersionParamsOSMacosx, radar.HTTPSummaryTLSVersionParamsOSIos}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
