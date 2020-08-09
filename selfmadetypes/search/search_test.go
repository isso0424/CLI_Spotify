package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToProcessResponse is test to process response.
func TestToProcessResponse(t *testing.T) {
	item := searchItem{
		Item: []ResultItem{
			{
				Name: "itemName",
				URI:  "itemURI",
			},
		},
	}

	result := toProcessResponse(item, []ResultItem{})

	assert.Equal(t, result[0].Name, "itemName")
	assert.Equal(t, result[0].URI, "itemURI")
}

// TestParseAndPrint is test for ParseandPrint
func TestParseAndPrint(t *testing.T) {
	response := Response{
		Artists: searchItem{
			Item: []ResultItem{
				{
					Name: "Artist",
					URI:  "ArtistURI",
				},
			},
		},
		Track: searchItem{
			Item: []ResultItem{
				{
					Name: "Track",
					URI:  "TrackURI",
				},
			},
		},
	}
	kinds := []string{"artist", "track"}

	searchResult := response.ParseAndPrint(kinds)

	assert.Equal(t, searchResult[0].Name, "Artist")
	assert.Equal(t, searchResult[0].URI, "ArtistURI")

	assert.Equal(t, searchResult[1].Name, "Track")
	assert.Equal(t, searchResult[1].URI, "TrackURI")
}
