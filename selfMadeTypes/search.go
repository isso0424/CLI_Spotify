package selfMadeTypes

import "fmt"

type SearchResponse struct {
	Album    searchItem `json:"albums"`
	Artists  searchItem `json:"artists"`
	Track    searchItem `json:"tracks"`
	Playlist searchItem `json:"playlists"`
	Show     searchItem `json:"shows"`
	Episode  searchItem `json:"episodes"`
}

func (response SearchResponse) ParseAndPrint(kinds []string) []searchResultItem {
  searchResults := []searchResultItem {}
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

func toProcessResponse(items searchItem, resultSlice []searchResultItem) []searchResultItem {
  for _, item := range(items.Item) {
    index := len(resultSlice)
    fmt.Printf("ID: %d\nName: %s\nURI: %s\n---------------\n", index, item.Name, item.Uri)
    resultSlice = append(resultSlice, item)
  }

  return resultSlice
}

type searchItem struct {
  Item []searchResultItem `json:"items"`
}

type searchResultItem struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}
