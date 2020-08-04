package command

import (
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
	"isso0424/spotify_CLI/selfmadetypes/commandtypes"
	request2 "isso0424/spotify_CLI/selfmadetypes/request"
	response2 "isso0424/spotify_CLI/selfmadetypes/response"
	"isso0424/spotify_CLI/util"
)

type welcome struct{}

// GetCommandName is getting command name function.
func (cmd welcome) GetCommandName() string {
	return "welcome"
}

// GetHelp is getting help function.
func (cmd welcome) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Other,
		Explain: "switch shuffle state",
	}
}

// Execute is excution command function.
func (cmd welcome) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, request2.GET, "/me", nil)
	if err != nil {
		return
	}

	var userInfo response2.User
	err = json.Unmarshal(response.GetBody(), &userInfo)
	if err != nil {
		return
	}

	util.Output(
		selfmadetypes.OutputMessage{
			Message: [][]string{
				{
					fmt.Sprintf("ようこそ! %sさん!", userInfo.DisplayName),
				},
			},
		},
	)

	return
}

type refresh struct{}

// GetCommandName is getting command name function.
func (cmd refresh) GetCommandName() string {
	return "refresh"
}

// GetHelp is getting help function.
func (cmd refresh) GetHelp() commandtypes.CommandHelp {
	return commandtypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    commandtypes.Other,
		Explain: "refresh access token",
	}
}

// Execute is excution command function.
func (cmd refresh) Execute(token *string) error {
	tokenPtr, err := auth.GetToken()
	if err != nil {
		return err
	}

	*token = *tokenPtr

	return nil
}

func help(commands []commandtypes.Command) {
	for _, command := range commands {
		commandHelp := command.GetHelp()
		util.Output(
			selfmadetypes.OutputMessage{
				Message: [][]string{
					{
						commandHelp.Name,
					},
					{
						"Kind: " + commandHelp.Kind.String(),
						"Description: " + commandHelp.Explain,
					},
				},
			},
		)
	}
}
