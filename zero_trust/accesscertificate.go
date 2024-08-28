// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// AccessCertificateService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessCertificateService] method instead.
type AccessCertificateService struct {
	Options  []option.RequestOption
	Settings *AccessCertificateSettingService
}

// NewAccessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessCertificateService(opts ...option.RequestOption) (r *AccessCertificateService) {
	r = &AccessCertificateService{}
	r.Options = opts
	r.Settings = NewAccessCertificateSettingService(opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *AccessCertificateService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessCertificateNewParams, opts ...option.RequestOption) (res *Certificate, err error) {
	var env AccessCertificateNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured mTLS certificate.
func (r *AccessCertificateService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, certificateID string, body AccessCertificateUpdateParams, opts ...option.RequestOption) (res *Certificate, err error) {
	var env AccessCertificateUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all mTLS root certificates.
func (r *AccessCertificateService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *pagination.SinglePage[Certificate], err error) {
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
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
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

// Lists all mTLS root certificates.
func (r *AccessCertificateService) ListAutoPaging(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Certificate] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountOrZone, accountOrZoneID, opts...))
}

// Deletes an mTLS certificate.
func (r *AccessCertificateService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, certificateID string, opts ...option.RequestOption) (res *AccessCertificateDeleteResponse, err error) {
	var env AccessCertificateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single mTLS certificate.
func (r *AccessCertificateService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, certificateID string, opts ...option.RequestOption) (res *Certificate, err error) {
	var env AccessCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AssociatedHostnames = string

type AssociatedHostnamesParam = string

type Certificate struct {
	// The ID of the application that will use this certificate.
	ID string `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []AssociatedHostnames `json:"associated_hostnames"`
	CreatedAt           time.Time             `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time             `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string          `json:"name"`
	UpdatedAt time.Time       `json:"updated_at" format:"date-time"`
	JSON      certificateJSON `json:"-"`
}

// certificateJSON contains the JSON metadata for the struct [Certificate]
type certificateJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Certificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateDeleteResponse struct {
	// UUID
	ID   string                              `json:"id"`
	JSON accessCertificateDeleteResponseJSON `json:"-"`
}

// accessCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [AccessCertificateDeleteResponse]
type accessCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateNewParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]AssociatedHostnamesParam] `json:"associated_hostnames"`
}

func (r AccessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Certificate                                 `json:"result"`
	JSON    accessCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateNewResponseEnvelope]
type accessCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateNewResponseEnvelopeSuccess bool

const (
	AccessCertificateNewResponseEnvelopeSuccessTrue AccessCertificateNewResponseEnvelopeSuccess = true
)

func (r AccessCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]AssociatedHostnamesParam] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r AccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Certificate                                    `json:"result"`
	JSON    accessCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateUpdateResponseEnvelope]
type accessCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateUpdateResponseEnvelopeSuccess bool

const (
	AccessCertificateUpdateResponseEnvelopeSuccessTrue AccessCertificateUpdateResponseEnvelopeSuccess = true
)

func (r AccessCertificateUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessCertificateDeleteResponse                `json:"result"`
	JSON    accessCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateDeleteResponseEnvelope]
type accessCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateDeleteResponseEnvelopeSuccess bool

const (
	AccessCertificateDeleteResponseEnvelopeSuccessTrue AccessCertificateDeleteResponseEnvelopeSuccess = true
)

func (r AccessCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Certificate                                 `json:"result"`
	JSON    accessCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateGetResponseEnvelope]
type accessCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateGetResponseEnvelopeSuccess bool

const (
	AccessCertificateGetResponseEnvelopeSuccessTrue AccessCertificateGetResponseEnvelopeSuccess = true
)

func (r AccessCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}