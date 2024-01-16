package netpbm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type PGM struct{
    Data [][]uint8
    Width, Height int
    MagicNumber string
    MaxValue    int
}

func ReadPGM(filename string) (*PGM, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Lecture du nombre magique
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read magic number")
	}
	magicNumber := scanner.Text()

	// Lecture de la largeur et de la hauteur
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read width and height")
	}
	width, height, err := parseWidthHeight(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("failed to parse width and height: %v", err)
	}

	// Lecture de la valeur maximale
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read max value")
	}
	maxValue, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("failed to parse max value: %v", err)
	}

	// Lecture des données de l'image
	data := make([][]uint8, height)
	for i := 0; i < height; i++ {
		data[i] = make([]uint8, width)
		for j := 0; j < width; j++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("failed to read image data")
			}
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, fmt.Errorf("failed to parse image data: %v", err)
			}
			data[i][j] = uint8(value)
		}
	}

	return &PGM{
		MagicNumber: magicNumber,
		Width:       width,
		Height:      height,
		MaxValue:    maxValue,
		Data:        data,
	}, nil
}

func parseWidthHeight(input string) (int, int, error) {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid width and height format")
	}
	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return width, height, nil
}
