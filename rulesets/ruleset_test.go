// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/rulesets"
)

func TestRulesetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.New(
		context.TODO(),
		"account_or_zone",
		"abf9b32d38c5f572afde3336ec0ce302",
		rulesets.RulesetNewParams{
			Kind:  cloudflare.F(rulesets.KindManaged),
			Name:  cloudflare.F("My ruleset"),
			Phase: cloudflare.F(rulesets.PhaseDDoSL4),
			Rules: cloudflare.F([]rulesets.RulesetNewParamsRuleUnion{rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}, rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}, rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
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

func TestRulesetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rulesets.Update(
		context.TODO(),
		"account_or_zone",
		"abf9b32d38c5f572afde3336ec0ce302",
		"2f2feab2026849078ba485f918791bdc",
		rulesets.RulesetUpdateParams{
			Rules: cloudflare.F([]rulesets.RulesetUpdateParamsRuleUnion{rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}, rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}, rulesets.BlockRuleParam{
				ID:     cloudflare.F("3a03d665bac047339bb530ecb439a90d"),
				Action: cloudflare.F(rulesets.BlockRuleActionBlock),
				ActionParameters: cloudflare.F(rulesets.BlockRuleActionParametersParam{
					Response: cloudflare.F(rulesets.BlockRuleActionParametersResponseParam{
						Content:     cloudflare.F("{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}"),
						ContentType: cloudflare.F("application/json"),
						StatusCode:  cloudflare.F(int64(400)),
					}),
				}),
				Description: cloudflare.F("Block when the IP address is not 1.1.1.1"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("ip.src ne 1.1.1.1"),
				Logging: cloudflare.F(rulesets.LoggingParam{
					Enabled: cloudflare.F(true),
				}),
				Ref: cloudflare.F("my_ref"),
			}}),
			Description: cloudflare.F("My ruleset to execute managed rulesets"),
			Kind:        cloudflare.F(rulesets.KindManaged),
			Name:        cloudflare.F("My ruleset"),
			Phase:       cloudflare.F(rulesets.PhaseDDoSL4),
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

func TestRulesetList(t *testing.T) {
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
	_, err := client.Rulesets.List(
		context.TODO(),
		"account_or_zone",
		"abf9b32d38c5f572afde3336ec0ce302",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRulesetDelete(t *testing.T) {
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
	err := client.Rulesets.Delete(
		context.TODO(),
		"account_or_zone",
		"abf9b32d38c5f572afde3336ec0ce302",
		"2f2feab2026849078ba485f918791bdc",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRulesetGet(t *testing.T) {
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
	_, err := client.Rulesets.Get(
		context.TODO(),
		"account_or_zone",
		"abf9b32d38c5f572afde3336ec0ce302",
		"2f2feab2026849078ba485f918791bdc",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
