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

func TestIdentityProviderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.New(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.IdentityProviderNewParams{
			IdentityProvider: zero_trust.AzureADParam{
				Config: cloudflare.F(zero_trust.AzureADConfigParam{
					Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
					ClientID:                 cloudflare.F("<your client id>"),
					ClientSecret:             cloudflare.F("<your client secret>"),
					ConditionalAccessEnabled: cloudflare.F(true),
					DirectoryID:              cloudflare.F("<your azure directory uuid>"),
					EmailClaimName:           cloudflare.F("custom_claim_name"),
					Prompt:                   cloudflare.F(zero_trust.AzureADConfigPromptLogin),
					SupportGroups:            cloudflare.F(true),
				}),
				Name: cloudflare.F("Widget Corps IDP"),
				Type: cloudflare.F(zero_trust.IdentityProviderTypeOnetimepin),
				ID:   cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				SCIMConfig: cloudflare.F(zero_trust.IdentityProviderSCIMConfigParam{
					Enabled:                cloudflare.F(true),
					GroupMemberDeprovision: cloudflare.F(true),
					SeatDeprovision:        cloudflare.F(true),
					Secret:                 cloudflare.F("secret"),
					UserDeprovision:        cloudflare.F(true),
				}),
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

func TestIdentityProviderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Update(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.IdentityProviderUpdateParams{
			IdentityProvider: zero_trust.AzureADParam{
				Config: cloudflare.F(zero_trust.AzureADConfigParam{
					Claims:                   cloudflare.F([]string{"email_verified", "preferred_username", "custom_claim_name"}),
					ClientID:                 cloudflare.F("<your client id>"),
					ClientSecret:             cloudflare.F("<your client secret>"),
					ConditionalAccessEnabled: cloudflare.F(true),
					DirectoryID:              cloudflare.F("<your azure directory uuid>"),
					EmailClaimName:           cloudflare.F("custom_claim_name"),
					Prompt:                   cloudflare.F(zero_trust.AzureADConfigPromptLogin),
					SupportGroups:            cloudflare.F(true),
				}),
				Name: cloudflare.F("Widget Corps IDP"),
				Type: cloudflare.F(zero_trust.IdentityProviderTypeOnetimepin),
				ID:   cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
				SCIMConfig: cloudflare.F(zero_trust.IdentityProviderSCIMConfigParam{
					Enabled:                cloudflare.F(true),
					GroupMemberDeprovision: cloudflare.F(true),
					SeatDeprovision:        cloudflare.F(true),
					Secret:                 cloudflare.F("secret"),
					UserDeprovision:        cloudflare.F(true),
				}),
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

func TestIdentityProviderList(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.List(
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

func TestIdentityProviderDelete(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Delete(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
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

func TestIdentityProviderGet(t *testing.T) {
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
	_, err := client.ZeroTrust.IdentityProviders.Get(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
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
