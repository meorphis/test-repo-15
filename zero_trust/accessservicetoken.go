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

// AccessServiceTokenService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessServiceTokenService] method instead.
type AccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewAccessServiceTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessServiceTokenService(opts ...option.RequestOption) (r *AccessServiceTokenService) {
	r = &AccessServiceTokenService{}
	r.Options = opts
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *AccessServiceTokenService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessServiceTokenNewParams, opts ...option.RequestOption) (res *AccessServiceTokenNewResponse, err error) {
	var env AccessServiceTokenNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured service token.
func (r *AccessServiceTokenService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, serviceTokenID string, body AccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all service tokens.
func (r *AccessServiceTokenService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *pagination.SinglePage[ServiceToken], err error) {
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
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
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

// Lists all service tokens.
func (r *AccessServiceTokenService) ListAutoPaging(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ServiceToken] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountOrZone, accountOrZoneID, opts...))
}

// Deletes a service token.
func (r *AccessServiceTokenService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, serviceTokenID string, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single service token.
func (r *AccessServiceTokenService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, serviceTokenID string, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Refreshes the expiration of a service token.
func (r *AccessServiceTokenService) Refresh(ctx context.Context, accountID string, serviceTokenID string, opts ...option.RequestOption) (res *ServiceToken, err error) {
	var env AccessServiceTokenRefreshResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccessServiceTokenService) Rotate(ctx context.Context, accountID string, serviceTokenID string, opts ...option.RequestOption) (res *AccessServiceTokenRotateResponse, err error) {
	var env AccessServiceTokenRotateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceTokenID == "" {
		err = errors.New("missing required service_token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", accountID, serviceTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceToken struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string           `json:"name"`
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	JSON      serviceTokenJSON `json:"-"`
}

// serviceTokenJSON contains the JSON metadata for the struct [ServiceToken]
type serviceTokenJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceTokenJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewResponse struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                            `json:"name"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenNewResponseJSON `json:"-"`
}

// accessServiceTokenNewResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenNewResponse]
type accessServiceTokenNewResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenRotateResponse struct {
	// The ID of the service token.
	ID string `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                               `json:"name"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      accessServiceTokenRotateResponseJSON `json:"-"`
}

// accessServiceTokenRotateResponseJSON contains the JSON metadata for the struct
// [AccessServiceTokenRotateResponse]
type accessServiceTokenRotateResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessServiceTokenNewParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r AccessServiceTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessServiceTokenNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessServiceTokenNewResponse                `json:"result"`
	JSON    accessServiceTokenNewResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenNewResponseEnvelope]
type accessServiceTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessServiceTokenNewResponseEnvelopeSuccess bool

const (
	AccessServiceTokenNewResponseEnvelopeSuccessTrue AccessServiceTokenNewResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenUpdateParams struct {
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r AccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessServiceTokenUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessServiceTokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                    `json:"result"`
	JSON    accessServiceTokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenUpdateResponseEnvelope]
type accessServiceTokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessServiceTokenUpdateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenUpdateResponseEnvelopeSuccessTrue AccessServiceTokenUpdateResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessServiceTokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                    `json:"result"`
	JSON    accessServiceTokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenDeleteResponseEnvelope]
type accessServiceTokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessServiceTokenDeleteResponseEnvelopeSuccess bool

const (
	AccessServiceTokenDeleteResponseEnvelopeSuccessTrue AccessServiceTokenDeleteResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessServiceTokenGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                 `json:"result"`
	JSON    accessServiceTokenGetResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenGetResponseEnvelope]
type accessServiceTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessServiceTokenGetResponseEnvelopeSuccess bool

const (
	AccessServiceTokenGetResponseEnvelopeSuccessTrue AccessServiceTokenGetResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenRefreshResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessServiceTokenRefreshResponseEnvelopeSuccess `json:"success,required"`
	Result  ServiceToken                                     `json:"result"`
	JSON    accessServiceTokenRefreshResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRefreshResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenRefreshResponseEnvelope]
type accessServiceTokenRefreshResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRefreshResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRefreshResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessServiceTokenRefreshResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRefreshResponseEnvelopeSuccessTrue AccessServiceTokenRefreshResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenRefreshResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenRefreshResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessServiceTokenRotateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessServiceTokenRotateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessServiceTokenRotateResponse                `json:"result"`
	JSON    accessServiceTokenRotateResponseEnvelopeJSON    `json:"-"`
}

// accessServiceTokenRotateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessServiceTokenRotateResponseEnvelope]
type accessServiceTokenRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessServiceTokenRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessServiceTokenRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessServiceTokenRotateResponseEnvelopeSuccess bool

const (
	AccessServiceTokenRotateResponseEnvelopeSuccessTrue AccessServiceTokenRotateResponseEnvelopeSuccess = true
)

func (r AccessServiceTokenRotateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessServiceTokenRotateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
