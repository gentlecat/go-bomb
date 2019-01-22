/*
Package giantbomb is a wrapper for the Giant Bomb API.

Giant Bomb is a video game website and wiki. Wiki contains a ton of information
about video games and things related to them. This information is collected by
volunteers, video game fans. Check https://www.giantbomb.com/.

API provides interface to access all this information. You will need to get an
API key to be able to use it. Information about getting your API key, terms of
use and other useful info is available at https://www.giantbomb.com/api/.

Now go build something, duder!

This library is meant to simplify interactons with Giant Bomb API. After you
get your API key, make sure to set Key variable to match it:

	giantbomb.Key = "YOUR_API_KEY"

There's a lot of data for you to get. It is usually a good idea to set a list
of fields that you need:

	giantbomb.FieldList = []string{
		"name",
		"platforms",
	}
*/
package giantbomb

import (
	"net/http"
	"strings"
)

var (
	Host      = "https://www.giantbomb.com"
	Key       string   // Your API key. Make sure to set this variable to match your key.
	FieldList []string // List of fields that determines data that you get in responses.
)

type Response struct {
	Error                string   `json:"error"`
	StatusCode           int      `json:"status_code"`
	Version              string   `json:"version"`
	Limit                int      `json:"limit"`
	Offset               int      `json:"offset"`
	NumberOfPageResults  int      `json:"number_of_page_results"`
	NumberOfTotalResults int      `json:"number_of_total_results"`
	Results              []Result `json:"results"`
}
type Result interface{}

type GiantBomb struct {
	Client  *http.Client
	baseURL string
}

// Pass empty string to resourceID if you don't need to specify it.
func getResourcePath(baseURL, resourceType, resourceID string) string {
	url := baseURL + "/api/" + resourceType + "/"
	if resourceID != "" {
		url += resourceID + "/"
	}
	url += "?format=json&api_key=" + Key
	if len(FieldList) > 0 {
		url += "&field_list=" + strings.Join(FieldList, ",")
	}
	return url
}
