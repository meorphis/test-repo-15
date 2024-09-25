// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPPayloadLogService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPPayloadLogService] method instead.
type DLPPayloadLogService struct {
	Options []option.RequestOption
}

// NewDLPPayloadLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPayloadLogService(opts ...option.RequestOption) (r *DLPPayloadLogService) {
	r = &DLPPayloadLogService{}
	r.Options = opts
	return
}

// Updates the DLP payload log settings for this account.
func (r *DLPPayloadLogService) Update(ctx context.Context, accountID string, body DLPPayloadLogUpdateParams, opts ...option.RequestOption) (res *DLPPayloadLogUpdateResponse, err error) {
	var env DLPPayloadLogUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the current DLP payload log settings for this account.
func (r *DLPPayloadLogService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DLPPayloadLogGetResponse, err error) {
	var env DLPPayloadLogGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPayloadLogUpdateResponse struct {
	PublicKey string                          `json:"public_key,required,nullable"`
	JSON      dlpPayloadLogUpdateResponseJSON `json:"-"`
}

// dlpPayloadLogUpdateResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogUpdateResponse]
type dlpPayloadLogUpdateResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogGetResponse struct {
	PublicKey string                       `json:"public_key,required,nullable"`
	JSON      dlpPayloadLogGetResponseJSON `json:"-"`
}

// dlpPayloadLogGetResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponse]
type dlpPayloadLogGetResponseJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateParams struct {
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r DLPPayloadLogUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPayloadLogUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   DLPPayloadLogUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success DLPPayloadLogUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPPayloadLogUpdateResponseEnvelope]
type dlpPayloadLogUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPPayloadLogUpdateResponseEnvelopeSuccess bool

const (
	DLPPayloadLogUpdateResponseEnvelopeSuccessTrue DLPPayloadLogUpdateResponseEnvelopeSuccess = true
)

func (r DLPPayloadLogUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPayloadLogUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPPayloadLogGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   DLPPayloadLogGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success DLPPayloadLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogGetResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponseEnvelope]
type dlpPayloadLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPPayloadLogGetResponseEnvelopeSuccess bool

const (
	DLPPayloadLogGetResponseEnvelopeSuccessTrue DLPPayloadLogGetResponseEnvelopeSuccess = true
)

func (r DLPPayloadLogGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPayloadLogGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
