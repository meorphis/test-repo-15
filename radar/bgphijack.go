// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// BGPHijackService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPHijackService] method instead.
type BGPHijackService struct {
	Options []option.RequestOption
	Events  *BGPHijackEventService
}

// NewBGPHijackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPHijackService(opts ...option.RequestOption) (r *BGPHijackService) {
	r = &BGPHijackService{}
	r.Options = opts
	r.Events = NewBGPHijackEventService(opts...)
	return
}