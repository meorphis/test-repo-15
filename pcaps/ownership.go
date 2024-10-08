// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pcaps

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

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *OwnershipService) New(ctx context.Context, accountID string, body OwnershipNewParams, opts ...option.RequestOption) (res *Ownership, err error) {
	var env OwnershipNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes buckets added to the packet captures API.
func (r *OwnershipService) Delete(ctx context.Context, accountID string, ownershipID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ownershipID == "" {
		err = errors.New("missing required ownership_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", accountID, ownershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *OwnershipService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]Ownership, err error) {
	var env OwnershipGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates buckets added to the packet captures API.
func (r *OwnershipService) Validate(ctx context.Context, accountID string, body OwnershipValidateParams, opts ...option.RequestOption) (res *Ownership, err error) {
	var env OwnershipValidateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Ownership struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status OwnershipStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string        `json:"validated"`
	JSON      ownershipJSON `json:"-"`
}

// ownershipJSON contains the JSON metadata for the struct [Ownership]
type ownershipJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Ownership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipJSON) RawJSON() string {
	return r.raw
}

// The status of the ownership challenge. Can be pending, success or failed.
type OwnershipStatus string

const (
	OwnershipStatusPending OwnershipStatus = "pending"
	OwnershipStatusSuccess OwnershipStatus = "success"
	OwnershipStatusFailed  OwnershipStatus = "failed"
)

func (r OwnershipStatus) IsKnown() bool {
	switch r {
	case OwnershipStatusPending, OwnershipStatusSuccess, OwnershipStatusFailed:
		return true
	}
	return false
}

type OwnershipNewParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r OwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Ownership             `json:"result,required"`
	// Whether the API call was successful
	Success OwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ownershipNewResponseEnvelopeJSON    `json:"-"`
}

// ownershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipNewResponseEnvelope]
type ownershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type OwnershipGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []Ownership           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    OwnershipGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo OwnershipGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ownershipGetResponseEnvelopeJSON       `json:"-"`
}

// ownershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipGetResponseEnvelope]
type ownershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OwnershipGetResponseEnvelopeSuccess bool

const (
	OwnershipGetResponseEnvelopeSuccessTrue OwnershipGetResponseEnvelopeSuccess = true
)

func (r OwnershipGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OwnershipGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OwnershipGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       ownershipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ownershipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [OwnershipGetResponseEnvelopeResultInfo]
type ownershipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OwnershipGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ownershipGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type OwnershipValidateParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r OwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OwnershipValidateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Ownership             `json:"result,required"`
	// Whether the API call was successful
	Success OwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ownershipValidateResponseEnvelopeJSON    `json:"-"`
}

// ownershipValidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OwnershipValidateResponseEnvelope]
type ownershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
