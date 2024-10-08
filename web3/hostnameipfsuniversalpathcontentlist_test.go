// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/web3"
)

func TestHostnameIPFSUniversalPathContentListUpdate(t *testing.T) {
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
	_, err := client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		web3.HostnameIPFSUniversalPathContentListUpdateParams{
			Action: cloudflare.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsActionBlock),
			Entries: cloudflare.F([]web3.HostnameIPFSUniversalPathContentListUpdateParamsEntry{{
				Content:     cloudflare.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: cloudflare.F("this is my content list entry"),
				Type:        cloudflare.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid),
			}, {
				Content:     cloudflare.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: cloudflare.F("this is my content list entry"),
				Type:        cloudflare.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid),
			}, {
				Content:     cloudflare.F("QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB"),
				Description: cloudflare.F("this is my content list entry"),
				Type:        cloudflare.F(web3.HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid),
			}}),
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

func TestHostnameIPFSUniversalPathContentListGet(t *testing.T) {
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
	_, err := client.Web3.Hostnames.IPFSUniversalPaths.ContentLists.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
