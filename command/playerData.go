package command

import (
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/command/parse"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/selfmadetypes/commandtypes"
	"isso0424/spotify_CLI/selfmadetypes/requesttypes"
	"isso0424/spotify_CLI/selfmadetypes/responsetypes"
	"isso0424/spotify_CLI/util"
	"strings"
)

type recent struct{}

// GetCommandName is getting command name function.
func (cmd recent) GetCommandName() string {
	return "recent"
}

// GetHelp is getting help function.
func (cmd recent) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Show recently played track",
		Kind:    commandtypes.PlayerData,
	}
}

// Execute is excution command function.
func (cmd recent) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, requesttypes.GET, "/me/player/recently-played?limit=1", nil)
	if err != nil {
		return
	}

	var recentPlayedTracks responsetypes.RecentPlayedTracks
	err = json.Unmarshal(response.GetBody(), &recentPlayedTracks)
	if err != nil {
		return
	}

	recentPlayedTrack := recentPlayedTracks.Items[0]
	artistNames := parse.GetArtistNames(recentPlayedTrack.Track.Artists)

	util.Output(
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					"TrackName: " + recentPlayedTrack.Track.Name,
					"Artist:    " + artistNames,
				},
			},
		},
	)

	return
}

type playlist struct{}

// GetCommandName is getting command name function.
func (cmd playlist) GetCommandName() string {
	return "playlist"
}

// GetHelp is getting help function.
func (cmd playlist) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Show playing playlist detail",
	}
}

// Execute is excution command function.
func (cmd playlist) Execute(token *string) (err error) {
	playlistID, err := request.GetPlayingPlaylistID(token)
	if err != nil {
		return
	}

	response, err := request.CreateRequest(
		token,
		requesttypes.GET,
		fmt.Sprintf("/playlists/%s?fields=name,owner,followers,tracks.total", *playlistID),
		nil,
	)
	if err != nil {
		return
	}

	var playlistDetails responsetypes.PlayList
	err = json.Unmarshal(response.GetBody(), &playlistDetails)
	if err != nil {
		return
	}

	util.Output(
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					"Playlist detail",
				},
				{
					"Name: " + playlistDetails.Name,
					"Owner: " + playlistDetails.Owner.DisplayName,
					fmt.Sprintf("Followers: %d users", playlistDetails.Followers.Total),
					fmt.Sprintf("Tracks: %d track(s)", playlistDetails.Tracks.Total),
				},
			},
		},
	)

	return err
}

type favoriteTrack struct{}

// GetCommandName is getting command name function.
func (cmd favoriteTrack) GetCommandName() string {
	return "favoriteTrack"
}

// GetHelp is getting help function.
func (cmd favoriteTrack) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "To be favorite playing track.",
		Kind:    commandtypes.PlayerData,
	}
}

// Execute is excution command function.
func (cmd favoriteTrack) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, requesttypes.GET, "/me/player/currently-playing", nil)
	if err != nil {
		return
	}

	var playingStatus responsetypes.CurrentPlayStatus

	err = json.Unmarshal(response.GetBody(), &playingStatus)
	if err != nil {
		return
	}

	id := strings.Split(playingStatus.Item.URI, ":")[2]
	_, err = request.CreateRequest(token, requesttypes.PUT, fmt.Sprintf("/me/tracks?ids=%s", id), nil)
	if err != nil {
		return
	}

	util.Output(
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					fmt.Sprintf("Success add '%s' to your favorite song!!!\n", playingStatus.Item.Name),
				},
			},
		},
	)

	return
}
