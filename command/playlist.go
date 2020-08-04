package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes/commandTypes"
	request2 "isso0424/spotify_CLI/selfmadetypes/request"
	response2 "isso0424/spotify_CLI/selfmadetypes/response"
	"isso0424/spotify_CLI/util"
)

type addToPlaylist struct{}

func (cmd addToPlaylist) GetCommandName() string {
	return "addToPlaylist"
}

func (cmd addToPlaylist) GetHelp() commandTypes.CommandHelp {
	return commandTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandTypes.Playlist,
		Explain: "Edit user's playlist",
	}
}

func (cmd addToPlaylist) Execute(token *string) (err error) {
	playlistID := util.Input("Please input playlist id", "PlaylistID")

	addTrackID := util.Input("Please input track id", "TrackID")
	addTrackURI := fmt.Sprintf("spotify:track:%s", addTrackID)

	response, err := request.CreateRequest(
		token,
		request2.POST,
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

	if response.GetStatusCode() == request.Created {
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

func (cmd createPlaylist) GetHelp() commandTypes.CommandHelp {
	return commandTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandTypes.Playlist,
		Explain: "Create new playlist",
	}
}

func (cmd createPlaylist) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, request2.GET, "/me", nil)
	if err != nil {
		return
	}

	var user response2.User
	err = json.Unmarshal(response.GetBody(), &user)
	if err != nil {
		return
	}

	userID := user.ID

	playlistName := util.Input("Please input new playlist name.", "Playlist name")

	values, err := json.Marshal(map[string]string{"name": playlistName})
	if err != nil {
		return
	}

	response, err = request.CreateRequest(
		token,
		request2.POST,
		fmt.Sprintf("/users/%s/playlists", userID),
		bytes.NewBuffer(values),
	)
	if err != nil {
		return
	}

	if response.GetStatusCode() == request.Ok || response.GetStatusCode() == request.Created {
		fmt.Println("Successful created playlist!!!")
	} else {
		fmt.Println("Failed create playlist.")
	}

	return err
}

type deleteTrackFromPlaylist struct{}

func (cmd deleteTrackFromPlaylist) GetCommandName() string {
	return "deleteTrackFromPlaylist"
}

func (cmd deleteTrackFromPlaylist) GetHelp() commandTypes.CommandHelp {
	return commandTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandTypes.Playlist,
		Explain: "Delete track from playlist",
	}
}

func (cmd deleteTrackFromPlaylist) Execute(token *string) (err error) {
	playlistID := util.Input("Please input playlist id", "PlaylistID")

	addTrackID := util.Input("Please input track id", "TrackID")
	addTrackURI := fmt.Sprintf("spotify:track:%s", addTrackID)

	body, err := json.Marshal(map[string][]map[string]string{"tracks": {{"uri": addTrackURI}}})
	if err != nil {
		return
	}
	response, err := request.CreateRequest(
		token,
		request2.DELETE,
		fmt.Sprintf(
			"/playlists/%s/tracks",
			playlistID,
		),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return
	}

	if response.GetStatusCode() == request.Ok {
		fmt.Println("Successful delete track!!!")
	} else {
		fmt.Println("Failed delete track.")
	}

	return err
}
