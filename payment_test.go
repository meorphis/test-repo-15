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

func TestPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.New(context.TODO(), meorphistest40.PaymentNewParams{
		Cart: meorphistest40.F(meorphistest40.PaymentNewParamsCart{
			Amounts: meorphistest40.F(meorphistest40.PaymentNewParamsCartAmounts{
				Total:    meorphistest40.F(int64(10000)),
				Currency: meorphistest40.F("USD"),
				Tax:      meorphistest40.F(int64(1000)),
			}),
			OrderReference:   meorphistest40.F("order_100"),
			OrderDescription: meorphistest40.F("Order #1234567890"),
			DisplayID:        meorphistest40.F("215614191"),
			Shipments: meorphistest40.F([]meorphistest40.PaymentNewParamsCartShipment{{
				Address: meorphistest40.F[meorphistest40.PaymentNewParamsCartShipmentsAddressUnion](meorphistest40.PaymentNewParamsCartShipmentsAddressAddressReferenceID{
					Tag: meorphistest40.F(meorphistest40.PaymentNewParamsCartShipmentsAddressAddressReferenceIDTagID),
					ID:  meorphistest40.F("D4g3h5tBuVYK9"),
				}),
				Cost: meorphistest40.F(meorphistest40.PaymentNewParamsCartShipmentsCost{
					Total:    meorphistest40.F(int64(10000)),
					Currency: meorphistest40.F("USD"),
					Tax:      meorphistest40.F(int64(900)),
				}),
				Carrier: meorphistest40.F("FedEx"),
			}}),
			Discounts: meorphistest40.F([]meorphistest40.PaymentNewParamsCartDiscount{{
				Amounts: meorphistest40.F(meorphistest40.PaymentNewParamsCartDiscountsAmounts{
					Total:    meorphistest40.F(int64(10000)),
					Currency: meorphistest40.F("USD"),
					Tax:      meorphistest40.F(int64(900)),
				}),
				Code:       meorphistest40.F("SUMMER10DISCOUNT"),
				DetailsURL: meorphistest40.F("https://www.example.com/SUMMER-SALE"),
			}}),
			Items: meorphistest40.F([]meorphistest40.PaymentNewParamsCartItem{{
				Name:        meorphistest40.F("Bolt Swag Bag"),
				Reference:   meorphistest40.F("item_100"),
				Description: meorphistest40.F("Large tote with Bolt logo."),
				TotalAmount: meorphistest40.F(int64(9000)),
				UnitPrice:   meorphistest40.F(int64(1000)),
				Quantity:    meorphistest40.F(int64(9)),
				ImageURL:    meorphistest40.F("https://www.example.com/products/123456/images/1.png"),
			}}),
		}),
		PaymentMethod: meorphistest40.F(meorphistest40.PaymentNewParamsPaymentMethod{
			Tag: meorphistest40.F(meorphistest40.PaymentNewParamsPaymentMethodTagSavedPaymentMethod),
			ID:  meorphistest40.F("id"),
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
