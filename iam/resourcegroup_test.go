// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/iam"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestResourceGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.New(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		iam.ResourceGroupNewParams{
			Scope: cloudflare.F(iam.ResourceGroupNewParamsScope{
				Key: cloudflare.F("com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
				Objects: cloudflare.F([]iam.ResourceGroupNewParamsScopeObject{{
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}}),
			}),
			Meta: cloudflare.F[any](map[string]interface{}{
				"editable": "false",
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

func TestResourceGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Update(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
		iam.ResourceGroupUpdateParams{
			Scope: cloudflare.F(iam.ResourceGroupUpdateParamsScope{
				Key: cloudflare.F("com.cloudflare.api.account.eb78d65290b24279ba6f44721b3ea3c4"),
				Objects: cloudflare.F([]iam.ResourceGroupUpdateParamsScopeObject{{
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}, {
					Key: cloudflare.F("com.cloudflare.api.account.zone.23f8d65290b24279ba6f44721b3eaad5"),
				}}),
			}),
			Meta: cloudflare.F[any](map[string]interface{}{
				"editable": "false",
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

func TestResourceGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.List(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		iam.ResourceGroupListParams{
			ID:      cloudflare.F("6d7f2f5f5b1d4a0e9081fdc98d432fd1"),
			Name:    cloudflare.F("NameOfTheResourceGroup"),
			Page:    cloudflare.F(1.000000),
			PerPage: cloudflare.F(5.000000),
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

func TestResourceGroupDelete(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Delete(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResourceGroupGet(t *testing.T) {
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
	_, err := client.IAM.ResourceGroups.Get(
		context.TODO(),
		"eb78d65290b24279ba6f44721b3ea3c4",
		"6d7f2f5f5b1d4a0e9081fdc98d432fd1",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
