package google

import (
	"time"

	"cloud.google.com/go/storage"
)

// GetSignedURL returns a signed URL to access a Google Cloud Storage resource.
// The parameter bucket is the bucket name.
// The parameter method indicates the HTTP method (eg. "GET") to allow access to.
// The parameter serviceAccountEmail must hold the email value of the
// service account used for signing the URL - you could use
// `metadata.Email()` to retrieve this value for a service account.
// The signBytes parameter is a function that takes care of the signing.
// One way to implement it is using this library:
//
//  saEmail, _ := metadata.Email(google.DefaultAccount)
//  name := "path/to/my/file"
//  signBytes := func(payload []byte) ([]byte, error) {
//    return google.SignBytes(payload, "", saEmail)
//  }
//	expires := time.Now().Add(time.Minute * 60)
//  signedURL, _ := google.GetSignedURL("bucket1", name, saEmail, "GET",
//    expires, signBytes)
//
// This method does network I/O and could introduce latency.
func GetSignedURL(bucket, name, serviceAccountEmail, method string,
	expires time.Time,
	signBytes func(payload []byte) ([]byte, error)) (string, error) {

	opts := &storage.SignedURLOptions{
		GoogleAccessID: serviceAccountEmail,
		Scheme:         storage.SigningSchemeV4,
		SignBytes:      signBytes,
		Method:         method,
		Expires:        expires,
	}

	return storage.SignedURL(bucket, name, opts)
}
