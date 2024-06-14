// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// TestingShipmentService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestingShipmentService] method instead.
type TestingShipmentService struct {
	Options []option.RequestOption
}

// NewTestingShipmentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTestingShipmentService(opts ...option.RequestOption) (r *TestingShipmentService) {
	r = &TestingShipmentService{}
	r.Options = opts
	return
}

// Simulate a shipment tracking update, such as those that Bolt would ingest from
// third-party shipping providers that would allow updating tracking and delivery
// information to shipments associated with orders.
func (r *TestingShipmentService) TestingShipmentTrackingNew(ctx context.Context, body TestingShipmentTestingShipmentTrackingNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "testing/shipments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type TestingShipmentTestingShipmentTrackingNewParams struct {
	// The shipment's status.
	Status param.Field[TestingShipmentTestingShipmentTrackingNewParamsStatus] `json:"status,required"`
	// A list of tracking updates that contain the shipment's status, location, and any
	// unique messages.
	TrackingDetails param.Field[[]TestingShipmentTestingShipmentTrackingNewParamsTrackingDetail] `json:"tracking_details,required"`
	// The carrier's tracking number for the shipment. Must be prefixed with
	// `MockBolt`.
	TrackingNumber param.Field[string] `json:"tracking_number,required"`
	// The shipment's actual or estimated delivery date.
	DeliveryDate param.Field[time.Time] `json:"delivery_date" format:"date-time"`
}

func (r TestingShipmentTestingShipmentTrackingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The shipment's status.
type TestingShipmentTestingShipmentTrackingNewParamsStatus string

const (
	TestingShipmentTestingShipmentTrackingNewParamsStatusInTransit TestingShipmentTestingShipmentTrackingNewParamsStatus = "in_transit"
	TestingShipmentTestingShipmentTrackingNewParamsStatusCancelled TestingShipmentTestingShipmentTrackingNewParamsStatus = "cancelled"
	TestingShipmentTestingShipmentTrackingNewParamsStatusFailure   TestingShipmentTestingShipmentTrackingNewParamsStatus = "failure"
	TestingShipmentTestingShipmentTrackingNewParamsStatusDelivered TestingShipmentTestingShipmentTrackingNewParamsStatus = "delivered"
)

func (r TestingShipmentTestingShipmentTrackingNewParamsStatus) IsKnown() bool {
	switch r {
	case TestingShipmentTestingShipmentTrackingNewParamsStatusInTransit, TestingShipmentTestingShipmentTrackingNewParamsStatusCancelled, TestingShipmentTestingShipmentTrackingNewParamsStatusFailure, TestingShipmentTestingShipmentTrackingNewParamsStatusDelivered:
		return true
	}
	return false
}

type TestingShipmentTestingShipmentTrackingNewParamsTrackingDetail struct {
	// The country associated this this set of tracking details, if any.
	CountryCode param.Field[string] `json:"country_code"`
	// The tracking detail's timestamp.
	EventDate param.Field[string] `json:"event_date"`
	// The locality associated this this set of tracking details, if any.
	Locality param.Field[string] `json:"locality"`
	// An arbitrary, human-readable message associated with this set of tracking
	// details.
	Message param.Field[string] `json:"message"`
	// The postal associated this this set of tracking details, if any.
	PostalCode param.Field[string] `json:"postal_code"`
	// The region associated this this set of tracking details, if any.
	Region param.Field[string]                                                               `json:"region"`
	Status param.Field[TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus] `json:"status"`
}

func (r TestingShipmentTestingShipmentTrackingNewParamsTrackingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus string

const (
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusUnknown            TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "unknown"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusPreTransit         TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "pre_transit"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusInTransit          TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "in_transit"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusOutForDelivery     TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "out_for_delivery"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusDelivered          TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "delivered"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusAvailableForPickup TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "available_for_pickup"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusReturnToSender     TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "return_to_sender"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusFailure            TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "failure"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusCancelled          TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "cancelled"
	TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusError              TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus = "error"
)

func (r TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatus) IsKnown() bool {
	switch r {
	case TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusUnknown, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusPreTransit, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusInTransit, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusOutForDelivery, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusDelivered, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusAvailableForPickup, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusReturnToSender, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusFailure, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusCancelled, TestingShipmentTestingShipmentTrackingNewParamsTrackingDetailsStatusError:
		return true
	}
	return false
}
