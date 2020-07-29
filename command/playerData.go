package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"strings"
)

type recent struct{}

// GetCommandName is getting command name function.
func (cmd recent) GetCommandName() string {
	return "recent"
}

// GetHelp is getting help function.
func (cmd recent) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Show recently played track",
		Kind:    "request",
	}
}

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
	artistNames := parse.GetArtistNames(recentPlayedTrack.Track.Artists)

	fmt.Printf(
		"TrackName: %s\n"+
			"Artist:    %s\n",
		recentPlayedTrack.Track.Name,
		artistNames,
	)

	return
}

type playlist struct{}

// GetCommandName is getting command name function.
func (cmd playlist) GetCommandName() string {
	return "playlist"
}

// GetHelp is getting help function.
func (cmd playlist) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Show playing playlist detail",
	}
}

// Execute is excution command function.
func (cmd playlist) Execute(token *string) (err error) {
	playlistID, err := parse.GetPlayingPlaylistID(token)
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

type favoriteTrack struct{}

// GetCommandName is getting command name function.
func (cmd favoriteTrack) GetCommandName() string {
	return "favoriteTrack"
}

// GetHelp is getting help function.
func (cmd favoriteTrack) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "To be favorite playing track.",
		Kind:    "request",
	}
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
