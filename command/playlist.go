package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/util"
)

type addToPlaylist struct{}

func (cmd addToPlaylist) GetCommandName() string {
	return "addToPlaylist"
}

func (cmd addToPlaylist) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "Edit user's playlist",
	}
}

func (cmd addToPlaylist) Execute(token *string) (err error) {
	var playlistID string
	util.Input("Please input playlist id", "PlaylistID", &playlistID)

	var addTrackID string
	util.Input("Please input track id", "TrackID", &addTrackID)
	addTrackURI := fmt.Sprintf("spotify:track:%s", addTrackID)

	_, statusCode, err := request.CreateRequest(
		token,
		selfmadetypes.POST,
		fmt.Sprintf(
			"/playlists/%s/tracks?uris=%s",
			playlistID,
			addTrackURI,
		),
		nil,
	)
	if err != nil {
		return
	}

	if statusCode == 201 {
		fmt.Println("Successful added!!!")
	} else {
		fmt.Println("Track add failed")
	}

	return err
}

type createPlaylist struct{}

func (cmd createPlaylist) GetCommandName() string {
	return "createPlaylist"
}

func (cmd createPlaylist) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "Create new playlist",
	}
}

func (cmd createPlaylist) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me", nil)
	if err != nil {
		return
	}

	var user selfmadetypes.User
	err = json.Unmarshal(response, &user)
	if err != nil {
		return
	}

	userID := user.ID

	var playlistName string
	util.Input("Please input new playlist name.", "Playlist name", &playlistName)

	values, err := json.Marshal(map[string]string{"name": playlistName})
	if err != nil {
		return
	}

	_, statusCode, err := request.CreateRequest(token, selfmadetypes.POST, fmt.Sprintf("/users/%s/playlists", userID), bytes.NewBuffer(values))
	if err != nil {
		return
	}

	if statusCode == 200 || statusCode == 201 {
		fmt.Println("Successful created playlist!!!")
	} else {
		fmt.Println("Failed create playlist.")
	}

	return
}
