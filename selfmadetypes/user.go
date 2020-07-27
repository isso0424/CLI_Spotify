package selfmadetypes

// User is response when GET request to user
type User struct {
	DisplayName  string       `json:"display_name"`
	ExternalUrls externalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}
