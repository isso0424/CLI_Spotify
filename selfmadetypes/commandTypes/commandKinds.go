package commandTypes

// CommandKind is kind of commands
// ex) loadfile, player, playlist and other...
type CommandKind int

const (
	// LoadFile is command kind that access to local file.
	LoadFile = iota
	// Other is command kind that is not match any kind.
	Other
	// Player is command kind that controle player.
	Player
	// PlayerData is command kind that controle player data.
	PlayerData
	// Playlist is command kind that edit playlist.
	Playlist
	// Search is command kind that search on spotify.
	Search
)

func (kind CommandKind) String() string {
	switch kind {
	case LoadFile:
		return "loadfile"
	case Other:
		return "other"
	case Player:
		return "player"
	case Playlist:
		return "playlist"
	case Search:
		return "search"
	default:
		return "invalid"
	}
}
