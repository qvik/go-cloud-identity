package google

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/pkg/errors"

	"golang.org/x/oauth2/google"
	iam "google.golang.org/api/iam/v1"
	iamcredentials "google.golang.org/api/iamcredentials/v1"
)

// SignBytes signs the given bytes using the given service account.
// Specify `google.DefaultAccount` as serviceAccount parameter to use the
// default account.
// You may specify "-" or empty string ("") for the projectID parameter
// to use the current project's ID.
// This method does network I/O and could introduce latency.
func SignBytes(bytes []byte, serviceAccount string) (string, error) {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, iam.CloudPlatformScope)
	if err != nil {
		return "", errors.Wrap(err, "failed to create default Google client")
	}

	credService, err := iamcredentials.New(client)
	if err != nil {
		return "", errors.Wrap(err, "failed to create IAM credentials service")
	}
	accountsService := iamcredentials.NewProjectsServiceAccountsService(credService)

	name := fmt.Sprintf("projects/-/serviceAccounts/%v", serviceAccount)
	encoded := base64.StdEncoding.EncodeToString(bytes)
	req := &iamcredentials.SignBlobRequest{
		Payload: encoded,
	}

	res, err := accountsService.SignBlob(name, req).Do()
	if err != nil {
		return "", errors.Wrap(err, "failed to sign bytes")
	}

	return res.SignedBlob, nil
}
