// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

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

// OwnershipService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOwnershipService] method instead.
type OwnershipService struct {
	Options []option.RequestOption
}

// NewOwnershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOwnershipService(opts ...option.RequestOption) (r *OwnershipService) {
	r = &OwnershipService{}
	r.Options = opts
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *OwnershipService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body OwnershipNewParams, opts ...option.RequestOption) (res *OwnershipNewResponse, err error) {
	var env OwnershipNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates ownership challenge of the destination.
func (r *OwnershipService) Validate(ctx context.Context, accountOrZone string, accountOrZoneID string, body OwnershipValidateParams, opts ...option.RequestOption) (res *OwnershipValidation, err error) {
	var env OwnershipValidateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership/validate", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OwnershipValidation struct {
	Valid bool                    `json:"valid"`
	JSON  ownershipValidationJSON `json:"-"`
}

// ownershipValidationJSON contains the JSON metadata for the struct
// [OwnershipValidation]
type ownershipValidationJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidationJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewResponse struct {
	Filename string                   `json:"filename"`
	Message  string                   `json:"message"`
	Valid    bool                     `json:"valid"`
	JSON     ownershipNewResponseJSON `json:"-"`
}

// ownershipNewResponseJSON contains the JSON metadata for the struct
// [OwnershipNewResponse]
type ownershipNewResponseJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseJSON) RawJSON() string {
	return r.raw
}

type OwnershipNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r OwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OwnershipNewResponse                `json:"result,nullable"`
	JSON    ownershipNewResponseEnvelopeJSON    `json:"-"`
}

// ownershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipNewResponseEnvelope]
type ownershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipNewResponseEnvelopeSuccess bool

const (
	OwnershipNewResponseEnvelopeSuccessTrue OwnershipNewResponseEnvelopeSuccess = true
)

func (r OwnershipNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r OwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipValidateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	Result  OwnershipValidation                      `json:"result,nullable"`
	JSON    ownershipValidateResponseEnvelopeJSON    `json:"-"`
}

// ownershipValidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipValidateResponseEnvelope]
type ownershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipValidateResponseEnvelopeSuccess bool

const (
	OwnershipValidateResponseEnvelopeSuccessTrue OwnershipValidateResponseEnvelopeSuccess = true
)

func (r OwnershipValidateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipValidateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
