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

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.New(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.OrganizationNewParams{
			AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
			Name:                     cloudflare.F("Widget Corps Internal Applications"),
			AllowAuthenticateViaWARP: cloudflare.F(true),
			AutoRedirectToIdentity:   cloudflare.F(true),
			IsUIReadOnly:             cloudflare.F(true),
			LoginDesign: cloudflare.F(zero_trust.LoginDesignParam{
				BackgroundColor: cloudflare.F("#c5ed1b"),
				FooterText:      cloudflare.F("This is an example description."),
				HeaderText:      cloudflare.F("This is an example description."),
				LogoPath:        cloudflare.F("https://example.com/logo.png"),
				TextColor:       cloudflare.F("#c5ed1b"),
			}),
			SessionDuration:                cloudflare.F("24h"),
			UIReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cloudflare.F("720h"),
			WARPAuthSessionDuration:        cloudflare.F("24h"),
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

func TestOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.Update(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.OrganizationUpdateParams{
			AllowAuthenticateViaWARP: cloudflare.F(true),
			AuthDomain:               cloudflare.F("test.cloudflareaccess.com"),
			AutoRedirectToIdentity:   cloudflare.F(true),
			CustomPages: cloudflare.F(zero_trust.OrganizationUpdateParamsCustomPages{
				Forbidden:      cloudflare.F("699d98642c564d2e855e9661899b7252"),
				IdentityDenied: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			}),
			IsUIReadOnly: cloudflare.F(true),
			LoginDesign: cloudflare.F(zero_trust.LoginDesignParam{
				BackgroundColor: cloudflare.F("#c5ed1b"),
				FooterText:      cloudflare.F("This is an example description."),
				HeaderText:      cloudflare.F("This is an example description."),
				LogoPath:        cloudflare.F("https://example.com/logo.png"),
				TextColor:       cloudflare.F("#c5ed1b"),
			}),
			Name:                           cloudflare.F("Widget Corps Internal Applications"),
			SessionDuration:                cloudflare.F("24h"),
			UIReadOnlyToggleReason:         cloudflare.F("Temporarily turn off the UI read only lock to make a change via the UI"),
			UserSeatExpirationInactiveTime: cloudflare.F("720h"),
			WARPAuthSessionDuration:        cloudflare.F("24h"),
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

func TestOrganizationList(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.List(
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

func TestOrganizationRevokeUsers(t *testing.T) {
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
	_, err := client.ZeroTrust.Organizations.RevokeUsers(
		context.TODO(),
		"account_or_zone",
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.OrganizationRevokeUsersParams{
			Email: cloudflare.F("test@example.com"),
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
