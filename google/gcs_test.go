package google

import (
	"log"
	"time"

	"github.com/qvik/go-cloud-identity/google"
	"cloud.google.com/go/compute/metadata"
)

func ExampleGetSignedURL() {
	saEmail, _ := metadata.Email(google.DefaultAccount)
	name := "path/to/my/file"
	signBytes := func(payload []byte) ([]byte, error) {
		return google.SignBytes(payload, "", saEmail)
	}
	signedURL, _ := google.GetSignedURL("bucket1", name, saEmail, "GET",
		time.Minute*60, signBytes)

	log.Printf("Got signed URL: %v", signedURL)
}
