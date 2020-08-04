package selfmadetypes

type row string

type lump []row

// OutputMessage is output message struct.
type OutputMessage struct {
	Message []lump
}
