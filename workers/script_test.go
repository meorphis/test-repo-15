// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stainless-sdks/meorphis-test-40-go"
	"github.com/stainless-sdks/meorphis-test-40-go/internal/testutil"
	"github.com/stainless-sdks/meorphis-test-40-go/option"
	"github.com/stainless-sdks/meorphis-test-40-go/workers"
)

func TestScriptUpdateWithOptionalParams(t *testing.T) {
	t.Skip("TODO: investigate broken test")
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
	_, err := client.Workers.Scripts.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		workers.ScriptUpdateParams{
			Body: workers.ScriptUpdateParamsBodyObject{
				AnyPartName: cloudflare.F([]io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents"))), io.Reader(bytes.NewBuffer([]byte("some file contents")))}),
				Metadata: cloudflare.F(workers.ScriptUpdateParamsBodyObjectMetadata{
					Bindings: cloudflare.F([]interface{}{map[string]interface{}{
						"name": "MY_ENV_VAR",
						"text": "my_data",
						"type": "plain_text",
					}}),
					BodyPart:           cloudflare.F("worker.js"),
					CompatibilityDate:  cloudflare.F("2023-07-25"),
					CompatibilityFlags: cloudflare.F([]string{"string", "string", "string"}),
					KeepBindings:       cloudflare.F([]string{"string", "string", "string"}),
					Logpush:            cloudflare.F(false),
					MainModule:         cloudflare.F("worker.js"),
					Migrations: cloudflare.F[workers.ScriptUpdateParamsBodyObjectMetadataMigrationsUnion](workers.SingleStepMigrationParam{
						DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
						NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
						NewTag:         cloudflare.F("v2"),
						OldTag:         cloudflare.F("v1"),
						RenamedClasses: cloudflare.F([]workers.SingleStepMigrationRenamedClassParam{{
							From: cloudflare.F("from"),
							To:   cloudflare.F("to"),
						}, {
							From: cloudflare.F("from"),
							To:   cloudflare.F("to"),
						}, {
							From: cloudflare.F("from"),
							To:   cloudflare.F("to"),
						}}),
						TransferredClasses: cloudflare.F([]workers.SingleStepMigrationTransferredClassParam{{
							From:       cloudflare.F("from"),
							FromScript: cloudflare.F("from_script"),
							To:         cloudflare.F("to"),
						}, {
							From:       cloudflare.F("from"),
							FromScript: cloudflare.F("from_script"),
							To:         cloudflare.F("to"),
						}, {
							From:       cloudflare.F("from"),
							FromScript: cloudflare.F("from_script"),
							To:         cloudflare.F("to"),
						}}),
					}),
					Placement: cloudflare.F(workers.PlacementConfigurationParam{
						Mode: cloudflare.F(workers.PlacementConfigurationModeSmart),
					}),
					Tags: cloudflare.F([]string{"string", "string", "string"}),
					TailConsumers: cloudflare.F([]workers.ConsumerScriptParam{{
						Service:     cloudflare.F("my-log-consumer"),
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
					}, {
						Service:     cloudflare.F("my-log-consumer"),
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
					}, {
						Service:     cloudflare.F("my-log-consumer"),
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
					}}),
					UsageModel:  cloudflare.F(workers.ScriptUpdateParamsBodyObjectMetadataUsageModelBundled),
					VersionTags: cloudflare.F[any](map[string]interface{}{}),
				}),
			},
			RollbackTo: cloudflare.F("f174e90a-fafe-4643-bbbc-4a0ed4fc8415"),
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

func TestScriptList(t *testing.T) {
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
	_, err := client.Workers.Scripts.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestScriptDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Workers.Scripts.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		workers.ScriptDeleteParams{
			Force: cloudflare.F(true),
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

func TestScriptGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	resp, err := client.Workers.Scripts.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
