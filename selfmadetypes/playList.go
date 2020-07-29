package selfmadetypes

// PlayList is playlist information struct.
type PlayList struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}

// PlayListFromRequest is playlist information from request.
type PlayListFromRequest struct {
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
