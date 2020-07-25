package selfMadeTypes

type SearchResponse struct {
  Album    searchAlbum    `json:"albums"`
  Artists  searchArtists  `json:"artists"`
  Track    searchAlbum    `json:"tracks"`
  Playlist searchPlayList `json:"playlists"`
  Show     searchShow     `json:"shows"`
  Episode  searchEpisode  `json:"episodes"`
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
