// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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

// PreviewService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPreviewService] method instead.
type PreviewService struct {
	Options []option.RequestOption
}

// NewPreviewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPreviewService(opts ...option.RequestOption) (r *PreviewService) {
	r = &PreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *PreviewService) Get(ctx context.Context, accountID string, previewID string, opts ...option.RequestOption) (res *PreviewGetResponse, err error) {
	var env PreviewGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if previewID == "" {
		err = errors.New("missing required preview_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%s", accountID, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PreviewGetResponse map[string]PreviewGetResponseItem

type PreviewGetResponseItem struct {
	Healthy bool                                      `json:"healthy"`
	Origins []map[string]PreviewGetResponseItemOrigin `json:"origins"`
	JSON    previewGetResponseItemJSON                `json:"-"`
}

// previewGetResponseItemJSON contains the JSON metadata for the struct
// [PreviewGetResponseItem]
type previewGetResponseItemJSON struct {
	Healthy     apijson.Field
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseItemJSON) RawJSON() string {
	return r.raw
}

// The origin ipv4/ipv6 address or domain name mapped to it's health data.
type PreviewGetResponseItemOrigin struct {
	FailureReason string                           `json:"failure_reason"`
	Healthy       bool                             `json:"healthy"`
	ResponseCode  float64                          `json:"response_code"`
	RTT           string                           `json:"rtt"`
	JSON          previewGetResponseItemOriginJSON `json:"-"`
}

// previewGetResponseItemOriginJSON contains the JSON metadata for the struct
// [PreviewGetResponseItemOrigin]
type previewGetResponseItemOriginJSON struct {
	FailureReason apijson.Field
	Healthy       apijson.Field
	ResponseCode  apijson.Field
	RTT           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PreviewGetResponseItemOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseItemOriginJSON) RawJSON() string {
	return r.raw
}

type PreviewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Resulting health data from a preview operation.
	Result PreviewGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success PreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewGetResponseEnvelopeJSON    `json:"-"`
}

// previewGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewGetResponseEnvelope]
type previewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewGetResponseEnvelopeSuccess bool

const (
	PreviewGetResponseEnvelopeSuccessTrue PreviewGetResponseEnvelopeSuccess = true
)

func (r PreviewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
