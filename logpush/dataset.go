// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// DatasetService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
	Fields  *DatasetFieldService
	Jobs    *DatasetJobService
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r *DatasetService) {
	r = &DatasetService{}
	r.Options = opts
	r.Fields = NewDatasetFieldService(opts...)
	r.Jobs = NewDatasetJobService(opts...)
	return
}
