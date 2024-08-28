// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/firewall"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestAccessRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.New(
		context.TODO(),
		"account_or_zone",
		map[string]interface{}{},
		firewall.AccessRuleNewParams{
			Configuration: cloudflare.F[firewall.AccessRuleNewParamsConfigurationUnion](firewall.AccessRuleIPConfigurationParam{
				Target: cloudflare.F(firewall.AccessRuleIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			}),
			Mode:  cloudflare.F(firewall.AccessRuleNewParamsModeBlock),
			Notes: cloudflare.F("This rule is enabled because of an event that occurred on date X."),
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

func TestAccessRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.List(
		context.TODO(),
		"account_or_zone",
		map[string]interface{}{},
		firewall.AccessRuleListParams{
			Direction: cloudflare.F(firewall.AccessRuleListParamsDirectionAsc),
			EgsPagination: cloudflare.F(firewall.AccessRuleListParamsEgsPagination{
				Json: cloudflare.F(firewall.AccessRuleListParamsEgsPaginationJson{
					Page:    cloudflare.F(1.000000),
					PerPage: cloudflare.F(1.000000),
				}),
			}),
			Filters: cloudflare.F(firewall.AccessRuleListParamsFilters{
				ConfigurationTarget: cloudflare.F(firewall.AccessRuleListParamsFiltersConfigurationTargetIP),
				ConfigurationValue:  cloudflare.F("198.51.100.4"),
				Match:               cloudflare.F(firewall.AccessRuleListParamsFiltersMatchAny),
				Mode:                cloudflare.F(firewall.AccessRuleListParamsFiltersModeBlock),
				Notes:               cloudflare.F("my note"),
			}),
			Order:   cloudflare.F(firewall.AccessRuleListParamsOrderConfigurationTarget),
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(20.000000),
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

func TestAccessRuleDelete(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Delete(
		context.TODO(),
		"account_or_zone",
		map[string]interface{}{},
		map[string]interface{}{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessRuleEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Edit(
		context.TODO(),
		"account_or_zone",
		map[string]interface{}{},
		map[string]interface{}{},
		firewall.AccessRuleEditParams{
			Configuration: cloudflare.F[firewall.AccessRuleEditParamsConfigurationUnion](firewall.AccessRuleIPConfigurationParam{
				Target: cloudflare.F(firewall.AccessRuleIPConfigurationTargetIP),
				Value:  cloudflare.F("198.51.100.4"),
			}),
			Mode:  cloudflare.F(firewall.AccessRuleEditParamsModeBlock),
			Notes: cloudflare.F("This rule is enabled because of an event that occurred on date X."),
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

func TestAccessRuleGet(t *testing.T) {
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
	_, err := client.Firewall.AccessRules.Get(
		context.TODO(),
		"account_or_zone",
		map[string]interface{}{},
		map[string]interface{}{},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
