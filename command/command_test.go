package command

import (
	"github.com/stretchr/testify/assert"
	"isso0424/spotify_CLI/util"
	"testing"
)

func TestUpdateRepeatStatus(t *testing.T) {
	resetFunc := setRepeatFunc(
		func(state *string) (string, error) {
			newState := util.SwitchRepeatState(*state)

			return newState, nil
		},
	)

	defer resetFunc()

	track := "track"
	off := "off"
	context := "context"
	state, _ := updateRepeatStatus(&track)
	assert.Equal(t, state, "context")

	state, _ = updateRepeatStatus(&off)
	assert.Equal(t, state, "track")

	state, _ = updateRepeatStatus(&context)
	assert.Equal(t, state, "off")
}
