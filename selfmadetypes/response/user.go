package response

// User is response when GET request to user
type User struct {
	DisplayName  string       `json:"display_name"`
	ExternalUrls externalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

// UserPlaylists is user's all playlists.
type UserPlaylists struct {
	Item []SearchResultItem `json:"items"`
}
