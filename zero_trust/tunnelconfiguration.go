// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// TunnelConfigurationService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelConfigurationService] method instead.
type TunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewTunnelConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelConfigurationService(opts ...option.RequestOption) (r *TunnelConfigurationService) {
	r = &TunnelConfigurationService{}
	r.Options = opts
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *TunnelConfigurationService) Update(ctx context.Context, accountID string, tunnelID string, body TunnelConfigurationUpdateParams, opts ...option.RequestOption) (res *TunnelConfigurationUpdateResponseUnion, err error) {
	var env TunnelConfigurationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *TunnelConfigurationService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelConfigurationGetResponseUnion, err error) {
	var env TunnelConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.TunnelConfigurationUpdateResponseUnknown],
// [zero_trust.TunnelConfigurationUpdateResponseArray] or [shared.UnionString].
type TunnelConfigurationUpdateResponseUnion interface {
	ImplementsZeroTrustTunnelConfigurationUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConfigurationUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelConfigurationUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConfigurationUpdateResponseArray []interface{}

func (r TunnelConfigurationUpdateResponseArray) ImplementsZeroTrustTunnelConfigurationUpdateResponseUnion() {
}

// Union satisfied by [zero_trust.TunnelConfigurationGetResponseUnknown],
// [zero_trust.TunnelConfigurationGetResponseArray] or [shared.UnionString].
type TunnelConfigurationGetResponseUnion interface {
	ImplementsZeroTrustTunnelConfigurationGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelConfigurationGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelConfigurationGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TunnelConfigurationGetResponseArray []interface{}

func (r TunnelConfigurationGetResponseArray) ImplementsZeroTrustTunnelConfigurationGetResponseUnion() {
}

type TunnelConfigurationUpdateParams struct {
	// The tunnel configuration and ingress rules.
	Config param.Field[TunnelConfigurationUpdateParamsConfig] `json:"config"`
}

func (r TunnelConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type TunnelConfigurationUpdateParamsConfig struct {
	// List of public hostname definitions
	Ingress param.Field[[]TunnelConfigurationUpdateParamsConfigIngress] `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[TunnelConfigurationUpdateParamsConfigOriginRequest] `json:"originRequest"`
	// Enable private network access from WARP users to private network routes
	WARPRouting param.Field[TunnelConfigurationUpdateParamsConfigWARPRouting] `json:"warp-routing"`
}

func (r TunnelConfigurationUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type TunnelConfigurationUpdateParamsConfigIngress struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[TunnelConfigurationUpdateParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r TunnelConfigurationUpdateParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationUpdateParamsConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r TunnelConfigurationUpdateParamsConfigIngressOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelConfigurationUpdateParamsConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[TunnelConfigurationUpdateParamsConfigOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r TunnelConfigurationUpdateParamsConfigOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelConfigurationUpdateParamsConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelConfigurationUpdateParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes
type TunnelConfigurationUpdateParamsConfigWARPRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TunnelConfigurationUpdateParamsConfigWARPRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   TunnelConfigurationUpdateResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// tunnelConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConfigurationUpdateResponseEnvelope]
type tunnelConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConfigurationUpdateResponseEnvelopeSuccess bool

const (
	TunnelConfigurationUpdateResponseEnvelopeSuccessTrue TunnelConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r TunnelConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors,required"`
	Messages []shared.ResponseInfo               `json:"messages,required"`
	Result   TunnelConfigurationGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success TunnelConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelConfigurationGetResponseEnvelope]
type tunnelConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelConfigurationGetResponseEnvelopeSuccess bool

const (
	TunnelConfigurationGetResponseEnvelopeSuccessTrue TunnelConfigurationGetResponseEnvelopeSuccess = true
)

func (r TunnelConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
