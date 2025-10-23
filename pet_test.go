// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetest_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks-staging/bruce-test-go"
	"github.com/stainless-sdks-staging/bruce-test-go/internal/testutil"
	"github.com/stainless-sdks-staging/bruce-test-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.New(context.TODO(), brucetest.PetNewParams{
		Pet: brucetest.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        brucetest.Int(10),
			Category: brucetest.CategoryParam{
				ID:   brucetest.Int(1),
				Name: brucetest.String("Dogs"),
			},
			Status: brucetest.PetStatusAvailable,
			Tags: []brucetest.PetTagParam{{
				ID:   brucetest.Int(0),
				Name: brucetest.String("name"),
			}},
		},
	})
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.Get(context.TODO(), 0)
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.Update(context.TODO(), brucetest.PetUpdateParams{
		Pet: brucetest.PetParam{
			Name:      "doggie",
			PhotoURLs: []string{"string"},
			ID:        brucetest.Int(10),
			Category: brucetest.CategoryParam{
				ID:   brucetest.Int(1),
				Name: brucetest.String("Dogs"),
			},
			Status: brucetest.PetStatusAvailable,
			Tags: []brucetest.PetTagParam{{
				ID:   brucetest.Int(0),
				Name: brucetest.String("name"),
			}},
		},
	})
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Pets.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.FindByStatus(context.TODO(), brucetest.PetFindByStatusParams{
		Status: brucetest.PetFindByStatusParamsStatusAvailable,
	})
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.FindByTags(context.TODO(), brucetest.PetFindByTagsParams{
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Pets.UpdateByID(
		context.TODO(),
		0,
		brucetest.PetUpdateByIDParams{
			Name:   brucetest.String("name"),
			Status: brucetest.String("status"),
		},
	)
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := brucetest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Pets.UploadImage(
		context.TODO(),
		0,
		io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		brucetest.PetUploadImageParams{
			AdditionalMetadata: brucetest.String("additionalMetadata"),
		},
	)
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
