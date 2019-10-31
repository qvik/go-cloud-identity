package google

import (
	"log"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/qvik/go-cloud-identity/google"
)

func ExampleGetSignedURL() {
	saEmail, _ := metadata.Email(google.DefaultAccount)
	name := "path/to/my/file"
	signBytes := func(payload []byte) ([]byte, error) {
		return google.SignBytes(payload, "", saEmail)
	}
	expires := time.Now().Add(time.Minute * 60)
	signedURL, _ := google.GetSignedURL("bucket1", name, saEmail, "GET",
		expires, signBytes)

	log.Printf("Got signed URL: %v", signedURL)
}
