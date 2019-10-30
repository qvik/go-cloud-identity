# Google Cloud Platform Identity Utility

[![GoDoc](https://godoc.org/github.com/qvik/go-cloud-identity?status.svg)](https://godoc.org/github.com/qvik/go-cloud-identity)

This library provides mechanisms for acquiring an identity token using Google's GCE metadata server and verifying it. It can be used eg. for facilitating authentication and authorization in service-to-service calls in Google Cloud Platform (GCP) environments.

The usual flow is:

1. The calling service creates and ID token with specified AUD value to match the service to be called. 
2. The calling service incorporates this token in the method call -- typically, in a HTTP request, in an `Authorization: Bearer` header
3. The called service extracts the token from the call
4. The called service verifies the token against its expected AUD value.

## Installing the library dependency

```sh
go get -u github.com/qvik/go-gcp-identity
```

## Acquiring an identity token

Retrieval of the identity token from a GCE metadata server is available for Google Compute Engine, Google AppEngine standard second generation and flexible runtimes.

To acquire an identity token:

```go
import (
    "github.com/qvik/go-cloud-identity/google"
    "log"
)

aud := "https://myapp/myservice" // Free-form string
identity, err := google.FetchMetadataIDToken(aud, "")
if err != nil {
    log.Fatalf("got error: %v", err)
    return
}
```

## Verifying an identity token

Verification of the identity token is available on any platform. It is highly recommended to cache the `IDTokenVerifier` object for performance reasons.

To verify an identity token:

```go
import (
    "github.com/qvik/go-cloud-identity/google"
    "log"
    "context"
)

ctx := context.Background()
verifier := google.NewVerifier(ctx, aud)
if _, err := verifier.VerifyIDToken(ctx, identity); err != nil {
    log.Fatalf("failed to verify token: %v", err)
}
```

## License

This library is released under the MIT license.

## Contributing 

Contributions to this library are welcomed. Any contributions have to meet the following criteria:

* Meaningfulness. Discuss whether what you are about to contribute indeed belongs to this library in the first place before submitting a pull request.
* Code style. Use gofmt and golint and you cannot go wrong with this. Generally do not exceed a line length of 80 characters.
* Testing. Try and include tests for your code.

## Contact

Any questions? Contact matti@qvik.fi.
