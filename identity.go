package gcpidentity

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"cloud.google.com/go/compute/metadata"
	"github.com/coreos/go-oidc"
)

const (
	// DefaultAccount is the constant for the default service account
	DefaultAccount = "default"

	// Google's issuer (iss) URL
	googleIssuerURL = "https://accounts.google.com"
)

// IDTokenVerifier provides a method for verifying a Google ID token.
// Internally it caches the public key set used for the verification so
// that the operation is as efficient as possible.
type IDTokenVerifier struct {
	oidcVerifier *oidc.IDTokenVerifier
}

// FetchGoogleMetadataIDToken retrieves your current identity from the GCE metadata
// server. It is available for Google Compute Engine, Google AppEngine standard
// second generation runtimes and Google AppEngine flexible.
// Parameter aud should contain a free-form string (usually an url) that
// indicates the target audience (receiver) of the request that the
// id token is used to authenticate to.
// Parameter account should indicate the service account identifier to use;
// use empty string for default account
func FetchGoogleMetadataIDToken(aud string, account string) (string, error) {
	if !metadata.OnGCE() {
		return "", errors.New("not running on GCE or compatible environment")
	}

	if aud == "" {
		return "", errors.New("must specify a value for the aud parameter")
	}

	if account == "" {
		account = DefaultAccount
	}

	v := url.Values{}
	v.Set("audience", aud)

	uri := fmt.Sprintf("instance/service-accounts/%v/identity?%v", account, v.Encode())

	response, err := metadata.Get(uri)
	if err != nil {
		return "", err
	}

	return response, nil
}

//TODO implement extracting this ID token from a Service ACcount file.
// See: https://github.com/salrashid123/google_id_token/blob/master/golang/GoogleIDToken.go

// NewVerifier creates a new IDTokenVerifier that internally caches the
// remote key set used for ID token verification.
func NewVerifier(ctx context.Context, aud string) (*IDTokenVerifier, error) {
	provider, err := oidc.NewProvider(ctx, googleIssuerURL)
	if err != nil {
		return nil, err
	}

	var config = &oidc.Config{
		SkipClientIDCheck: false,
		ClientID:          aud,
	}

	idTokenVerifier := &IDTokenVerifier{
		oidcVerifier: provider.Verifier(config),
	}

	return idTokenVerifier, nil
}

// VerifyGoogleIDToken verifies an Google ID token.
// The parameter token is the JWT string.
// Return IDToken and nil error when verify succeeds.
func (v *IDTokenVerifier) VerifyGoogleIDToken(ctx context.Context, token string) (*oidc.IDToken, error) {
	return v.oidcVerifier.Verify(ctx, token)
}
