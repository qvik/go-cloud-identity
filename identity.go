package gcpidentity

import (
	"context"
	"fmt"
	"net/url"

	"cloud.google.com/go/compute/metadata"
	"github.com/coreos/go-oidc"
)

const (
	// DefaultAccount is the constant for the default service account
	DefaultAccount = "default"

	// Google's root certificates URL
	googleRootCertURL = "https://www.googleapis.com/oauth2/v3/certs"

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
// to use the default account, specify `gceidentity.DefaultAccount´.
func FetchGoogleMetadataIDToken(aud string, account string) (string, error) {
	if !metadata.OnGCE() {
		return "", fmt.Errorf("not running on GCE or compatible environment")
	}

	if aud == "" {
		return "", fmt.Errorf("must specify a value for the aud parameter")
	}

	v := url.Values{}
	v.Set("audience", aud)

	uri := fmt.Sprintf("instance/service-accounts/%v/identity?%v",
		account, v.Encode())

	response, err := metadata.Get(uri)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve metadata: %v", err)
	}

	return response, nil
}

//TODO implement extracting this ID token from a Service ACcount file.
// See: https://github.com/salrashid123/google_id_token/blob/master/golang/GoogleIDToken.go

// NewVerifier creates a new IDTokenVerifier that internally caches the
// remote key set used for ID token verification.
func NewVerifier(ctx context.Context, aud string) *IDTokenVerifier {
	keySet := oidc.NewRemoteKeySet(ctx, googleRootCertURL)

	var config = &oidc.Config{
		SkipClientIDCheck: false,
		ClientID:          aud,
	}
	oidcVerifier := oidc.NewVerifier(googleIssuerURL, keySet, config)

	return &IDTokenVerifier{
		oidcVerifier: oidcVerifier,
	}
}

// VerifyGoogleIDToken verifies an Google ID token.
// The parameter token is the raw ID token string.
func (v *IDTokenVerifier) VerifyGoogleIDToken(ctx context.Context,
	token string) error {

	_, err := v.oidcVerifier.Verify(ctx, token)

	return err
}
