package responsetypes

import "fmt"

// SearchResponse is searched responseTypes
type SearchResponse struct {
	Album    searchItem `json:"albums"`
	Artists  searchItem `json:"artists"`
	Track    searchItem `json:"tracks"`
	Playlist searchItem `json:"playlists"`
	Show     searchItem `json:"shows"`
	Episode  searchItem `json:"episodes"`
}

// ParseAndPrint is function that responseTypes parse and print
func (response SearchResponse) ParseAndPrint(kinds []string) []SearchResultItem {
	searchResults := []SearchResultItem{}
	for _, kind := range kinds {
		switch kind {
		case "album":
			fmt.Println("-----Albums-----")
			searchResults = toProcessResponse(response.Album, searchResults)
		case "artist":
			fmt.Println("-----Artists-----")
			searchResults = toProcessResponse(response.Artists, searchResults)
		case "track":
			fmt.Println("-----Tracks-----")
			searchResults = toProcessResponse(response.Track, searchResults)
		case "playlist":
			fmt.Println("-----Playlists-----")
			searchResults = toProcessResponse(response.Playlist, searchResults)
		case "show":
			fmt.Println("-----Shows-----")
			searchResults = toProcessResponse(response.Show, searchResults)
		case "episode":
			fmt.Println("-----Episodes-----")
			searchResults = toProcessResponse(response.Episode, searchResults)
		}
	}

	return searchResults
}

func toProcessResponse(items searchItem, resultSlice []SearchResultItem) []SearchResultItem {
	for _, item := range items.Item {
		index := len(resultSlice)
		fmt.Printf("ID: %d\nName: %s\nURI: %s\n---------------\n", index, item.Name, item.URI)
		resultSlice = append(resultSlice, item)
	}

	return resultSlice
}

type searchItem struct {
	Item []SearchResultItem `json:"items"`
}

// SearchResultItem is SearchResultItem's item
type SearchResultItem struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}
