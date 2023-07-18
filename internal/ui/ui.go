package ui

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type ASCIIArtGenerator struct {
	ASCIIArtMap map[int]string
}

func NewASCIIArtGenerator() *ASCIIArtGenerator {
	const n int = 6
	const prefix string = "ASCIIArts/ascii"
	const postfix string = ".txt"
	asciiMap := make(map[int]string)

	// assigning file path
	for i := 1; i <= n; i++ {
		path := prefix + strconv.Itoa(i) + postfix
		asciiMap[i] = path
	}
	return &ASCIIArtGenerator{
		ASCIIArtMap: asciiMap,
	}
}

func (ascii *ASCIIArtGenerator) GetRandASCIIArt() (string, error) {
	rand.Seed(time.Now().UnixNano())
	// Generate a random number from 1 to 6
	randomNumber := rand.Intn(6) + 1
	path, ok := ascii.ASCIIArtMap[randomNumber]
	if !ok {
		return "", errors.New("failed to read ASCII file")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
}
