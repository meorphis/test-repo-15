// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// SettingService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Returns a list of settings (and their details) that Page Rules can apply to
// matching requests.
func (r *SettingService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]SettingListResponse, err error) {
	var env SettingListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingListResponse = interface{}

type SettingListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Settings available for the zone.
	Result []SettingListResponse `json:"result,required"`
	// Whether the API call was successful
	Success SettingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingListResponseEnvelopeJSON    `json:"-"`
}

// settingListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingListResponseEnvelope]
type settingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingListResponseEnvelopeSuccess bool

const (
	SettingListResponseEnvelopeSuccessTrue SettingListResponseEnvelopeSuccess = true
)

func (r SettingListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
