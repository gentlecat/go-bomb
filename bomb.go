/*
Package giantbomb is a wrapper for the Giant Bomb API.

Giant Bomb is a video game website and wiki. Wiki contains a ton of information
about video games and things related to them. This information is collected by
volunteers, video game fans. Check https://www.giantbomb.com/.

API provides interface to access all this information. You will need to get an
API key to be able to use it. Information about getting your API key, terms of
use and other useful info is available at https://www.giantbomb.com/api/.

Now go build something, duder!
*/
package giantbomb

import (
	"net/http"
	"net/url"
)

// Response object contains information returned from the API.
type Response struct {
	Results              []Result `json:"results"`                 // Zero or more items that match the filters specified
	NumberOfTotalResults int      `json:"number_of_total_results"` // The number of total results matching the filter conditions specified
	NumberOfPageResults  int      `json:"number_of_page_results"`  // The number of results on this page
	Limit                int      `json:"limit"`
	Offset               int      `json:"offset"`
	Version              string   `json:"version"`
	StatusCode           int      `json:"status_code"`
	Error                string   `json:"error"` // A text string representing the status_code
}

// Result is an item returned from the API that matches the filters you specify.
type Result interface{}

type GBClient struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient returns a new GBClient instance that can be used for calling the API.
func NewClient(apiKey string) *GBClient {
	return &GBClient{
		httpClient: &http.Client{},
		baseURL: (&url.URL{
			Scheme: "HTTPS",
			Host:   "www.giantbomb.com",
		}).String(),
		apiKey: apiKey,
	}
}

// Pass empty string to resourceID if you don't need to specify it.
func (api *GBClient) generateRequestURL(resourceType, resourceID string, queryParams url.Values) (string, error) {
	u, err := url.Parse(api.baseURL)
	if err != nil {
		return api.baseURL, err
	}

	// Overwriting mandatory parameters
	queryParams["format"] = []string{"json"}
	queryParams["api_key"] = []string{api.apiKey}
	u.RawQuery = queryParams.Encode()

	// and the path...
	u.Path = "/api/" + resourceType + "/"
	if resourceID != "" {
		u.Path += resourceID + "/"
	}

	return u.String(), nil
}
