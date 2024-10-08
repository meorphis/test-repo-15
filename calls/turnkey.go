// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// TURNKeyService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTURNKeyService] method instead.
type TURNKeyService struct {
	Options []option.RequestOption
}

// NewTURNKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTURNKeyService(opts ...option.RequestOption) (r *TURNKeyService) {
	r = &TURNKeyService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare Calls TURN key.
func (r *TURNKeyService) New(ctx context.Context, accountID string, body TURNKeyNewParams, opts ...option.RequestOption) (res *TURNKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Edit details for a single TURN key.
func (r *TURNKeyService) Update(ctx context.Context, accountID string, keyID string, body TURNKeyUpdateParams, opts ...option.RequestOption) (res *string, err error) {
	var env TURNKeyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", accountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all TURN keys in the Cloudflare account
func (r *TURNKeyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", accountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all TURN keys in the Cloudflare account
func (r *TURNKeyService) ListAutoPaging(ctx context.Context, accountID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, opts...))
}

// Deletes a TURN key from Cloudflare Calls
func (r *TURNKeyService) Delete(ctx context.Context, accountID string, keyID string, opts ...option.RequestOption) (res *string, err error) {
	var env TURNKeyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", accountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single TURN key.
func (r *TURNKeyService) Get(ctx context.Context, accountID string, keyID string, opts ...option.RequestOption) (res *string, err error) {
	var env TURNKeyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", accountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TURNKeyNewResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// Bearer token
	Key string `json:"key"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of a TURN key, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string                 `json:"uid"`
	JSON turnKeyNewResponseJSON `json:"-"`
}

// turnKeyNewResponseJSON contains the JSON metadata for the struct
// [TURNKeyNewResponse]
type turnKeyNewResponseJSON struct {
	Created     apijson.Field
	Key         apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type TURNKeyNewParams struct {
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r TURNKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TURNKeyUpdateParams struct {
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r TURNKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TURNKeyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TURNKeyUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Bearer token
	Result string                            `json:"result"`
	JSON   turnKeyUpdateResponseEnvelopeJSON `json:"-"`
}

// turnKeyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TURNKeyUpdateResponseEnvelope]
type turnKeyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNKeyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TURNKeyUpdateResponseEnvelopeSuccess bool

const (
	TURNKeyUpdateResponseEnvelopeSuccessTrue TURNKeyUpdateResponseEnvelopeSuccess = true
)

func (r TURNKeyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TURNKeyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TURNKeyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TURNKeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	// Bearer token
	Result string                            `json:"result"`
	JSON   turnKeyDeleteResponseEnvelopeJSON `json:"-"`
}

// turnKeyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TURNKeyDeleteResponseEnvelope]
type turnKeyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNKeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TURNKeyDeleteResponseEnvelopeSuccess bool

const (
	TURNKeyDeleteResponseEnvelopeSuccessTrue TURNKeyDeleteResponseEnvelopeSuccess = true
)

func (r TURNKeyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TURNKeyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TURNKeyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TURNKeyGetResponseEnvelopeSuccess `json:"success,required"`
	// Bearer token
	Result string                         `json:"result"`
	JSON   turnKeyGetResponseEnvelopeJSON `json:"-"`
}

// turnKeyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TURNKeyGetResponseEnvelope]
type turnKeyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TURNKeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TURNKeyGetResponseEnvelopeSuccess bool

const (
	TURNKeyGetResponseEnvelopeSuccessTrue TURNKeyGetResponseEnvelopeSuccess = true
)

func (r TURNKeyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TURNKeyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
