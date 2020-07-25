package command

type status struct {
}

func(_ status) GetCommandName() string {
  return "status"
}

type next struct {
}

func(_ next) GetCommandName() string {
  return "next"
}

type pause struct {
}

func(_ pause) GetCommandName() string {
  return "pause"
}

type play struct {
}

func(_ play) GetCommandName() string {
  return "play"
}

type prev struct {
}

func(_ prev) GetCommandName() string {
  return "prev"
}

type repeat struct {
}

func(_ repeat) GetCommandName() string {
  return "repeat"
}

type resume struct {
}

func(_ resume) GetCommandName() string {
  return "resume"
}

type shuffle struct {
}

func(_ shuffle) GetCommandName() string {
  return "shuffle"
}

type welcome struct {
}

func(_ welcome) GetCommandName() string {
  return "welcome"
}
