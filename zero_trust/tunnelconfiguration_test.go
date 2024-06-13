// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/zero_trust"
)

func TestTunnelConfigurationUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Tunnels.Configurations.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
		zero_trust.TunnelConfigurationUpdateParams{
			Config: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfig{
				Ingress: cloudflare.F([]zero_trust.TunnelConfigurationUpdateParamsConfigIngress{{
					Hostname: cloudflare.F("tunnel.example.com"),
					Service:  cloudflare.F("https://localhost:8001"),
					OriginRequest: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequest{
						Access: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess{
							AUDTag:   cloudflare.F([]string{"string", "string", "string"}),
							TeamName: cloudflare.F("teamName"),
							Required: cloudflare.F(true),
						}),
						CAPool:                 cloudflare.F("caPool"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("httpHostHeader"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTLSVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("originServerName"),
						ProxyType:              cloudflare.F("proxyType"),
						TCPKeepAlive:           cloudflare.F(int64(0)),
						TLSTimeout:             cloudflare.F(int64(0)),
					}),
					Path: cloudflare.F("subpath"),
				}, {
					Hostname: cloudflare.F("tunnel.example.com"),
					Service:  cloudflare.F("https://localhost:8001"),
					OriginRequest: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequest{
						Access: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess{
							AUDTag:   cloudflare.F([]string{"string", "string", "string"}),
							TeamName: cloudflare.F("teamName"),
							Required: cloudflare.F(true),
						}),
						CAPool:                 cloudflare.F("caPool"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("httpHostHeader"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTLSVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("originServerName"),
						ProxyType:              cloudflare.F("proxyType"),
						TCPKeepAlive:           cloudflare.F(int64(0)),
						TLSTimeout:             cloudflare.F(int64(0)),
					}),
					Path: cloudflare.F("subpath"),
				}, {
					Hostname: cloudflare.F("tunnel.example.com"),
					Service:  cloudflare.F("https://localhost:8001"),
					OriginRequest: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequest{
						Access: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigIngressOriginRequestAccess{
							AUDTag:   cloudflare.F([]string{"string", "string", "string"}),
							TeamName: cloudflare.F("teamName"),
							Required: cloudflare.F(true),
						}),
						CAPool:                 cloudflare.F("caPool"),
						ConnectTimeout:         cloudflare.F(int64(0)),
						DisableChunkedEncoding: cloudflare.F(true),
						HTTP2Origin:            cloudflare.F(true),
						HTTPHostHeader:         cloudflare.F("httpHostHeader"),
						KeepAliveConnections:   cloudflare.F(int64(0)),
						KeepAliveTimeout:       cloudflare.F(int64(0)),
						NoHappyEyeballs:        cloudflare.F(true),
						NoTLSVerify:            cloudflare.F(true),
						OriginServerName:       cloudflare.F("originServerName"),
						ProxyType:              cloudflare.F("proxyType"),
						TCPKeepAlive:           cloudflare.F(int64(0)),
						TLSTimeout:             cloudflare.F(int64(0)),
					}),
					Path: cloudflare.F("subpath"),
				}}),
				OriginRequest: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigOriginRequest{
					Access: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigOriginRequestAccess{
						AUDTag:   cloudflare.F([]string{"string", "string", "string"}),
						TeamName: cloudflare.F("teamName"),
						Required: cloudflare.F(true),
					}),
					CAPool:                 cloudflare.F("caPool"),
					ConnectTimeout:         cloudflare.F(int64(0)),
					DisableChunkedEncoding: cloudflare.F(true),
					HTTP2Origin:            cloudflare.F(true),
					HTTPHostHeader:         cloudflare.F("httpHostHeader"),
					KeepAliveConnections:   cloudflare.F(int64(0)),
					KeepAliveTimeout:       cloudflare.F(int64(0)),
					NoHappyEyeballs:        cloudflare.F(true),
					NoTLSVerify:            cloudflare.F(true),
					OriginServerName:       cloudflare.F("originServerName"),
					ProxyType:              cloudflare.F("proxyType"),
					TCPKeepAlive:           cloudflare.F(int64(0)),
					TLSTimeout:             cloudflare.F(int64(0)),
				}),
				WARPRouting: cloudflare.F(zero_trust.TunnelConfigurationUpdateParamsConfigWARPRouting{
					Enabled: cloudflare.F(true),
				}),
			}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTunnelConfigurationGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Tunnels.Configurations.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"f70ff985-a4ef-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
