// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestMerchantCallbackList(t *testing.T) {
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
	_, err := client.Merchants.Callbacks.List(context.TODO(), meorphistest40.MerchantCallbackListParams{
		XPublishableKey: meorphistest40.F("string"),
	})
	if err != nil {
		var apierr *meorphistest40.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMerchantCallbackMerchantCallbacksUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Merchants.Callbacks.MerchantCallbacksUpdate(context.TODO(), meorphistest40.MerchantCallbackMerchantCallbacksUpdateParams{
		XPublishableKey:               meorphistest40.F("string"),
		AccountPage:                   meorphistest40.F("https://www.example.com/account"),
		BaseDomain:                    meorphistest40.F("https://www.example.com/"),
		ConfirmationRedirect:          meorphistest40.F("https://www.example.com/bolt/redirect"),
		CreateOrder:                   meorphistest40.F("https://www.example.com/bolt/order"),
		Debug:                         meorphistest40.F("https://www.example.com/bolt/debug"),
		GetAccount:                    meorphistest40.F("https://www.example.com/bolt/account"),
		MobileAppDomain:               meorphistest40.F("https://m.example.com/"),
		OAuthLogout:                   meorphistest40.F("https://www.example.com/bolt/logout"),
		OAuthRedirect:                 meorphistest40.F("https://www.example.com/bolt/oauth"),
		PrivacyPolicy:                 meorphistest40.F("https://www.example.com/privacy-policy"),
		ProductInfo:                   meorphistest40.F("https://www.example.com/bolt/product"),
		RemoteAPI:                     meorphistest40.F("https://www.example.com/bolt/remote-api"),
		Shipping:                      meorphistest40.F("https://www.example.com/bolt/shipping"),
		SupportPage:                   meorphistest40.F("https://www.example.com/help"),
		Tax:                           meorphistest40.F("https://www.example.com/bolt/tax"),
		TermsOfService:                meorphistest40.F("https://www.example.com/terms-of-service"),
		UniversalMerchantAPI:          meorphistest40.F("https://www.example.com/bolt/merchant-api"),
		UpdateCart:                    meorphistest40.F("https://www.example.com/bolt/cart"),
		ValidateAdditionalAccountData: meorphistest40.F("https://www.example.com/bolt/validate-account"),
	})
	if err != nil {
		var apierr *meorphistest40.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
