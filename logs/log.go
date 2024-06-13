// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// LogService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options  []option.RequestOption
	Control  *ControlService
	RayID    *RayIDService
	Received *ReceivedService
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r *LogService) {
	r = &LogService{}
	r.Options = opts
	r.Control = NewControlService(opts...)
	r.RayID = NewRayIDService(opts...)
	r.Received = NewReceivedService(opts...)
	return
}
