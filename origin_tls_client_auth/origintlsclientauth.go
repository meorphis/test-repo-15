// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// OriginTLSClientAuthService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginTLSClientAuthService] method instead.
type OriginTLSClientAuthService struct {
	Options   []option.RequestOption
	Hostnames *HostnameService
	Settings  *SettingService
}

// NewOriginTLSClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginTLSClientAuthService(opts ...option.RequestOption) (r *OriginTLSClientAuthService) {
	r = &OriginTLSClientAuthService{}
	r.Options = opts
	r.Hostnames = NewHostnameService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *OriginTLSClientAuthService) New(ctx context.Context, zoneID string, body OriginTLSClientAuthNewParams, opts ...option.RequestOption) (res *OriginTLSClientAuthNewResponseUnion, err error) {
	var env OriginTLSClientAuthNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *OriginTLSClientAuthService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *pagination.SinglePage[ZoneAuthenticatedOriginPull], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneID)
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

// List Certificates
func (r *OriginTLSClientAuthService) ListAutoPaging(ctx context.Context, zoneID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZoneAuthenticatedOriginPull] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, zoneID, opts...))
}

// Delete Certificate
func (r *OriginTLSClientAuthService) Delete(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *OriginTLSClientAuthDeleteResponseUnion, err error) {
	var env OriginTLSClientAuthDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Certificate Details
func (r *OriginTLSClientAuthService) Get(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *OriginTLSClientAuthGetResponseUnion, err error) {
	var env OriginTLSClientAuthGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneAuthenticatedOriginPull struct {
	// Identifier
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate activation.
	Status ZoneAuthenticatedOriginPullStatus `json:"status"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                       `json:"uploaded_on" format:"date-time"`
	JSON       zoneAuthenticatedOriginPullJSON `json:"-"`
}

// zoneAuthenticatedOriginPullJSON contains the JSON metadata for the struct
// [ZoneAuthenticatedOriginPull]
type zoneAuthenticatedOriginPullJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	ExpiresOn   apijson.Field
	Issuer      apijson.Field
	Signature   apijson.Field
	Status      apijson.Field
	UploadedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAuthenticatedOriginPull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAuthenticatedOriginPullJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate activation.
type ZoneAuthenticatedOriginPullStatus string

const (
	ZoneAuthenticatedOriginPullStatusInitializing       ZoneAuthenticatedOriginPullStatus = "initializing"
	ZoneAuthenticatedOriginPullStatusPendingDeployment  ZoneAuthenticatedOriginPullStatus = "pending_deployment"
	ZoneAuthenticatedOriginPullStatusPendingDeletion    ZoneAuthenticatedOriginPullStatus = "pending_deletion"
	ZoneAuthenticatedOriginPullStatusActive             ZoneAuthenticatedOriginPullStatus = "active"
	ZoneAuthenticatedOriginPullStatusDeleted            ZoneAuthenticatedOriginPullStatus = "deleted"
	ZoneAuthenticatedOriginPullStatusDeploymentTimedOut ZoneAuthenticatedOriginPullStatus = "deployment_timed_out"
	ZoneAuthenticatedOriginPullStatusDeletionTimedOut   ZoneAuthenticatedOriginPullStatus = "deletion_timed_out"
)

func (r ZoneAuthenticatedOriginPullStatus) IsKnown() bool {
	switch r {
	case ZoneAuthenticatedOriginPullStatusInitializing, ZoneAuthenticatedOriginPullStatusPendingDeployment, ZoneAuthenticatedOriginPullStatusPendingDeletion, ZoneAuthenticatedOriginPullStatusActive, ZoneAuthenticatedOriginPullStatusDeleted, ZoneAuthenticatedOriginPullStatusDeploymentTimedOut, ZoneAuthenticatedOriginPullStatusDeletionTimedOut:
		return true
	}
	return false
}

// Union satisfied by
// [origin_tls_client_auth.OriginTLSClientAuthNewResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthNewResponseUnion interface {
	ImplementsOriginTLSClientAuthOriginTLSClientAuthNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [origin_tls_client_auth.OriginTLSClientAuthDeleteResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthDeleteResponseUnion interface {
	ImplementsOriginTLSClientAuthOriginTLSClientAuthDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [origin_tls_client_auth.OriginTLSClientAuthGetResponseUnknown] or
// [shared.UnionString].
type OriginTLSClientAuthGetResponseUnion interface {
	ImplementsOriginTLSClientAuthOriginTLSClientAuthGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginTLSClientAuthGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginTLSClientAuthNewParams struct {
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r OriginTLSClientAuthNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginTLSClientAuthNewResponseUnion           `json:"result"`
	JSON    originTLSClientAuthNewResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthNewResponseEnvelope]
type originTLSClientAuthNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthNewResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthNewResponseEnvelopeSuccessTrue OriginTLSClientAuthNewResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSClientAuthDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginTLSClientAuthDeleteResponseUnion           `json:"result"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthDeleteResponseEnvelope]
type originTLSClientAuthDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthDeleteResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue OriginTLSClientAuthDeleteResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSClientAuthGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginTLSClientAuthGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginTLSClientAuthGetResponseUnion           `json:"result"`
	JSON    originTLSClientAuthGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthGetResponseEnvelope]
type originTLSClientAuthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginTLSClientAuthGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthGetResponseEnvelopeSuccessTrue OriginTLSClientAuthGetResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
