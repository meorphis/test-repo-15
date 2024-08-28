// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
)

// DatasetFieldService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetFieldService] method instead.
type DatasetFieldService struct {
	Options []option.RequestOption
}

// NewDatasetFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetFieldService(opts ...option.RequestOption) (r *DatasetFieldService) {
	r = &DatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *DatasetFieldService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, datasetID string, opts ...option.RequestOption) (res *DatasetFieldGetResponse, err error) {
	var env DatasetFieldGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountOrZone == "" {
		err = errors.New("missing required account_or_zone parameter")
		return
	}
	if accountOrZoneID == "" {
		err = errors.New("missing required account_or_zone_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/fields", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatasetFieldGetResponse = interface{}

type DatasetFieldGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DatasetFieldGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DatasetFieldGetResponse                `json:"result"`
	JSON    datasetFieldGetResponseEnvelopeJSON    `json:"-"`
}

// datasetFieldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetFieldGetResponseEnvelope]
type datasetFieldGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatasetFieldGetResponseEnvelopeSuccess bool

const (
	DatasetFieldGetResponseEnvelopeSuccessTrue DatasetFieldGetResponseEnvelopeSuccess = true
)

func (r DatasetFieldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatasetFieldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
