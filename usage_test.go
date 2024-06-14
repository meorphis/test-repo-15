// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestUsage(t *testing.T) {
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
	accountAccountGetResponse, err := client.Accounts.AccountGet(context.TODO(), meorphistest40.AccountAccountGetParams{
		XPublishableKey: meorphistest40.F("string"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", accountAccountGetResponse.Addresses)
}
