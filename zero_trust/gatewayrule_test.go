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

func TestGatewayRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		zero_trust.GatewayRuleNewParams{
			Action:        cloudflare.F(zero_trust.GatewayRuleNewParamsActionOn),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block bad websites based on their host name."),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]zero_trust.GatewayFilter{zero_trust.GatewayFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(zero_trust.RuleSettingParam{
				AddHeaders: cloudflare.F[any](map[string]interface{}{
					"My-Next-Header": map[string]interface{}{
						"0": "foo",
						"1": "bar",
					},
					"X-Custom-Header-Name": map[string]interface{}{
						"0": "somecustomvalue",
					},
				}),
				AllowChildBypass: cloudflare.F(false),
				AuditSSH: cloudflare.F(zero_trust.RuleSettingAuditSSHParam{
					CommandLogging: cloudflare.F(false),
				}),
				BISOAdminControls: cloudflare.F(zero_trust.RuleSettingBISOAdminControlsParam{
					DCP: cloudflare.F(false),
					DD:  cloudflare.F(false),
					DK:  cloudflare.F(false),
					DP:  cloudflare.F(false),
					DU:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(zero_trust.RuleSettingCheckSessionParam{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				DNSResolvers: cloudflare.F(zero_trust.RuleSettingDNSResolversParam{
					IPV4: cloudflare.F([]zero_trust.DNSResolverSettingsV4Param{{
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					IPV6: cloudflare.F([]zero_trust.DNSResolverSettingsV6Param{{
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cloudflare.F(zero_trust.RuleSettingEgressParam{
					IPV4:         cloudflare.F("192.0.2.2"),
					IPV4Fallback: cloudflare.F("192.0.2.3"),
					IPV6:         cloudflare.F("2001:DB8::/64"),
				}),
				IgnoreCNAMECategoryMatches:      cloudflare.F(true),
				InsecureDisableDNSSECValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				IPIndicatorFeeds:                cloudflare.F(true),
				L4override: cloudflare.F(zero_trust.RuleSettingL4overrideParam{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				NotificationSettings: cloudflare.F(zero_trust.RuleSettingNotificationSettingsParam{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("msg"),
					SupportURL: cloudflare.F("support_url"),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(zero_trust.RuleSettingPayloadLogParam{
					Enabled: cloudflare.F(true),
				}),
				ResolveDNSThroughCloudflare: cloudflare.F(true),
				UntrustedCERT: cloudflare.F(zero_trust.RuleSettingUntrustedCERTParam{
					Action: cloudflare.F(zero_trust.RuleSettingUntrustedCERTActionPassThrough),
				}),
			}),
			Schedule: cloudflare.F(zero_trust.ScheduleParam{
				Fri:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Mon:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sat:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sun:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Thu:      cloudflare.F("08:00-12:30,13:30-17:00"),
				TimeZone: cloudflare.F("America/New York"),
				Tue:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Wed:      cloudflare.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cloudflare.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
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

func TestGatewayRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.GatewayRuleUpdateParams{
			Action:        cloudflare.F(zero_trust.GatewayRuleUpdateParamsActionOn),
			Name:          cloudflare.F("block bad websites"),
			Description:   cloudflare.F("Block bad websites based on their host name."),
			DevicePosture: cloudflare.F("any(device_posture.checks.passed[*] in {\"1308749e-fcfb-4ebc-b051-fe022b632644\"})"),
			Enabled:       cloudflare.F(true),
			Filters:       cloudflare.F([]zero_trust.GatewayFilter{zero_trust.GatewayFilterHTTP}),
			Identity:      cloudflare.F("any(identity.groups.name[*] in {\"finance\"})"),
			Precedence:    cloudflare.F(int64(0)),
			RuleSettings: cloudflare.F(zero_trust.RuleSettingParam{
				AddHeaders: cloudflare.F[any](map[string]interface{}{
					"My-Next-Header": map[string]interface{}{
						"0": "foo",
						"1": "bar",
					},
					"X-Custom-Header-Name": map[string]interface{}{
						"0": "somecustomvalue",
					},
				}),
				AllowChildBypass: cloudflare.F(false),
				AuditSSH: cloudflare.F(zero_trust.RuleSettingAuditSSHParam{
					CommandLogging: cloudflare.F(false),
				}),
				BISOAdminControls: cloudflare.F(zero_trust.RuleSettingBISOAdminControlsParam{
					DCP: cloudflare.F(false),
					DD:  cloudflare.F(false),
					DK:  cloudflare.F(false),
					DP:  cloudflare.F(false),
					DU:  cloudflare.F(false),
				}),
				BlockPageEnabled: cloudflare.F(true),
				BlockReason:      cloudflare.F("This website is a security risk"),
				BypassParentRule: cloudflare.F(false),
				CheckSession: cloudflare.F(zero_trust.RuleSettingCheckSessionParam{
					Duration: cloudflare.F("300s"),
					Enforce:  cloudflare.F(true),
				}),
				DNSResolvers: cloudflare.F(zero_trust.RuleSettingDNSResolversParam{
					IPV4: cloudflare.F([]zero_trust.DNSResolverSettingsV4Param{{
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2.2.2.2"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
					IPV6: cloudflare.F([]zero_trust.DNSResolverSettingsV6Param{{
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}, {
						IP:                         cloudflare.F("2001:DB8::"),
						Port:                       cloudflare.F(int64(5053)),
						RouteThroughPrivateNetwork: cloudflare.F(true),
						VnetID:                     cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					}}),
				}),
				Egress: cloudflare.F(zero_trust.RuleSettingEgressParam{
					IPV4:         cloudflare.F("192.0.2.2"),
					IPV4Fallback: cloudflare.F("192.0.2.3"),
					IPV6:         cloudflare.F("2001:DB8::/64"),
				}),
				IgnoreCNAMECategoryMatches:      cloudflare.F(true),
				InsecureDisableDNSSECValidation: cloudflare.F(false),
				IPCategories:                    cloudflare.F(true),
				IPIndicatorFeeds:                cloudflare.F(true),
				L4override: cloudflare.F(zero_trust.RuleSettingL4overrideParam{
					IP:   cloudflare.F("1.1.1.1"),
					Port: cloudflare.F(int64(0)),
				}),
				NotificationSettings: cloudflare.F(zero_trust.RuleSettingNotificationSettingsParam{
					Enabled:    cloudflare.F(true),
					Msg:        cloudflare.F("msg"),
					SupportURL: cloudflare.F("support_url"),
				}),
				OverrideHost: cloudflare.F("example.com"),
				OverrideIPs:  cloudflare.F([]string{"1.1.1.1", "2.2.2.2"}),
				PayloadLog: cloudflare.F(zero_trust.RuleSettingPayloadLogParam{
					Enabled: cloudflare.F(true),
				}),
				ResolveDNSThroughCloudflare: cloudflare.F(true),
				UntrustedCERT: cloudflare.F(zero_trust.RuleSettingUntrustedCERTParam{
					Action: cloudflare.F(zero_trust.RuleSettingUntrustedCERTActionPassThrough),
				}),
			}),
			Schedule: cloudflare.F(zero_trust.ScheduleParam{
				Fri:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Mon:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sat:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Sun:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Thu:      cloudflare.F("08:00-12:30,13:30-17:00"),
				TimeZone: cloudflare.F("America/New York"),
				Tue:      cloudflare.F("08:00-12:30,13:30-17:00"),
				Wed:      cloudflare.F("08:00-12:30,13:30-17:00"),
			}),
			Traffic: cloudflare.F("http.request.uri matches \".*a/partial/uri.*\" and http.request.host in $01302951-49f9-47c9-a400-0297e60b6a10"),
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

func TestGatewayRuleList(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGatewayRuleGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Gateway.Rules.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}