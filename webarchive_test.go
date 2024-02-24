// The full code is still under testing
package webarchive

import (
	"net/url"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasParamWithParam(t *testing.T) {
	// Create a URL with parameters
	u, err := url.Parse("https://example.com/path?param=value")
	assert.NoError(t, err)

	// Call the hasParam function
	result := hasParams(u)

	// Assertions
	assert.True(t, result)
}
func TestHasParamWithoutParam(t *testing.T) {
	// Create a URL without parameters
	u, err := url.Parse("https://example.com/path")
	assert.NoError(t, err)

	// Call the hasParam function
	result := hasParams(u)

	// Assertions
	assert.False(t, result)
}

func TestMethodHasParam(t *testing.T) {
	// Create an instance of the Result
	result := &Result{
		URLs: []*url.URL{
			{RawQuery: "param1=value1"},
			{RawQuery: ""},
			{RawQuery: "param2=value2"},
		},
	}

	filteredURLs, err := result.HasParams()

	// Assert that no error
	assert.NoError(t, err)

	// Assert that the number of filteredURLs is exepected
	assert.Len(t, filteredURLs, 2)

	// Assert that the filtered parameters is actually valid
	assert.Equal(t, "param1=value1", filteredURLs[0].RawQuery)
	assert.Equal(t, "param2=value2", filteredURLs[1].RawQuery)
}
