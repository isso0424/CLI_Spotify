package selfMadeTypes

type RequestCommand interface {
	Execute(*string) error
	GetCommandName() string
	GetHelp() CommandHelp
}

type FileloadCommand interface {
	Execute() error
	GetCommandName() string
	GetHelp() CommandHelp
}

type RequestAndFileloadCommand interface {
	Execute(*string) error
	GetCommandName() string
	GetHelp() CommandHelp
}

type CommandHelp struct {
	Name    string
	Kind    string
	Explain string
}

type Command interface {
  Execute(interface{}) error
  GetCommandName() string
  GetHelp() CommandHelp
}
