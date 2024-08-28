// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/spectrum"
)

func TestAnalyticsEventBytimeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Spectrum.Analytics.Events.Bytimes.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		spectrum.AnalyticsEventBytimeGetParams{
			Dimensions: cloudflare.F([]spectrum.Dimension{spectrum.DimensionEvent, spectrum.DimensionAppID}),
			Filters:    cloudflare.F("event==disconnect%20AND%20coloName!=SFO"),
			Metrics:    cloudflare.F([]spectrum.AnalyticsEventBytimeGetParamsMetric{spectrum.AnalyticsEventBytimeGetParamsMetricCount, spectrum.AnalyticsEventBytimeGetParamsMetricBytesIngress}),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F([]interface{}{"+count", "-bytesIngress"}),
			TimeDelta:  cloudflare.F(spectrum.AnalyticsEventBytimeGetParamsTimeDeltaYear),
			Until:      cloudflare.F(time.Now()),
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
