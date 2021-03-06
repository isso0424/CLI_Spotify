package responsetypes

import "isso0424/spotify_CLI/selfmadetypes/search"

// Content is data struct when you execute GET to player
type Content struct {
	IsPlaying            bool                   `json:"is_playing"`
	Device               device                 `json:"device"`
	ShuffleState         bool                   `json:"shuffle_state"`
	RepeatState          string                 `json:"repeat_state"`
	Timestamp            int64                  `json:"timestamp"`
	ProgressMs           int32                  `json:"progress_ms"`
	Item                 Item                   `json:"item"`
	CurrentlyPlayingType string                 `json:"currently_playing_type"`
	Action               map[string]interface{} `json:"actions"`
	Context              context                `json:"context"`
}

type device struct {
	ID               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	Volume           int32  `json:"volume_percent"`
}

type context struct {
	ExternalUrls externalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type externalUrls struct {
	Spotify string `json:"spotify"`
}

// Item is Content's item
type Item struct {
	Album            Album        `json:"album"`
	Artists          []Artists    `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int32        `json:"disc_number"`
	DurationsMs      int64        `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      externalIds  `json:"external_ids"`
	ExternalUrls     externalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsLocal          bool         `json:"is_local"`
	Name             string       `json:"name"`
	Popularity       int32        `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int32        `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}

type externalIds struct {
	Isrc string `json:"isrc"`
}

type image struct {
	Height int32  `json:"height"`
	URL    string `json:"url"`
	Width  int32  `json:"width"`
}

// Artists is Content's artists value
type Artists struct {
	ExternalUrls externalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

// CurrentPlayStatus is current playing song status.
type CurrentPlayStatus struct {
	Item search.ResultItem `json:"item"`
}

// RecentPlayedTracks is responseTypes of Get recently played track.
type RecentPlayedTracks struct {
	Items []recentPlayedItem `json:"items"`
}

type recentPlayedItem struct {
	Track trackSimplified `json:"track"`
}

type trackSimplified struct {
	Artists          []Artists    `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	Duration         int          `json:"duration_ms"`
	ExternalUrls     externalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	Name             string       `json:"name"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}
