// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/intel"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestMiscategorizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.Miscategorizations.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		intel.MiscategorizationNewParams{
			ContentAdds:     cloudflare.F([]float64{82.000000}),
			ContentRemoves:  cloudflare.F([]float64{155.000000}),
			IndicatorType:   cloudflare.F(intel.MiscategorizationNewParamsIndicatorTypeDomain),
			IP:              cloudflare.F[any](map[string]interface{}{}),
			SecurityAdds:    cloudflare.F([]float64{117.000000, 131.000000}),
			SecurityRemoves: cloudflare.F([]float64{83.000000}),
			URL:             cloudflare.F("url"),
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
