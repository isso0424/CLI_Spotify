package selfmadetypes

// PlayList is playlist information struct
type PlayList struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}

// PlayListFromRequest is playlist information from request
type PlayListFromRequest struct {
	Name  string `json:"name"`
	Owner User   `json:"owner"`
}
