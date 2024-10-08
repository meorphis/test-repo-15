// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/dns"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestSettingEditWithOptionalParams(t *testing.T) {
	t.Skip("HTTP 422 from prism")
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
	_, err := client.DNS.Settings.Edit(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.SettingEditParams{
			ZoneDefaults: cloudflare.F(dns.DNSSettingParam{
				FoundationDNS: cloudflare.F(false),
				MultiProvider: cloudflare.F(false),
				Nameservers: cloudflare.F(dns.NameserverParam{
					Type: cloudflare.F(dns.NameserverTypeCloudflareStandard),
				}),
				NSTTL:              cloudflare.F(86400.000000),
				SecondaryOverrides: cloudflare.F(false),
				SOA: cloudflare.F(dns.DNSSettingSOAParam{
					Expire:  cloudflare.F(604800.000000),
					MinTTL:  cloudflare.F(1800.000000),
					MNAME:   cloudflare.F("kristina.ns.cloudflare.com"),
					Refresh: cloudflare.F(10000.000000),
					Retry:   cloudflare.F(2400.000000),
					RNAME:   cloudflare.F("admin.example.com"),
					TTL:     cloudflare.F(3600.000000),
				}),
				ZoneMode: cloudflare.F(dns.DNSSettingZoneModeStandard),
			}),
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

func TestSettingGet(t *testing.T) {
	t.Skip("HTTP 422 from prism")
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
	_, err := client.DNS.Settings.Get(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
