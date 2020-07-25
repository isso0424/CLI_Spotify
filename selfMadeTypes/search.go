package selfMadeTypes

type SearchResponse struct {
  Album    []album    `json:"albums"`
  Artists  []Artists  `json:"artists"`
  Track    []track    `json:"tracks"`
  Playlist []PlayList `json:"playlists"`
  Show     []show     `json:"shows"`
  Episode  []episode  `json:"episodes"`
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
