// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// TunnelConnectorService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelConnectorService] method instead.
type TunnelConnectorService struct {
	Options []option.RequestOption
}

// NewTunnelConnectorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTunnelConnectorService(opts ...option.RequestOption) (r *TunnelConnectorService) {
	r = &TunnelConnectorService{}
	r.Options = opts
	return
}

// Fetches connector and connection details for a Cloudflare Tunnel.
func (r *TunnelConnectorService) Get(ctx context.Context, accountID string, tunnelID string, connectorID string, opts ...option.RequestOption) (res *Client, err error) {
	var env TunnelConnectorGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connectors/%s", accountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelConnectorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A client (typically cloudflared) that maintains connections to a Cloudflare data
	// center.
	Result Client `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelConnectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelConnectorGetResponseEnvelope]
type tunnelConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConnectorGetResponseEnvelopeSuccess bool

const (
	TunnelConnectorGetResponseEnvelopeSuccessTrue TunnelConnectorGetResponseEnvelopeSuccess = true
)

func (r TunnelConnectorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConnectorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
