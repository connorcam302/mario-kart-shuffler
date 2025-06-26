package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func secureShuffle(slice []string) error {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		j := int(jBig.Int64())
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}

func main() {
	var tracks = []string{
		"Acorn Heights",
		"Airship Fortress",
		"Boo Cinema",
		"Bowser's Castle",
		"Cheep Cheep Falls",
		"Choco Mountain",
		"Crown City",
		"Dandelion Depths",
		"Desert Hills",
		"Dino Dino Jungle",
		"DK Pass",
		"DK Spaceport",
		"Dry Bones Burnout",
		"Faraway Oasis",
		"Great ? Block Ruins",
		"Koopa Troopa Beach",
		"Mario Bros. Circuit",
		"Mario Circuit",
		"Moo Moo Meadows",
		"Peach Beach",
		"Peach Stadium",
		"Salty Salty Speedway",
		"Shy Guy Bazaar",
		"Sky-High Sundae",
		"Starview Peak",
		"Toad's Factory",
		"Wario Stadium",
		"Wario's Galleon",
		"Whistlestop Summit",
	}

	err := secureShuffle(tracks)
	if err != nil {
		panic(err)
	}

	fmt.Println("Shuffled tracks:")
	for i, track := range tracks[:16] {
		fmt.Printf("%d: %s \n", i+1, track)
	}
}
