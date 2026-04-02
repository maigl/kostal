package handler

import (
	"os"
	"testing"
)

func TestPaletteManager(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "palette-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	tmpfile.Close()

	InitPaletteManager(tmpfile.Name())

	colors := [4]string{"111111", "222222", "333333", "444444"}
	GlobalPaletteManager.SetPalette(colors)

	got := GlobalPaletteManager.GetPalette()
	for i, c := range got {
		if c != colors[i] {
			t.Errorf("color[%d]: got %q, want %q", i, c, colors[i])
		}
	}

	InitPaletteManager(tmpfile.Name())
	got = GlobalPaletteManager.GetPalette()
	for i, c := range got {
		if c != colors[i] {
			t.Errorf("after reload color[%d]: got %q, want %q", i, c, colors[i])
		}
	}
}
