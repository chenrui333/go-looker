package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/chenrui333/go-looker/client"
	"github.com/chenrui333/go-looker/client/api_auth"

	httptransport "github.com/go-openapi/runtime/client"
)

// Config is used to configure the creation of a Looker client.
type Config struct {
	Token string
}

// Meta is used by the provider to access the Looker API.
type Meta struct {
	Looker   *apiclient.Looker
	AuthInfo runtime.ClientAuthInfoWriter
}

func main() {
	fmt.Println("Hello")

	clientID := os.Getenv("LOOKER_CLIENT_ID")
	clientSecret := os.Getenv("LOOKER_CLIENT_SECRET")

	fmt.Println("clientID is", clientID)
	fmt.Println("clientSecret is", clientSecret)

	// loginParams := api_auth.NewLoginParams()
	loginParams := api_auth.LoginParams{
		ClientID:     &clientID,
		ClientSecret: &clientSecret,
	}

	fmt.Println("loginParams is", *loginParams.ClientID)

	// create the API client
	client := apiclient.New(httptransport.New("", "", nil), strfmt.Default)

	// make the authenticated request to get all items
	bearerTokenAuth := httptransport.BearerToken(os.Getenv("API_ACCESS_TOKEN"))
	// basicAuth := httptransport.BasicAuth(os.Getenv("API_USER"), os.Getenv("API_PASSWORD"))
	// apiKeyQueryAuth := httptransport.APIKeyAuth("apiKey", "query", os.Getenv("API_KEY"))
	// apiKeyHeaderAuth := httptransport.APIKeyAuth("X-API-TOKEN", "header", os.Getenv("API_KEY"))
	resp, err := client.Operations.All(operations.AllParams{}, bearerTokenAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, basicAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyQueryAuth)
	// resp, err := client.Operations.All(operations.AllParams{}, apiKeyHeaderAuth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)

}
