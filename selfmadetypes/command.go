package selfmadetypes

// RequestCommand is interface that request commands.
type RequestCommand interface {
	Execute(*string) error
	GetCommandName() string
	GetHelp() CommandHelp
}

// FileloadCommand is interface that fileLoad commands.
type FileloadCommand interface {
	Execute() error
	GetCommandName() string
	GetHelp() CommandHelp
}

// RequestAndFileloadCommand is interface that request and fileLoad commands.
type RequestAndFileloadCommand interface {
	Execute(*string) error
	GetCommandName() string
	GetHelp() CommandHelp
}

// CommandHelp is command's help.
type CommandHelp struct {
	Name    string
	Kind    string
	Explain string
}

// Command is all commands interface.
type Command interface {
	GetCommandName() string
	GetHelp() CommandHelp
}
