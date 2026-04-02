package handler

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var hexRegex = regexp.MustCompile(`#?([0-9a-fA-F]{6})`)

var coolorsPaletteRegex = regexp.MustCompile(`"colors":\["(#[0-9a-fA-F]{6})`)

func ParsePalette(input string) ([4]string, error) {
	matches := hexRegex.FindAllStringSubmatch(input, -1)
	if len(matches) < 4 {
		return [4]string{}, fmt.Errorf("expected 4 colors, found %d", len(matches))
	}
	colors := [4]string{}
	for i := 0; i < 4; i++ {
		colors[i] = strings.ToLower(matches[i][1])
	}
	return colors, nil
}

func ParseAutoColorDuration(input string) (time.Duration, error) {
	if input == "" {
		return 0, nil
	}
	return time.ParseDuration(input)
}

func FetchPalette() ([4]string, error) {
	resp, err := http.Get("https://coolors.co/generate")
	if err != nil {
		return generateRandomPalette()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return generateRandomPalette()
	}

	matches := coolorsPaletteRegex.FindSubmatch(body)
	if len(matches) < 2 {
		return generateRandomPalette()
	}

	return ParsePalette(string(matches[1]))
}

func generateRandomPalette() ([4]string, error) {
	colors := [4]string{}
	for i := range colors {
		r := rand.Intn(256)
		g := rand.Intn(256)
		b := rand.Intn(256)
		colors[i] = fmt.Sprintf("%02x%02x%02x", r, g, b)
	}
	return colors, nil
}
