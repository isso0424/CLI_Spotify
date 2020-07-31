package command

import (
	"encoding/json"
	"fmt"
	"isso0424/spotify_CLI/auth"
	"isso0424/spotify_CLI/command/request"
	"isso0424/spotify_CLI/selfmadetypes"
)

type welcome struct{}

// GetCommandName is getting command name function.
func (cmd welcome) GetCommandName() string {
	return "welcome"
}

// GetHelp is getting help function.
func (cmd welcome) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
		Explain: "switch shuffle state",
	}
}

// Execute is excution command function.
func (cmd welcome) Execute(token *string) (err error) {
	response, err := request.CreateRequest(token, selfmadetypes.GET, "/me", nil)
	if err != nil {
		return
	}

	var userInfo selfmadetypes.User
	err = json.Unmarshal(response.GetBody(), &userInfo)
	if err != nil {
		return
	}

	fmt.Printf("ようこそ! %sさん!\n", userInfo.DisplayName)

	return
}

type refresh struct{}

// GetCommandName is getting command name function.
func (cmd refresh) GetCommandName() string {
	return "refresh"
}

// GetHelp is getting help function.
func (cmd refresh) GetHelp() selfmadetypes.CommandHelp {
	return selfmadetypes.CommandHelp{
		Name:    cmd.GetCommandName(),
		Kind:    "request",
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
