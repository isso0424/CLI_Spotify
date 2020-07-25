package command

import "isso0424/spotify_CLI/selfMadeTypes"

func (cmd status) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "get playing status in spotify",
	}
}

func (cmd next) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play next queuing track",
	}
}

func (cmd pause) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "stop playing track",
	}
}

func (cmd play) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play track from url",
	}
}

func (cmd prev) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "play previous track",
	}
}

func (cmd repeat) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch repeat mode",
	}
}

func (cmd resume) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "resume stopping track",
	}
}

func (cmd shuffle) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch shuffle state",
	}
}

func (cmd welcome) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch shuffle state",
	}
}

func (cmd save) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "loadfile",
		Explain: "save playlist to file",
	}
}

func (cmd show) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "loadfile",
		Explain: "show saved all playlists",
	}
}

func (cmd random) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "requestAndLoadfile",
		Explain: "play random playlist from play",
	}
}

func (cmd load) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "requestAndLoadfile",
		Explain: "play saved playlist",
	}
}

func (cmd refresh) GetHelp() selfMadeTypes.CommandHelp {
	return selfMadeTypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "refresh access token",
	}
}
