package selfMadeTypes

type PlayList struct {
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

type PlayListFromRequest struct {
	Name  string `json:"name"`
	Owner User   `json:"owner"`
}
