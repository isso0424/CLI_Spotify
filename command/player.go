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

// Execute is excution command function.
func (cmd next) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.POST, "/me/player/next", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
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

// Execute is excution command function.
func (cmd prev) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.POST, "/me/player/previous", nil)

	if err != nil {
		return
	}

	err = status{}.Execute(token)

	return
}

// Execute is excution command function.
func (cmd status) Execute(token *string) (err error) {
	err = request.PrintPlayingStatus(token)

	return
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

// Execute is excution command function.
func (cmd resume) Execute(token *string) (err error) {
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, "/me/player/play", nil)

	if err != nil {
		return
	}
	fmt.Println("resumed!!!")

	return
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
