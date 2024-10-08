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

// GatewayCertificateService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayCertificateService] method instead.
type GatewayCertificateService struct {
	Options []option.RequestOption
}

// NewGatewayCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayCertificateService(opts ...option.RequestOption) (r *GatewayCertificateService) {
	r = &GatewayCertificateService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust certificate.
func (r *GatewayCertificateService) New(ctx context.Context, accountID string, body GatewayCertificateNewParams, opts ...option.RequestOption) (res *GatewayCertificateNewResponse, err error) {
	var env GatewayCertificateNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust certificates for an account.
func (r *GatewayCertificateService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayCertificateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates", accountID)
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

// Fetches all Zero Trust certificates for an account.
func (r *GatewayCertificateService) ListAutoPaging(ctx context.Context, accountID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayCertificateListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, opts...))
}

// Deletes a gateway-managed Zero Trust certificate. A certificate must be
// deactivated from the edge (inactive) before it is deleted.
func (r *GatewayCertificateService) Delete(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *GatewayCertificateDeleteResponse, err error) {
	var env GatewayCertificateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Binds a single Zero Trust certificate to the edge.
func (r *GatewayCertificateService) Activate(ctx context.Context, accountID string, certificateID string, body GatewayCertificateActivateParams, opts ...option.RequestOption) (res *GatewayCertificateActivateResponse, err error) {
	var env GatewayCertificateActivateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s/activate", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Unbinds a single Zero Trust certificate from the edge
func (r *GatewayCertificateService) Deactivate(ctx context.Context, accountID string, certificateID string, body GatewayCertificateDeactivateParams, opts ...option.RequestOption) (res *GatewayCertificateDeactivateResponse, err error) {
	var env GatewayCertificateDeactivateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s/deactivate", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust certificate.
func (r *GatewayCertificateService) Get(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *GatewayCertificateGetResponse, err error) {
	var env GatewayCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayCertificateNewResponse struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'active' state may be used for Gateway TLS interception.
	BindingStatus GatewayCertificateNewResponseBindingStatus `json:"binding_status"`
	CreatedAt     time.Time                                  `json:"created_at" format:"date-time"`
	// Use this certificate for Gateway TLS interception
	Enabled   bool      `json:"enabled"`
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       GatewayCertificateNewResponseType `json:"type"`
	UpdatedAt  time.Time                         `json:"updated_at" format:"date-time"`
	UploadedOn time.Time                         `json:"uploaded_on" format:"date-time"`
	JSON       gatewayCertificateNewResponseJSON `json:"-"`
}

// gatewayCertificateNewResponseJSON contains the JSON metadata for the struct
// [GatewayCertificateNewResponse]
type gatewayCertificateNewResponseJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	ExpiresOn     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'active' state may be used for Gateway TLS interception.
type GatewayCertificateNewResponseBindingStatus string

const (
	GatewayCertificateNewResponseBindingStatusPendingDeployment GatewayCertificateNewResponseBindingStatus = "pending_deployment"
	GatewayCertificateNewResponseBindingStatusActive            GatewayCertificateNewResponseBindingStatus = "active"
	GatewayCertificateNewResponseBindingStatusPendingDeletion   GatewayCertificateNewResponseBindingStatus = "pending_deletion"
	GatewayCertificateNewResponseBindingStatusInactive          GatewayCertificateNewResponseBindingStatus = "inactive"
)

func (r GatewayCertificateNewResponseBindingStatus) IsKnown() bool {
	switch r {
	case GatewayCertificateNewResponseBindingStatusPendingDeployment, GatewayCertificateNewResponseBindingStatusActive, GatewayCertificateNewResponseBindingStatusPendingDeletion, GatewayCertificateNewResponseBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type GatewayCertificateNewResponseType string

const (
	GatewayCertificateNewResponseTypeCustom         GatewayCertificateNewResponseType = "custom"
	GatewayCertificateNewResponseTypeGatewayManaged GatewayCertificateNewResponseType = "gateway_managed"
)

func (r GatewayCertificateNewResponseType) IsKnown() bool {
	switch r {
	case GatewayCertificateNewResponseTypeCustom, GatewayCertificateNewResponseTypeGatewayManaged:
		return true
	}
	return false
}

type GatewayCertificateListResponse struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'active' state may be used for Gateway TLS interception.
	BindingStatus GatewayCertificateListResponseBindingStatus `json:"binding_status"`
	CreatedAt     time.Time                                   `json:"created_at" format:"date-time"`
	// Use this certificate for Gateway TLS interception
	Enabled   bool      `json:"enabled"`
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       GatewayCertificateListResponseType `json:"type"`
	UpdatedAt  time.Time                          `json:"updated_at" format:"date-time"`
	UploadedOn time.Time                          `json:"uploaded_on" format:"date-time"`
	JSON       gatewayCertificateListResponseJSON `json:"-"`
}

// gatewayCertificateListResponseJSON contains the JSON metadata for the struct
// [GatewayCertificateListResponse]
type gatewayCertificateListResponseJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	ExpiresOn     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'active' state may be used for Gateway TLS interception.
type GatewayCertificateListResponseBindingStatus string

const (
	GatewayCertificateListResponseBindingStatusPendingDeployment GatewayCertificateListResponseBindingStatus = "pending_deployment"
	GatewayCertificateListResponseBindingStatusActive            GatewayCertificateListResponseBindingStatus = "active"
	GatewayCertificateListResponseBindingStatusPendingDeletion   GatewayCertificateListResponseBindingStatus = "pending_deletion"
	GatewayCertificateListResponseBindingStatusInactive          GatewayCertificateListResponseBindingStatus = "inactive"
)

func (r GatewayCertificateListResponseBindingStatus) IsKnown() bool {
	switch r {
	case GatewayCertificateListResponseBindingStatusPendingDeployment, GatewayCertificateListResponseBindingStatusActive, GatewayCertificateListResponseBindingStatusPendingDeletion, GatewayCertificateListResponseBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type GatewayCertificateListResponseType string

const (
	GatewayCertificateListResponseTypeCustom         GatewayCertificateListResponseType = "custom"
	GatewayCertificateListResponseTypeGatewayManaged GatewayCertificateListResponseType = "gateway_managed"
)

func (r GatewayCertificateListResponseType) IsKnown() bool {
	switch r {
	case GatewayCertificateListResponseTypeCustom, GatewayCertificateListResponseTypeGatewayManaged:
		return true
	}
	return false
}

type GatewayCertificateDeleteResponse struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'active' state may be used for Gateway TLS interception.
	BindingStatus GatewayCertificateDeleteResponseBindingStatus `json:"binding_status"`
	CreatedAt     time.Time                                     `json:"created_at" format:"date-time"`
	// Use this certificate for Gateway TLS interception
	Enabled   bool      `json:"enabled"`
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       GatewayCertificateDeleteResponseType `json:"type"`
	UpdatedAt  time.Time                            `json:"updated_at" format:"date-time"`
	UploadedOn time.Time                            `json:"uploaded_on" format:"date-time"`
	JSON       gatewayCertificateDeleteResponseJSON `json:"-"`
}

// gatewayCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [GatewayCertificateDeleteResponse]
type gatewayCertificateDeleteResponseJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	ExpiresOn     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'active' state may be used for Gateway TLS interception.
type GatewayCertificateDeleteResponseBindingStatus string

const (
	GatewayCertificateDeleteResponseBindingStatusPendingDeployment GatewayCertificateDeleteResponseBindingStatus = "pending_deployment"
	GatewayCertificateDeleteResponseBindingStatusActive            GatewayCertificateDeleteResponseBindingStatus = "active"
	GatewayCertificateDeleteResponseBindingStatusPendingDeletion   GatewayCertificateDeleteResponseBindingStatus = "pending_deletion"
	GatewayCertificateDeleteResponseBindingStatusInactive          GatewayCertificateDeleteResponseBindingStatus = "inactive"
)

func (r GatewayCertificateDeleteResponseBindingStatus) IsKnown() bool {
	switch r {
	case GatewayCertificateDeleteResponseBindingStatusPendingDeployment, GatewayCertificateDeleteResponseBindingStatusActive, GatewayCertificateDeleteResponseBindingStatusPendingDeletion, GatewayCertificateDeleteResponseBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type GatewayCertificateDeleteResponseType string

const (
	GatewayCertificateDeleteResponseTypeCustom         GatewayCertificateDeleteResponseType = "custom"
	GatewayCertificateDeleteResponseTypeGatewayManaged GatewayCertificateDeleteResponseType = "gateway_managed"
)

func (r GatewayCertificateDeleteResponseType) IsKnown() bool {
	switch r {
	case GatewayCertificateDeleteResponseTypeCustom, GatewayCertificateDeleteResponseTypeGatewayManaged:
		return true
	}
	return false
}

type GatewayCertificateActivateResponse struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'active' state may be used for Gateway TLS interception.
	BindingStatus GatewayCertificateActivateResponseBindingStatus `json:"binding_status"`
	CreatedAt     time.Time                                       `json:"created_at" format:"date-time"`
	// Use this certificate for Gateway TLS interception
	Enabled   bool      `json:"enabled"`
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       GatewayCertificateActivateResponseType `json:"type"`
	UpdatedAt  time.Time                              `json:"updated_at" format:"date-time"`
	UploadedOn time.Time                              `json:"uploaded_on" format:"date-time"`
	JSON       gatewayCertificateActivateResponseJSON `json:"-"`
}

// gatewayCertificateActivateResponseJSON contains the JSON metadata for the struct
// [GatewayCertificateActivateResponse]
type gatewayCertificateActivateResponseJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	ExpiresOn     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCertificateActivateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateActivateResponseJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'active' state may be used for Gateway TLS interception.
type GatewayCertificateActivateResponseBindingStatus string

const (
	GatewayCertificateActivateResponseBindingStatusPendingDeployment GatewayCertificateActivateResponseBindingStatus = "pending_deployment"
	GatewayCertificateActivateResponseBindingStatusActive            GatewayCertificateActivateResponseBindingStatus = "active"
	GatewayCertificateActivateResponseBindingStatusPendingDeletion   GatewayCertificateActivateResponseBindingStatus = "pending_deletion"
	GatewayCertificateActivateResponseBindingStatusInactive          GatewayCertificateActivateResponseBindingStatus = "inactive"
)

func (r GatewayCertificateActivateResponseBindingStatus) IsKnown() bool {
	switch r {
	case GatewayCertificateActivateResponseBindingStatusPendingDeployment, GatewayCertificateActivateResponseBindingStatusActive, GatewayCertificateActivateResponseBindingStatusPendingDeletion, GatewayCertificateActivateResponseBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type GatewayCertificateActivateResponseType string

const (
	GatewayCertificateActivateResponseTypeCustom         GatewayCertificateActivateResponseType = "custom"
	GatewayCertificateActivateResponseTypeGatewayManaged GatewayCertificateActivateResponseType = "gateway_managed"
)

func (r GatewayCertificateActivateResponseType) IsKnown() bool {
	switch r {
	case GatewayCertificateActivateResponseTypeCustom, GatewayCertificateActivateResponseTypeGatewayManaged:
		return true
	}
	return false
}

type GatewayCertificateDeactivateResponse struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'active' state may be used for Gateway TLS interception.
	BindingStatus GatewayCertificateDeactivateResponseBindingStatus `json:"binding_status"`
	CreatedAt     time.Time                                         `json:"created_at" format:"date-time"`
	// Use this certificate for Gateway TLS interception
	Enabled   bool      `json:"enabled"`
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       GatewayCertificateDeactivateResponseType `json:"type"`
	UpdatedAt  time.Time                                `json:"updated_at" format:"date-time"`
	UploadedOn time.Time                                `json:"uploaded_on" format:"date-time"`
	JSON       gatewayCertificateDeactivateResponseJSON `json:"-"`
}

// gatewayCertificateDeactivateResponseJSON contains the JSON metadata for the
// struct [GatewayCertificateDeactivateResponse]
type gatewayCertificateDeactivateResponseJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	ExpiresOn     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCertificateDeactivateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateDeactivateResponseJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'active' state may be used for Gateway TLS interception.
type GatewayCertificateDeactivateResponseBindingStatus string

const (
	GatewayCertificateDeactivateResponseBindingStatusPendingDeployment GatewayCertificateDeactivateResponseBindingStatus = "pending_deployment"
	GatewayCertificateDeactivateResponseBindingStatusActive            GatewayCertificateDeactivateResponseBindingStatus = "active"
	GatewayCertificateDeactivateResponseBindingStatusPendingDeletion   GatewayCertificateDeactivateResponseBindingStatus = "pending_deletion"
	GatewayCertificateDeactivateResponseBindingStatusInactive          GatewayCertificateDeactivateResponseBindingStatus = "inactive"
)

func (r GatewayCertificateDeactivateResponseBindingStatus) IsKnown() bool {
	switch r {
	case GatewayCertificateDeactivateResponseBindingStatusPendingDeployment, GatewayCertificateDeactivateResponseBindingStatusActive, GatewayCertificateDeactivateResponseBindingStatusPendingDeletion, GatewayCertificateDeactivateResponseBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type GatewayCertificateDeactivateResponseType string

const (
	GatewayCertificateDeactivateResponseTypeCustom         GatewayCertificateDeactivateResponseType = "custom"
	GatewayCertificateDeactivateResponseTypeGatewayManaged GatewayCertificateDeactivateResponseType = "gateway_managed"
)

func (r GatewayCertificateDeactivateResponseType) IsKnown() bool {
	switch r {
	case GatewayCertificateDeactivateResponseTypeCustom, GatewayCertificateDeactivateResponseTypeGatewayManaged:
		return true
	}
	return false
}

type GatewayCertificateGetResponse struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'active' state may be used for Gateway TLS interception.
	BindingStatus GatewayCertificateGetResponseBindingStatus `json:"binding_status"`
	CreatedAt     time.Time                                  `json:"created_at" format:"date-time"`
	// Use this certificate for Gateway TLS interception
	Enabled   bool      `json:"enabled"`
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       GatewayCertificateGetResponseType `json:"type"`
	UpdatedAt  time.Time                         `json:"updated_at" format:"date-time"`
	UploadedOn time.Time                         `json:"uploaded_on" format:"date-time"`
	JSON       gatewayCertificateGetResponseJSON `json:"-"`
}

// gatewayCertificateGetResponseJSON contains the JSON metadata for the struct
// [GatewayCertificateGetResponse]
type gatewayCertificateGetResponseJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	ExpiresOn     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateGetResponseJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'active' state may be used for Gateway TLS interception.
type GatewayCertificateGetResponseBindingStatus string

const (
	GatewayCertificateGetResponseBindingStatusPendingDeployment GatewayCertificateGetResponseBindingStatus = "pending_deployment"
	GatewayCertificateGetResponseBindingStatusActive            GatewayCertificateGetResponseBindingStatus = "active"
	GatewayCertificateGetResponseBindingStatusPendingDeletion   GatewayCertificateGetResponseBindingStatus = "pending_deletion"
	GatewayCertificateGetResponseBindingStatusInactive          GatewayCertificateGetResponseBindingStatus = "inactive"
)

func (r GatewayCertificateGetResponseBindingStatus) IsKnown() bool {
	switch r {
	case GatewayCertificateGetResponseBindingStatusPendingDeployment, GatewayCertificateGetResponseBindingStatusActive, GatewayCertificateGetResponseBindingStatusPendingDeletion, GatewayCertificateGetResponseBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type GatewayCertificateGetResponseType string

const (
	GatewayCertificateGetResponseTypeCustom         GatewayCertificateGetResponseType = "custom"
	GatewayCertificateGetResponseTypeGatewayManaged GatewayCertificateGetResponseType = "gateway_managed"
)

func (r GatewayCertificateGetResponseType) IsKnown() bool {
	switch r {
	case GatewayCertificateGetResponseTypeCustom, GatewayCertificateGetResponseTypeGatewayManaged:
		return true
	}
	return false
}

type GatewayCertificateNewParams struct {
	// Number of days the generated certificate will be valid, minimum 1 day and
	// maximum 30 years. Defaults to 5 years.
	ValidityPeriodDays param.Field[int64] `json:"validity_period_days"`
}

func (r GatewayCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayCertificateNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayCertificateNewResponse                `json:"result"`
	JSON    gatewayCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayCertificateNewResponseEnvelope]
type gatewayCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayCertificateNewResponseEnvelopeSuccess bool

const (
	GatewayCertificateNewResponseEnvelopeSuccessTrue GatewayCertificateNewResponseEnvelopeSuccess = true
)

func (r GatewayCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayCertificateDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayCertificateDeleteResponse                `json:"result"`
	JSON    gatewayCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayCertificateDeleteResponseEnvelope]
type gatewayCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayCertificateDeleteResponseEnvelopeSuccess bool

const (
	GatewayCertificateDeleteResponseEnvelopeSuccessTrue GatewayCertificateDeleteResponseEnvelopeSuccess = true
)

func (r GatewayCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayCertificateActivateParams struct {
	Body interface{} `json:"body,required"`
}

func (r GatewayCertificateActivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GatewayCertificateActivateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayCertificateActivateResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayCertificateActivateResponse                `json:"result"`
	JSON    gatewayCertificateActivateResponseEnvelopeJSON    `json:"-"`
}

// gatewayCertificateActivateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayCertificateActivateResponseEnvelope]
type gatewayCertificateActivateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCertificateActivateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateActivateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayCertificateActivateResponseEnvelopeSuccess bool

const (
	GatewayCertificateActivateResponseEnvelopeSuccessTrue GatewayCertificateActivateResponseEnvelopeSuccess = true
)

func (r GatewayCertificateActivateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayCertificateActivateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayCertificateDeactivateParams struct {
	Body interface{} `json:"body,required"`
}

func (r GatewayCertificateDeactivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GatewayCertificateDeactivateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayCertificateDeactivateResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayCertificateDeactivateResponse                `json:"result"`
	JSON    gatewayCertificateDeactivateResponseEnvelopeJSON    `json:"-"`
}

// gatewayCertificateDeactivateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayCertificateDeactivateResponseEnvelope]
type gatewayCertificateDeactivateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCertificateDeactivateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateDeactivateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayCertificateDeactivateResponseEnvelopeSuccess bool

const (
	GatewayCertificateDeactivateResponseEnvelopeSuccessTrue GatewayCertificateDeactivateResponseEnvelopeSuccess = true
)

func (r GatewayCertificateDeactivateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayCertificateDeactivateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayCertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayCertificateGetResponse                `json:"result"`
	JSON    gatewayCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayCertificateGetResponseEnvelope]
type gatewayCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayCertificateGetResponseEnvelopeSuccess bool

const (
	GatewayCertificateGetResponseEnvelopeSuccessTrue GatewayCertificateGetResponseEnvelopeSuccess = true
)

func (r GatewayCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
