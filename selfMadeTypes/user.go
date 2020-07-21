package selfMadeTypes

type user struct {
	DisplayName  string       `json:"display_name"`
	ExternalUrls externalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Id           string       `json:"id"`
	Type         string       `json:"type"`
	Uri          string       `json:"uri"`
}

type followers struct {
	Href  string `json:"href"`
	Total int32  `json:"total"`
}
