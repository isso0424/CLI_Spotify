package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
	"net/url"
	"strconv"
	"strings"
)

type search struct{}

// GetCommandName is getting command name function.
func (cmd search) GetCommandName() string {
	return "search"
}

// GetHelp is getting help function.
func (cmd search) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "search with spotify",
		Kind:    selfmadetypes.Search,
	}
}

// Execute is excution command function.
func (cmd search) Execute(token *string) (err error) {
	kind := util.Input(
		"please input search kind\n\n"+
			"search kinds: album artist playlist track show episode\n\n"+
			"if input over 2 types, please enter with a colon\n"+
			"------------------------",
		"Kind",
	)
	kinds := strings.Split(kind, ",")
	for _, kind := range kinds {
		if util.ExistTarget(kind, []string{"album", "artist", "playlist", "track", "show", "episode"}) {
			return fmt.Errorf("search type %s is not found", kind)
		}
	}

	keyword := util.Input("Please input search keyword\n------------------------", "Keyword")
	keyword = url.QueryEscape(keyword)

	response, err := request.CreateRequest(
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
	err = json.Unmarshal(response.GetBody(), &searchResponse)
	if err != nil {
		return
	}

	searchResultItems := searchResponse.ParseAndPrint(kinds)

	err = saveSearchResult(searchResultItems)

	return err
}

func saveSearchResult(searchResults []selfmadetypes.SearchResultItem) (err error) {
	isSave := util.Input("Want to save result?\n------------------------", "Want to save?")

	if isSave != "yes" {
		return
	}

	rawIndex := util.Input("Please input index\n------------------------", "Index")

	index, err := strconv.Atoi(rawIndex)
	if err != nil {
		return
	}

	if index >= len(searchResults) {
		return errors.New("index is out of range")
	}

	item := searchResults[index]

	err = file.SavePlayList(item)

	return
}
