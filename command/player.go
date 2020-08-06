package command

import (
	"errors"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/selfmadetypes/commandtypes"
	"isso0424/spotify_CLI/selfmadetypes/requesttypes"
	"isso0424/spotify_CLI/util"
	"strconv"
)

type next struct{}

// GetCommandName is getting command name function.
func (cmd next) GetCommandName() string {
	return "next"
}

// GetHelp is getting help function.
func (cmd next) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "play next queuing track",
	}
}

// Execute is excution command function.
func (cmd next) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, requesttypes.POST, "/me/player/next", nil)

	if err != nil {
		return
	}

	return
}

type pause struct{}

// GetCommandName is getting command name function.
func (cmd pause) GetCommandName() string {
	return "pause"
}

// GetHelp is getting help function.
func (cmd pause) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "stop playing track",
	}
}

// Execute is excution command function.
func (cmd pause) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, requesttypes.PUT, "/me/player/pause", nil)

	if err != nil {
		return
	}

	return
}

type play struct{}

// GetCommandName is getting command name function.
func (cmd play) GetCommandName() string {
	return "play"
}

// GetHelp is getting help function.
func (cmd play) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "play track from url",
	}
}

// Execute is excution command function.
func (cmd play) Execute(token *string) (err error) {
	href := util.Input("please input playlist href\n------------------------", "PlayListURL")

	uri, err := parse.CreateContextURI(href)
	if err != nil {
		return
	}
	err = request.PlayFromURL(token, uri)

	return
}

type prev struct{}

// GetCommandName is getting command name function.
func (cmd prev) GetCommandName() string {
	return "prev"
}

// GetHelp is getting help function.
func (cmd prev) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "play previous track",
	}
}

// Execute is excution command function.
func (cmd prev) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, requesttypes.POST, "/me/player/previous", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

type status struct{}

// GetCommandName is getting command name function.
func (cmd status) GetCommandName() string {
	return "status"
}

// GetHelp is getting help function.
func (cmd status) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "get playing status in spotify",
	}
}

// Execute is excution command function.
func (cmd status) Execute(_ *string) error {
	return nil
}

// GetCommandName is getting command name function.
type repeat struct{}

// GetCommandName is getting command name function.
func (cmd repeat) GetCommandName() string {
	return "repeat"
}

// GetHelp is getting help function.
func (cmd repeat) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "switch repeat mode",
	}
}

// Execute is excution command function.
func (cmd repeat) Execute(token *string) (err error) {
	status, err := request.GetStatus(token)

	if err != nil {
		return
	}

	state := util.SwitchRepeatState(status.RepeatState)

	_, err = request.CreateRequest(token, requesttypes.PUT, fmt.Sprintf("/me/player/repeat?state=%s", state), nil)

	if err != nil {
		return
	}

	util.Output(
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					fmt.Sprintf("Repeat state change to `%s`\n", state),
				},
			},
		},
	)

	return
}

type resume struct{}

// GetCommandName is getting command name function.
func (cmd resume) GetCommandName() string {
	return "resume"
}

// GetHelp is getting help function.
func (cmd resume) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "resume stopping track",
	}
}

// Execute is excution command function.
func (cmd resume) Execute(token *string) (err error) {
	_, err = request.CreateRequest(token, requesttypes.PUT, "/me/player/play", nil)

	if err != nil {
		return
	}

	return
}

type shuffle struct{}

// GetCommandName is getting command name function.
func (cmd shuffle) GetCommandName() string {
	return "shuffle"
}

// GetHelp is getting help function.
func (cmd shuffle) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "switch shuffle state",
	}
}

// Execute is excution command function.
func (cmd shuffle) Execute(token *string) (err error) {
	status, err := request.GetStatus(token)
	if err != nil {
		return
	}

	state := !status.ShuffleState

	_, err = request.CreateRequest(token, requesttypes.PUT, fmt.Sprintf("/me/player/shuffle?state=%v", state), nil)
	if err != nil {
		return
	}

	util.Output(
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					fmt.Sprintf("Switch shuffle state to %v", state),
				},
			},
		},
	)

	return
}

type volume struct{}

// GetCommandName is getting command name function.
func (cmd volume) GetCommandName() string {
	return "volume"
}

// GetHelp is getting help function.
func (cmd volume) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Player,
		Explain: "set volume percent",
	}
}

// Execute is excution command function.
func (cmd volume) Execute(token *string) (err error) {
	percent := util.Input("please volume percent\n------------------------", "Volume")

	percentInt, err := strconv.Atoi(percent)
	if err != nil {
		return
	}

	if percentInt < 0 || percentInt > 100 {
		return errors.New("percent range is 0 to 100")
	}

	_, err = request.CreateRequest(
		token,
		requesttypes.PUT,
		fmt.Sprintf(
			"/me/player/volume?volume_percent=%s",
			percent,
		),
		nil,
	)

	return
}
