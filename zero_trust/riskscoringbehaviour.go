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
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// RiskScoringBehaviourService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRiskScoringBehaviourService] method instead.
type RiskScoringBehaviourService struct {
	Options []option.RequestOption
}

// NewRiskScoringBehaviourService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRiskScoringBehaviourService(opts ...option.RequestOption) (r *RiskScoringBehaviourService) {
	r = &RiskScoringBehaviourService{}
	r.Options = opts
	return
}

// Update configuration for risk behaviors
func (r *RiskScoringBehaviourService) Update(ctx context.Context, accountIdentifier string, body RiskScoringBehaviourUpdateParams, opts ...option.RequestOption) (res *RiskScoringBehaviourUpdateResponse, err error) {
	var env RiskScoringBehaviourUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/behaviors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get all behaviors and associated configuration
func (r *RiskScoringBehaviourService) Get(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *RiskScoringBehaviourGetResponse, err error) {
	var env RiskScoringBehaviourGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/behaviors", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RiskScoringBehaviourUpdateResponse struct {
	Behaviors map[string]RiskScoringBehaviourUpdateResponseBehavior `json:"behaviors"`
	JSON      riskScoringBehaviourUpdateResponseJSON                `json:"-"`
}

// riskScoringBehaviourUpdateResponseJSON contains the JSON metadata for the struct
// [RiskScoringBehaviourUpdateResponse]
type riskScoringBehaviourUpdateResponseJSON struct {
	Behaviors   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseBehavior struct {
	Description string                                               `json:"description"`
	Enabled     bool                                                 `json:"enabled"`
	Name        string                                               `json:"name"`
	RiskLevel   RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel `json:"risk_level"`
	JSON        riskScoringBehaviourUpdateResponseBehaviorJSON       `json:"-"`
}

// riskScoringBehaviourUpdateResponseBehaviorJSON contains the JSON metadata for
// the struct [RiskScoringBehaviourUpdateResponseBehavior]
type riskScoringBehaviourUpdateResponseBehaviorJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseBehavior) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseBehaviorJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel string

const (
	RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelLow    RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel = "low"
	RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelMedium RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel = "medium"
	RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelHigh   RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel = "high"
)

func (r RiskScoringBehaviourUpdateResponseBehaviorsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelLow, RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelMedium, RiskScoringBehaviourUpdateResponseBehaviorsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringBehaviourGetResponse struct {
	Behaviors map[string]RiskScoringBehaviourGetResponseBehavior `json:"behaviors"`
	JSON      riskScoringBehaviourGetResponseJSON                `json:"-"`
}

// riskScoringBehaviourGetResponseJSON contains the JSON metadata for the struct
// [RiskScoringBehaviourGetResponse]
type riskScoringBehaviourGetResponseJSON struct {
	Behaviors   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseBehavior struct {
	Description string                                            `json:"description"`
	Enabled     bool                                              `json:"enabled"`
	Name        string                                            `json:"name"`
	RiskLevel   RiskScoringBehaviourGetResponseBehaviorsRiskLevel `json:"risk_level"`
	JSON        riskScoringBehaviourGetResponseBehaviorJSON       `json:"-"`
}

// riskScoringBehaviourGetResponseBehaviorJSON contains the JSON metadata for the
// struct [RiskScoringBehaviourGetResponseBehavior]
type riskScoringBehaviourGetResponseBehaviorJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseBehavior) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseBehaviorJSON) RawJSON() string {
	return r.raw
}

type RiskScoringBehaviourGetResponseBehaviorsRiskLevel string

const (
	RiskScoringBehaviourGetResponseBehaviorsRiskLevelLow    RiskScoringBehaviourGetResponseBehaviorsRiskLevel = "low"
	RiskScoringBehaviourGetResponseBehaviorsRiskLevelMedium RiskScoringBehaviourGetResponseBehaviorsRiskLevel = "medium"
	RiskScoringBehaviourGetResponseBehaviorsRiskLevelHigh   RiskScoringBehaviourGetResponseBehaviorsRiskLevel = "high"
)

func (r RiskScoringBehaviourGetResponseBehaviorsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourGetResponseBehaviorsRiskLevelLow, RiskScoringBehaviourGetResponseBehaviorsRiskLevelMedium, RiskScoringBehaviourGetResponseBehaviorsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringBehaviourUpdateParams struct {
	Behaviors param.Field[map[string]RiskScoringBehaviourUpdateParamsBehaviors] `json:"behaviors"`
}

func (r RiskScoringBehaviourUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RiskScoringBehaviourUpdateParamsBehaviors struct {
	Enabled   param.Field[bool]                                               `json:"enabled,required"`
	RiskLevel param.Field[RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel] `json:"risk_level,required"`
}

func (r RiskScoringBehaviourUpdateParamsBehaviors) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel string

const (
	RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelLow    RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel = "low"
	RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelMedium RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel = "medium"
	RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelHigh   RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel = "high"
)

func (r RiskScoringBehaviourUpdateParamsBehaviorsRiskLevel) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelLow, RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelMedium, RiskScoringBehaviourUpdateParamsBehaviorsRiskLevelHigh:
		return true
	}
	return false
}

type RiskScoringBehaviourUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   RiskScoringBehaviourUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success RiskScoringBehaviourUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    riskScoringBehaviourUpdateResponseEnvelopeJSON    `json:"-"`
}

// riskScoringBehaviourUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [RiskScoringBehaviourUpdateResponseEnvelope]
type riskScoringBehaviourUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringBehaviourUpdateResponseEnvelopeSuccess bool

const (
	RiskScoringBehaviourUpdateResponseEnvelopeSuccessTrue RiskScoringBehaviourUpdateResponseEnvelopeSuccess = true
)

func (r RiskScoringBehaviourUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RiskScoringBehaviourGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   RiskScoringBehaviourGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success RiskScoringBehaviourGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    riskScoringBehaviourGetResponseEnvelopeJSON    `json:"-"`
}

// riskScoringBehaviourGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RiskScoringBehaviourGetResponseEnvelope]
type riskScoringBehaviourGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RiskScoringBehaviourGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r riskScoringBehaviourGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RiskScoringBehaviourGetResponseEnvelopeSuccess bool

const (
	RiskScoringBehaviourGetResponseEnvelopeSuccessTrue RiskScoringBehaviourGetResponseEnvelopeSuccess = true
)

func (r RiskScoringBehaviourGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RiskScoringBehaviourGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
