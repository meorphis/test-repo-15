// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// AccountSettingService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSettingService] method instead.
type AccountSettingService struct {
	Options []option.RequestOption
}

// NewAccountSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountSettingService(opts ...option.RequestOption) (r *AccountSettingService) {
	r = &AccountSettingService{}
	r.Options = opts
	return
}

// Creates Worker account settings for an account.
func (r *AccountSettingService) Update(ctx context.Context, accountID string, body AccountSettingUpdateParams, opts ...option.RequestOption) (res *AccountSettingUpdateResponse, err error) {
	var env AccountSettingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Worker account settings for an account.
func (r *AccountSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountSettingGetResponse, err error) {
	var env AccountSettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountSettingUpdateResponse struct {
	DefaultUsageModel interface{}                      `json:"default_usage_model"`
	GreenCompute      interface{}                      `json:"green_compute"`
	JSON              accountSettingUpdateResponseJSON `json:"-"`
}

// accountSettingUpdateResponseJSON contains the JSON metadata for the struct
// [AccountSettingUpdateResponse]
type accountSettingUpdateResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSettingGetResponse struct {
	DefaultUsageModel interface{}                   `json:"default_usage_model"`
	GreenCompute      interface{}                   `json:"green_compute"`
	JSON              accountSettingGetResponseJSON `json:"-"`
}

// accountSettingGetResponseJSON contains the JSON metadata for the struct
// [AccountSettingGetResponse]
type accountSettingGetResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSettingUpdateParams struct {
	Body string `json:"body,required"`
}

func (r AccountSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountSettingUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccountSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccountSettingUpdateResponse                `json:"result"`
	JSON    accountSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// accountSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccountSettingUpdateResponseEnvelope]
type accountSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSettingUpdateResponseEnvelopeSuccess bool

const (
	AccountSettingUpdateResponseEnvelopeSuccessTrue AccountSettingUpdateResponseEnvelopeSuccess = true
)

func (r AccountSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccountSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccountSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccountSettingGetResponse                `json:"result"`
	JSON    accountSettingGetResponseEnvelopeJSON    `json:"-"`
}

// accountSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountSettingGetResponseEnvelope]
type accountSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSettingGetResponseEnvelopeSuccess bool

const (
	AccountSettingGetResponseEnvelopeSuccessTrue AccountSettingGetResponseEnvelopeSuccess = true
)

func (r AccountSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}