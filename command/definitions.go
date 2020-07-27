package command

type status struct{}

// GetCommandName is getting command name function.
func (cmd status) GetCommandName() string {
	return "status"
}

type next struct{}

// GetCommandName is getting command name function.
func (cmd next) GetCommandName() string {
	return "next"
}

type pause struct{}

// GetCommandName is getting command name function.
func (cmd pause) GetCommandName() string {
	return "pause"
}

type play struct{}

// GetCommandName is getting command name function.
func (cmd play) GetCommandName() string {
	return "play"
}

type prev struct{}

// GetCommandName is getting command name function.
func (cmd prev) GetCommandName() string {
	return "prev"
}

// GetCommandName is getting command name function.
type repeat struct{}

// GetCommandName is getting command name function.
func (cmd repeat) GetCommandName() string {
	return "repeat"
}

type resume struct{}

// GetCommandName is getting command name function.
func (cmd resume) GetCommandName() string {
	return "resume"
}

type shuffle struct{}

// GetCommandName is getting command name function.
func (cmd shuffle) GetCommandName() string {
	return "shuffle"
}

type welcome struct{}

// GetCommandName is getting command name function.
func (cmd welcome) GetCommandName() string {
	return "welcome"
}

type save struct{}

// GetCommandName is getting command name function.
func (cmd save) GetCommandName() string {
	return "save"
}

type show struct{}

// GetCommandName is getting command name function.
func (cmd show) GetCommandName() string {
	return "show"
}

type random struct {
}

// GetCommandName is getting command name function.
func (cmd random) GetCommandName() string {
	return "random"
}

type load struct{}

// GetCommandName is getting command name function.
func (cmd load) GetCommandName() string {
	return "load"
}

type refresh struct{}

// GetCommandName is getting command name function.
func (cmd refresh) GetCommandName() string {
	return "refresh"
}

type volume struct{}

// GetCommandName is getting command name function.
func (cmd volume) GetCommandName() string {
	return "volume"
}

type search struct{}

// GetCommandName is getting command name function.
func (cmd search) GetCommandName() string {
	return "search"
}

type favoriteTrack struct{}

// GetCommandName is getting command name function.
func (cmd favoriteTrack) GetCommandName() string {
	return "favoriteTrack"
}
