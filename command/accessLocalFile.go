package command

import (
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
)

type save struct{}

// GetCommandName is getting command name function.
func (cmd save) GetCommandName() string {
	return "save"
}

// GetHelp is getting help function.
func (cmd save) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    selfmadetypes.LoadFile,
		Explain: "save playlist to file",
	}
}

// Execute is excution command function.
func (cmd save) Execute() (err error) {
	var url string
	util.Input("please input playlist url\n", "PlayListURL", &url)

	uri, err := parse.CreateContextURI(url)
	if err != nil {
		return
	}

	var name string
	util.Input("\nplease input playlist name\n", "PlayListName", &name)

	list := selfmadetypes.SearchResultItem{URI: *uri, Name: name}

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	if util.CheckDuplicateName(name, playlistList) {
		err = file.SavePlayList(list)
	} else {
		err = &selfmadetypes.NameDuplicateError{Target: name}
	}

	return
}

type show struct{}

// GetCommandName is getting command name function.
func (cmd show) GetCommandName() string {
	return "show"
}

// GetHelp is getting help function.
func (cmd show) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    selfmadetypes.LoadFile,
		Explain: "show saved all playlists",
	}
}

// Execute is execution command function.
func (cmd show) Execute() (err error) {
	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for index, target := range playlistList {
		fmt.Printf(
			"id: %d\n------------------------------------------------\nname: %s\nuri: %s\n\n",
			index,
			target.Name,
			target.URI,
		)
	}

	return
}

type random struct {
}

// GetCommandName is getting command name function.
func (cmd random) GetCommandName() string {
	return "random"
}

// GetHelp is getting help function.
func (cmd random) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    selfmadetypes.LoadFile,
		Explain: "play random playlist from play",
	}
}

// Execute is excution command function.
func (cmd random) Execute(token *string) (err error) {
	playlists, err := file.LoadPlayList()
	if err != nil {
		return
	}

	choicePlaylist := util.Choose(playlists)
	err = request.PlayFromURL(token, choicePlaylist.URI)
	err = request.PrintPlayingStatus(token)

	return
}

type load struct{}

// GetCommandName is getting command name function.
func (cmd load) GetCommandName() string {
	return "load"
}

func (cmd load) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    selfmadetypes.LoadFile,
		Explain: "play saved playlist",
	}
}

// Execute is excution command function.
func (cmd load) Execute(token *string) (err error) {
	var name string
	util.Input("please input playlist name", "PlayListName", &name)

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for _, target := range playlistList {
		if target.Name == name {
			fmt.Printf("play %s\n", target.Name)
			err = request.PlayFromURL(token, target.URI)
			err = request.PrintPlayingStatus(token)
			return
		}
	}

	err = &selfmadetypes.NotFound{Target: name}

	return
}

type importOwnPlaylists struct{}

// GetCommandName is getting command name function.
func (cmd importOwnPlaylists) GetCommandName() string {
	return "importOwnPlaylists"
}

// GetHelp is getting help function.
func (cmd importOwnPlaylists) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Import user playlists",
		Kind:    selfmadetypes.LoadFile,
	}
}

// Execute is execution command function.
func (cmd importOwnPlaylists) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, selfmadetypes.GET, "/me/playlists", nil)
	if err != nil {
		return
	}

	var userPlayLists selfmadetypes.UserPlaylists
	err = json.Unmarshal(response.GetBody(), &userPlayLists)
	if err != nil {
		return
	}

	for _, playlist := range userPlayLists.Item {
		err = file.SavePlayList(playlist)
		if err != nil {
			return
		}
	}

	return
}
