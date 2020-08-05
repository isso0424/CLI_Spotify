package responsetypes

// PlayList is playlist information from requestTypes.
type PlayList struct {
	Name      string         `json:"name"`
	Owner     User           `json:"owner"`
	Href      string         `json:"href"`
	Followers followers      `json:"followers"`
	Tracks    playlistPaging `json:"tracks"`
}

type followers struct {
	Href  string `json:"href"`
	Total int    `json:"Total"`
}

type playlistPaging struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
