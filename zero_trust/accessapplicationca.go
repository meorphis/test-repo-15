// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// AccessApplicationCAService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationCAService] method instead.
type AccessApplicationCAService struct {
	Options []option.RequestOption
}

// NewAccessApplicationCAService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationCAService(opts ...option.RequestOption) (r *AccessApplicationCAService) {
	r = &AccessApplicationCAService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccessApplicationCAService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, appID string, opts ...option.RequestOption) (res *AccessApplicationCANewResponseUnion, err error) {
	var env AccessApplicationCANewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessApplicationCAService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *pagination.SinglePage[CA], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/ca", accountOrZone, accountOrZoneID)
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

// Lists short-lived certificate CAs and their public keys.
func (r *AccessApplicationCAService) ListAutoPaging(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CA] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountOrZone, accountOrZoneID, opts...))
}

// Deletes a short-lived certificate CA.
func (r *AccessApplicationCAService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, appID string, opts ...option.RequestOption) (res *AccessApplicationCADeleteResponse, err error) {
	var env AccessApplicationCADeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessApplicationCAService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, appID string, opts ...option.RequestOption) (res *AccessApplicationCAGetResponseUnion, err error) {
	var env AccessApplicationCAGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CA struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	AUD string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string `json:"public_key"`
	JSON      caJSON `json:"-"`
}

// caJSON contains the JSON metadata for the struct [CA]
type caJSON struct {
	ID          apijson.Field
	AUD         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.AccessApplicationCANewResponseUnknown] or
// [shared.UnionString].
type AccessApplicationCANewResponseUnion interface {
	ImplementsZeroTrustAccessApplicationCANewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationCANewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessApplicationCADeleteResponse struct {
	// The ID of the CA.
	ID   string                                `json:"id"`
	JSON accessApplicationCADeleteResponseJSON `json:"-"`
}

// accessApplicationCADeleteResponseJSON contains the JSON metadata for the struct
// [AccessApplicationCADeleteResponse]
type accessApplicationCADeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.AccessApplicationCAGetResponseUnknown] or
// [shared.UnionString].
type AccessApplicationCAGetResponseUnion interface {
	ImplementsZeroTrustAccessApplicationCAGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationCAGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessApplicationCANewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationCANewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationCANewResponseUnion           `json:"result"`
	JSON    accessApplicationCANewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCANewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCANewResponseEnvelope]
type accessApplicationCANewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCANewResponseEnvelopeSuccess bool

const (
	AccessApplicationCANewResponseEnvelopeSuccessTrue AccessApplicationCANewResponseEnvelopeSuccess = true
)

func (r AccessApplicationCANewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCANewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCADeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationCADeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationCADeleteResponse                `json:"result"`
	JSON    accessApplicationCADeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCADeleteResponseEnvelope]
type accessApplicationCADeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCADeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationCADeleteResponseEnvelopeSuccessTrue AccessApplicationCADeleteResponseEnvelopeSuccess = true
)

func (r AccessApplicationCADeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCADeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCAGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessApplicationCAGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationCAGetResponseUnion           `json:"result"`
	JSON    accessApplicationCAGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCAGetResponseEnvelope]
type accessApplicationCAGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCAGetResponseEnvelopeSuccess bool

const (
	AccessApplicationCAGetResponseEnvelopeSuccessTrue AccessApplicationCAGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationCAGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCAGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
