// Package commandtypes is types for command
package commandtypes

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

// CommandHelp is command's help.
type CommandHelp struct {
	Name    string
	Kind    CommandKind
	Explain string
}

// Command is all commands interface.
type Command interface {
	GetCommandName() string
	GetHelp() CommandHelp
}
