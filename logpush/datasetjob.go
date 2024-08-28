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

// DatasetJobService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetJobService] method instead.
type DatasetJobService struct {
	Options []option.RequestOption
}

// NewDatasetJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetJobService(opts ...option.RequestOption) (r *DatasetJobService) {
	r = &DatasetJobService{}
	r.Options = opts
	return
}

// Lists Logpush jobs for an account or zone for a dataset.
func (r *DatasetJobService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, datasetID string, opts ...option.RequestOption) (res *[]LogpushJob, err error) {
	var env DatasetJobGetResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/jobs", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatasetJobGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DatasetJobGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []LogpushJob                         `json:"result"`
	JSON    datasetJobGetResponseEnvelopeJSON    `json:"-"`
}

// datasetJobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetJobGetResponseEnvelope]
type datasetJobGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetJobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatasetJobGetResponseEnvelopeSuccess bool

const (
	DatasetJobGetResponseEnvelopeSuccessTrue DatasetJobGetResponseEnvelopeSuccess = true
)

func (r DatasetJobGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatasetJobGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
