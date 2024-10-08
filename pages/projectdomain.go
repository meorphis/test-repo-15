// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/stainless-sdks/meorphis-test-40-go/internal/apijson"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/pagination"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/requestconfig"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/shared"
	"github.com/tidwall/gjson"
)

// ProjectDomainService contains methods and other services that help with
// interacting with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDomainService] method instead.
type ProjectDomainService struct {
	Options []option.RequestOption
}

// NewProjectDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectDomainService(opts ...option.RequestOption) (r *ProjectDomainService) {
	r = &ProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *ProjectDomainService) New(ctx context.Context, accountID string, projectName string, body ProjectDomainNewParams, opts ...option.RequestOption) (res *ProjectDomainNewResponseUnion, err error) {
	var env ProjectDomainNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) List(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) (res *pagination.SinglePage[ProjectDomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountID, projectName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) ListAutoPaging(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProjectDomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, accountID, projectName, opts...))
}

// Delete a Pages project's domain.
func (r *ProjectDomainService) Delete(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *ProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *ProjectDomainService) Edit(ctx context.Context, accountID string, projectName string, domainName string, body ProjectDomainEditParams, opts ...option.RequestOption) (res *ProjectDomainEditResponseUnion, err error) {
	var env ProjectDomainEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single domain.
func (r *ProjectDomainService) Get(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *ProjectDomainGetResponseUnion, err error) {
	var env ProjectDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [pages.ProjectDomainNewResponseUnknown],
// [pages.ProjectDomainNewResponseArray] or [shared.UnionString].
type ProjectDomainNewResponseUnion interface {
	ImplementsPagesProjectDomainNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDomainNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDomainNewResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDomainNewResponseArray []interface{}

func (r ProjectDomainNewResponseArray) ImplementsPagesProjectDomainNewResponseUnion() {}

type ProjectDomainListResponse = interface{}

type ProjectDomainDeleteResponse = interface{}

// Union satisfied by [pages.ProjectDomainEditResponseUnknown],
// [pages.ProjectDomainEditResponseArray] or [shared.UnionString].
type ProjectDomainEditResponseUnion interface {
	ImplementsPagesProjectDomainEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDomainEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDomainEditResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDomainEditResponseArray []interface{}

func (r ProjectDomainEditResponseArray) ImplementsPagesProjectDomainEditResponseUnion() {}

// Union satisfied by [pages.ProjectDomainGetResponseUnknown],
// [pages.ProjectDomainGetResponseArray] or [shared.UnionString].
type ProjectDomainGetResponseUnion interface {
	ImplementsPagesProjectDomainGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDomainGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProjectDomainGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ProjectDomainGetResponseArray []interface{}

func (r ProjectDomainGetResponseArray) ImplementsPagesProjectDomainGetResponseUnion() {}

type ProjectDomainNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r ProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   ProjectDomainNewResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainNewResponseEnvelopeJSON    `json:"-"`
}

// projectDomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainNewResponseEnvelope]
type projectDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainNewResponseEnvelopeSuccess bool

const (
	ProjectDomainNewResponseEnvelopeSuccessTrue ProjectDomainNewResponseEnvelopeSuccess = true
)

func (r ProjectDomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainEditParams struct {
	Body interface{} `json:"body,required"`
}

func (r ProjectDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ProjectDomainEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ProjectDomainEditResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainEditResponseEnvelopeJSON    `json:"-"`
}

// projectDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainEditResponseEnvelope]
type projectDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainEditResponseEnvelopeSuccess bool

const (
	ProjectDomainEditResponseEnvelopeSuccessTrue ProjectDomainEditResponseEnvelopeSuccess = true
)

func (r ProjectDomainEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   ProjectDomainGetResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ProjectDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainGetResponseEnvelopeJSON    `json:"-"`
}

// projectDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainGetResponseEnvelope]
type projectDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ProjectDomainGetResponseEnvelopeSuccess bool

const (
	ProjectDomainGetResponseEnvelopeSuccessTrue ProjectDomainGetResponseEnvelopeSuccess = true
)

func (r ProjectDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
