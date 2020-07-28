package selfmadetypes

// RecentPlayedTracks is response of Get recently played track.
type RecentPlayedTracks struct{
	Items []recentPlayedItem `json:"items"`
}

type recentPlayedItem struct{
	Track trackSimplified `json:"track"`
}

type trackSimplified struct{
	Artists []Artists `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	DiscNumber int `json:"disc_number"`
	Duration int `json:"duration_ms"`
	ExternalUrls externalUrls `json:"external_urls"`
	Href string `json:"href"`
	ID string `json:"id"`
	IsPlayable bool `json:"is_playable"`
	Name string `json:"name"`
	PreviewURL string `json:"preview_url"`
	TrackNumber int `json:"track_number"`
	Type string `json:"type"`
	URI string `json:"uri"`
}
