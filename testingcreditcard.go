// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// TestingCreditCardService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestingCreditCardService] method instead.
type TestingCreditCardService struct {
	Options []option.RequestOption
}

// NewTestingCreditCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTestingCreditCardService(opts ...option.RequestOption) (r *TestingCreditCardService) {
	r = &TestingCreditCardService{}
	r.Options = opts
	return
}

// Retrieve test credit card information. This includes its token, which is
// generated against the `4111 1111 1111 1004` test card.
func (r *TestingCreditCardService) List(ctx context.Context, opts ...option.RequestOption) (res *TestingCreditCardListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "testing/credit-cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TestingCreditCardListResponse struct {
	// The expiration date of the credit card. TODO TO MAKE EXPIRATION REUSABLE
	Expiration string `json:"expiration,required" format:"^\\d{4}-\\d{2}$"`
	// The last 4 digits of the credit card number.
	Last4 string `json:"last4,required" format:"^\\d{4}$"`
	// The credit card network.
	Network TestingCreditCardListResponseNetwork `json:"network,required"`
	JSON    testingCreditCardListResponseJSON    `json:"-"`
}

// testingCreditCardListResponseJSON contains the JSON metadata for the struct
// [TestingCreditCardListResponse]
type testingCreditCardListResponseJSON struct {
	Expiration  apijson.Field
	Last4       apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestingCreditCardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testingCreditCardListResponseJSON) RawJSON() string {
	return r.raw
}

// The credit card network.
type TestingCreditCardListResponseNetwork string

const (
	TestingCreditCardListResponseNetworkVisa         TestingCreditCardListResponseNetwork = "visa"
	TestingCreditCardListResponseNetworkMastercard   TestingCreditCardListResponseNetwork = "mastercard"
	TestingCreditCardListResponseNetworkAmex         TestingCreditCardListResponseNetwork = "amex"
	TestingCreditCardListResponseNetworkDiscover     TestingCreditCardListResponseNetwork = "discover"
	TestingCreditCardListResponseNetworkJcb          TestingCreditCardListResponseNetwork = "jcb"
	TestingCreditCardListResponseNetworkUnionpay     TestingCreditCardListResponseNetwork = "unionpay"
	TestingCreditCardListResponseNetworkAlliancedata TestingCreditCardListResponseNetwork = "alliancedata"
	TestingCreditCardListResponseNetworkCitiplcc     TestingCreditCardListResponseNetwork = "citiplcc"
)

func (r TestingCreditCardListResponseNetwork) IsKnown() bool {
	switch r {
	case TestingCreditCardListResponseNetworkVisa, TestingCreditCardListResponseNetworkMastercard, TestingCreditCardListResponseNetworkAmex, TestingCreditCardListResponseNetworkDiscover, TestingCreditCardListResponseNetworkJcb, TestingCreditCardListResponseNetworkUnionpay, TestingCreditCardListResponseNetworkAlliancedata, TestingCreditCardListResponseNetworkCitiplcc:
		return true
	}
	return false
}
