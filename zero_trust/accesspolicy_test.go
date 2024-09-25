// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/zero_trust"
)

func TestAccessPolicyNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Policies.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		zero_trust.AccessPolicyNewParams{
			Decision: cloudflare.F(zero_trust.DecisionAllow),
			Include: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]zero_trust.ApprovalGroupParam{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("email_list_uuid"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			IsolationRequired:            cloudflare.F(false),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			SessionDuration: cloudflare.F("24h"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessPolicyUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Policies.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		zero_trust.AccessPolicyUpdateParams{
			Decision: cloudflare.F(zero_trust.DecisionAllow),
			Include: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			Name: cloudflare.F("Allow devs"),
			ApprovalGroups: cloudflare.F([]zero_trust.ApprovalGroupParam{{
				ApprovalsNeeded: cloudflare.F(1.000000),
				EmailAddresses:  cloudflare.F([]string{"test1@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("email_list_uuid"),
			}, {
				ApprovalsNeeded: cloudflare.F(3.000000),
				EmailAddresses:  cloudflare.F([]string{"test@cloudflare.com", "test2@cloudflare.com"}),
				EmailListUUID:   cloudflare.F("597147a1-976b-4ef2-9af0-81d5d007fc34"),
			}}),
			ApprovalRequired: cloudflare.F(true),
			Exclude: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			IsolationRequired:            cloudflare.F(false),
			PurposeJustificationPrompt:   cloudflare.F("Please enter a justification for entering this protected domain."),
			PurposeJustificationRequired: cloudflare.F(true),
			Require: cloudflare.F([]zero_trust.AccessRuleUnionParam{zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}, zero_trust.EmailRuleParam{
				Email: cloudflare.F(zero_trust.EmailRuleEmailParam{
					Email: cloudflare.F("test@example.com"),
				}),
			}}),
			SessionDuration: cloudflare.F("24h"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessPolicyList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Policies.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessPolicyDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Policies.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccessPolicyGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Access.Policies.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
