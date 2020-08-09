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

// TestParseAndPrint is test for ParseandPrint
func TestParseAndPrint(t *testing.T) {
	response := SearchResponse{
		Artists: searchItem{
			Item: []SearchResultItem{
				{
					Name: "Artist",
					URI:  "ArtistURI",
				},
			},
		},
		Track: searchItem{
			Item: []SearchResultItem{
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
