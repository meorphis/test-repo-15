// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// TunnelConnectionService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelConnectionService] method instead.
type TunnelConnectionService struct {
	Options []option.RequestOption
}

// NewTunnelConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelConnectionService(opts ...option.RequestOption) (r *TunnelConnectionService) {
	r = &TunnelConnectionService{}
	r.Options = opts
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *TunnelConnectionService) Delete(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelConnectionDeleteResponseUnion, err error) {
	var env TunnelConnectionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tunnels/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *TunnelConnectionService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *[]Client, err error) {
	var env TunnelConnectionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type Client struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []ClientConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string     `json:"version"`
	JSON    clientJSON `json:"-"`
}

// clientJSON contains the JSON metadata for the struct [Client]
type clientJSON struct {
	ID            apijson.Field
	Arch          apijson.Field
	ConfigVersion apijson.Field
	Conns         apijson.Field
	Features      apijson.Field
	RunAt         apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Client) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientJSON) RawJSON() string {
	return r.raw
}

type ClientConn struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string         `json:"uuid" format:"uuid"`
	JSON clientConnJSON `json:"-"`
}

// clientConnJSON contains the JSON metadata for the struct [ClientConn]
type clientConnJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ClientConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientConnJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.TunnelConnectionDeleteResponseUnknown],
// [zero_trust.TunnelConnectionDeleteResponseArray] or [shared.UnionString].
type TunnelConnectionDeleteResponseUnion interface {
	ImplementsZeroTrustTunnelConnectionDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConnectionDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelConnectionDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConnectionDeleteResponseArray []interface{}

func (r TunnelConnectionDeleteResponseArray) ImplementsZeroTrustTunnelConnectionDeleteResponseUnion() {
}

type TunnelConnectionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors,required"`
	Messages []shared.ResponseInfo               `json:"messages,required"`
	Result   TunnelConnectionDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConnectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConnectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelConnectionDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConnectionDeleteResponseEnvelope]
type tunnelConnectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConnectionDeleteResponseEnvelopeSuccess bool

const (
	TunnelConnectionDeleteResponseEnvelopeSuccessTrue TunnelConnectionDeleteResponseEnvelopeSuccess = true
)

func (r TunnelConnectionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConnectionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelConnectionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []Client              `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TunnelConnectionGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TunnelConnectionGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tunnelConnectionGetResponseEnvelopeJSON       `json:"-"`
}

// tunnelConnectionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConnectionGetResponseEnvelope]
type tunnelConnectionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConnectionGetResponseEnvelopeSuccess bool

const (
	TunnelConnectionGetResponseEnvelopeSuccessTrue TunnelConnectionGetResponseEnvelopeSuccess = true
)

func (r TunnelConnectionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConnectionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelConnectionGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       tunnelConnectionGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// tunnelConnectionGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [TunnelConnectionGetResponseEnvelopeResultInfo]
type tunnelConnectionGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectionGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}