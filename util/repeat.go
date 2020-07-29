package util

const (
	track = "track"
	off   = "off"
)

func SwitchRepeatState(state string) string {
	switch state {
	case track:
		return off
	case "context":
		return track
	case off:
		return "context"
	}

	return off
}
