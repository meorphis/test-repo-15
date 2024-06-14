// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// MerchantService contains methods and other services that help with interacting
// with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMerchantService] method instead.
type MerchantService struct {
	Options     []option.RequestOption
	Callbacks   *MerchantCallbackService
	Identifiers *MerchantIdentifierService
}

// NewMerchantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMerchantService(opts ...option.RequestOption) (r *MerchantService) {
	r = &MerchantService{}
	r.Options = opts
	r.Callbacks = NewMerchantCallbackService(opts...)
	r.Identifiers = NewMerchantIdentifierService(opts...)
	return
}
