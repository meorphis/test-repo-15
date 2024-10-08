// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

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

// NamespaceBulkService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceBulkService] method instead.
type NamespaceBulkService struct {
	Options []option.RequestOption
}

// NewNamespaceBulkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceBulkService(opts ...option.RequestOption) (r *NamespaceBulkService) {
	r = &NamespaceBulkService{}
	r.Options = opts
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *NamespaceBulkService) Update(ctx context.Context, accountID string, namespaceID string, body NamespaceBulkUpdateParams, opts ...option.RequestOption) (res *NamespaceBulkUpdateResponse, err error) {
	var env NamespaceBulkUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *NamespaceBulkService) Delete(ctx context.Context, accountID string, namespaceID string, opts ...option.RequestOption) (res *NamespaceBulkDeleteResponse, err error) {
	var env NamespaceBulkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NamespaceBulkUpdateResponse struct {
	JSON namespaceBulkUpdateResponseJSON `json:"-"`
}

// namespaceBulkUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceBulkUpdateResponse]
type namespaceBulkUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkDeleteResponse struct {
	JSON namespaceBulkDeleteResponseJSON `json:"-"`
}

// namespaceBulkDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceBulkDeleteResponse]
type namespaceBulkDeleteResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceBulkUpdateParams struct {
	Body []NamespaceBulkUpdateParamsBody `json:"body,required"`
}

func (r NamespaceBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type NamespaceBulkUpdateParamsBody struct {
	// Whether or not the server should base64 decode the value before storing it.
	// Useful for writing values that wouldn't otherwise be valid JSON strings, such as
	// images.
	Base64 param.Field[bool] `json:"base64"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// should expire.
	Expiration param.Field[float64] `json:"expiration"`
	// The number of seconds for which the key should be visible before it expires. At
	// least 60.
	ExpirationTTL param.Field[float64] `json:"expiration_ttl"`
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key"`
	// Arbitrary JSON that is associated with a key.
	Metadata param.Field[map[string]interface{}] `json:"metadata"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value"`
}

func (r NamespaceBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceBulkUpdateResponse                `json:"result,nullable"`
	JSON    namespaceBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceBulkUpdateResponseEnvelope]
type namespaceBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceBulkUpdateResponseEnvelopeSuccess bool

const (
	NamespaceBulkUpdateResponseEnvelopeSuccessTrue NamespaceBulkUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceBulkDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  NamespaceBulkDeleteResponse                `json:"result,nullable"`
	JSON    namespaceBulkDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceBulkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceBulkDeleteResponseEnvelope]
type namespaceBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceBulkDeleteResponseEnvelopeSuccess bool

const (
	NamespaceBulkDeleteResponseEnvelopeSuccessTrue NamespaceBulkDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
