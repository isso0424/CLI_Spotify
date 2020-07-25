package command

import "isso0424/spotify_CLI/selfMadeTypes"

func(cmd status) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "get playing status in spotify",
  }
}

func(cmd next) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "play next queuing track",
  }
}

func(cmd pause) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "stop playing track",
  }
}

func(cmd play) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "play track from url",
  }
}

func (cmd prev) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "play previous track",
  }
}

func(cmd repeat) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "switch repeat mode",
  }
}

func(cmd resume) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "resume stopping track",
  }
}

func(cmd shuffle) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "switch shuffle state",
  }
}

func(cmd welcome) getHelp() selfMadeTypes.CommandHelp {
  return selfMadeTypes.CommandHelp{
    Name: cmd.getCommandName(),
    Kind: "request",
    Explain: "switch shuffle state",
  }
}
