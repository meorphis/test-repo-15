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

// IncomingService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIncomingService] method instead.
type IncomingService struct {
	Options []option.RequestOption
}

// NewIncomingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIncomingService(opts ...option.RequestOption) (r *IncomingService) {
	r = &IncomingService{}
	r.Options = opts
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *IncomingService) New(ctx context.Context, zoneID string, body IncomingNewParams, opts ...option.RequestOption) (res *IncomingNewResponse, err error) {
	var env IncomingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *IncomingService) Update(ctx context.Context, zoneID string, body IncomingUpdateParams, opts ...option.RequestOption) (res *IncomingUpdateResponse, err error) {
	var env IncomingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *IncomingService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *IncomingDeleteResponse, err error) {
	var env IncomingDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *IncomingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *IncomingGetResponse, err error) {
	var env IncomingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IncomingNewResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                 `json:"soa_serial"`
	JSON      incomingNewResponseJSON `json:"-"`
}

// incomingNewResponseJSON contains the JSON metadata for the struct
// [IncomingNewResponse]
type incomingNewResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IncomingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingNewResponseJSON) RawJSON() string {
	return r.raw
}

type IncomingUpdateResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                    `json:"soa_serial"`
	JSON      incomingUpdateResponseJSON `json:"-"`
}

// incomingUpdateResponseJSON contains the JSON metadata for the struct
// [IncomingUpdateResponse]
type incomingUpdateResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IncomingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IncomingDeleteResponse struct {
	ID   string                     `json:"id"`
	JSON incomingDeleteResponseJSON `json:"-"`
}

// incomingDeleteResponseJSON contains the JSON metadata for the struct
// [IncomingDeleteResponse]
type incomingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncomingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IncomingGetResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                 `json:"soa_serial"`
	JSON      incomingGetResponseJSON `json:"-"`
}

// incomingGetResponseJSON contains the JSON metadata for the struct
// [IncomingGetResponse]
type incomingGetResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IncomingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingGetResponseJSON) RawJSON() string {
	return r.raw
}

type IncomingNewParams struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r IncomingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IncomingNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IncomingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  IncomingNewResponse                `json:"result"`
	JSON    incomingNewResponseEnvelopeJSON    `json:"-"`
}

// incomingNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IncomingNewResponseEnvelope]
type incomingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncomingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IncomingNewResponseEnvelopeSuccess bool

const (
	IncomingNewResponseEnvelopeSuccessTrue IncomingNewResponseEnvelopeSuccess = true
)

func (r IncomingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IncomingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IncomingUpdateParams struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r IncomingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IncomingUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IncomingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  IncomingUpdateResponse                `json:"result"`
	JSON    incomingUpdateResponseEnvelopeJSON    `json:"-"`
}

// incomingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [IncomingUpdateResponseEnvelope]
type incomingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncomingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IncomingUpdateResponseEnvelopeSuccess bool

const (
	IncomingUpdateResponseEnvelopeSuccessTrue IncomingUpdateResponseEnvelopeSuccess = true
)

func (r IncomingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IncomingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IncomingDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IncomingDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  IncomingDeleteResponse                `json:"result"`
	JSON    incomingDeleteResponseEnvelopeJSON    `json:"-"`
}

// incomingDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [IncomingDeleteResponseEnvelope]
type incomingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncomingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IncomingDeleteResponseEnvelopeSuccess bool

const (
	IncomingDeleteResponseEnvelopeSuccessTrue IncomingDeleteResponseEnvelopeSuccess = true
)

func (r IncomingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IncomingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IncomingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IncomingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  IncomingGetResponse                `json:"result"`
	JSON    incomingGetResponseEnvelopeJSON    `json:"-"`
}

// incomingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IncomingGetResponseEnvelope]
type incomingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IncomingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r incomingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IncomingGetResponseEnvelopeSuccess bool

const (
	IncomingGetResponseEnvelopeSuccessTrue IncomingGetResponseEnvelopeSuccess = true
)

func (r IncomingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IncomingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
