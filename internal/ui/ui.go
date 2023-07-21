package ui

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Serenity0204/LVCS/resources"
)

type ASCIIArtGenerator struct {
	ASCIIArtMap map[int]string
}

func NewASCIIArtGenerator() *ASCIIArtGenerator {
	asciiMap := make(map[int]string)

	asciiMap[1] = resources.ASCII1
	asciiMap[2] = resources.ASCII2
	asciiMap[3] = resources.ASCII3
	asciiMap[4] = resources.ASCII4
	asciiMap[5] = resources.ASCII5
	return &ASCIIArtGenerator{
		ASCIIArtMap: asciiMap,
	}
}

func (ascii *ASCIIArtGenerator) GetRandASCIIArt() (string, error) {
	rand.Seed(time.Now().UnixNano())
	// Generate a random number from 1 to 5
	randomNumber := rand.Intn(5) + 1
	asciiArt, ok := ascii.ASCIIArtMap[randomNumber]
	if !ok {
		return "", errors.New("failed to read ASCII file")
	}
	return asciiArt, nil
}
