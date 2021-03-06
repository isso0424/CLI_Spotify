package command

import (
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/selfmadetypes/commandtypes"
	"isso0424/spotify_CLI/selfmadetypes/requesttypes"
	"isso0424/spotify_CLI/selfmadetypes/responsetypes"
	search2 "isso0424/spotify_CLI/selfmadetypes/search"
	commanderrors "isso0424/spotify_CLI/selfmadetypes/selfmadeerrors"
	"isso0424/spotify_CLI/util"
)

type save struct{}

// GetCommandName is getting command name function.
func (cmd save) GetCommandName() string {
	return "save"
}

// GetHelp is getting help function.
func (cmd save) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.LoadFile,
		Explain: "save playlist to file",
	}
}

// Execute is excution command function.
func (cmd save) Execute() (err error) {
	url := util.Input("please input playlist url\n", "PlayListURL")

	uri, err := parse.CreateContextURI(url)
	if err != nil {
		return
	}

	name := util.Input("\nplease input playlist name\n", "PlayListName")

	list := search2.ResultItem{URI: uri, Name: name}

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	if util.CheckDuplicateName(name, playlistList) {
		err = file.SavePlayList(list)
	} else {
		err = &commanderrors.NameDuplicateError{Target: name}
	}

	return
}

type show struct{}

// GetCommandName is getting command name function.
func (cmd show) GetCommandName() string {
	return "show"
}

// GetHelp is getting help function.
func (cmd show) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.LoadFile,
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
		util.Output(
			selfmadetypes.OutputMessage{
				Message: [][]string{
					{
						fmt.Sprintf("id: %d", index),
					},
					{
						fmt.Sprintf("name: %s", target.Name),
						fmt.Sprintf("uri:  %s", target.URI),
					},
				},
			},
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
func (cmd random) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.LoadFile,
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
	if err != nil {
		return
	}
	err = request.PrintPlayingStatus(token)

	return
}

type load struct{}

// GetCommandName is getting command name function.
func (cmd load) GetCommandName() string {
	return "load"
}

func (cmd load) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.LoadFile,
		Explain: "play saved playlist",
	}
}

// Execute is excution command function.
func (cmd load) Execute(token *string) (err error) {
	name := util.Input("please input playlist name", "PlayListName")

	playlistList, err := file.LoadPlayList()

	if err != nil {
		return
	}

	for _, target := range playlistList {
		if target.Name == name {
			util.Output(
				selfmadetypes.OutputMessage{
					Message: [][]string{
						{
							fmt.Sprintf("play %s", target.Name),
						},
					},
				},
			)
			err = request.PlayFromURL(token, target.URI)
			if err != nil {
				return
			}
			err = request.PrintPlayingStatus(token)
			return
		}
	}

	err = &commanderrors.NotFound{Target: name}

	return err
}

type importOwnPlaylists struct{}

// GetCommandName is getting command name function.
func (cmd importOwnPlaylists) GetCommandName() string {
	return "importOwnPlaylists"
}

// GetHelp is getting help function.
func (cmd importOwnPlaylists) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Import user playlists",
		Kind:    commandtypes.LoadFile,
	}
}

// Execute is execution command function.
func (cmd importOwnPlaylists) Execute(token *string) (err error) {
	res, err := request.CreateRequest(token, requesttypes.GET, "/me/playlists", nil)
	if err != nil {
		return
	}

	var userPlayLists responsetypes.UserPlaylists
	err = json.Unmarshal(res.GetBody(), &userPlayLists)
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
