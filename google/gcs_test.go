package google

import (
	"encoding/base64"
	"log"
	"time"

	"github.com/pkg/errors"

	"cloud.google.com/go/compute/metadata"

	googleidentity "github.com/qvik/go-cloud-identity/google"
)

func ExampleGetSignedURL() {
	saEmail, _ := metadata.Email(googleidentity.DefaultAccount)
	name := "path/to/my/file"
	signBytes := func(payload []byte) ([]byte, error) {
		signature, _, err := googleidentity.SignBytes(payload, saEmail)
		if err != nil {
			return nil, err
		}

		signatureBytes, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			return nil, errors.Wrap(err, "base64 decoding failed")
		}

		return signatureBytes, nil
	}
	expires := time.Now().Add(time.Minute * 60)
	signedURL, _ := googleidentity.GetSignedURL("bucket1", name, saEmail, "GET",
		expires, signBytes)

	log.Printf("Got signed URL: %v", signedURL)
}
