package selfmadetypes

// Album is response type for album.
type Album struct {
	AlbumType        string    `json:"album_type"`
	Artists          []Artists `json:"artists"`
	AvailableMarkets []string  `json:"available_markets"`
	Name             string    `json:"name"`
}
