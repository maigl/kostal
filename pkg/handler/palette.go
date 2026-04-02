package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var hexRegex = regexp.MustCompile(`#?([0-9a-fA-F]{6})`)

const (
	AutoColorSourceColormind = "colormind"
	AutoColorSourceLocal     = "local"
	AutoColorSourceHybrid    = "hybrid"
)

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

func GenerateLocalPalette() ([4]string, error) {
	h := float64(time.Now().UnixNano()%1000000) / 1000000.0
	colors := [4]string{}
	for i := 0; i < 4; i++ {
		hue := math.Mod(h+float64(i)*0.25, 1.0)
		r, g, b := hslToRgb(hue, 0.6, 0.5)
		colors[i] = fmt.Sprintf("%02x%02x%02x", r, g, b)
	}
	return colors, nil
}

func hslToRgb(h, s, l float64) (r, g, b uint8) {
	if s == 0 {
		r, g, b = uint8(l*255), uint8(l*255), uint8(l*255)
		return
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	p := 2*l - q

	r = uint8(hueToRgb(p, q, h+1.0/3.0) * 255)
	g = uint8(hueToRgb(p, q, h) * 255)
	b = uint8(hueToRgb(p, q, h-1.0/3.0) * 255)
	return
}

func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

func FetchPalette(source string) ([4]string, error) {
	switch source {
	case AutoColorSourceLocal:
		return GenerateLocalPalette()
	case AutoColorSourceHybrid:
		colors, err := fetchColormindPalette()
		if err == nil {
			return colors, nil
		}
		return GenerateLocalPalette()
	default:
		return fetchColormindPalette()
	}
}

func fetchColormindPalette() ([4]string, error) {
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
