// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pcaps

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// DownloadService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDownloadService] method instead.
type DownloadService struct {
	Options []option.RequestOption
}

// NewDownloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDownloadService(opts ...option.RequestOption) (r *DownloadService) {
	r = &DownloadService{}
	r.Options = opts
	return
}

// Download PCAP information into a file. Response is a binary PCAP file.
func (r *DownloadService) Get(ctx context.Context, accountID string, pcapID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.tcpdump.pcap")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pcapID == "" {
		err = errors.New("missing required pcap_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pcaps/%s/download", accountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
