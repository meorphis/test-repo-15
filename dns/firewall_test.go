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

func TestFirewallNewWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Firewall.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.FirewallNewParams{
			Name:        cloudflare.F("My Awesome DNS Firewall cluster"),
			UpstreamIPs: cloudflare.F([]dns.UpstreamIPsParam{"192.0.2.1", "198.51.100.1", "string"}),
			AttackMitigation: cloudflare.F(dns.AttackMitigationParam{
				Enabled:                   cloudflare.F(true),
				OnlyWhenUpstreamUnhealthy: cloudflare.F(false),
			}),
			DeprecateAnyRequests: cloudflare.F(true),
			ECSFallback:          cloudflare.F(false),
			MaximumCacheTTL:      cloudflare.F(900.000000),
			MinimumCacheTTL:      cloudflare.F(60.000000),
			NegativeCacheTTL:     cloudflare.F(900.000000),
			Ratelimit:            cloudflare.F(600.000000),
			Retries:              cloudflare.F(2.000000),
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

func TestFirewallListWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Firewall.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.FirewallListParams{
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(1.000000),
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

func TestFirewallDelete(t *testing.T) {
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
	_, err := client.DNS.Firewall.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestFirewallEditWithOptionalParams(t *testing.T) {
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
	_, err := client.DNS.Firewall.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		dns.FirewallEditParams{
			Firewall: dns.FirewallParam{
				ID:                   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
				DeprecateAnyRequests: cloudflare.F(true),
				DNSFirewallIPs:       cloudflare.F([]dns.FirewallIPsParam{"203.0.113.1", "203.0.113.254", "string", "string"}),
				ECSFallback:          cloudflare.F(false),
				MaximumCacheTTL:      cloudflare.F(900.000000),
				MinimumCacheTTL:      cloudflare.F(60.000000),
				Name:                 cloudflare.F("My Awesome DNS Firewall cluster"),
				UpstreamIPs:          cloudflare.F([]dns.UpstreamIPsParam{"192.0.2.1", "198.51.100.1", "string"}),
				AttackMitigation: cloudflare.F(dns.AttackMitigationParam{
					Enabled:                   cloudflare.F(true),
					OnlyWhenUpstreamUnhealthy: cloudflare.F(false),
				}),
				NegativeCacheTTL: cloudflare.F(900.000000),
				Ratelimit:        cloudflare.F(600.000000),
				Retries:          cloudflare.F(2.000000),
			},
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

func TestFirewallGet(t *testing.T) {
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
	_, err := client.DNS.Firewall.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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
