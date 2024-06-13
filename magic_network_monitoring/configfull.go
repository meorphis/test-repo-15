// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

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

// ConfigFullService contains methods and other services that help with interacting
// with the testcloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigFullService] method instead.
type ConfigFullService struct {
	Options []option.RequestOption
}

// NewConfigFullService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigFullService(opts ...option.RequestOption) (r *ConfigFullService) {
	r = &ConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *ConfigFullService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Configuration, err error) {
	var env ConfigFullGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config/full", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigFullGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Configuration         `json:"result,required"`
	// Whether the API call was successful
	Success ConfigFullGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configFullGetResponseEnvelopeJSON    `json:"-"`
}

// configFullGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigFullGetResponseEnvelope]
type configFullGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFullGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigFullGetResponseEnvelopeSuccess bool

const (
	ConfigFullGetResponseEnvelopeSuccessTrue ConfigFullGetResponseEnvelopeSuccess = true
)

func (r ConfigFullGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigFullGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
