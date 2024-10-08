// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

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

// CertificatePackOrderService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificatePackOrderService] method instead.
type CertificatePackOrderService struct {
	Options []option.RequestOption
}

// NewCertificatePackOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCertificatePackOrderService(opts ...option.RequestOption) (r *CertificatePackOrderService) {
	r = &CertificatePackOrderService{}
	r.Options = opts
	return
}

// For a given zone, order an advanced certificate pack.
func (r *CertificatePackOrderService) New(ctx context.Context, zoneID string, body CertificatePackOrderNewParams, opts ...option.RequestOption) (res *CertificatePackOrderNewResponse, err error) {
	var env CertificatePackOrderNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificatePackOrderNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority CertificatePackOrderNewResponseCertificateAuthority `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding bool `json:"cloudflare_branding"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts []Host `json:"hosts"`
	// Status of certificate pack.
	Status Status `json:"status"`
	// Type of certificate pack.
	Type CertificatePackOrderNewResponseType `json:"type"`
	// Validation Method selected for the order.
	ValidationMethod CertificatePackOrderNewResponseValidationMethod `json:"validation_method"`
	// Validity Days selected for the order.
	ValidityDays CertificatePackOrderNewResponseValidityDays `json:"validity_days"`
	JSON         certificatePackOrderNewResponseJSON         `json:"-"`
}

// certificatePackOrderNewResponseJSON contains the JSON metadata for the struct
// [CertificatePackOrderNewResponse]
type certificatePackOrderNewResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CloudflareBranding   apijson.Field
	Hosts                apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	ValidationMethod     apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CertificatePackOrderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackOrderNewResponseJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackOrderNewResponseCertificateAuthority string

const (
	CertificatePackOrderNewResponseCertificateAuthorityGoogle      CertificatePackOrderNewResponseCertificateAuthority = "google"
	CertificatePackOrderNewResponseCertificateAuthorityLetsEncrypt CertificatePackOrderNewResponseCertificateAuthority = "lets_encrypt"
)

func (r CertificatePackOrderNewResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewResponseCertificateAuthorityGoogle, CertificatePackOrderNewResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

// Type of certificate pack.
type CertificatePackOrderNewResponseType string

const (
	CertificatePackOrderNewResponseTypeAdvanced CertificatePackOrderNewResponseType = "advanced"
)

func (r CertificatePackOrderNewResponseType) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewResponseTypeAdvanced:
		return true
	}
	return false
}

// Validation Method selected for the order.
type CertificatePackOrderNewResponseValidationMethod string

const (
	CertificatePackOrderNewResponseValidationMethodTXT   CertificatePackOrderNewResponseValidationMethod = "txt"
	CertificatePackOrderNewResponseValidationMethodHTTP  CertificatePackOrderNewResponseValidationMethod = "http"
	CertificatePackOrderNewResponseValidationMethodEmail CertificatePackOrderNewResponseValidationMethod = "email"
)

func (r CertificatePackOrderNewResponseValidationMethod) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewResponseValidationMethodTXT, CertificatePackOrderNewResponseValidationMethodHTTP, CertificatePackOrderNewResponseValidationMethodEmail:
		return true
	}
	return false
}

// Validity Days selected for the order.
type CertificatePackOrderNewResponseValidityDays int64

const (
	CertificatePackOrderNewResponseValidityDays14  CertificatePackOrderNewResponseValidityDays = 14
	CertificatePackOrderNewResponseValidityDays30  CertificatePackOrderNewResponseValidityDays = 30
	CertificatePackOrderNewResponseValidityDays90  CertificatePackOrderNewResponseValidityDays = 90
	CertificatePackOrderNewResponseValidityDays365 CertificatePackOrderNewResponseValidityDays = 365
)

func (r CertificatePackOrderNewResponseValidityDays) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewResponseValidityDays14, CertificatePackOrderNewResponseValidityDays30, CertificatePackOrderNewResponseValidityDays90, CertificatePackOrderNewResponseValidityDays365:
		return true
	}
	return false
}

type CertificatePackOrderNewParams struct {
	// Certificate Authority selected for the order. For information on any certificate
	// authority specific details or restrictions
	// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
	CertificateAuthority param.Field[CertificatePackOrderNewParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]HostParam] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[CertificatePackOrderNewParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[CertificatePackOrderNewParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[CertificatePackOrderNewParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r CertificatePackOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. For information on any certificate
// authority specific details or restrictions
// [see this page for more details.](https://developers.cloudflare.com/ssl/reference/certificate-authorities)
type CertificatePackOrderNewParamsCertificateAuthority string

const (
	CertificatePackOrderNewParamsCertificateAuthorityGoogle      CertificatePackOrderNewParamsCertificateAuthority = "google"
	CertificatePackOrderNewParamsCertificateAuthorityLetsEncrypt CertificatePackOrderNewParamsCertificateAuthority = "lets_encrypt"
)

func (r CertificatePackOrderNewParamsCertificateAuthority) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewParamsCertificateAuthorityGoogle, CertificatePackOrderNewParamsCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

// Type of certificate pack.
type CertificatePackOrderNewParamsType string

const (
	CertificatePackOrderNewParamsTypeAdvanced CertificatePackOrderNewParamsType = "advanced"
)

func (r CertificatePackOrderNewParamsType) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewParamsTypeAdvanced:
		return true
	}
	return false
}

// Validation Method selected for the order.
type CertificatePackOrderNewParamsValidationMethod string

const (
	CertificatePackOrderNewParamsValidationMethodTXT   CertificatePackOrderNewParamsValidationMethod = "txt"
	CertificatePackOrderNewParamsValidationMethodHTTP  CertificatePackOrderNewParamsValidationMethod = "http"
	CertificatePackOrderNewParamsValidationMethodEmail CertificatePackOrderNewParamsValidationMethod = "email"
)

func (r CertificatePackOrderNewParamsValidationMethod) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewParamsValidationMethodTXT, CertificatePackOrderNewParamsValidationMethodHTTP, CertificatePackOrderNewParamsValidationMethodEmail:
		return true
	}
	return false
}

// Validity Days selected for the order.
type CertificatePackOrderNewParamsValidityDays int64

const (
	CertificatePackOrderNewParamsValidityDays14  CertificatePackOrderNewParamsValidityDays = 14
	CertificatePackOrderNewParamsValidityDays30  CertificatePackOrderNewParamsValidityDays = 30
	CertificatePackOrderNewParamsValidityDays90  CertificatePackOrderNewParamsValidityDays = 90
	CertificatePackOrderNewParamsValidityDays365 CertificatePackOrderNewParamsValidityDays = 365
)

func (r CertificatePackOrderNewParamsValidityDays) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewParamsValidityDays14, CertificatePackOrderNewParamsValidityDays30, CertificatePackOrderNewParamsValidityDays90, CertificatePackOrderNewParamsValidityDays365:
		return true
	}
	return false
}

type CertificatePackOrderNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CertificatePackOrderNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CertificatePackOrderNewResponse                `json:"result"`
	JSON    certificatePackOrderNewResponseEnvelopeJSON    `json:"-"`
}

// certificatePackOrderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CertificatePackOrderNewResponseEnvelope]
type certificatePackOrderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificatePackOrderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatePackOrderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificatePackOrderNewResponseEnvelopeSuccess bool

const (
	CertificatePackOrderNewResponseEnvelopeSuccessTrue CertificatePackOrderNewResponseEnvelopeSuccess = true
)

func (r CertificatePackOrderNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CertificatePackOrderNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
