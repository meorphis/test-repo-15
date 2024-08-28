// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// PermissionGroupService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPermissionGroupService] method instead.
type PermissionGroupService struct {
	Options []option.RequestOption
}

// NewPermissionGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPermissionGroupService(opts ...option.RequestOption) (r *PermissionGroupService) {
	r = &PermissionGroupService{}
	r.Options = opts
	return
}

// List all the permissions groups for an account.
func (r *PermissionGroupService) List(ctx context.Context, accountID string, query PermissionGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[PermissionGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/permission_groups", accountID)
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

// List all the permissions groups for an account.
func (r *PermissionGroupService) ListAutoPaging(ctx context.Context, accountID string, query PermissionGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[PermissionGroupListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Get information about a specific permission group in an account.
func (r *PermissionGroupService) Get(ctx context.Context, accountID string, permissionGroupID string, opts ...option.RequestOption) (res *PermissionGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if permissionGroupID == "" {
		err = errors.New("missing required permission_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/permission_groups/%s", accountID, permissionGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PermissionGroupListResponse = interface{}

// A named group of permissions that map to a group of operations against
// resources.
type PermissionGroupGetResponse struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta interface{} `json:"meta"`
	// Name of the group.
	Name string                         `json:"name"`
	JSON permissionGroupGetResponseJSON `json:"-"`
}

// permissionGroupGetResponseJSON contains the JSON metadata for the struct
// [PermissionGroupGetResponse]
type permissionGroupGetResponseJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionGroupListParams struct {
	// ID of the permission group to be fetched.
	ID param.Field[string] `query:"id"`
	// Label of the permission group to be fetched.
	Label param.Field[string] `query:"label"`
	// Name of the permission group to be fetched.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [PermissionGroupListParams]'s query parameters as
// `url.Values`.
func (r PermissionGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
