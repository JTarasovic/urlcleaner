package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanURL(t *testing.T) {
	var testCases = []struct {
		url      string
		expected string
	}{
		{
			url:      "https://example.com",
			expected: "https://example.com",
		},
		{
			url:      "http://feedproxy.google.com/~r/ServeTheHome/~3/Lj0FaF7ZHcQ/",
			expected: "https://www.servethehome.com/microchip-nvme-sas-4-sata-smartroc-3200-smartioc-2200-launched/",
		},
	}

	for _, tc := range testCases {
		t.Run("invalid", func(t *testing.T) {
			actual, err := cleanURL(tc.url)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
func TestMain_invalid(t *testing.T) {
	var testCases = []struct {
		url string
	}{
		{
			url: "not a real url",
		},
		{
			url: "https://ihopeandpraytothetestinggodsthatthisurldoesntexist.tld",
		},
	}

	for _, tc := range testCases {
		t.Run("invalid", func(t *testing.T) {
			actual, err := cleanURL(tc.url)
			assert.Equal(t, "", actual)
			assert.NotNil(t, err)
		})
	}
}
