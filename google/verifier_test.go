package google

import (
	"context"
	"log"
)

func ExampleIDTokenVerifier_VerifyIDToken() {
	ctx := context.Background()
	aud := "https://myapp/myservice" // Free-form string
	token := "<token-redacted>"
	verifier := MustNewVerifier(ctx, aud)
	if _, err := verifier.VerifyIDToken(ctx, token); err != nil {
		log.Fatalf("failed to verify token: %v", err)
	}
}
