// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/workers"
)

func TestScriptSettingEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Scripts.Settings.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		workers.ScriptSettingEditParams{
			ScriptSetting: workers.ScriptSettingParam{
				Logpush: cloudflare.F(false),
				TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}, {
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}, {
					Service:     cloudflare.F("my-log-consumer"),
					Environment: cloudflare.F("production"),
					Namespace:   cloudflare.F("my-namespace"),
				}}),
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

func TestScriptSettingGet(t *testing.T) {
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
	_, err := client.Workers.Scripts.Settings.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
