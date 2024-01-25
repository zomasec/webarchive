package webarchive

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

// Archive represents the Wayback Machine archive utility.
type Archive struct {
	Query      string       // Query represents the domain or URL to query in the Wayback Machine.
	HTTPClient *http.Client // HTTPClient is an optional custom HTTP client. If nil, a default client with a 10-second timeout is used.
}

// Result represents the result of fetching URLs from the Wayback Machine.
type Result struct {
	URLs []*url.URL // URLs is a slice of parsed URLs retrieved from the Wayback Machine.
}

// NewArchive creates a new Archive instance with the specified query and optional HTTP client.
func NewArchive(query string, client *http.Client) (*Archive, error) {
	defaultClient := &http.Client{
		Timeout: time.Second * 10,
	}
	if client == nil {
		client = defaultClient
	}

	return &Archive{
		Query:      query,
		HTTPClient: client,
	}, nil
}

// FetchURLs fetches URLs from the Wayback Machine for the specified query.
// It queries the Wayback Machine CDX API and parses the retrieved URLs.
func (a *Archive) FetchURLs() (*Result, error) {
	var result []*url.URL
	waybackURL := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=%s/*&output=txt&collapse=urlkey&fl=original&page=/", a.Query)
	req, err := http.NewRequest("GET", waybackURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	URLs := strings.Split(string(body), "\n")

	for _, u := range URLs {
		if u != "" {
			parsedURL, err := url.ParseRequestURI(u)
			if err == nil {
				result = append(result, parsedURL)
			}
		}
	}

	return &Result{URLs: result}, nil
}

// FormatAsJSON formats the result of URLs as JSON.
// It returns a JSON representation of the Result struct.
func (r *Result) FormatAsJSON() ([]byte, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// hasParams checks if a URL has parameters.
// It is used internally for local filtering.
func hasParams(u *url.URL) bool {
	return u.RawQuery != ""
}

// HasParams filters URLs with parameters from the result.
// It returns a slice of URLs with parameters.
func (r *Result) HasParams() ([]*url.URL, error) {
	var filteredURLs []*url.URL
	for _, u := range r.URLs {
		if hasParams(u) {
			filteredURLs = append(filteredURLs, u)
		}
	}
	return filteredURLs, nil
}

// FilterByExtension filters URLs by a specific file extension from the result.
// It returns a slice of URLs with the specified extension.
func (r *Result) FilterByExtension(ext string) ([]*url.URL, error) {
	var filteredURLs []*url.URL
	for _, u := range r.URLs {
		if hasParams(u) {
			if filepath.Ext(u.Path) == ext {
				filteredURLs = append(filteredURLs, u)
			}
		}
	}
	return filteredURLs, nil
}
