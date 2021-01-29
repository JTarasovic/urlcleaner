package main

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/gobwas/glob"
	"github.com/stretchr/testify/assert"
)

func TestFilterQueryParams(t *testing.T) {
	var testCases = []struct {
		name     string
		url      string
		expected map[string]string
	}{
		{
			name: "no params",
			url:  "https://example.com",
			expected: map[string]string{
				"*":           "https://example.com",
				"utm*":        "https://example.com",
				"doesntmatch": "https://example.com",
			},
		},
		{
			name: "utm params",
			url:  "https://www.servethehome.com/server-industry-takeaways-from-intel-q4-2020-earnings/?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+ServeTheHome+%28ServeTheHome.com%29",
			expected: map[string]string{
				"*":    "https://www.servethehome.com/server-industry-takeaways-from-intel-q4-2020-earnings/",
				"utm*": "https://www.servethehome.com/server-industry-takeaways-from-intel-q4-2020-earnings/",
				// params are re-writen in order regardless of match
				"*doesntmatch": "https://www.servethehome.com/server-industry-takeaways-from-intel-q4-2020-earnings/?utm_campaign=Feed%3A+ServeTheHome+%28ServeTheHome.com%29&utm_medium=feed&utm_source=feedburner",
			},
		},
	}

	for _, tc := range testCases {
		for glb, expected := range tc.expected {
			t.Run(fmt.Sprintf("%s %s", tc.name, glb), func(t *testing.T) {
				u, err := url.Parse(tc.url)
				assert.Nil(t, err)
				filterQueryParams(u, glob.MustCompile(glb))
				assert.Equal(t, expected, u.String(), "failed for glob: %s", glb)
			})
		}
	}
}
