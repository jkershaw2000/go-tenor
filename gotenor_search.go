package gotenor

import (
	"net/url"
	"strconv"
)

// GetSearch provides an implementation of the tenor search endpoint.
// It requires a query, filter(off, low, medium, high),
// locale(2 letter ISO country codes) and limit(-1 for unlimited)
// Returns tenorData struct.
func (t *Tenor) GetSearch(query, filter, locale string, limit int) (*tenorData, error) {
	url := t._urlBuilder("search") + "&q=" + url.QueryEscape(query)

	if limit > 0 {
		url += "&limit=" + strconv.Itoa(limit)
	}

	if filter != "" {
		url += "&contentfilter=" + filter
	}

	if locale != "" {
		url += "&locale=" + locale
	}

	body, err := t._fetch(url)
	if err != nil {
		return nil, err
	}

	return t._parseData(body)

}
