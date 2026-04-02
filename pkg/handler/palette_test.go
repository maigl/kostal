package handler

import (
	"testing"
	"time"
)

func TestParsePalette(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [4]string
		hasError bool
	}{
		{
			name:     "dash-separated with hash",
			input:    "#c41b5c-#08415c-#6b818c-#f1bf98",
			expected: [4]string{"c41b5c", "08415c", "6b818c", "f1bf98"},
		},
		{
			name:     "dash-separated without hash",
			input:    "c41b5c-08415c-6b818c-f1bf98",
			expected: [4]string{"c41b5c", "08415c", "6b818c", "f1bf98"},
		},
		{
			name:     "five colors ignores fifth",
			input:    "c41b5c-08415c-6b818c-f1bf98-eee5e9",
			expected: [4]string{"c41b5c", "08415c", "6b818c", "f1bf98"},
			hasError: false,
		},
		{
			name:     "only three colors returns error",
			input:    "c41b5c-08415c-6b818c",
			expected: [4]string{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParsePalette(tt.input)
			if tt.hasError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			for i, c := range tt.expected {
				if result[i] != c {
					t.Errorf("color[%d]: got %q, want %q", i, result[i], c)
				}
			}
		})
	}
}

func TestFetchPalette(t *testing.T) {
	palette, err := FetchPalette()
	if err != nil {
		t.Fatalf("FetchPalette failed: %v", err)
	}
	if len(palette) != 4 {
		t.Errorf("expected 4 colors, got %d", len(palette))
	}
	for i, color := range palette {
		if len(color) != 6 {
			t.Errorf("color[%d] %q is not valid hex", i, color)
		}
	}
}

func TestParseAutoColorDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		hasError bool
	}{
		{input: "5m", expected: 5 * time.Minute, hasError: false},
		{input: "1h", expected: 1 * time.Hour, hasError: false},
		{input: "30s", expected: 30 * time.Second, hasError: false},
		{input: "invalid", expected: 0, hasError: true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseAutoColorDuration(tt.input)
			if tt.hasError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
