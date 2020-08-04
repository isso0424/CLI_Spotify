package util

import (
	"fmt"
	"isso0424/spotify_CLI/selfmadetypes"
)

// Output print to stdout.
func Output(message selfmadetypes.OutputMessage) {
	length := getMaxLengthRow(message)
	splitter := createSplitter(length)

	for _, lump := range message.Message {
		fmt.Println(splitter)
		for _, row := range lump {
			fmt.Println(row)
		}
		fmt.Println(splitter)
	}
}

func getMaxLengthRow(message selfmadetypes.OutputMessage) (maxLength int) {
	for _, lump := range message.Message {
		for _, row := range lump {
			if len(row) > maxLength {
				maxLength = len(row)
			}
		}
	}

	return
}

func createSplitter(length int) (splitter string) {
	const splitCharactor = "-"
	for i := 0; i < length; i++ {
		splitter += splitCharactor
	}

	return
}
