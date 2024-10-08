// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_certificates_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/custom_certificates"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestPrioritizeUpdate(t *testing.T) {
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
	_, err := client.CustomCertificates.Prioritize.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		custom_certificates.PrioritizeUpdateParams{
			Certificates: cloudflare.F([]custom_certificates.PrioritizeUpdateParamsCertificate{{
				ID:       cloudflare.F("5a7805061c76ada191ed06f989cc3dac"),
				Priority: cloudflare.F(2.000000),
			}, {
				ID:       cloudflare.F("9a7806061c88ada191ed06f989cc3dac"),
				Priority: cloudflare.F(1.000000),
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
