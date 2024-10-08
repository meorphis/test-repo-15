// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// ProjectDeploymentHistoryLogService contains methods and other services that help
// with interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDeploymentHistoryLogService] method instead.
type ProjectDeploymentHistoryLogService struct {
	Options []option.RequestOption
}

// NewProjectDeploymentHistoryLogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewProjectDeploymentHistoryLogService(opts ...option.RequestOption) (r *ProjectDeploymentHistoryLogService) {
	r = &ProjectDeploymentHistoryLogService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *ProjectDeploymentHistoryLogService) Get(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *ProjectDeploymentHistoryLogGetResponseUnion, err error) {
	var env ProjectDeploymentHistoryLogGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [pages.ProjectDeploymentHistoryLogGetResponseUnknown],
// [pages.ProjectDeploymentHistoryLogGetResponseArray] or [shared.UnionString].
type ProjectDeploymentHistoryLogGetResponseUnion interface {
	ImplementsPagesProjectDeploymentHistoryLogGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentHistoryLogGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDeploymentHistoryLogGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDeploymentHistoryLogGetResponseArray []interface{}

func (r ProjectDeploymentHistoryLogGetResponseArray) ImplementsPagesProjectDeploymentHistoryLogGetResponseUnion() {
}

type ProjectDeploymentHistoryLogGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   ProjectDeploymentHistoryLogGetResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeploymentHistoryLogGetResponseEnvelopeJSON    `json:"-"`
}

// projectDeploymentHistoryLogGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ProjectDeploymentHistoryLogGetResponseEnvelope]
type projectDeploymentHistoryLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentHistoryLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentHistoryLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess bool

const (
	ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess = true
)

func (r ProjectDeploymentHistoryLogGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeploymentHistoryLogGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
