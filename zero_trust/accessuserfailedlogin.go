// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// AccessUserFailedLoginService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserFailedLoginService] method instead.
type AccessUserFailedLoginService struct {
	Options []option.RequestOption
}

// NewAccessUserFailedLoginService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessUserFailedLoginService(opts ...option.RequestOption) (r *AccessUserFailedLoginService) {
	r = &AccessUserFailedLoginService{}
	r.Options = opts
	return
}

// Get all failed login attempts for a single user.
func (r *AccessUserFailedLoginService) List(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) (res *pagination.SinglePage[AccessUserFailedLoginListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", accountID, userID)
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

// Get all failed login attempts for a single user.
func (r *AccessUserFailedLoginService) ListAutoPaging(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessUserFailedLoginListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, userID, opts...))
}

type AccessUserFailedLoginListResponse struct {
	Expiration int64                                 `json:"expiration"`
	Metadata   interface{}                           `json:"metadata"`
	JSON       accessUserFailedLoginListResponseJSON `json:"-"`
}

// accessUserFailedLoginListResponseJSON contains the JSON metadata for the struct
// [AccessUserFailedLoginListResponse]
type accessUserFailedLoginListResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserFailedLoginListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserFailedLoginListResponseJSON) RawJSON() string {
	return r.raw
}
