// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

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

// OutgoingStatusService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOutgoingStatusService] method instead.
type OutgoingStatusService struct {
	Options []option.RequestOption
}

// NewOutgoingStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOutgoingStatusService(opts ...option.RequestOption) (r *OutgoingStatusService) {
	r = &OutgoingStatusService{}
	r.Options = opts
	return
}

// Get primary zone transfer status.
func (r *OutgoingStatusService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *EnableTransfer, err error) {
	var env OutgoingStatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/status", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OutgoingStatusGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OutgoingStatusGetResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone
	Result EnableTransfer                        `json:"result"`
	JSON   outgoingStatusGetResponseEnvelopeJSON `json:"-"`
}

// outgoingStatusGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OutgoingStatusGetResponseEnvelope]
type outgoingStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OutgoingStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outgoingStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OutgoingStatusGetResponseEnvelopeSuccess bool

const (
	OutgoingStatusGetResponseEnvelopeSuccessTrue OutgoingStatusGetResponseEnvelopeSuccess = true
)

func (r OutgoingStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OutgoingStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
