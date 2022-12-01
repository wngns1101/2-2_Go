package main

import "gadget"

// func playList(device gadget.TapePlayer, songs []string) {
// 	for _, song := range songs {
// 		device.Play(song)
// 	}
// 	device.Stop()
// }

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

func main() {
	player := gadget.TapePlayer{}
	recordPlayer := gadget.TapeRecorder{}
	mixtape := []string{"거짓말", "너랑 나", "거짓말"}
	mcrtap := []string{"Welcome to my PlayGround", "wow", "BetterBetter"}

	playList(player, mixtape)
	playList(recordPlayer, mcrtap)
}
