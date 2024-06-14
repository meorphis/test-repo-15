// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestTestingAccountTestingAccountNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := meorphistest40.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Testings.Accounts.TestingAccountNew(context.TODO(), meorphistest40.TestingAccountTestingAccountNewParams{
		DeactivateAt: meorphistest40.F(time.Now()),
		EmailState:   meorphistest40.F(meorphistest40.TestingAccountTestingAccountNewParamsEmailStateUnverified),
		PhoneState:   meorphistest40.F(meorphistest40.TestingAccountTestingAccountNewParamsPhoneStateVerified),
		HasAddress:   meorphistest40.F(true),
		IsMigrated:   meorphistest40.F(true),
	})
	if err != nil {
		var apierr *meorphistest40.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
