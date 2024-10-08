// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/apiquery"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/param"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// AIModelSchemaService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelSchemaService] method instead.
type AIModelSchemaService struct {
	Options []option.RequestOption
}

// NewAIModelSchemaService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIModelSchemaService(opts ...option.RequestOption) (r *AIModelSchemaService) {
	r = &AIModelSchemaService{}
	r.Options = opts
	return
}

// Get Model Schema
func (r *AIModelSchemaService) Get(ctx context.Context, accountID string, query AIModelSchemaGetParams, opts ...option.RequestOption) (res *AIModelSchemaGetResponse, err error) {
	var env AIModelSchemaGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/models/schema", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIModelSchemaGetResponse = interface{}

type AIModelSchemaGetParams struct {
	// Model Name
	Model param.Field[string] `query:"model,required"`
}

// URLQuery serializes [AIModelSchemaGetParams]'s query parameters as `url.Values`.
func (r AIModelSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AIModelSchemaGetResponseEnvelope struct {
	Result  AIModelSchemaGetResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    aiModelSchemaGetResponseEnvelopeJSON `json:"-"`
}

// aiModelSchemaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIModelSchemaGetResponseEnvelope]
type aiModelSchemaGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiModelSchemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
