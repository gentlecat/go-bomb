package giantbomb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func (api *GiantBomb) Search(query string, limit int, page int, resources []string) (*Response, error) {
	url := getResourcePath(api.baseURL, "search", "") +
		"&query=\"" + query + "\"" +
		"&limit=" + strconv.Itoa(limit) +
		"&page=" + strconv.Itoa(page) +
		"&resources=" + strings.Join(resources, ",")
	resp, err := api.Client.Get(url)
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
func (api *GiantBomb) Platforms(limit int, offset int) (*Response, error) {
	url := getResourcePath(api.baseURL, "platforms", "") +
		"&limit=" + strconv.Itoa(limit) +
		"&offset=" + strconv.Itoa(offset)
	resp, err := api.Client.Get(url)
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
