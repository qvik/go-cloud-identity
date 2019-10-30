// Package google contains methods specific for Google Cloud Platform.
package google

import (
	"context"
	"fmt"

	cloudidentity "github.com/qvik/go-cloud-identity"
)

const (
	// Google's issuer (iss) URL
	googleIssuerURL = "https://accounts.google.com"

	// DefaultAccount is the constant for the default service account
	DefaultAccount = "default"
)

//TODO implement extracting this ID token from a Service Account file.
// See: https://github.com/salrashid123/google_id_token/blob/master/golang/GoogleIDToken.go

// NewVerifier creates a new IDTokenVerifier that uses Google's
// issuer URL.
func NewVerifier(ctx context.Context,
	aud string) (*cloudidentity.IDTokenVerifier, error) {

	return cloudidentity.NewVerifier(ctx, googleIssuerURL, aud)
}

// MustNewVerifier creates a new IDTokenVerifier that uses Google's
// issuer URL.
// Panics on errors.
func MustNewVerifier(ctx context.Context,
	aud string) *cloudidentity.IDTokenVerifier {

	verifier, err := cloudidentity.NewVerifier(ctx, googleIssuerURL, aud)
	if err != nil {
		panic(fmt.Sprintf("failed to create verifier: %v", err))
	}

	return verifier
}
