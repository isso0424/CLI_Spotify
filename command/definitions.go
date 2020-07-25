package command

type status struct {
}

func(_ status) getCommandName() string {
  return "status"
}

type next struct {
}

func(_ next) getCommandName() string {
  return "next"
}

type pause struct {
}

func(_ pause) getCommandName() string {
  return "pause"
}

type play struct {
}

func(_ play) getCommandName() string {
  return "play"
}

type prev struct {
}

func(_ prev) getCommandName() string {
  return "prev"
}

type repeat struct {
}

func(_ repeat) getCommandName() string {
  return "repeat"
}

type resume struct {
}

func(_ resume) getCommandName() string {
  return "resume"
}

type shuffle struct {
}

func(_ shuffle) getCommandName() string {
  return "shuffle"
}

type welcome struct {
}

func(_ welcome) getCommandName() string {
  return "welcome"
}
