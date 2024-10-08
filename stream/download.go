// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// DownloadService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDownloadService] method instead.
type DownloadService struct {
	Options []option.RequestOption
}

// NewDownloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDownloadService(opts ...option.RequestOption) (r *DownloadService) {
	r = &DownloadService{}
	r.Options = opts
	return
}

// Creates a download for a video when a video is ready to view.
func (r *DownloadService) New(ctx context.Context, accountID string, identifier string, body DownloadNewParams, opts ...option.RequestOption) (res *DownloadNewResponseUnion, err error) {
	var env DownloadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the downloads for a video.
func (r *DownloadService) Delete(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *DownloadDeleteResponseUnion, err error) {
	var env DownloadDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the downloads created for a video.
func (r *DownloadService) Get(ctx context.Context, accountID string, identifier string, opts ...option.RequestOption) (res *DownloadGetResponseUnion, err error) {
	var env DownloadGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [stream.DownloadNewResponseUnknown] or [shared.UnionString].
type DownloadNewResponseUnion interface {
	ImplementsStreamDownloadNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DownloadNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.DownloadDeleteResponseUnknown] or
// [shared.UnionString].
type DownloadDeleteResponseUnion interface {
	ImplementsStreamDownloadDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DownloadDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.DownloadGetResponseUnknown] or [shared.UnionString].
type DownloadGetResponseUnion interface {
	ImplementsStreamDownloadGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DownloadGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DownloadNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r DownloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DownloadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DownloadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DownloadNewResponseUnion           `json:"result"`
	JSON    downloadNewResponseEnvelopeJSON    `json:"-"`
}

// downloadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DownloadNewResponseEnvelope]
type downloadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DownloadNewResponseEnvelopeSuccess bool

const (
	DownloadNewResponseEnvelopeSuccessTrue DownloadNewResponseEnvelopeSuccess = true
)

func (r DownloadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DownloadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DownloadDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DownloadDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DownloadDeleteResponseUnion           `json:"result"`
	JSON    downloadDeleteResponseEnvelopeJSON    `json:"-"`
}

// downloadDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DownloadDeleteResponseEnvelope]
type downloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DownloadDeleteResponseEnvelopeSuccess bool

const (
	DownloadDeleteResponseEnvelopeSuccessTrue DownloadDeleteResponseEnvelopeSuccess = true
)

func (r DownloadDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DownloadDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DownloadGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DownloadGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DownloadGetResponseUnion           `json:"result"`
	JSON    downloadGetResponseEnvelopeJSON    `json:"-"`
}

// downloadGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DownloadGetResponseEnvelope]
type downloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DownloadGetResponseEnvelopeSuccess bool

const (
	DownloadGetResponseEnvelopeSuccessTrue DownloadGetResponseEnvelopeSuccess = true
)

func (r DownloadGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DownloadGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
