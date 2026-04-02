package handler

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type PaletteManager struct {
	mu         sync.RWMutex
	palette    [4]string
	configFile string
}

var GlobalPaletteManager *PaletteManager

func InitPaletteManager(configFile string) {
	GlobalPaletteManager = &PaletteManager{
		configFile: configFile,
		palette:    [4]string{"08415c", "6b818c", "f1bf98", "c41b5c"},
	}
	GlobalPaletteManager.loadFromFile()
}

func (pm *PaletteManager) GetPalette() [4]string {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.palette
}

func (pm *PaletteManager) SetPalette(colors [4]string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.palette = colors
	pm.saveToFile()
}

func (pm *PaletteManager) loadFromFile() {
	data, err := os.ReadFile(pm.configFile)
	if err != nil {
		log.Printf("loadFromFile: %v", err)
		return
	}
	var cfg struct {
		Palette string `json:"palette"`
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Printf("loadFromFile: unmarshal error: %v", err)
		return
	}
	colors, err := ParsePalette(cfg.Palette)
	if err != nil {
		log.Printf("loadFromFile: parse error: %v", err)
		return
	}
	pm.palette = colors
}

func (pm *PaletteManager) saveToFile() {
	data, err := os.ReadFile(pm.configFile)
	var cfg map[string]string
	if err == nil {
		json.Unmarshal(data, &cfg)
	}
	if cfg == nil {
		cfg = make(map[string]string)
	}
	cfg["palette"] = pm.palette[0] + "-" + pm.palette[1] + "-" + pm.palette[2] + "-" + pm.palette[3]

	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Printf("saveToFile: marshal error: %v", err)
		return
	}
	if err := os.WriteFile(pm.configFile, out, 0644); err != nil {
		log.Printf("saveToFile: write error: %v", err)
	}
}
