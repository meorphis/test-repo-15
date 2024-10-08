// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/event_notifications"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestR2ConfigurationQueueUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.EventNotifications.R2.Configuration.Queues.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		event_notifications.R2ConfigurationQueueUpdateParams{
			Rules: cloudflare.F([]event_notifications.R2ConfigurationQueueUpdateParamsRule{{
				Actions: cloudflare.F([]event_notifications.R2ConfigurationQueueUpdateParamsRulesAction{event_notifications.R2ConfigurationQueueUpdateParamsRulesActionPutObject, event_notifications.R2ConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  cloudflare.F("img/"),
				Suffix:  cloudflare.F(".jpeg"),
			}, {
				Actions: cloudflare.F([]event_notifications.R2ConfigurationQueueUpdateParamsRulesAction{event_notifications.R2ConfigurationQueueUpdateParamsRulesActionPutObject, event_notifications.R2ConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  cloudflare.F("img/"),
				Suffix:  cloudflare.F(".jpeg"),
			}, {
				Actions: cloudflare.F([]event_notifications.R2ConfigurationQueueUpdateParamsRulesAction{event_notifications.R2ConfigurationQueueUpdateParamsRulesActionPutObject, event_notifications.R2ConfigurationQueueUpdateParamsRulesActionCopyObject}),
				Prefix:  cloudflare.F("img/"),
				Suffix:  cloudflare.F(".jpeg"),
			}}),
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

func TestR2ConfigurationQueueDelete(t *testing.T) {
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
	_, err := client.EventNotifications.R2.Configuration.Queues.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
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
