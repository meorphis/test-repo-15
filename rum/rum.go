// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rum

import (
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

// RUMService contains methods and other services that help with interacting with
// the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRUMService] method instead.
type RUMService struct {
	Options  []option.RequestOption
	SiteInfo *SiteInfoService
	Rules    *RuleService
}

// NewRUMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRUMService(opts ...option.RequestOption) (r *RUMService) {
	r = &RUMService{}
	r.Options = opts
	r.SiteInfo = NewSiteInfoService(opts...)
	r.Rules = NewRuleService(opts...)
	return
}
