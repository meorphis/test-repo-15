// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// TestingService contains methods and other services that help with interacting
// with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestingService] method instead.
type TestingService struct {
	Options     []option.RequestOption
	Accounts    *TestingAccountService
	CreditCards *TestingCreditCardService
	Shipments   *TestingShipmentService
}

// NewTestingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestingService(opts ...option.RequestOption) (r *TestingService) {
	r = &TestingService{}
	r.Options = opts
	r.Accounts = NewTestingAccountService(opts...)
	r.CreditCards = NewTestingCreditCardService(opts...)
	r.Shipments = NewTestingShipmentService(opts...)
	return
}
