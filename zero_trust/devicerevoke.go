// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// DeviceRevokeService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceRevokeService] method instead.
type DeviceRevokeService struct {
	Options []option.RequestOption
}

// NewDeviceRevokeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceRevokeService(opts ...option.RequestOption) (r *DeviceRevokeService) {
	r = &DeviceRevokeService{}
	r.Options = opts
	return
}

// Revokes a list of devices.
func (r *DeviceRevokeService) New(ctx context.Context, accountID string, body DeviceRevokeNewParams, opts ...option.RequestOption) (res *DeviceRevokeNewResponseUnion, err error) {
	var env DeviceRevokeNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/revoke", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.DeviceRevokeNewResponseUnknown] or
// [shared.UnionString].
type DeviceRevokeNewResponseUnion interface {
	ImplementsZeroTrustDeviceRevokeNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceRevokeNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DeviceRevokeNewParams struct {
	// A list of device ids to revoke.
	Body []string `json:"body,required"`
}

func (r DeviceRevokeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DeviceRevokeNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   DeviceRevokeNewResponseUnion `json:"result,required"`
	// Whether the API call was successful.
	Success DeviceRevokeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceRevokeNewResponseEnvelopeJSON    `json:"-"`
}

// deviceRevokeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceRevokeNewResponseEnvelope]
type deviceRevokeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRevokeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceRevokeNewResponseEnvelopeSuccess bool

const (
	DeviceRevokeNewResponseEnvelopeSuccessTrue DeviceRevokeNewResponseEnvelopeSuccess = true
)

func (r DeviceRevokeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceRevokeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
