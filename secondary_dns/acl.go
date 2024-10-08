// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// ACLService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewACLService] method instead.
type ACLService struct {
	Options []option.RequestOption
}

// NewACLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewACLService(opts ...option.RequestOption) (r *ACLService) {
	r = &ACLService{}
	r.Options = opts
	return
}

// Create ACL.
func (r *ACLService) New(ctx context.Context, accountID string, body ACLNewParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env ACLNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify ACL.
func (r *ACLService) Update(ctx context.Context, accountID string, aclID string, body ACLUpdateParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env ACLUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", accountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List ACLs.
func (r *ACLService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *pagination.SinglePage[ACL], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", accountID)
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

// List ACLs.
func (r *ACLService) ListAutoPaging(ctx context.Context, accountID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ACL] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, opts...))
}

// Delete ACL.
func (r *ACLService) Delete(ctx context.Context, accountID string, aclID string, opts ...option.RequestOption) (res *ACLDeleteResponse, err error) {
	var env ACLDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", accountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get ACL.
func (r *ACLService) Get(ctx context.Context, accountID string, aclID string, opts ...option.RequestOption) (res *ACL, err error) {
	var env ACLGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", accountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ACL struct {
	ID string `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string  `json:"name,required"`
	JSON aclJSON `json:"-"`
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

type ACLParam struct {
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r ACLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ACLDeleteResponse struct {
	ID   string                `json:"id"`
	JSON aclDeleteResponseJSON `json:"-"`
}

// aclDeleteResponseJSON contains the JSON metadata for the struct
// [ACLDeleteResponse]
type aclDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ACLNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r ACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ACLNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ACLNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ACL                           `json:"result"`
	JSON    aclNewResponseEnvelopeJSON    `json:"-"`
}

// aclNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLNewResponseEnvelope]
type aclNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLNewResponseEnvelopeSuccess bool

const (
	ACLNewResponseEnvelopeSuccessTrue ACLNewResponseEnvelopeSuccess = true
)

func (r ACLNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ACLUpdateParams struct {
	ACL ACLParam `json:"acl,required"`
}

func (r ACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ACL)
}

type ACLUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ACLUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ACL                              `json:"result"`
	JSON    aclUpdateResponseEnvelopeJSON    `json:"-"`
}

// aclUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLUpdateResponseEnvelope]
type aclUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLUpdateResponseEnvelopeSuccess bool

const (
	ACLUpdateResponseEnvelopeSuccessTrue ACLUpdateResponseEnvelopeSuccess = true
)

func (r ACLUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ACLDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ACLDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ACLDeleteResponse                `json:"result"`
	JSON    aclDeleteResponseEnvelopeJSON    `json:"-"`
}

// aclDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLDeleteResponseEnvelope]
type aclDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLDeleteResponseEnvelopeSuccess bool

const (
	ACLDeleteResponseEnvelopeSuccessTrue ACLDeleteResponseEnvelopeSuccess = true
)

func (r ACLDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ACLGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ACLGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ACL                           `json:"result"`
	JSON    aclGetResponseEnvelopeJSON    `json:"-"`
}

// aclGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ACLGetResponseEnvelope]
type aclGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ACLGetResponseEnvelopeSuccess bool

const (
	ACLGetResponseEnvelopeSuccessTrue ACLGetResponseEnvelopeSuccess = true
)

func (r ACLGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ACLGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
