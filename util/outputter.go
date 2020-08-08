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
			if stringLengthCounter(row) > maxLength {
				maxLength = stringLengthCounter(row)
			}
		}
	}

	return
}

func stringLengthCounter(str string) (length int) {
	runes := []rune(str)
	for _, charactor := range runes {
		if charactor > 255 {
			length++
		}
		length++
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
