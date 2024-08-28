// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthchecks

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

// Create a new preview health check.
func (r *PreviewService) New(ctx context.Context, zoneID string, body PreviewNewParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	var env PreviewNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/preview", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *PreviewService) Delete(ctx context.Context, zoneID string, healthcheckID string, opts ...option.RequestOption) (res *PreviewDeleteResponse, err error) {
	var env PreviewDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check preview.
func (r *PreviewService) Get(ctx context.Context, zoneID string, healthcheckID string, opts ...option.RequestOption) (res *Healthcheck, err error) {
	var env PreviewGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PreviewDeleteResponse struct {
	// Identifier
	ID   string                    `json:"id"`
	JSON previewDeleteResponseJSON `json:"-"`
}

// previewDeleteResponseJSON contains the JSON metadata for the struct
// [PreviewDeleteResponse]
type previewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PreviewNewParams struct {
	QueryHealthcheck QueryHealthcheckParam `json:"query_healthcheck,required"`
}

func (r PreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.QueryHealthcheck)
}

type PreviewNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Healthcheck           `json:"result,required"`
	// Whether the API call was successful
	Success PreviewNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewNewResponseEnvelopeJSON    `json:"-"`
}

// previewNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewNewResponseEnvelope]
type previewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewNewResponseEnvelopeSuccess bool

const (
	PreviewNewResponseEnvelopeSuccessTrue PreviewNewResponseEnvelopeSuccess = true
)

func (r PreviewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PreviewDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PreviewDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success PreviewDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewDeleteResponseEnvelopeJSON    `json:"-"`
}

// previewDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewDeleteResponseEnvelope]
type previewDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewDeleteResponseEnvelopeSuccess bool

const (
	PreviewDeleteResponseEnvelopeSuccessTrue PreviewDeleteResponseEnvelopeSuccess = true
)

func (r PreviewDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PreviewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Healthcheck           `json:"result,required"`
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
