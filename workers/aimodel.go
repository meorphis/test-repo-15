// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// AIModelService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelService] method instead.
type AIModelService struct {
	Options []option.RequestOption
	Schema  *AIModelSchemaService
}

// NewAIModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIModelService(opts ...option.RequestOption) (r *AIModelService) {
	r = &AIModelService{}
	r.Options = opts
	r.Schema = NewAIModelSchemaService(opts...)
	return
}
