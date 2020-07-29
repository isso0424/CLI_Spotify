package command

import (
	"errors"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
	"strconv"
)

type next struct{}

// GetCommandName is getting command name function.
func (cmd next) GetCommandName() string {
	return "next"
}

// GetHelp is getting help function.
func (cmd next) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play next queuing track",
	}
}

// Execute is excution command function.
func (cmd next) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.POST, "/me/player/next", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

type pause struct{}

// GetCommandName is getting command name function.
func (cmd pause) GetCommandName() string {
	return "pause"
}

// GetHelp is getting help function.
func (cmd pause) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "stop playing track",
	}
}

// Execute is excution command function.
func (cmd pause) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, "/me/player/pause", nil)

	if err != nil {
		return
	}
	fmt.Println("paused!!!")

	return
}

type play struct{}

// GetCommandName is getting command name function.
func (cmd play) GetCommandName() string {
	return "play"
}

// GetHelp is getting help function.
func (cmd play) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play track from url",
	}
}

// Execute is excution command function.
func (cmd play) Execute(token *string) (err error) {
	var href string
	util.Input("please input playlist href\n------------------------", "PlayListURL", &href)

	uri, err := parse.CreateContextURI(href)
	if err != nil {
		return
	}
	err = request.PlayFromURL(token, *uri)

	return
}

type prev struct{}

// GetCommandName is getting command name function.
func (cmd prev) GetCommandName() string {
	return "prev"
}

// GetHelp is getting help function.
func (cmd prev) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play previous track",
	}
}

// Execute is excution command function.
func (cmd prev) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.POST, "/me/player/previous", nil)

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
func (cmd status) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "get playing status in spotify",
	}
}

// Execute is excution command function.
func (cmd status) Execute(token *string) (err error) {
	err = request.PrintPlayingStatus(token)

	return
}

// GetCommandName is getting command name function.
type repeat struct{}

// GetCommandName is getting command name function.
func (cmd repeat) GetCommandName() string {
	return "repeat"
}

// GetHelp is getting help function.
func (cmd repeat) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
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

	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, fmt.Sprintf("/me/player/repeat?state=%s", state), nil)

	if err != nil {
		return
	}

	fmt.Printf("Repeat state change to `%s`\n", state)

	return
}

type resume struct{}

// GetCommandName is getting command name function.
func (cmd resume) GetCommandName() string {
	return "resume"
}

// GetHelp is getting help function.
func (cmd resume) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "resume stopping track",
	}
}

// Execute is excution command function.
func (cmd resume) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, "/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
}

type shuffle struct{}

// GetCommandName is getting command name function.
func (cmd shuffle) GetCommandName() string {
	return "shuffle"
}

// GetHelp is getting help function.
func (cmd shuffle) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
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

	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, fmt.Sprintf("/me/player/shuffle?state=%v", state), nil)
	if err != nil {
		return
	}

	fmt.Printf("Switch shuffle state to %v\n", state)

	return
}

type volume struct{}

// GetCommandName is getting command name function.
func (cmd volume) GetCommandName() string {
	return "volume"
}

// GetHelp is getting help function.
func (cmd volume) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "set volume percent",
	}
}

// Execute is excution command function.
func (cmd volume) Execute(token *string) (err error) {
	var percent string
	util.Input("please volume percent\n------------------------", "Volume", &percent)

	percentInt, err := strconv.Atoi(percent)
	if err != nil {
		return
	}

	if percentInt < 0 || percentInt > 100 {
		return errors.New("percent range is 0 to 100")
	}

	_, _, err = request.CreateRequest(
		token,
		selfmadetypes.PUT,
		fmt.Sprintf(
			"/me/player/volume?volume_percent=%s",
			percent,
		),
		nil,
	)

	return
}