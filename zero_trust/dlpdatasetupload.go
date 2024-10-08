// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPDatasetUploadService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDatasetUploadService] method instead.
type DLPDatasetUploadService struct {
	Options []option.RequestOption
}

// NewDLPDatasetUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPDatasetUploadService(opts ...option.RequestOption) (r *DLPDatasetUploadService) {
	r = &DLPDatasetUploadService{}
	r.Options = opts
	return
}

// Prepare to upload a new version of a dataset.
func (r *DLPDatasetUploadService) New(ctx context.Context, accountID string, datasetID string, opts ...option.RequestOption) (res *NewVersion, err error) {
	var env DLPDatasetUploadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload", accountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Upload a new version of a dataset.
func (r *DLPDatasetUploadService) Edit(ctx context.Context, accountID string, datasetID string, version int64, body DLPDatasetUploadEditParams, opts ...option.RequestOption) (res *Dataset, err error) {
	var env DLPDatasetUploadEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/datasets/%s/upload/%v", accountID, datasetID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NewVersion struct {
	MaxCells int64          `json:"max_cells,required"`
	Version  int64          `json:"version,required"`
	Secret   string         `json:"secret" format:"password"`
	JSON     newVersionJSON `json:"-"`
}

// newVersionJSON contains the JSON metadata for the struct [NewVersion]
type newVersionJSON struct {
	MaxCells    apijson.Field
	Version     apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newVersionJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelope struct {
	Errors     []shared.ResponseInfo                         `json:"errors,required"`
	Messages   []shared.ResponseInfo                         `json:"messages,required"`
	Success    bool                                          `json:"success,required"`
	Result     NewVersion                                    `json:"result"`
	ResultInfo DLPDatasetUploadNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadNewResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadNewResponseEnvelope]
type dlpDatasetUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadNewResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                             `json:"total_count,required"`
	JSON       dlpDatasetUploadNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DLPDatasetUploadNewResponseEnvelopeResultInfo]
type dlpDatasetUploadNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditParams struct {
	Body string `json:"body,required"`
}

func (r DLPDatasetUploadEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPDatasetUploadEditResponseEnvelope struct {
	Errors     []shared.ResponseInfo                          `json:"errors,required"`
	Messages   []shared.ResponseInfo                          `json:"messages,required"`
	Success    bool                                           `json:"success,required"`
	Result     Dataset                                        `json:"result"`
	ResultInfo DLPDatasetUploadEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpDatasetUploadEditResponseEnvelopeJSON       `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPDatasetUploadEditResponseEnvelope]
type dlpDatasetUploadEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDatasetUploadEditResponseEnvelopeResultInfo struct {
	// total number of pages
	Count int64 `json:"count,required"`
	// current page
	Page int64 `json:"page,required"`
	// number of items per page
	PerPage int64 `json:"per_page,required"`
	// total number of items
	TotalCount int64                                              `json:"total_count,required"`
	JSON       dlpDatasetUploadEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpDatasetUploadEditResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DLPDatasetUploadEditResponseEnvelopeResultInfo]
type dlpDatasetUploadEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDatasetUploadEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDatasetUploadEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
