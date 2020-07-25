package command

type status struct {
}

func (_ status) GetCommandName() string {
	return "status"
}

type next struct {
}

func (_ next) GetCommandName() string {
	return "next"
}

type pause struct {
}

func (_ pause) GetCommandName() string {
	return "pause"
}

type play struct {
}

func (_ play) GetCommandName() string {
	return "play"
}

type prev struct {
}

func (_ prev) GetCommandName() string {
	return "prev"
}

type repeat struct {
}

func (_ repeat) GetCommandName() string {
	return "repeat"
}

type resume struct {
}

func (_ resume) GetCommandName() string {
	return "resume"
}

type shuffle struct {
}

func (_ shuffle) GetCommandName() string {
	return "shuffle"
}

type welcome struct {
}

func (_ welcome) GetCommandName() string {
	return "welcome"
}

type save struct {
}

func (_ save) GetCommandName() string {
	return "save"
}

type show struct {
}

func (_ show) GetCommandName() string {
	return "show"
}

type random struct {
}

func (_ random) GetCommandName() string {
	return "random"
}

type load struct {
}

func (_ load) GetCommandName() string {
	return "load"
}

type refresh struct {
}

func (_ refresh) GetCommandName() string {
	return "refresh"
}

type volume struct {
}

func (_ volume) GetCommandName() string {
  return "volume"
}
