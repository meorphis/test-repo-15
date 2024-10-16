// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/cloud_connector"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestRuleUpdate(t *testing.T) {
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
	_, err := client.CloudConnector.Rules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloud_connector.RuleUpdateParams{
			Body: []cloud_connector.RuleUpdateParamsBody{{
				ID:          cloudflare.F("95c365e17e1b46599cd99e5b231fac4e"),
				Description: cloudflare.F("Rule description"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("http.cookie eq \"a=b\""),
				Parameters: cloudflare.F(cloud_connector.RuleUpdateParamsBodyParameters{
					Host: cloudflare.F("examplebucket.s3.eu-north-1.amazonaws.com"),
				}),
				Provider: cloudflare.F(cloud_connector.RuleUpdateParamsBodyProviderAwsS3),
			}, {
				ID:          cloudflare.F("95c365e17e1b46599cd99e5b231fac4e"),
				Description: cloudflare.F("Rule description"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("http.cookie eq \"a=b\""),
				Parameters: cloudflare.F(cloud_connector.RuleUpdateParamsBodyParameters{
					Host: cloudflare.F("examplebucket.s3.eu-north-1.amazonaws.com"),
				}),
				Provider: cloudflare.F(cloud_connector.RuleUpdateParamsBodyProviderAwsS3),
			}, {
				ID:          cloudflare.F("95c365e17e1b46599cd99e5b231fac4e"),
				Description: cloudflare.F("Rule description"),
				Enabled:     cloudflare.F(true),
				Expression:  cloudflare.F("http.cookie eq \"a=b\""),
				Parameters: cloudflare.F(cloud_connector.RuleUpdateParamsBodyParameters{
					Host: cloudflare.F("examplebucket.s3.eu-north-1.amazonaws.com"),
				}),
				Provider: cloudflare.F(cloud_connector.RuleUpdateParamsBodyProviderAwsS3),
			}},
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

func TestRuleList(t *testing.T) {
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
	_, err := client.CloudConnector.Rules.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}