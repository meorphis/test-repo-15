// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// NetworkVirtualNetworkService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkVirtualNetworkService] method instead.
type NetworkVirtualNetworkService struct {
	Options []option.RequestOption
}

// NewNetworkVirtualNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkVirtualNetworkService(opts ...option.RequestOption) (r *NetworkVirtualNetworkService) {
	r = &NetworkVirtualNetworkService{}
	r.Options = opts
	return
}

// Adds a new virtual network to an account.
func (r *NetworkVirtualNetworkService) New(ctx context.Context, accountID string, body NetworkVirtualNetworkNewParams, opts ...option.RequestOption) (res *NetworkVirtualNetworkNewResponseUnion, err error) {
	var env NetworkVirtualNetworkNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters virtual networks in an account.
func (r *NetworkVirtualNetworkService) List(ctx context.Context, accountID string, query NetworkVirtualNetworkListParams, opts ...option.RequestOption) (res *pagination.SinglePage[VirtualNetwork], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Lists and filters virtual networks in an account.
func (r *NetworkVirtualNetworkService) ListAutoPaging(ctx context.Context, accountID string, query NetworkVirtualNetworkListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[VirtualNetwork] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, query, opts...))
}

// Deletes an existing virtual network.
func (r *NetworkVirtualNetworkService) Delete(ctx context.Context, accountID string, virtualNetworkID string, opts ...option.RequestOption) (res *NetworkVirtualNetworkDeleteResponseUnion, err error) {
	var env NetworkVirtualNetworkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if virtualNetworkID == "" {
		err = errors.New("missing required virtual_network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing virtual network.
func (r *NetworkVirtualNetworkService) Edit(ctx context.Context, accountID string, virtualNetworkID string, body NetworkVirtualNetworkEditParams, opts ...option.RequestOption) (res *NetworkVirtualNetworkEditResponseUnion, err error) {
	var env NetworkVirtualNetworkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if virtualNetworkID == "" {
		err = errors.New("missing required virtual_network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VirtualNetwork struct {
	// UUID of the virtual network.
	ID string `json:"id,required" format:"uuid"`
	// Optional remark describing the virtual network.
	Comment string `json:"comment,required"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork bool `json:"is_default_network,required"`
	// A user-friendly name for the virtual network.
	Name string `json:"name,required"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time          `json:"deleted_at" format:"date-time"`
	JSON      virtualNetworkJSON `json:"-"`
}

// virtualNetworkJSON contains the JSON metadata for the struct [VirtualNetwork]
type virtualNetworkJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VirtualNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r virtualNetworkJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.NetworkVirtualNetworkNewResponseUnknown],
// [zero_trust.NetworkVirtualNetworkNewResponseArray] or [shared.UnionString].
type NetworkVirtualNetworkNewResponseUnion interface {
	ImplementsZeroTrustNetworkVirtualNetworkNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NetworkVirtualNetworkNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkVirtualNetworkNewResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NetworkVirtualNetworkNewResponseArray []interface{}

func (r NetworkVirtualNetworkNewResponseArray) ImplementsZeroTrustNetworkVirtualNetworkNewResponseUnion() {
}

// Union satisfied by [zero_trust.NetworkVirtualNetworkDeleteResponseUnknown],
// [zero_trust.NetworkVirtualNetworkDeleteResponseArray] or [shared.UnionString].
type NetworkVirtualNetworkDeleteResponseUnion interface {
	ImplementsZeroTrustNetworkVirtualNetworkDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NetworkVirtualNetworkDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkVirtualNetworkDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NetworkVirtualNetworkDeleteResponseArray []interface{}

func (r NetworkVirtualNetworkDeleteResponseArray) ImplementsZeroTrustNetworkVirtualNetworkDeleteResponseUnion() {
}

// Union satisfied by [zero_trust.NetworkVirtualNetworkEditResponseUnknown],
// [zero_trust.NetworkVirtualNetworkEditResponseArray] or [shared.UnionString].
type NetworkVirtualNetworkEditResponseUnion interface {
	ImplementsZeroTrustNetworkVirtualNetworkEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NetworkVirtualNetworkEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NetworkVirtualNetworkEditResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NetworkVirtualNetworkEditResponseArray []interface{}

func (r NetworkVirtualNetworkEditResponseArray) ImplementsZeroTrustNetworkVirtualNetworkEditResponseUnion() {
}

type NetworkVirtualNetworkNewParams struct {
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefault param.Field[bool] `json:"is_default"`
}

func (r NetworkVirtualNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkVirtualNetworkNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []shared.ResponseInfo                 `json:"messages,required"`
	Result   NetworkVirtualNetworkNewResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success NetworkVirtualNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkVirtualNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// networkVirtualNetworkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkVirtualNetworkNewResponseEnvelope]
type networkVirtualNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkNewResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkNewResponseEnvelopeSuccessTrue NetworkVirtualNetworkNewResponseEnvelopeSuccess = true
)

func (r NetworkVirtualNetworkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkVirtualNetworkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkVirtualNetworkListParams struct {
	// UUID of the virtual network.
	ID param.Field[string] `query:"id" format:"uuid"`
	// If `true`, only include the default virtual network. If `false`, exclude the
	// default virtual network. If empty, all virtual networks will be included.
	IsDefault param.Field[bool] `query:"is_default"`
	// If `true`, only include deleted virtual networks. If `false`, exclude deleted
	// virtual networks. If empty, all virtual networks will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [NetworkVirtualNetworkListParams]'s query parameters as
// `url.Values`.
func (r NetworkVirtualNetworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NetworkVirtualNetworkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                    `json:"errors,required"`
	Messages []shared.ResponseInfo                    `json:"messages,required"`
	Result   NetworkVirtualNetworkDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success NetworkVirtualNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkVirtualNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkVirtualNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [NetworkVirtualNetworkDeleteResponseEnvelope]
type networkVirtualNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkDeleteResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkDeleteResponseEnvelopeSuccessTrue NetworkVirtualNetworkDeleteResponseEnvelopeSuccess = true
)

func (r NetworkVirtualNetworkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkVirtualNetworkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkVirtualNetworkEditParams struct {
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name"`
}

func (r NetworkVirtualNetworkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkVirtualNetworkEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   NetworkVirtualNetworkEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success NetworkVirtualNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkVirtualNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// networkVirtualNetworkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkVirtualNetworkEditResponseEnvelope]
type networkVirtualNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkVirtualNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkVirtualNetworkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkVirtualNetworkEditResponseEnvelopeSuccess bool

const (
	NetworkVirtualNetworkEditResponseEnvelopeSuccessTrue NetworkVirtualNetworkEditResponseEnvelopeSuccess = true
)

func (r NetworkVirtualNetworkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkVirtualNetworkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
