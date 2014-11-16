/*
Package giantbomb implements a wrapper for Giant Bomb API.

Giant Bomb is a video game website and wiki. Wiki contains a ton of information
about video games and things related to them. This information is collected by
volunteers, video game fans.

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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	ResourceTypeGame      = "game"
	ResourceTypeFranchise = "franchise"
	ResourceTypeCharacter = "character"
	ResourceTypeConcept   = "concept"
	ResourceTypeObject    = "object"
	ResourceTypeLocation  = "location"
	ResourceTypePerson    = "person"
	ResourceTypeCompany   = "company"
	ResourceTypeVideo     = "video"
)

var (
	Host      = "https://www.giantbomb.com/api/"
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

// Pass empty string to resourceID if you don't need to specify it.
func getBaseURL(resourceType string, resourceID string) string {
	url := Host + resourceType + "/"
	if resourceID != "" {
		url += resourceID + "/"
	}
	url += "?format=json&api_key=" + Key
	if len(FieldList) > 0 {
		url += "&field_list=" + strings.Join(FieldList, ",")
	}
	return url
}

func Search(query string, limit int, page int, resources []string) (*Response, error) {
	url := getBaseURL("search", "") +
		"&query=\"" + query + "\"" +
		"&limit=" + strconv.Itoa(limit) +
		"&page=" + strconv.Itoa(page) +
		"&resources=" + strings.Join(resources, ",")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request to %s failed (%s)!", url, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Platforms returns list of existing gaming platforms.
func Platforms(limit int, offset int) (*Response, error) {
	url := getBaseURL("platforms", "") +
		"&limit=" + strconv.Itoa(limit) +
		"&offset=" + strconv.Itoa(offset)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request to %s failed (%s)!", url, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
