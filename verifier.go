// Package cloudidentity provides functionality for dealing with software
// identities in cloud environments, such as OpenID Connect ID token
// acquirement and authorization.
package cloudidentity

import (
	"context"

	oidc "github.com/coreos/go-oidc"
)

// IDTokenVerifier provides a method for verifying an OpenID Connect ID token.
// Internally it caches the public key set used for the verification so
// that the operation is as efficient as possible.
type IDTokenVerifier struct {
	oidcVerifier *oidc.IDTokenVerifier
}

// VerifyIDToken verifies an ID token.
// The parameter token is the JWT string.
// Returns IDToken and nil error when verify succeeds.
func (v *IDTokenVerifier) VerifyIDToken(ctx context.Context,
	token string) (*oidc.IDToken, error) {

	return v.oidcVerifier.Verify(ctx, token)
}

// NewVerifier creates a new IDTokenVerifier that internally caches the
// remote key set used for ID token verification.
func NewVerifier(ctx context.Context,
	issuerURL, aud string) (*IDTokenVerifier, error) {

	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		return nil, err
	}

	var config = &oidc.Config{
		ClientID: aud,
	}

	idTokenVerifier := &IDTokenVerifier{
		oidcVerifier: provider.Verifier(config),
	}

	return idTokenVerifier, nil
}
