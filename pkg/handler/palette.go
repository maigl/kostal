package handler

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var hexRegex = regexp.MustCompile(`#?([0-9a-fA-F]{6})`)

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
