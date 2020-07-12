package selfMadeTypes

type Content struct {
  IsPlaying bool `json:"is_playing"`
  Device device `json:"device"`
  ShuffleState bool `json:"shuffle_state"`
  RepeatState string `json:"repeat_state"`
  Timestamp int64 `json:"timestamp"`
  ProgressMs int32 `json:"progress_ms"`
  Item item `json:"item"`
  CurrentlyPlayingType string `json:"currently_playing_type"`
  Action map[string]interface{} `json:"actions"`
  Context context `json:"context"`
}

type device struct {
  Id string `json:"id"`
  IsActive bool `json:"is_active"`
  IsPrivateSession bool `json:"is_private_session"`
  IsRestricted bool `json:"is_restricted"`
  Name string `json:"name"`
  Type string `json:"type"`
  Volume int32 `json:"volume_percent"`
}

type context struct {
  ExternalUrls externalUrls `json:"external_urls"`
  Href string `json:"href"`
  Type string `json:"type"`
  Uri string `json:"uri"`
}

type externalUrls struct {
  Spotify string `json:"spotify"`
}

type item struct {
  Album album `json:"album"`
  Artists []artists `json:"artists"`
  AvailableMarkets []string `json:"available_markets"`
  DiscNumber int32 `json:"disc_number"`
  DurationsMs int64 `json:"duration_ms"`
  Explicit bool `json:"explicit"`
  ExternalIds externalIds `json:"external_ids"`
  ExternalUrls externalUrls `json:"external_urls"`
  Href string `json:"href"`
  Id string `json:"id"`
  IsLocal bool `json:"is_local"`
  Name string `json:"name"`
  Popularity int32 `json:"popularity"`
  PreviewUrl string `json:"preview_url"`
  TrackNumber int32 `json:"track_number"`
  Type string `json:"type"`
  Uri string `json:"uri"`
}

type externalIds struct {
  Isrc string `json:"isrc"`
}

type album struct {
  AlbumType string `json:"album_type"`
  Artists []artists `json:"artists"`
  AvailableMarkets []string `json:"available_markets"`
  ExternalUrls externalUrls `json:"external_urls"`
  Href string `json:"href"`
  Id string `json:"id"`
  Images []image `json:"images"`
  Name string `json:"name"`
  ReleaseDate string `json:"release_date"`
  ReleaseDatePrecision string `json:"release_date_precision"`
  TotalTracks int32 `json:"total_tracks"`
  Type string `json:"type"`
  Uri string `json:"uri"`
}

type image struct {
  Height int32 `json:"height"`
  Url string `json:"url"`
  Width int32 `json:"width"`
}

type artists struct {
  ExternalUrls externalUrls `json:"external_urls"`
  Href string `json:"href"`
  Id string `json:"id"`
  Name string `json:"name"`
  Type string `json:"type"`
  Uri string `json:"uri"`
}
