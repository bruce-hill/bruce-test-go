// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brucetest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks-staging/bruce-test-go"
	"github.com/stainless-sdks-staging/bruce-test-go/internal/testutil"
	"github.com/stainless-sdks-staging/bruce-test-go/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.New(context.TODO(), brucetest.UserNewParams{
		User: brucetest.UserParam{
			ID:         brucetest.Int(10),
			Email:      brucetest.String("john@email.com"),
			FirstName:  brucetest.String("John"),
			LastName:   brucetest.String("James"),
			Password:   brucetest.String("12345"),
			Phone:      brucetest.String("12345"),
			Username:   brucetest.String("theUser"),
			UserStatus: brucetest.Int(1),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), "username")
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Users.Update(
		context.TODO(),
		"username",
		brucetest.UserUpdateParams{
			User: brucetest.UserParam{
				ID:         brucetest.Int(10),
				Email:      brucetest.String("john@email.com"),
				FirstName:  brucetest.String("John"),
				LastName:   brucetest.String("James"),
				Password:   brucetest.String("12345"),
				Phone:      brucetest.String("12345"),
				Username:   brucetest.String("theUser"),
				UserStatus: brucetest.Int(1),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.Users.Delete(context.TODO(), "username")
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.NewWithList(context.TODO(), brucetest.UserNewWithListParams{
		Items: []brucetest.UserParam{{
			ID:         brucetest.Int(10),
			Email:      brucetest.String("john@email.com"),
			FirstName:  brucetest.String("John"),
			LastName:   brucetest.String("James"),
			Password:   brucetest.String("12345"),
			Phone:      brucetest.String("12345"),
			Username:   brucetest.String("theUser"),
			UserStatus: brucetest.Int(1),
		}},
	})
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Login(context.TODO(), brucetest.UserLoginParams{
		Password: brucetest.String("password"),
		Username: brucetest.String("username"),
	})
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.Users.Logout(context.TODO())
	if err != nil {
		var apierr *brucetest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
