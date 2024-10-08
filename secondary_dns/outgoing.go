// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// OutgoingService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOutgoingService] method instead.
type OutgoingService struct {
	Options []option.RequestOption
	Status  *OutgoingStatusService
}

// NewOutgoingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOutgoingService(opts ...option.RequestOption) (r *OutgoingService) {
	r = &OutgoingService{}
	r.Options = opts
	r.Status = NewOutgoingStatusService(opts...)
	return
}

// Create primary zone configuration for outgoing zone transfers.
func (r *OutgoingService) New(ctx context.Context, zoneID string, body OutgoingNewParams, opts ...option.RequestOption) (res *OutgoingNewResponse, err error) {
	var env OutgoingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *OutgoingService) Update(ctx context.Context, zoneID string, body OutgoingUpdateParams, opts ...option.RequestOption) (res *OutgoingUpdateResponse, err error) {
	var env OutgoingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete primary zone configuration for outgoing zone transfers.
func (r *OutgoingService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *OutgoingDeleteResponse, err error) {
	var env OutgoingDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable outgoing zone transfers for primary zone and clears IXFR backlog of
// primary zone.
func (r *OutgoingService) Disable(ctx context.Context, zoneID string, body OutgoingDisableParams, opts ...option.RequestOption) (res *DisableTransfer, err error) {
	var env OutgoingDisableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/disable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable outgoing zone transfers for primary zone.
func (r *OutgoingService) Enable(ctx context.Context, zoneID string, body OutgoingEnableParams, opts ...option.RequestOption) (res *EnableTransfer, err error) {
	var env OutgoingEnableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
func (r *OutgoingService) ForceNotify(ctx context.Context, zoneID string, body OutgoingForceNotifyParams, opts ...option.RequestOption) (res *string, err error) {
	var env OutgoingForceNotifyResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/force_notify", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *OutgoingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *OutgoingGetResponse, err error) {
	var env OutgoingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DisableTransfer = string

type EnableTransfer = string

type OutgoingNewResponse struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                 `json:"soa_serial"`
	JSON      outgoingNewResponseJSON `json:"-"`
}

// outgoingNewResponseJSON contains the JSON metadata for the struct
// [OutgoingNewResponse]
type outgoingNewResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SOASerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OutgoingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingNewResponseJSON) RawJSON() string {
	return r.raw
}

type OutgoingUpdateResponse struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                    `json:"soa_serial"`
	JSON      outgoingUpdateResponseJSON `json:"-"`
}

// outgoingUpdateResponseJSON contains the JSON metadata for the struct
// [OutgoingUpdateResponse]
type outgoingUpdateResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SOASerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OutgoingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OutgoingDeleteResponse struct {
	ID   string                     `json:"id"`
	JSON outgoingDeleteResponseJSON `json:"-"`
}

// outgoingDeleteResponseJSON contains the JSON metadata for the struct
// [OutgoingDeleteResponse]
type outgoingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OutgoingGetResponse struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                 `json:"soa_serial"`
	JSON      outgoingGetResponseJSON `json:"-"`
}

// outgoingGetResponseJSON contains the JSON metadata for the struct
// [OutgoingGetResponse]
type outgoingGetResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SOASerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OutgoingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingGetResponseJSON) RawJSON() string {
	return r.raw
}

type OutgoingNewParams struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r OutgoingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OutgoingNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OutgoingNewResponse                `json:"result"`
	JSON    outgoingNewResponseEnvelopeJSON    `json:"-"`
}

// outgoingNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingNewResponseEnvelope]
type outgoingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingNewResponseEnvelopeSuccess bool

const (
	OutgoingNewResponseEnvelopeSuccessTrue OutgoingNewResponseEnvelopeSuccess = true
)

func (r OutgoingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OutgoingUpdateParams struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r OutgoingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OutgoingUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  OutgoingUpdateResponse                `json:"result"`
	JSON    outgoingUpdateResponseEnvelopeJSON    `json:"-"`
}

// outgoingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingUpdateResponseEnvelope]
type outgoingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingUpdateResponseEnvelopeSuccess bool

const (
	OutgoingUpdateResponseEnvelopeSuccessTrue OutgoingUpdateResponseEnvelopeSuccess = true
)

func (r OutgoingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OutgoingDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  OutgoingDeleteResponse                `json:"result"`
	JSON    outgoingDeleteResponseEnvelopeJSON    `json:"-"`
}

// outgoingDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingDeleteResponseEnvelope]
type outgoingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingDeleteResponseEnvelopeSuccess bool

const (
	OutgoingDeleteResponseEnvelopeSuccessTrue OutgoingDeleteResponseEnvelopeSuccess = true
)

func (r OutgoingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OutgoingDisableParams struct {
	Body interface{} `json:"body,required"`
}

func (r OutgoingDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OutgoingDisableResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingDisableResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone
	Result DisableTransfer                     `json:"result"`
	JSON   outgoingDisableResponseEnvelopeJSON `json:"-"`
}

// outgoingDisableResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingDisableResponseEnvelope]
type outgoingDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingDisableResponseEnvelopeSuccess bool

const (
	OutgoingDisableResponseEnvelopeSuccessTrue OutgoingDisableResponseEnvelopeSuccess = true
)

func (r OutgoingDisableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingDisableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OutgoingEnableParams struct {
	Body interface{} `json:"body,required"`
}

func (r OutgoingEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OutgoingEnableResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingEnableResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone
	Result EnableTransfer                     `json:"result"`
	JSON   outgoingEnableResponseEnvelopeJSON `json:"-"`
}

// outgoingEnableResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingEnableResponseEnvelope]
type outgoingEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingEnableResponseEnvelopeSuccess bool

const (
	OutgoingEnableResponseEnvelopeSuccessTrue OutgoingEnableResponseEnvelopeSuccess = true
)

func (r OutgoingEnableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingEnableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OutgoingForceNotifyParams struct {
	Body interface{} `json:"body,required"`
}

func (r OutgoingForceNotifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OutgoingForceNotifyResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingForceNotifyResponseEnvelopeSuccess `json:"success,required"`
	// When force_notify query parameter is set to true, the response is a simple
	// string
	Result string                                  `json:"result"`
	JSON   outgoingForceNotifyResponseEnvelopeJSON `json:"-"`
}

// outgoingForceNotifyResponseEnvelopeJSON contains the JSON metadata for the
// struct [OutgoingForceNotifyResponseEnvelope]
type outgoingForceNotifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingForceNotifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingForceNotifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingForceNotifyResponseEnvelopeSuccess bool

const (
	OutgoingForceNotifyResponseEnvelopeSuccessTrue OutgoingForceNotifyResponseEnvelopeSuccess = true
)

func (r OutgoingForceNotifyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingForceNotifyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OutgoingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OutgoingGetResponse                `json:"result"`
	JSON    outgoingGetResponseEnvelopeJSON    `json:"-"`
}

// outgoingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingGetResponseEnvelope]
type outgoingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingGetResponseEnvelopeSuccess bool

const (
	OutgoingGetResponseEnvelopeSuccessTrue OutgoingGetResponseEnvelopeSuccess = true
)

func (r OutgoingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
