//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"errors"
	"testing"
	"time"
)

var (
	mockAzTokenProviderSuccess = func(ctx context.Context, scopes []string, tenant string) ([]byte, error) {
		return []byte(`{
  "accessToken": "mocktoken",
  "expiresOn": "2001-02-03 04:05:06.000007",
  "subscription": "mocksub",
  "tenant": "mocktenant",
  "tokenType": "Bearer"
}
`), nil
	}
	mockAzTokenProviderFailure = func(ctx context.Context, scopes []string, tenant string) ([]byte, error) {
		return nil, newAuthenticationFailedError(credNameAzureCLI, "mock provider error", nil, nil)
	}
)

func TestAzureCLICredential_DefaultChainError(t *testing.T) {
	cred, err := NewAzureCLICredential(&AzureCLICredentialOptions{
		inDefaultChain: true,
		tokenProvider:  mockAzTokenProviderFailure,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = cred.GetToken(context.Background(), testTRO)
	var ue *credentialUnavailableError
	if !errors.As(err, &ue) {
		t.Fatalf("expected credentialUnavailableError, got %T: %q", err, err)
	}
}

func TestAzureCLICredential_Error(t *testing.T) {
	// GetToken shouldn't invoke the CLI a second time after a failure
	authNs := 0
	expected := newCredentialUnavailableError(credNameAzureCLI, "it didn't work")
	o := AzureCLICredentialOptions{
		tokenProvider: func(context.Context, []string, string) ([]byte, error) {
			authNs++
			return nil, expected
		},
	}
	cred, err := NewAzureCLICredential(&o)
	if err != nil {
		t.Fatal(err)
	}
	_, err = cred.GetToken(context.Background(), testTRO)
	if err == nil {
		t.Fatal("expected an error")
	}
	if err != expected {
		t.Fatalf("expected %v, got %v", expected, err)
	}
	if authNs != 1 {
		t.Fatalf("expected 1 authN, got %d", authNs)
	}
}

func TestAzureCLICredential_GetTokenSuccess(t *testing.T) {
	options := AzureCLICredentialOptions{}
	options.tokenProvider = mockAzTokenProviderSuccess
	cred, err := NewAzureCLICredential(&options)
	if err != nil {
		t.Fatal(err)
	}
	at, err := cred.GetToken(context.Background(), testTRO)
	if err != nil {
		t.Fatal(err)
	}
	if at.Token != "mocktoken" {
		t.Fatalf("unexpected access token %q", at.Token)
	}
	expected := time.Date(2001, 2, 3, 4, 5, 6, 7000, time.Local).UTC()
	if actual := at.ExpiresOn; !actual.Equal(expected) || actual.Location() != time.UTC {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestAzureCLICredential_GetTokenInvalidToken(t *testing.T) {
	options := AzureCLICredentialOptions{}
	options.tokenProvider = mockAzTokenProviderFailure
	cred, err := NewAzureCLICredential(&options)
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	_, err = cred.GetToken(context.Background(), testTRO)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one.")
	}
}

func TestAzureCLICredential_TenantID(t *testing.T) {
	expected := "expected-tenant-id"
	called := false
	options := AzureCLICredentialOptions{
		TenantID: expected,
		tokenProvider: func(ctx context.Context, scopes []string, tenantID string) ([]byte, error) {
			called = true
			if tenantID != expected {
				t.Fatal("Unexpected tenant ID: " + tenantID)
			}
			return mockAzTokenProviderSuccess(ctx, scopes, tenantID)
		},
	}
	cred, err := NewAzureCLICredential(&options)
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	_, err = cred.GetToken(context.Background(), testTRO)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !called {
		t.Fatal("token provider wasn't called")
	}
}
