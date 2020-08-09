package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestToProcessResponse is test to process response.
func TestToProcessResponse(t *testing.T) {
	item := searchItem{
		Item: []SearchResultItem{
			{
				Name: "itemName",
				URI:  "itemURI",
			},
		},
	}

	result := toProcessResponse(item, []SearchResultItem{})

	assert.Equal(t, result[0].Name, "itemName")
	assert.Equal(t, result[0].URI, "itemURI")
}
