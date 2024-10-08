// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_post_quantum_encryption

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

// OriginPostQuantumEncryptionService contains methods and other services that help
// with interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginPostQuantumEncryptionService] method instead.
type OriginPostQuantumEncryptionService struct {
	Options []option.RequestOption
}

// NewOriginPostQuantumEncryptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOriginPostQuantumEncryptionService(opts ...option.RequestOption) (r *OriginPostQuantumEncryptionService) {
	r = &OriginPostQuantumEncryptionService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Update(ctx context.Context, zoneID string, body OriginPostQuantumEncryptionUpdateParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionUpdateResponseUnion, err error) {
	var env OriginPostQuantumEncryptionUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionGetResponseUnion, err error) {
	var env OriginPostQuantumEncryptionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [origin_post_quantum_encryption.OriginPostQuantumEncryptionUpdateResponseUnknown]
// or [shared.UnionString].
type OriginPostQuantumEncryptionUpdateResponseUnion interface {
	ImplementsOriginPostQuantumEncryptionOriginPostQuantumEncryptionUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPostQuantumEncryptionUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [origin_post_quantum_encryption.OriginPostQuantumEncryptionGetResponseUnknown]
// or [shared.UnionString].
type OriginPostQuantumEncryptionGetResponseUnion interface {
	ImplementsOriginPostQuantumEncryptionOriginPostQuantumEncryptionGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPostQuantumEncryptionGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginPostQuantumEncryptionUpdateParams struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionUpdateParamsValue] `json:"value,required"`
}

func (r OriginPostQuantumEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Post Quantum Encryption Setting.
type OriginPostQuantumEncryptionUpdateParamsValue string

const (
	OriginPostQuantumEncryptionUpdateParamsValuePreferred OriginPostQuantumEncryptionUpdateParamsValue = "preferred"
	OriginPostQuantumEncryptionUpdateParamsValueSupported OriginPostQuantumEncryptionUpdateParamsValue = "supported"
	OriginPostQuantumEncryptionUpdateParamsValueOff       OriginPostQuantumEncryptionUpdateParamsValue = "off"
)

func (r OriginPostQuantumEncryptionUpdateParamsValue) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateParamsValuePreferred, OriginPostQuantumEncryptionUpdateParamsValueSupported, OriginPostQuantumEncryptionUpdateParamsValueOff:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                          `json:"errors,required"`
	Messages []shared.ResponseInfo                          `json:"messages,required"`
	Result   OriginPostQuantumEncryptionUpdateResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    originPostQuantumEncryptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionUpdateResponseEnvelope]
type originPostQuantumEncryptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess = true
)

func (r OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   OriginPostQuantumEncryptionGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionGetResponseEnvelope]
type originPostQuantumEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginPostQuantumEncryptionGetResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionGetResponseEnvelopeSuccess = true
)

func (r OriginPostQuantumEncryptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
