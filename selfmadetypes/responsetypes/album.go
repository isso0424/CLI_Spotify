package responsetypes

// Album is responseTypes type for album.
type Album struct {
	AlbumType            string       `json:"album_type"`
	Artists              []Artists    `json:"artists"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalUrls         externalUrls `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	TotalTracks          int32        `json:"total_tracks"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
}
