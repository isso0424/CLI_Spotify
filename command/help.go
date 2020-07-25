package command

import "isso0424/spotify_CLI/selfMadeTypes"

func(_ getPlayStatus) getHelp() selfMadeTypes.CommandHelp {
  name := "status"
  kind := "request"
  explain := "get playing status in spotify"

  return selfMadeTypes.CommandHelp{
    Name: name,
    Kind: kind,
    Explain: explain,
  }
}
