package google

import "log"

func ExampleFetchMetadataIDToken() {
	aud := "https://myapp/myservice" // Free-form string
	identity, err := FetchMetadataIDToken(aud, "")
	if err != nil {
		log.Fatalf("got error: %v", err)
		return
	}

	log.Printf("Got identity token: %v", identity)
	// Output: Got identity token: <token>
}
