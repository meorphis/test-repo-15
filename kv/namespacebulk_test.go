// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/kv"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestNamespaceBulkUpdate(t *testing.T) {
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
	_, err := client.KV.Namespaces.Bulk.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0f2ac74b498b48028cb68387c421e279",
		kv.NamespaceBulkUpdateParams{
			Body: []kv.NamespaceBulkUpdateParamsBody{{
				Base64:        cloudflare.F(true),
				Expiration:    cloudflare.F(1578435000.000000),
				ExpirationTTL: cloudflare.F(300.000000),
				Key:           cloudflare.F("My-Key"),
				Metadata: cloudflare.F(map[string]interface{}{
					"someMetadataKey": "bar",
				}),
				Value: cloudflare.F("Some string"),
			}, {
				Base64:        cloudflare.F(true),
				Expiration:    cloudflare.F(1578435000.000000),
				ExpirationTTL: cloudflare.F(300.000000),
				Key:           cloudflare.F("My-Key"),
				Metadata: cloudflare.F(map[string]interface{}{
					"someMetadataKey": "bar",
				}),
				Value: cloudflare.F("Some string"),
			}, {
				Base64:        cloudflare.F(true),
				Expiration:    cloudflare.F(1578435000.000000),
				ExpirationTTL: cloudflare.F(300.000000),
				Key:           cloudflare.F("My-Key"),
				Metadata: cloudflare.F(map[string]interface{}{
					"someMetadataKey": "bar",
				}),
				Value: cloudflare.F("Some string"),
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

func TestNamespaceBulkDelete(t *testing.T) {
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
	_, err := client.KV.Namespaces.Bulk.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"0f2ac74b498b48028cb68387c421e279",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
