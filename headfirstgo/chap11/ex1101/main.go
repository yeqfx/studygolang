package main

import (
	"fmt"

	"./gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
		// } else {
		// 	fmt.Println("Player was not a TapeRecorder")
	}
}

// func playListWithRecorder(device gadget.TapeRecorder, songs []string) {
// 	for _, song := range songs {
// 		device.Play(song)
// 	}
// 	device.Stop()
// }

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
	fmt.Println("")
	player2 := gadget.TapeRecorder{}
	playList(player2, mixtape)
	fmt.Println("----------------------")
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
