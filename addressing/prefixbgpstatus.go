// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// PrefixBGPStatusService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixBGPStatusService] method instead.
type PrefixBGPStatusService struct {
	Options []option.RequestOption
}

// NewPrefixBGPStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrefixBGPStatusService(opts ...option.RequestOption) (r *PrefixBGPStatusService) {
	r = &PrefixBGPStatusService{}
	r.Options = opts
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *PrefixBGPStatusService) Edit(ctx context.Context, accountID string, prefixID string, body PrefixBGPStatusEditParams, opts ...option.RequestOption) (res *PrefixBGPStatusEditResponse, err error) {
	var env PrefixBGPStatusEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the current advertisement state for a prefix.
func (r *PrefixBGPStatusService) Get(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *PrefixBGPStatusGetResponse, err error) {
	var env PrefixBGPStatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrefixBGPStatusEditResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                       `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixBGPStatusEditResponseJSON `json:"-"`
}

// prefixBGPStatusEditResponseJSON contains the JSON metadata for the struct
// [PrefixBGPStatusEditResponse]
type prefixBGPStatusEditResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusGetResponse struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                      `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixBGPStatusGetResponseJSON `json:"-"`
}

// prefixBGPStatusGetResponseJSON contains the JSON metadata for the struct
// [PrefixBGPStatusGetResponse]
type prefixBGPStatusGetResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixBGPStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPStatusEditParams struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r PrefixBGPStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPStatusEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixBGPStatusEditResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixBGPStatusEditResponse                `json:"result"`
	JSON    prefixBGPStatusEditResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPStatusEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPStatusEditResponseEnvelope]
type prefixBGPStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPStatusEditResponseEnvelopeSuccess bool

const (
	PrefixBGPStatusEditResponseEnvelopeSuccessTrue PrefixBGPStatusEditResponseEnvelopeSuccess = true
)

func (r PrefixBGPStatusEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPStatusEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixBGPStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PrefixBGPStatusGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixBGPStatusGetResponse                `json:"result"`
	JSON    prefixBGPStatusGetResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPStatusGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixBGPStatusGetResponseEnvelope]
type prefixBGPStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrefixBGPStatusGetResponseEnvelopeSuccess bool

const (
	PrefixBGPStatusGetResponseEnvelopeSuccessTrue PrefixBGPStatusGetResponseEnvelopeSuccess = true
)

func (r PrefixBGPStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
