// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/intel"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
)

func TestAttackSurfaceReportIssueListWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		intel.AttackSurfaceReportIssueListParams{
			Dismissed:     cloudflare.F(false),
			IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			Page:          cloudflare.F(int64(1)),
			PerPage:       cloudflare.F(int64(25)),
			Product:       cloudflare.F([]string{"access", "dns"}),
			ProductNeq:    cloudflare.F([]string{"access", "dns"}),
			Severity:      cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			SeverityNeq:   cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
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

func TestAttackSurfaceReportIssueClassWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Class(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		intel.AttackSurfaceReportIssueClassParams{
			Dismissed:     cloudflare.F(false),
			IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			Product:       cloudflare.F([]string{"access", "dns"}),
			ProductNeq:    cloudflare.F([]string{"access", "dns"}),
			Severity:      cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			SeverityNeq:   cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
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

func TestAttackSurfaceReportIssueDismissWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Dismiss(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"issue_id",
		intel.AttackSurfaceReportIssueDismissParams{
			Dismiss: cloudflare.F(true),
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

func TestAttackSurfaceReportIssueSeverityWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Severity(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		intel.AttackSurfaceReportIssueSeverityParams{
			Dismissed:     cloudflare.F(false),
			IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			Product:       cloudflare.F([]string{"access", "dns"}),
			ProductNeq:    cloudflare.F([]string{"access", "dns"}),
			Severity:      cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			SeverityNeq:   cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
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

func TestAttackSurfaceReportIssueTypeWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.AttackSurfaceReport.Issues.Type(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		intel.AttackSurfaceReportIssueTypeParams{
			Dismissed:     cloudflare.F(false),
			IssueClass:    cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueClassNeq: cloudflare.F([]string{"a_record_dangling", "always_use_https_not_enabled"}),
			IssueType:     cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			IssueTypeNeq:  cloudflare.F([]intel.IssueType{intel.IssueTypeComplianceViolation, intel.IssueTypeEmailSecurity}),
			Product:       cloudflare.F([]string{"access", "dns"}),
			ProductNeq:    cloudflare.F([]string{"access", "dns"}),
			Severity:      cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			SeverityNeq:   cloudflare.F([]intel.SeverityQueryParam{intel.SeverityQueryParamLow, intel.SeverityQueryParamModerate}),
			Subject:       cloudflare.F([]string{"example.com", "example.com", "example.com"}),
			SubjectNeq:    cloudflare.F([]string{"example.com", "example.com", "example.com"}),
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
