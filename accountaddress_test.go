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

func TestAccountAddressUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Addresses.Update(
		context.TODO(),
		"D4g3h5tBuVYK9",
		meorphistest40.AccountAddressUpdateParams{
			CountryCode:     meorphistest40.F("US"),
			FirstName:       meorphistest40.F("Alice"),
			LastName:        meorphistest40.F("Baker"),
			Locality:        meorphistest40.F("San Francisco"),
			PostalCode:      meorphistest40.F("94105"),
			StreetAddress1:  meorphistest40.F("535 Mission St, Ste 1401"),
			XPublishableKey: meorphistest40.F("string"),
			Company:         meorphistest40.F("ACME Corporation"),
			Email:           meorphistest40.F("alice@example.com"),
			IsDefault:       meorphistest40.F(true),
			Phone:           meorphistest40.F("+14155550199"),
			Region:          meorphistest40.F("CA"),
			StreetAddress2:  meorphistest40.F("c/o Shipping Department"),
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

func TestAccountAddressDelete(t *testing.T) {
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
	err := client.Accounts.Addresses.Delete(
		context.TODO(),
		"D4g3h5tBuVYK9",
		meorphistest40.AccountAddressDeleteParams{
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

func TestAccountAddressAccountAddressNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Addresses.AccountAddressNew(context.TODO(), meorphistest40.AccountAddressAccountAddressNewParams{
		CountryCode:     meorphistest40.F("US"),
		FirstName:       meorphistest40.F("Alice"),
		LastName:        meorphistest40.F("Baker"),
		Locality:        meorphistest40.F("San Francisco"),
		PostalCode:      meorphistest40.F("94105"),
		StreetAddress1:  meorphistest40.F("535 Mission St, Ste 1401"),
		XPublishableKey: meorphistest40.F("string"),
		Company:         meorphistest40.F("ACME Corporation"),
		Email:           meorphistest40.F("alice@example.com"),
		IsDefault:       meorphistest40.F(true),
		Phone:           meorphistest40.F("+14155550199"),
		Region:          meorphistest40.F("CA"),
		StreetAddress2:  meorphistest40.F("c/o Shipping Department"),
	})
	if err != nil {
		var apierr *meorphistest40.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
