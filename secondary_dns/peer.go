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

// PeerService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPeerService] method instead.
type PeerService struct {
	Options []option.RequestOption
}

// NewPeerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPeerService(opts ...option.RequestOption) (r *PeerService) {
	r = &PeerService{}
	r.Options = opts
	return
}

// Create Peer.
func (r *PeerService) New(ctx context.Context, accountID string, body PeerNewParams, opts ...option.RequestOption) (res *Peer, err error) {
	var env PeerNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify Peer.
func (r *PeerService) Update(ctx context.Context, accountID string, peerID string, body PeerUpdateParams, opts ...option.RequestOption) (res *Peer, err error) {
	var env PeerUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", accountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Peers.
func (r *PeerService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *pagination.SinglePage[Peer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", accountID)
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

// List Peers.
func (r *PeerService) ListAutoPaging(ctx context.Context, accountID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Peer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, opts...))
}

// Delete Peer.
func (r *PeerService) Delete(ctx context.Context, accountID string, peerID string, opts ...option.RequestOption) (res *PeerDeleteResponse, err error) {
	var env PeerDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", accountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Peer.
func (r *PeerService) Get(ctx context.Context, accountID string, peerID string, opts ...option.RequestOption) (res *Peer, err error) {
	var env PeerGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", accountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Peer struct {
	ID string `json:"id,required"`
	// The name of the peer.
	Name string `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP string `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable bool `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port float64 `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TSIGID string   `json:"tsig_id"`
	JSON   peerJSON `json:"-"`
}

// peerJSON contains the JSON metadata for the struct [Peer]
type peerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TSIGID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Peer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerJSON) RawJSON() string {
	return r.raw
}

type PeerParam struct {
	// The name of the peer.
	Name param.Field[string] `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP param.Field[string] `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable param.Field[bool] `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port param.Field[float64] `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TSIGID param.Field[string] `json:"tsig_id"`
}

func (r PeerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PeerDeleteResponse struct {
	ID   string                 `json:"id"`
	JSON peerDeleteResponseJSON `json:"-"`
}

// peerDeleteResponseJSON contains the JSON metadata for the struct
// [PeerDeleteResponse]
type peerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PeerNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r PeerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PeerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PeerNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Peer                           `json:"result"`
	JSON    peerNewResponseEnvelopeJSON    `json:"-"`
}

// peerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerNewResponseEnvelope]
type peerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerNewResponseEnvelopeSuccess bool

const (
	PeerNewResponseEnvelopeSuccessTrue PeerNewResponseEnvelopeSuccess = true
)

func (r PeerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PeerUpdateParams struct {
	Peer PeerParam `json:"peer,required"`
}

func (r PeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Peer)
}

type PeerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PeerUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Peer                              `json:"result"`
	JSON    peerUpdateResponseEnvelopeJSON    `json:"-"`
}

// peerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerUpdateResponseEnvelope]
type peerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerUpdateResponseEnvelopeSuccess bool

const (
	PeerUpdateResponseEnvelopeSuccessTrue PeerUpdateResponseEnvelopeSuccess = true
)

func (r PeerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PeerDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PeerDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  PeerDeleteResponse                `json:"result"`
	JSON    peerDeleteResponseEnvelopeJSON    `json:"-"`
}

// peerDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerDeleteResponseEnvelope]
type peerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerDeleteResponseEnvelopeSuccess bool

const (
	PeerDeleteResponseEnvelopeSuccessTrue PeerDeleteResponseEnvelopeSuccess = true
)

func (r PeerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PeerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success PeerGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Peer                           `json:"result"`
	JSON    peerGetResponseEnvelopeJSON    `json:"-"`
}

// peerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PeerGetResponseEnvelope]
type peerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PeerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PeerGetResponseEnvelopeSuccess bool

const (
	PeerGetResponseEnvelopeSuccessTrue PeerGetResponseEnvelopeSuccess = true
)

func (r PeerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PeerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
