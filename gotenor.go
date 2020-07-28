// Package gotenor provides a wrapper for the tenor gif API in go.
package gotenor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://api.tenor.com/v1/"

// tenorMediaObject is used to provide the fetched data about each individual
// media format provided.
type tenorMediaObject struct {
	URL, Preview   string
	Dims           []int
	Duration, Size float64
}

// tenorResult is a struct containing the data about each gif returned by the query.
type tenorResult struct {
	Tags                    []string `json:"tags"`
	URL, ItemURL, Title, ID string
	Media                   []map[string]tenorMediaObject
	Created                 float64
	Shares                  int
	HasAudio, Composite     bool
}

// tenorData is the top level struct
type tenorData struct {
	Results []tenorResult
	Next    string
}

// Tenor is the gerneal Tenor class
type Tenor struct {
	apiKey string
}

//_urlBuilder creates the url for the given requestType and parameters. It
// returns a string containing the URL.
func (t *Tenor) _urlBuilder(requestType string, params string) string {
	var query string
	if requestType == "singleId" {
		query = "gifs?ids=" + params
	}
	return baseURL + query + "&key=" + t.apiKey
}

// _fetch Fetches the  API data from the given url.
func (t *Tenor) _fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error HTTP status %d: %s", resp.StatusCode, resp.Status))
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//_parseData parses the JSON data from the request into usable structs.
func (t *Tenor) _parseData(body []byte) (*tenorData, error) {
	var data tenorData

	if len(body) <= 0 {
		return nil, errors.New("_parseData: no data in response body to parse")
	}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// NewTenor is a constructor to create a Tenor struct.
func NewTenor(apiKey string) *Tenor {
	return &Tenor{
		apiKey: apiKey,
	}
}

// GetById fetches GIF information based on the given ID
// More inforation: https://tenor.com/gifapi/documentation#endpoints-gifs
func (t *Tenor) GetById(id string) (*tenorData, error) {
	body, err := t._fetch(t._urlBuilder("singleId", id))
	if err != nil {
		return nil, err
	}

	return t._parseData(body)
}
