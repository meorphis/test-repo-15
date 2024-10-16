// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// RegistrarService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistrarService] method instead.
type RegistrarService struct {
	Options []option.RequestOption
	Domains *DomainService
}

// NewRegistrarService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistrarService(opts ...option.RequestOption) (r *RegistrarService) {
	r = &RegistrarService{}
	r.Options = opts
	r.Domains = NewDomainService(opts...)
	return
}