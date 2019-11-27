package google

import (
	"errors"
	"fmt"
	"net/url"

	"cloud.google.com/go/compute/metadata"
)

// FetchMetadataIDToken retrieves your current identity from the GCE metadata
// server. It is available for Google Compute Engine, Google AppEngine standard
// second generation runtimes and Google AppEngine flexible.
// Parameter aud should contain a free-form string (usually an url) that
// indicates the target audience (receiver) of the request that the
// id token is used to authenticate to.
// Parameter account should indicate the service account identifier to use;
// use empty string or google.DefaultAccount for default account.
func FetchMetadataIDToken(aud string, account string) (string, error) {
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

	uri := fmt.Sprintf("instance/service-accounts/%v/identity?%v",
		account, v.Encode())

	response, err := metadata.Get(uri)
	if err != nil {
		return "", err
	}

	return response, nil
}
