// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// SmartRoutingService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmartRoutingService] method instead.
type SmartRoutingService struct {
	Options []option.RequestOption
}

// NewSmartRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSmartRoutingService(opts ...option.RequestOption) (r *SmartRoutingService) {
	r = &SmartRoutingService{}
	r.Options = opts
	return
}

// Updates enablement of Argo Smart Routing.
func (r *SmartRoutingService) Edit(ctx context.Context, zoneID string, body SmartRoutingEditParams, opts ...option.RequestOption) (res *SmartRoutingEditResponseUnion, err error) {
	var env SmartRoutingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Argo Smart Routing setting
func (r *SmartRoutingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SmartRoutingGetResponseUnion, err error) {
	var env SmartRoutingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [argo.SmartRoutingEditResponseUnknown] or
// [shared.UnionString].
type SmartRoutingEditResponseUnion interface {
	ImplementsArgoSmartRoutingEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartRoutingEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [argo.SmartRoutingGetResponseUnknown] or
// [shared.UnionString].
type SmartRoutingGetResponseUnion interface {
	ImplementsArgoSmartRoutingGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartRoutingGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SmartRoutingEditParams struct {
	// Enables Argo Smart Routing.
	Value param.Field[SmartRoutingEditParamsValue] `json:"value,required"`
}

func (r SmartRoutingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Argo Smart Routing.
type SmartRoutingEditParamsValue string

const (
	SmartRoutingEditParamsValueOn  SmartRoutingEditParamsValue = "on"
	SmartRoutingEditParamsValueOff SmartRoutingEditParamsValue = "off"
)

func (r SmartRoutingEditParamsValue) IsKnown() bool {
	switch r {
	case SmartRoutingEditParamsValueOn, SmartRoutingEditParamsValueOff:
		return true
	}
	return false
}

type SmartRoutingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   SmartRoutingEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success SmartRoutingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartRoutingEditResponseEnvelopeJSON    `json:"-"`
}

// smartRoutingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SmartRoutingEditResponseEnvelope]
type smartRoutingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartRoutingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartRoutingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartRoutingEditResponseEnvelopeSuccess bool

const (
	SmartRoutingEditResponseEnvelopeSuccessTrue SmartRoutingEditResponseEnvelopeSuccess = true
)

func (r SmartRoutingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartRoutingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartRoutingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   SmartRoutingGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success SmartRoutingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartRoutingGetResponseEnvelopeJSON    `json:"-"`
}

// smartRoutingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SmartRoutingGetResponseEnvelope]
type smartRoutingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartRoutingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartRoutingGetResponseEnvelopeSuccess bool

const (
	SmartRoutingGetResponseEnvelopeSuccessTrue SmartRoutingGetResponseEnvelopeSuccess = true
)

func (r SmartRoutingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartRoutingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
