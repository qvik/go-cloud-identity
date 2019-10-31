package google

import (
	"context"
	"encoding/base64"
	"fmt"
	stdlog "log"

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
func SignBytes(payload []byte, projectID,
	serviceAccount string) ([]byte, error) {

	ctx := context.Background()
	client, err := google.DefaultClient(ctx, iam.CloudPlatformScope)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create default Google client")
	}

	credService, err := iamcredentials.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create IAM credentials service")
	}
	accountsService := iamcredentials.NewProjectsServiceAccountsService(credService)

	if projectID == "" {
		projectID = "-"
	}

	name := fmt.Sprintf("projects/%v/serviceAccounts/%v",
		projectID, serviceAccount)
	encoded := base64.StdEncoding.EncodeToString(payload)
	req := &iamcredentials.SignBlobRequest{
		Payload: encoded,
	}

	stdlog.Printf("name=%v", name)

	res, err := accountsService.SignBlob(name, req).Do()
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign bytes")
	}

	signedBytes, err := base64.StdEncoding.DecodeString(res.SignedBlob)
	if err != nil {
		return nil, errors.Wrap(err, "base64 decoding failed")
	}

	return signedBytes, nil
}
