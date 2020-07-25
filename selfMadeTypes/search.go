package selfMadeTypes

import "fmt"

type SearchResponse struct {
	Album    searchAlbum    `json:"albums"`
	Artists  searchArtists  `json:"artists"`
	Track    searchTrack    `json:"tracks"`
	Playlist searchPlayList `json:"playlists"`
	Show     searchShow     `json:"shows"`
	Episode  searchEpisode  `json:"episodes"`
}

func (response SearchResponse) ParseAndPrint(kinds []string) {
	for _, kind := range kinds {
		switch kind {
		case "album":
			fmt.Println("-----Albums-----")
			for _, album := range response.Album.Item {
				fmt.Printf("Name: %s\nURI: %s\n---------------\n", album.Name, album.Uri)
			}
		case "artist":
			fmt.Println("-----Artists-----")
			for _, artist := range response.Artists.Item {
				fmt.Printf("Name: %s\nURI: %s\n---------------\n", artist.Name, artist.Uri)
			}
		case "track":
			fmt.Println("-----Tracks-----")
			for _, track := range response.Track.Item {
				fmt.Printf("Name: %s\nURI: %s\n---------------\n", track.Name, track.Uri)
			}
		case "playlist":
			fmt.Println("-----Playlists-----")
			for _, playlist := range response.Playlist.Item {
				fmt.Printf("Name: %s\nURI: %s\n---------------\n", playlist.Name, playlist.Uri)
			}
		case "show":
			fmt.Println("-----Shows-----")
			for _, show := range response.Show.Item {
				fmt.Printf("Name: %s\nURI: %s\n---------------\n", show.Name, show.Uri)
			}
		case "episode":
			fmt.Println("-----Episodes-----")
			for _, episode := range response.Episode.Item {
				fmt.Printf("Name: %s\nURI: %s\n---------------\n", episode.Name, episode.Uri)
			}
		}
	}
}

type searchAlbum struct {
	Item []album `json:"items"`
}

type searchArtists struct {
	Item []Artists `json:"items"`
}

type searchPlayList struct {
	Item []PlayList `json:"items"`
}

type searchShow struct {
	Item []show `json:"items"`
}

type searchEpisode struct {
	Item []episode `json:"items"`
}

type searchTrack struct {
	Item []track `json:"items"`
}

type track struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type show struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type episode struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}
