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

func TestGuestPaymentGuestPaymentsInitializeWithOptionalParams(t *testing.T) {
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
	_, err := client.Guests.Payments.GuestPaymentsInitialize(context.TODO(), meorphistest40.GuestPaymentGuestPaymentsInitializeParams{
		Cart: meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCart{
			Amounts: meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartAmounts{
				Total:    meorphistest40.F(int64(10000)),
				Currency: meorphistest40.F("USD"),
				Tax:      meorphistest40.F(int64(1000)),
			}),
			OrderReference:   meorphistest40.F("order_100"),
			OrderDescription: meorphistest40.F("Order #1234567890"),
			DisplayID:        meorphistest40.F("215614191"),
			Shipments: meorphistest40.F([]meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartShipment{{
				Address: meorphistest40.F[meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressUnion](meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceID{
					Tag: meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  meorphistest40.F("D4g3h5tBuVYK9"),
				}),
				Cost: meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartShipmentsCost{
					Total:    meorphistest40.F(int64(10000)),
					Currency: meorphistest40.F("USD"),
					Tax:      meorphistest40.F(int64(900)),
				}),
				Carrier: meorphistest40.F("FedEx"),
			}}),
			Discounts: meorphistest40.F([]meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartDiscount{{
				Amounts: meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartDiscountsAmounts{
					Total:    meorphistest40.F(int64(10000)),
					Currency: meorphistest40.F("USD"),
					Tax:      meorphistest40.F(int64(900)),
				}),
				Code:       meorphistest40.F("SUMMER10DISCOUNT"),
				DetailsURL: meorphistest40.F("https://www.example.com/SUMMER-SALE"),
			}}),
			Items: meorphistest40.F([]meorphistest40.GuestPaymentGuestPaymentsInitializeParamsCartItem{{
				Name:        meorphistest40.F("Bolt Swag Bag"),
				Reference:   meorphistest40.F("item_100"),
				Description: meorphistest40.F("Large tote with Bolt logo."),
				TotalAmount: meorphistest40.F(int64(9000)),
				UnitPrice:   meorphistest40.F(int64(1000)),
				Quantity:    meorphistest40.F(int64(9)),
				ImageURL:    meorphistest40.F("https://www.example.com/products/123456/images/1.png"),
			}}),
		}),
		PaymentMethod: meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsPaymentMethod{
			Tag:     meorphistest40.F(meorphistest40.GuestPaymentGuestPaymentsInitializeParamsPaymentMethodTagPaypal),
			Success: meorphistest40.F("https://example.com"),
			Cancel:  meorphistest40.F("https://example.com"),
		}),
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
