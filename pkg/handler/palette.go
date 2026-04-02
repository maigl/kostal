package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func FetchPalette() ([4]string, error) {
	body := bytes.NewBufferString(`{"model":"default"}`)
	resp, err := http.Post("http://colormind.io/api/", "application/json", body)
	if err != nil {
		return [4]string{}, fmt.Errorf("colormind request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return [4]string{}, fmt.Errorf("colormind returned status %d", resp.StatusCode)
	}

	var result struct {
		Result [][]int `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return [4]string{}, fmt.Errorf("failed to parse colormind response: %w", err)
	}

	if len(result.Result) < 4 {
		return [4]string{}, fmt.Errorf("colormind returned fewer than 4 colors")
	}

	colors := [4]string{}
	for i := 0; i < 4; i++ {
		rgb := result.Result[i]
		colors[i] = fmt.Sprintf("%02x%02x%02x", rgb[0], rgb[1], rgb[2])
	}
	return colors, nil
}
