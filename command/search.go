package command

import (
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
	"net/url"
	"strings"
)

// Execute is excution command function.
func (cmd search) Execute(token *string) (err error) {
	var kind string
	util.Input(
		"please input search kind\n\n"+
			"search kinds: album artist playlist track show episode\n\n"+
			"if input over 2 types, please enter with a colon\n"+
			"------------------------",
		"Kind",
		&kind,
	)
	kinds := strings.Split(kind, ",")
	for _, kind := range kinds {
		if existTarget(kind, []string{"album", "artist", "playlist", "track", "show", "episode"}) {
			return fmt.Errorf("search type %s is not found", kind)
		}
	}

	var keyword string
	util.Input("Please input search keyword\n------------------------", "Keyword", &keyword)
	keyword = url.QueryEscape(keyword)

	response, _, err := request.CreateRequest(
		token,
		selfmadetypes.GET,
		fmt.Sprintf(
			"/search?q=%s&type=%s",
			keyword,
			kind,
		),
		nil,
	)
	if err != nil {
		return
	}

	var searchResponse selfmadetypes.SearchResponse
	err = json.Unmarshal(response, &searchResponse)
	if err != nil {
		return
	}

	searchResultItems := searchResponse.ParseAndPrint(kinds)

	err = saveSearchResult(searchResultItems)

	return err
}
