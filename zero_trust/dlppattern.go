// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/logpush"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// DLPPatternService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPPatternService] method instead.
type DLPPatternService struct {
	Options []option.RequestOption
}

// NewDLPPatternService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPatternService(opts ...option.RequestOption) (r *DLPPatternService) {
	r = &DLPPatternService{}
	r.Options = opts
	return
}

// Validates whether this pattern is a valid regular expression. Rejects it if the
// regular expression is too complex or can match an unbounded-length string. Your
// regex will be rejected if it uses the Kleene Star -- be sure to bound the
// maximum number of characters that can be matched.
func (r *DLPPatternService) Validate(ctx context.Context, accountID string, body DLPPatternValidateParams, opts ...option.RequestOption) (res *logpush.OwnershipValidation, err error) {
	var env DLPPatternValidateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPatternValidateParams struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
}

func (r DLPPatternValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPatternValidateResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   logpush.OwnershipValidation `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DLPPatternValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPatternValidateResponseEnvelopeJSON    `json:"-"`
}

// dlpPatternValidateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPPatternValidateResponseEnvelope]
type dlpPatternValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPatternValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPatternValidateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPPatternValidateResponseEnvelopeSuccess bool

const (
	DLPPatternValidateResponseEnvelopeSuccessTrue DLPPatternValidateResponseEnvelopeSuccess = true
)

func (r DLPPatternValidateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPatternValidateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
