package command

import (
	"fmt"
	"isso0424/spotify_CLI/selfmadetypes"
)

func help(commands []selfmadetypes.Command) {
	for _, command := range commands {
		commandHelp := command.GetHelp()
		fmt.Printf(
			"-------------------------------\n"+
				"%s\n"+
				"-------------------------------\n"+
				"Kind: %s\n"+
				"Description: %s\n\n",
			commandHelp.Name,
			commandHelp.Kind,
			commandHelp.Explain,
		)
	}
}

// GetHelp is getting help function.
func (cmd status) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "get playing status in spotify",
	}
}

// GetHelp is getting help function.
func (cmd next) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play next queuing track",
	}
}

// GetHelp is getting help function.
func (cmd pause) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "stop playing track",
	}
}

// GetHelp is getting help function.
func (cmd play) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play track from url",
	}
}

// GetHelp is getting help function.
func (cmd prev) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play previous track",
	}
}

// GetHelp is getting help function.
func (cmd repeat) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch repeat mode",
	}
}

// GetHelp is getting help function.
func (cmd resume) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "resume stopping track",
	}
}

// GetHelp is getting help function.
func (cmd shuffle) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch shuffle state",
	}
}

// GetHelp is getting help function.
func (cmd welcome) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch shuffle state",
	}
}

// GetHelp is getting help function.
func (cmd save) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "loadfile",
		Explain: "save playlist to file",
	}
}

// GetHelp is getting help function.
func (cmd show) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "loadfile",
		Explain: "show saved all playlists",
	}
}

// GetHelp is getting help function.
func (cmd random) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "requestAndLoadfile",
		Explain: "play random playlist from play",
	}
}

func (cmd load) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "requestAndLoadfile",
		Explain: "play saved playlist",
	}
}

// GetHelp is getting help function.
func (cmd refresh) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "refresh access token",
	}
}

// GetHelp is getting help function.
func (cmd volume) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "set volume percent",
	}
}

// GetHelp is getting help function.
func (cmd search) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "search with spotify",
		Kind:    "request",
	}
}

// GetHelp is getting help function.
func (cmd favoriteTrack) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "To be favorite playing track.",
		Kind:    "request",
	}
}

// GetHelp is getting help function.
func (cmd importOwnPlaylists) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Import user playlists",
		Kind:    "request",
	}
}

// GetHelp is getting help function.
func (cmd recent) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Show recently played track",
		Kind:    "request",
	}
}

// GetHelp is getting help function.
func (cmd playlist) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Explain: "Show playing playlist detail",
	}
}
