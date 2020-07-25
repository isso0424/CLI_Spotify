package selfMadeType

type RequestCommand interface {
  Execute(*string) error
  GetHelp() string
}

type FileloadCommand interface {
  Execute() error
  GetHelp() string
}

type RequestAndFileloadCommand interface {
  Execute(*string) error
  GetHelp() string
}

type CommandHelp struct {
  Name string
  Kind string
  Explain string
}
