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

func TestAccountPaymentMethodDelete(t *testing.T) {
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
	err := client.Accounts.PaymentMethods.Delete(
		context.TODO(),
		"D4g3h5tBuVYK9",
		meorphistest40.AccountPaymentMethodDeleteParams{
			XPublishableKey: meorphistest40.F("string"),
		},
	)
	if err != nil {
		var apierr *meorphistest40.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPaymentMethodAccountAddPaymentMethodWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.PaymentMethods.AccountAddPaymentMethod(context.TODO(), meorphistest40.AccountPaymentMethodAccountAddPaymentMethodParams{
		Token: meorphistest40.F("a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0"),
		Tag:   meorphistest40.F(meorphistest40.AccountPaymentMethodAccountAddPaymentMethodParamsTagCreditCard),
		BillingAddress: meorphistest40.F[meorphistest40.AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressUnion](meorphistest40.AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceID{
			Tag: meorphistest40.F(meorphistest40.AccountPaymentMethodAccountAddPaymentMethodParamsBillingAddressAddressReferenceIDTagID),
			ID:  meorphistest40.F("D4g3h5tBuVYK9"),
		}),
		Bin:             meorphistest40.F("411111"),
		Expiration:      meorphistest40.F("2025-03"),
		Last4:           meorphistest40.F("1004"),
		Network:         meorphistest40.F(meorphistest40.AccountPaymentMethodAccountAddPaymentMethodParamsNetworkVisa),
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
