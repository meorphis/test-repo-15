// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// KeyService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	Options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r *KeyService) {
	r = &KeyService{}
	r.Options = opts
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *KeyService) New(ctx context.Context, accountID string, body KeyNewParams, opts ...option.RequestOption) (res *Keys, err error) {
	var env KeyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *KeyService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *KeyDeleteResponseUnion, err error) {
	var env KeyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *KeyService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]KeyGetResponse, err error) {
	var env KeyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Keys struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string   `json:"pem"`
	JSON keysJSON `json:"-"`
}

// keysJSON contains the JSON metadata for the struct [Keys]
type keysJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Keys) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keysJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [stream.KeyDeleteResponseUnknown] or [shared.UnionString].
type KeyDeleteResponseUnion interface {
	ImplementsStreamKeyDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KeyDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type KeyGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time          `json:"created" format:"date-time"`
	JSON    keyGetResponseJSON `json:"-"`
}

// keyGetResponseJSON contains the JSON metadata for the struct [KeyGetResponse]
type keyGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseJSON) RawJSON() string {
	return r.raw
}

type KeyNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r KeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type KeyNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success KeyNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Keys                          `json:"result"`
	JSON    keyNewResponseEnvelopeJSON    `json:"-"`
}

// keyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelope]
type keyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeyNewResponseEnvelopeSuccess bool

const (
	KeyNewResponseEnvelopeSuccessTrue KeyNewResponseEnvelopeSuccess = true
)

func (r KeyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success KeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  KeyDeleteResponseUnion           `json:"result"`
	JSON    keyDeleteResponseEnvelopeJSON    `json:"-"`
}

// keyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelope]
type keyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeyDeleteResponseEnvelopeSuccess bool

const (
	KeyDeleteResponseEnvelopeSuccessTrue KeyDeleteResponseEnvelopeSuccess = true
)

func (r KeyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success KeyGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []KeyGetResponse              `json:"result"`
	JSON    keyGetResponseEnvelopeJSON    `json:"-"`
}

// keyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyGetResponseEnvelope]
type keyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeyGetResponseEnvelopeSuccess bool

const (
	KeyGetResponseEnvelopeSuccessTrue KeyGetResponseEnvelopeSuccess = true
)

func (r KeyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
