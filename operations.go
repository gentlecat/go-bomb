package giantbomb

import (
	"net/url"
	"strconv"
)

// Search does searching. ;)
// You can specify a list of resources that needs to be returned. See Resource*
// constants within the same package for reference.
func (api *GBClient) Search(query string, limit int, page int, resources []string, extraParams url.Values) (*Response, error) {
	if extraParams == nil {
		extraParams = make(url.Values)
	}
	extraParams["query"] = []string{query}
	extraParams["limit"] = []string{strconv.Itoa(limit)}
	extraParams["page"] = []string{strconv.Itoa(page)}
	extraParams["resources"] = resources

	return api.Get("search", "", extraParams)
}

// Platforms returns list of existing platforms.
func (api *GBClient) Platforms(limit int, offset int, extraParams url.Values) (*Response, error) {
	if extraParams == nil {
		extraParams = make(url.Values)
	}
	extraParams["limit"] = []string{strconv.Itoa(limit)}
	extraParams["offset"] = []string{strconv.Itoa(offset)}

	return api.Get("platforms", "", extraParams)
}
