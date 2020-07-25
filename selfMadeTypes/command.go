package selfMadeTypes

type RequestCommand interface {
  execute(*string) error
  getCommandName() string
  getHelp() CommandHelp
}

type FileloadCommand interface {
  execute() error
  getCommandName() string
  getHelp() CommandHelp
}

type RequestAndFileloadCommand interface {
  execute(*string) error
  getCommandName() string
  getHelp() CommandHelp
}

type CommandHelp struct {
  Name string
  Kind string
  Explain string
}
