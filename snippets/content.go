// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// ContentService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentService] method instead.
type ContentService struct {
	Options []option.RequestOption
}

// NewContentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewContentService(opts ...option.RequestOption) (r *ContentService) {
	r = &ContentService{}
	r.Options = opts
	return
}

// Snippet Content
func (r *ContentService) Get(ctx context.Context, zoneID string, snippetName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "multipart/form-data")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s/content", zoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
