package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/file"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"strings"
)

// Execute is excution command function.
func (cmd recent) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me/player/recently-played?limit=1", nil)
	if err != nil {
		return
	}

	var recentPlayedTracks selfmadetypes.RecentPlayedTracks
	err = json.Unmarshal(response, &recentPlayedTracks)
	if err != nil {
		return
	}

	recentPlayedTrack := recentPlayedTracks.Items[0]
	artistNames := getArtistsName(recentPlayedTrack.Track.Artists)

	fmt.Printf(
		"TrackName: %s\n"+
			"Artist:    %s\n",
		recentPlayedTrack.Track.Name,
		artistNames,
	)

	return
}

// Execute is excution command function.
func (cmd playlist) Execute(token *string) (err error) {
	playlistID, err := getPlayingPlaylistID(token)
	if err != nil {
		return
	}

	response, _, err := request.CreateRequest(
		token,
		selfmadetypes.GET,
		fmt.Sprintf("/playlists/%s?fields=name,owner,followers,tracks.total", *playlistID),
		nil,
	)
	if err != nil {
		return
	}

	var playlistDetails selfmadetypes.PlayListFromRequest
	err = json.Unmarshal(response, &playlistDetails)
	if err != nil {
		return
	}

	fmt.Printf(
		"Playlist detail\n"+
			"---------------\n"+
			"Name: %s\n"+
			"Owner: %s\n"+
			"Followers: %d users\n"+
			"Tracks: %d track(s)\n\n",
		playlistDetails.Name,
		playlistDetails.Owner.DisplayName,
		playlistDetails.Followers.Total,
		playlistDetails.Tracks.Total,
	)

	return err
}

// Execute is excution command function.
func (cmd favoriteTrack) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me/player/currently-playing", nil)
	if err != nil {
		return
	}

	var playingStatus selfmadetypes.CurrentPlayStatus

	response = bytes.Trim(response, "\x00")
	err = json.Unmarshal(response, &playingStatus)
	if err != nil {
		return
	}

	id := strings.Split(playingStatus.Item.URI, ":")[2]
	_, _, err = request.CreateRequest(token, selfmadetypes.PUT, fmt.Sprintf("/me/tracks?ids=%s", id), nil)
	if err != nil {
		return
	}

	fmt.Printf("Success add '%s' to your favorite song!!!\n", playingStatus.Item.Name)

	return
}

// Execute is excution command function.
func (cmd importOwnPlaylists) Execute(token *string) (err error) {
	response, _, err := request.CreateRequest(token, selfmadetypes.GET, "/me/playlists", nil)
	if err != nil {
		return
	}

	var userPlayLists selfmadetypes.UserPlaylists
	err = json.Unmarshal(response, &userPlayLists)
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
