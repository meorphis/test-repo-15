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

func TestTestingShipmentTestingShipmentTrackingNewWithOptionalParams(t *testing.T) {
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
	err := client.Testings.Shipments.TestingShipmentTrackingNew(context.TODO(), meorphistest40.TestingShipmentTestingShipmentTrackingNewParams{
		Status: meorphistest40.F(meorphistest40.TestingShipmentTestingShipmentTrackingNewParamsStatusInTransit),
		TrackingDetails: meorphistest40.F([]meorphistest40.TestingShipmentTestingShipmentTrackingNewParamsTrackingDetail{{
			EventDate:   meorphistest40.F("2014-08-21:T14:24:00Z"),
			Status:      meorphistest40.F(meorphistest40.TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusPreTransit),
			Message:     meorphistest40.F("Billing information received"),
			Locality:    meorphistest40.F("San Francisco"),
			Region:      meorphistest40.F("CA"),
			PostalCode:  meorphistest40.F("string"),
			CountryCode: meorphistest40.F("US"),
		}, {
			EventDate:   meorphistest40.F("2014-08-21:T14:24:00Z"),
			Status:      meorphistest40.F(meorphistest40.TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusPreTransit),
			Message:     meorphistest40.F("Billing information received"),
			Locality:    meorphistest40.F("San Francisco"),
			Region:      meorphistest40.F("CA"),
			PostalCode:  meorphistest40.F("string"),
			CountryCode: meorphistest40.F("US"),
		}, {
			EventDate:   meorphistest40.F("2014-08-21:T14:24:00Z"),
			Status:      meorphistest40.F(meorphistest40.TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusPreTransit),
			Message:     meorphistest40.F("Billing information received"),
			Locality:    meorphistest40.F("San Francisco"),
			Region:      meorphistest40.F("CA"),
			PostalCode:  meorphistest40.F("string"),
			CountryCode: meorphistest40.F("US"),
		}}),
		TrackingNumber: meorphistest40.F("MockBolt-143292"),
		DeliveryDate:   meorphistest40.F(time.Now()),
	})
	if err != nil {
		var apierr *meorphistest40.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
