package giantbomb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

func (api *GBClient) Search(query string, limit int, page int, resources []string, extraParams url.Values) (*Response, error) {
	if extraParams == nil { extraParams = make(url.Values)}
	extraParams["query"] = []string{query}
	extraParams["limit"] = []string{strconv.Itoa(limit)}
	extraParams["page"] = []string{strconv.Itoa(page)}
	extraParams["resources"] = resources

	u, err := api.generateRequestURL("search", "", extraParams)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpClient.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request to %s failed (%s)!", u, resp.Status))
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

// Platforms returns list of existing platforms.
func (api *GBClient) Platforms(limit int, offset int, extraParams url.Values) (*Response, error) {
	if extraParams == nil { extraParams = make(url.Values)}
	extraParams["limit"] = []string{strconv.Itoa(limit)}
	extraParams["offset"] = []string{strconv.Itoa(offset)}

	u, err := api.generateRequestURL("platforms", "", extraParams)
	if err != nil {
		return nil, err
	}

	resp, err := api.httpClient.Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Request to %s failed (%s)!", u, resp.Status))
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
