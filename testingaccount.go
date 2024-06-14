// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package meorphistest40

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// TestingAccountService contains methods and other services that help with
// interacting with the meorphis-test-40 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestingAccountService] method instead.
type TestingAccountService struct {
	Options []option.RequestOption
}

// NewTestingAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTestingAccountService(opts ...option.RequestOption) (r *TestingAccountService) {
	r = &TestingAccountService{}
	r.Options = opts
	return
}

// Create a Bolt shopper account for testing purposes.
func (r *TestingAccountService) TestingAccountNew(ctx context.Context, body TestingAccountTestingAccountNewParams, opts ...option.RequestOption) (res *TestingAccountTestingAccountNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "testing/accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TestingAccountTestingAccountNewResponse struct {
	DeactivateAt time.Time                                         `json:"deactivate_at,required" format:"date-time"`
	Email        string                                            `json:"email,required"`
	EmailState   TestingAccountTestingAccountNewResponseEmailState `json:"email_state,required"`
	OAuthCode    string                                            `json:"oauth_code,required"`
	OtpCode      string                                            `json:"otp_code,required"`
	Phone        string                                            `json:"phone,required"`
	PhoneState   TestingAccountTestingAccountNewResponsePhoneState `json:"phone_state,required"`
	JSON         testingAccountTestingAccountNewResponseJSON       `json:"-"`
}

// testingAccountTestingAccountNewResponseJSON contains the JSON metadata for the
// struct [TestingAccountTestingAccountNewResponse]
type testingAccountTestingAccountNewResponseJSON struct {
	DeactivateAt apijson.Field
	Email        apijson.Field
	EmailState   apijson.Field
	OAuthCode    apijson.Field
	OtpCode      apijson.Field
	Phone        apijson.Field
	PhoneState   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TestingAccountTestingAccountNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testingAccountTestingAccountNewResponseJSON) RawJSON() string {
	return r.raw
}

type TestingAccountTestingAccountNewResponseEmailState string

const (
	TestingAccountTestingAccountNewResponseEmailStateMissing    TestingAccountTestingAccountNewResponseEmailState = "missing"
	TestingAccountTestingAccountNewResponseEmailStateUnverified TestingAccountTestingAccountNewResponseEmailState = "unverified"
	TestingAccountTestingAccountNewResponseEmailStateVerified   TestingAccountTestingAccountNewResponseEmailState = "verified"
)

func (r TestingAccountTestingAccountNewResponseEmailState) IsKnown() bool {
	switch r {
	case TestingAccountTestingAccountNewResponseEmailStateMissing, TestingAccountTestingAccountNewResponseEmailStateUnverified, TestingAccountTestingAccountNewResponseEmailStateVerified:
		return true
	}
	return false
}

type TestingAccountTestingAccountNewResponsePhoneState string

const (
	TestingAccountTestingAccountNewResponsePhoneStateMissing    TestingAccountTestingAccountNewResponsePhoneState = "missing"
	TestingAccountTestingAccountNewResponsePhoneStateUnverified TestingAccountTestingAccountNewResponsePhoneState = "unverified"
	TestingAccountTestingAccountNewResponsePhoneStateVerified   TestingAccountTestingAccountNewResponsePhoneState = "verified"
)

func (r TestingAccountTestingAccountNewResponsePhoneState) IsKnown() bool {
	switch r {
	case TestingAccountTestingAccountNewResponsePhoneStateMissing, TestingAccountTestingAccountNewResponsePhoneStateUnverified, TestingAccountTestingAccountNewResponsePhoneStateVerified:
		return true
	}
	return false
}

type TestingAccountTestingAccountNewParams struct {
	DeactivateAt param.Field[time.Time]                                       `json:"deactivate_at,required" format:"date-time"`
	EmailState   param.Field[TestingAccountTestingAccountNewParamsEmailState] `json:"email_state,required"`
	PhoneState   param.Field[TestingAccountTestingAccountNewParamsPhoneState] `json:"phone_state,required"`
	HasAddress   param.Field[bool]                                            `json:"has_address"`
	IsMigrated   param.Field[bool]                                            `json:"is_migrated"`
}

func (r TestingAccountTestingAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TestingAccountTestingAccountNewParamsEmailState string

const (
	TestingAccountTestingAccountNewParamsEmailStateMissing    TestingAccountTestingAccountNewParamsEmailState = "missing"
	TestingAccountTestingAccountNewParamsEmailStateUnverified TestingAccountTestingAccountNewParamsEmailState = "unverified"
	TestingAccountTestingAccountNewParamsEmailStateVerified   TestingAccountTestingAccountNewParamsEmailState = "verified"
)

func (r TestingAccountTestingAccountNewParamsEmailState) IsKnown() bool {
	switch r {
	case TestingAccountTestingAccountNewParamsEmailStateMissing, TestingAccountTestingAccountNewParamsEmailStateUnverified, TestingAccountTestingAccountNewParamsEmailStateVerified:
		return true
	}
	return false
}

type TestingAccountTestingAccountNewParamsPhoneState string

const (
	TestingAccountTestingAccountNewParamsPhoneStateMissing    TestingAccountTestingAccountNewParamsPhoneState = "missing"
	TestingAccountTestingAccountNewParamsPhoneStateUnverified TestingAccountTestingAccountNewParamsPhoneState = "unverified"
	TestingAccountTestingAccountNewParamsPhoneStateVerified   TestingAccountTestingAccountNewParamsPhoneState = "verified"
)

func (r TestingAccountTestingAccountNewParamsPhoneState) IsKnown() bool {
	switch r {
	case TestingAccountTestingAccountNewParamsPhoneStateMissing, TestingAccountTestingAccountNewParamsPhoneStateUnverified, TestingAccountTestingAccountNewParamsPhoneStateVerified:
		return true
	}
	return false
}
