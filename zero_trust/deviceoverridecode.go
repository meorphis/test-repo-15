// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// DeviceOverrideCodeService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceOverrideCodeService] method instead.
type DeviceOverrideCodeService struct {
	Options []option.RequestOption
}

// NewDeviceOverrideCodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceOverrideCodeService(opts ...option.RequestOption) (r *DeviceOverrideCodeService) {
	r = &DeviceOverrideCodeService{}
	r.Options = opts
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *DeviceOverrideCodeService) List(ctx context.Context, accountID string, deviceID string, opts ...option.RequestOption) (res *DeviceOverrideCodeListResponse, err error) {
	var env DeviceOverrideCodeListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/%s/override_codes", accountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceOverrideCodeListResponse struct {
	DisableForTime DeviceOverrideCodeListResponseDisableForTime `json:"disable_for_time"`
	JSON           deviceOverrideCodeListResponseJSON           `json:"-"`
}

// deviceOverrideCodeListResponseJSON contains the JSON metadata for the struct
// [DeviceOverrideCodeListResponse]
type deviceOverrideCodeListResponseJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeListResponseDisableForTime struct {
	// Override code that is valid for 1 hour.
	Number1 string `json:"1"`
	// Override code that is valid for 12 hour2.
	Number12 string `json:"12"`
	// Override code that is valid for 24 hour.2.
	Number24 string `json:"24"`
	// Override code that is valid for 3 hours.
	Number3 string `json:"3"`
	// Override code that is valid for 6 hours.
	Number6 string                                           `json:"6"`
	JSON    deviceOverrideCodeListResponseDisableForTimeJSON `json:"-"`
}

// deviceOverrideCodeListResponseDisableForTimeJSON contains the JSON metadata for
// the struct [DeviceOverrideCodeListResponseDisableForTime]
type deviceOverrideCodeListResponseDisableForTimeJSON struct {
	Number1     apijson.Field
	Number12    apijson.Field
	Number24    apijson.Field
	Number3     apijson.Field
	Number6     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponseDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseDisableForTimeJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeListResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   DeviceOverrideCodeListResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceOverrideCodeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceOverrideCodeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceOverrideCodeListResponseEnvelopeJSON       `json:"-"`
}

// deviceOverrideCodeListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceOverrideCodeListResponseEnvelope]
type deviceOverrideCodeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceOverrideCodeListResponseEnvelopeSuccess bool

const (
	DeviceOverrideCodeListResponseEnvelopeSuccessTrue DeviceOverrideCodeListResponseEnvelopeSuccess = true
)

func (r DeviceOverrideCodeListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceOverrideCodeListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceOverrideCodeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       deviceOverrideCodeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceOverrideCodeListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DeviceOverrideCodeListResponseEnvelopeResultInfo]
type deviceOverrideCodeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
