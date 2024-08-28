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

func TestAccessApplicationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.New(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationNewParams{
			Body: zero_trust.AccessApplicationNewParamsBodySelfHostedApplication{
				Domain:                   cloudflare.F("test.example.com/admin"),
				Type:                     cloudflare.F("self_hosted"),
				AllowAuthenticateViaWARP: cloudflare.F(true),
				AllowedIdPs:              cloudflare.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
				AppLauncherVisible:       cloudflare.F(true),
				AutoRedirectToIdentity:   cloudflare.F(true),
				CORSHeaders: cloudflare.F(zero_trust.CORSHeadersParam{
					AllowAllHeaders:  cloudflare.F(true),
					AllowAllMethods:  cloudflare.F(true),
					AllowAllOrigins:  cloudflare.F(true),
					AllowCredentials: cloudflare.F(true),
					AllowedHeaders:   cloudflare.F([]zero_trust.AllowedHeadersParam{"string", "string", "string"}),
					AllowedMethods:   cloudflare.F([]zero_trust.AllowedMethods{zero_trust.AllowedMethodsGet}),
					AllowedOrigins:   cloudflare.F([]zero_trust.AllowedOriginsParam{"https://example.com"}),
					MaxAge:           cloudflare.F(-1.000000),
				}),
				CustomDenyMessage:        cloudflare.F("custom_deny_message"),
				CustomDenyURL:            cloudflare.F("custom_deny_url"),
				CustomNonIdentityDenyURL: cloudflare.F("custom_non_identity_deny_url"),
				CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
				EnableBindingCookie:      cloudflare.F(true),
				HTTPOnlyCookieAttribute:  cloudflare.F(true),
				LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
				Name:                     cloudflare.F("Admin Site"),
				OptionsPreflightBypass:   cloudflare.F(true),
				PathCookieAttribute:      cloudflare.F(true),
				Policies: cloudflare.F([]zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPolicyUnion{zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}, zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}, zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}}),
				SameSiteCookieAttribute: cloudflare.F("strict"),
				SCIMConfig: cloudflare.F(zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfig{
					IdPUID:    cloudflare.F("idp_uid"),
					RemoteURI: cloudflare.F("remote_uri"),
					Authentication: cloudflare.F[zero_trust.AccessApplicationNewParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
						Password: cloudflare.F("password"),
						Scheme:   cloudflare.F(zero_trust.SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic),
						User:     cloudflare.F("user"),
					}),
					DeactivateOnDelete: cloudflare.F(true),
					Enabled:            cloudflare.F(true),
					Mappings: cloudflare.F([]zero_trust.SCIMConfigMappingParam{{
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}, {
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}, {
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}}),
				}),
				SelfHostedDomains:      cloudflare.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
				ServiceAuth401Redirect: cloudflare.F(true),
				SessionDuration:        cloudflare.F("24h"),
				SkipInterstitial:       cloudflare.F(true),
				Tags:                   cloudflare.F([]string{"engineers", "engineers", "engineers"}),
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

func TestAccessApplicationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Update(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessApplicationUpdateParams{
			Body: zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplication{
				Domain:                   cloudflare.F("test.example.com/admin"),
				Type:                     cloudflare.F("self_hosted"),
				AllowAuthenticateViaWARP: cloudflare.F(true),
				AllowedIdPs:              cloudflare.F([]zero_trust.AllowedIdPsParam{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
				AppLauncherVisible:       cloudflare.F(true),
				AutoRedirectToIdentity:   cloudflare.F(true),
				CORSHeaders: cloudflare.F(zero_trust.CORSHeadersParam{
					AllowAllHeaders:  cloudflare.F(true),
					AllowAllMethods:  cloudflare.F(true),
					AllowAllOrigins:  cloudflare.F(true),
					AllowCredentials: cloudflare.F(true),
					AllowedHeaders:   cloudflare.F([]zero_trust.AllowedHeadersParam{"string", "string", "string"}),
					AllowedMethods:   cloudflare.F([]zero_trust.AllowedMethods{zero_trust.AllowedMethodsGet}),
					AllowedOrigins:   cloudflare.F([]zero_trust.AllowedOriginsParam{"https://example.com"}),
					MaxAge:           cloudflare.F(-1.000000),
				}),
				CustomDenyMessage:        cloudflare.F("custom_deny_message"),
				CustomDenyURL:            cloudflare.F("custom_deny_url"),
				CustomNonIdentityDenyURL: cloudflare.F("custom_non_identity_deny_url"),
				CustomPages:              cloudflare.F([]string{"699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252", "699d98642c564d2e855e9661899b7252"}),
				EnableBindingCookie:      cloudflare.F(true),
				HTTPOnlyCookieAttribute:  cloudflare.F(true),
				LogoURL:                  cloudflare.F("https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"),
				Name:                     cloudflare.F("Admin Site"),
				OptionsPreflightBypass:   cloudflare.F(true),
				PathCookieAttribute:      cloudflare.F(true),
				Policies: cloudflare.F([]zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPolicyUnion{zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}, zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}, zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationPoliciesAccessAppPolicyLink{
					ID:         cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
					Precedence: cloudflare.F(int64(0)),
				}}),
				SameSiteCookieAttribute: cloudflare.F("strict"),
				SCIMConfig: cloudflare.F(zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfig{
					IdPUID:    cloudflare.F("idp_uid"),
					RemoteURI: cloudflare.F("remote_uri"),
					Authentication: cloudflare.F[zero_trust.AccessApplicationUpdateParamsBodySelfHostedApplicationSCIMConfigAuthenticationUnion](zero_trust.SCIMConfigAuthenticationHTTPBasicParam{
						Password: cloudflare.F("password"),
						Scheme:   cloudflare.F(zero_trust.SCIMConfigAuthenticationHTTPBasicSchemeHttpbasic),
						User:     cloudflare.F("user"),
					}),
					DeactivateOnDelete: cloudflare.F(true),
					Enabled:            cloudflare.F(true),
					Mappings: cloudflare.F([]zero_trust.SCIMConfigMappingParam{{
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}, {
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}, {
						Schema:  cloudflare.F("urn:ietf:params:scim:schemas:core:2.0:User"),
						Enabled: cloudflare.F(true),
						Filter:  cloudflare.F("title pr or userType eq \"Intern\""),
						Operations: cloudflare.F(zero_trust.SCIMConfigMappingOperationsParam{
							Create: cloudflare.F(true),
							Delete: cloudflare.F(true),
							Update: cloudflare.F(true),
						}),
						TransformJsonata: cloudflare.F("$merge([$, {'userName': $substringBefore($.userName, '@') & '+test@' & $substringAfter($.userName, '@')}])"),
					}}),
				}),
				SelfHostedDomains:      cloudflare.F([]zero_trust.SelfHostedDomainsParam{"test.example.com/admin", "test.anotherexample.com/staff"}),
				ServiceAuth401Redirect: cloudflare.F(true),
				SessionDuration:        cloudflare.F("24h"),
				SkipInterstitial:       cloudflare.F(true),
				Tags:                   cloudflare.F([]string{"engineers", "engineers", "engineers"}),
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

func TestAccessApplicationList(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.List(
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

func TestAccessApplicationDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Delete(
		context.TODO(),
		"account_or_zone",
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

func TestAccessApplicationGet(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.Get(
		context.TODO(),
		"account_or_zone",
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

func TestAccessApplicationRevokeTokens(t *testing.T) {
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
	_, err := client.ZeroTrust.Access.Applications.RevokeTokens(
		context.TODO(),
		"account_or_zone",
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
