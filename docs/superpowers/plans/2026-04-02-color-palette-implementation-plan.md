# Color Palette Configuration Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add runtime color palette configuration with startup flags, env vars, config file persistence, POST /colors endpoint, and auto-color fetching from coolors.co.

**Architecture:** 
- Palette state managed in `pkg/config` package
- Palette parsing/fetching in new `pkg/handler/palette.go`
- Handler modified to inject CSS variables into template
- Auto-color runs as a goroutine ticker started in main

**Tech Stack:** Go stdlib (net/http, text/template, time), viper for config, robfig/cron already in use

---

## Task 1: Add Palette Config to Config Package

**Files:**
- Modify: `pkg/config/config.go`

- [ ] **Step 1: Add PaletteParams struct**

Add to `pkg/config/config.go`:

```go
type PaletteParams struct {
    Palette    string
    AutoColor  string
    ConfigFile string
}
```

- [ ] **Step 2: Update Init() to add palette flags**

Add after existing pflag definitions:

```go
pflag.String("palette", "", "Color palette (e.g. c41b5c-08415c-6b818c-f1bf98)")
pflag.String("auto-color", "", "Auto-fetch palette interval (e.g. 5m, 1h)")
pflag.String("config", "config.json", "Config file path")
```

- [ ] **Step 3: Update Params struct**

Add fields to `Params`:
```go
type Params struct {
    // ... existing fields ...
    Palette   *PaletteParams
}
```

- [ ] **Step 4: Initialize PaletteParams in Init()**

Add after Config assignment:
```go
Config.Palette = &PaletteParams{
    Palette:    viper.GetString("palette"),
    AutoColor:  viper.GetString("auto-color"),
    ConfigFile: viper.GetString("config"),
}
```

- [ ] **Step 5: Commit**

```bash
git add pkg/config/config.go
git commit -m "feat: add palette config params"
```

---

## Task 2: Create Palette Parser

**Files:**
- Create: `pkg/handler/palette.go`
- Test: `pkg/handler/palette_test.go`

- [ ] **Step 1: Write failing test for ParsePalette**

Create `pkg/handler/palette_test.go`:

```go
package handler

import "testing"

func TestParsePalette(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected [4]string
    }{
        {
            name:  "dash-separated with hash",
            input: "#c41b5c-#08415c-#6b818c-#f1bf98",
            expected: [4]string{"c41b5c", "08415c", "6b818c", "f1bf98"},
        },
        {
            name:  "dash-separated without hash",
            input: "c41b5c-08415c-6b818c-f1bf98",
            expected: [4]string{"c41b5c", "08415c", "6b818c", "f1bf98"},
        },
        {
            name:  "five colors ignores fifth",
            input: "c41b5c-08415c-6b818c-f1bf98-eee5e9",
            expected: [4]string{"c41b5c", "08415c", "6b818c", "f1bf98"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := ParsePalette(tt.input)
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
```

- [ ] **Step 2: Run test to verify it fails**

```bash
go test ./pkg/handler/... -run TestParsePalette -v
```
Expected: FAIL undefined: ParsePalette

- [ ] **Step 3: Write ParsePalette function**

Create `pkg/handler/palette.go`:

```go
package handler

import (
    "regexp"
    "strings"
)

var hexRegex = regexp.MustCompile(`#?([0-9a-fA-F]{6})`)

func ParsePalette(input string) ([4]string, error) {
    matches := hexRegex.FindAllStringSubmatch(input, -1)
    if len(matches) < 4 {
        return [4]string{}, nil
    }
    colors := [4]string{}
    for i := 0; i < 4; i++ {
        colors[i] = strings.ToLower(matches[i][1])
    }
    return colors, nil
}
```

- [ ] **Step 4: Run test to verify it passes**

```bash
go test ./pkg/handler/... -run TestParsePalette -v
```
Expected: PASS

- [ ] **Step 5: Add ParseAutoColorDuration test and function**

Add to `palette_test.go`:

```go
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
```

Add to `palette.go`:

```go
import "time"

func ParseAutoColorDuration(input string) (time.Duration, error) {
    if input == "" {
        return 0, nil
    }
    return time.ParseDuration(input)
}
```

- [ ] **Step 6: Run all palette tests**

```bash
go test ./pkg/handler/... -run "Palette" -v
```
Expected: PASS

- [ ] **Step 7: Commit**

```bash
git add pkg/handler/palette.go pkg/handler/palette_test.go
git commit -m "feat: add palette parsing utilities"
```

---

## Task 3: Create Coolors Fetcher

**Files:**
- Modify: `pkg/handler/palette.go`
- Create: `pkg/handler/palette_test.go` (add fetch test)

- [ ] **Step 1: Add test for FetchPalette**

Add to `palette_test.go`:

```go
func TestFetchPalette(t *testing.T) {
    palette, err := FetchPalette()
    if err != nil {
        t.Fatalf("FetchPalette failed: %v", err)
    }
    if len(palette) != 4 {
        t.Errorf("expected 4 colors, got %d", len(palette))
    }
}
```

- [ ] **Step 2: Run test to verify it fails**

```bash
go test ./pkg/handler/... -run TestFetchPalette -v
```
Expected: FAIL undefined: FetchPalette

- [ ] **Step 3: Write FetchPalette function**

Add to `palette.go`:

```go
import (
    "io"
    "net/http"
    "regexp"
)

var coolorsPaletteRegex = regexp.MustCompile(`"colors":\["(#[0-9a-fA-F]{6})`)

func FetchPalette() ([4]string, error) {
    resp, err := http.Get("https://coolors.co/generate")
    if err != nil {
        return [4]string{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return [4]string{}, err
    }

    matches := coolorsPaletteRegex.FindSubmatch(body)
    if len(matches) < 2 {
        return [4]string{}, nil
    }

    return ParsePalette(string(matches[1]))
}
```

- [ ] **Step 4: Run test to verify it passes**

```bash
go test ./pkg/handler/... -run TestFetchPalette -v
```
Expected: PASS (or network error - that's ok)

- [ ] **Step 5: Commit**

```bash
git add pkg/handler/palette.go
git commit -m "feat: add coolors.co palette fetcher"
```

---

## Task 4: Create Palette Manager with Config Persistence

**Files:**
- Create: `pkg/handler/palettemanager.go`
- Create: `pkg/handler/palettemanager_test.go`

- [ ] **Step 1: Write PaletteManager and tests**

Create `pkg/handler/palettemanager.go`:

```go
package handler

import (
    "encoding/json"
    "log"
    "os"
    "sync"
)

type PaletteManager struct {
    mu       sync.RWMutex
    palette  [4]string
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
        return
    }
    var cfg struct {
        Palette string `json:"palette"`
    }
    if err := json.Unmarshal(data, &cfg); err != nil {
        return
    }
    colors, err := ParsePalette(cfg.Palette)
    if err != nil {
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
    
    out, _ := json.MarshalIndent(cfg, "", "  ")
    os.WriteFile(pm.configFile, out, 0644)
}
```

Create `pkg/handler/palettemanager_test.go`:

```go
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
```

- [ ] **Step 2: Run test to verify it passes**

```bash
go test ./pkg/handler/... -run TestPaletteManager -v
```
Expected: PASS

- [ ] **Step 3: Commit**

```bash
git add pkg/handler/palettemanager.go pkg/handler/palettemanager_test.go
git commit -m "feat: add palette manager with config persistence"
```

---

## Task 5: Add POST /colors Handler

**Files:**
- Modify: `pkg/handler/handler.go`

- [ ] **Step 1: Read handler.go first 10 lines to add import**

```go
import (
    "io"
    // ... existing imports ...
)
```

- [ ] **Step 2: Add SetColors handler**

Add after `RenderForecast`:

```go
func SetColors(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    colors, err := ParsePalette(string(body))
    if err != nil || colors == [4]string{} {
        http.Error(w, "Invalid palette format", http.StatusBadRequest)
        return
    }
    GlobalPaletteManager.SetPalette(colors)
    w.WriteHeader(http.StatusOK)
}
```

- [ ] **Step 3: Commit**

```bash
git add pkg/handler/handler.go
git commit -m "feat: add POST /colors endpoint"
```

---

## Task 6: Modify Web Handler to Inject Palette

**Files:**
- Modify: `pkg/handler/handler.go`

- [ ] **Step 1: Create template data struct**

Add at top of handler.go:

```go
type PageData struct {
    Battery     kostalModbus.PowerItem
    Consumption kostalModbus.PowerItem
    Grid        kostalModbus.PowerItem
    Yield       kostalModbus.PowerItem
    Palette     [4]string
}
```

- [ ] **Step 2: Modify Web handler**

Replace Web function with:

```go
func Web(w http.ResponseWriter, r *http.Request) {
    power, err := kostalModbus.GetPower()
    if err != nil {
        power = map[string]kostalModbus.PowerItem{
            "battery":     {Label: "battery", Unit: "%"},
            "consumption": {Label: "consumption", Unit: "kW"},
            "grid":        {Label: "to grid", Unit: "kW"},
            "yield":       {Label: "yield", Unit: "kW"},
        }
    }

    tmpl, err := template.ParseFiles(config.Config.WebDirPath + "/frame.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := PageData{
        Battery:     power["battery"],
        Consumption: power["consumption"],
        Grid:        power["grid"],
        Yield:       power["yield"],
        Palette:     GlobalPaletteManager.GetPalette(),
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

- [ ] **Step 3: Run build to check for errors**

```bash
go build ./cmd/frame/...
```
Expected: success

- [ ] **Step 4: Commit**

```bash
git add pkg/handler/handler.go
git commit -m "feat: inject palette into web template"
```

---

## Task 7: Update HTML Template with Dynamic CSS

**Files:**
- Modify: `web/frame.html`

- [ ] **Step 1: Replace :root CSS block**

Find this block (lines 14-19):
```html
:root {
    --color1: #08415C;
    --color2: #6B818C;
    --color3: #F1BF98;
    --color4: #C41B5C;
}
```

Replace with:
```html
:root {
    --color1: #{{.Palette._0}};
    --color2: #{{.Palette._1}};
    --color3: #{{.Palette._2}};
    --color4: #{{.Palette._3}};
}
```

- [ ] **Step 2: Commit**

```bash
git add web/frame.html
git commit -m "feat: make CSS variables dynamic via template"
```

---

## Task 8: Wire Up Main and Start Auto-Color

**Files:**
- Modify: `cmd/frame/main.go`

- [ ] **Step 1: Modify main.go**

Replace main.go content:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/maigl/kostal/pkg/config"
    "github.com/maigl/kostal/pkg/handler"
    "github.com/maigl/kostal/pkg/solcast"

    cron "github.com/robfig/cron/v3"
)

func main() {
    config.Init()

    handler.InitPaletteManager(config.Config.Palette.ConfigFile)

    if config.Config.Palette.Palette != "" {
        colors, err := handler.ParsePalette(config.Config.Palette.Palette)
        if err == nil {
            handler.GlobalPaletteManager.SetPalette(colors)
        }
    }

    if config.Config.Palette.AutoColor != "" {
        duration, err := handler.ParseAutoColorDuration(config.Config.Palette.AutoColor)
        if err == nil && duration > 0 {
            go func() {
                ticker := time.NewTicker(duration)
                for range ticker.C {
                    colors, err := handler.FetchPalette()
                    if err != nil {
                        log.Printf("auto-color fetch failed: %v", err)
                        continue
                    }
                    handler.GlobalPaletteManager.SetPalette(colors)
                    log.Println("auto-color: fetched new palette")
                }
            }()
        }
    }

    c := cron.New()
    c.AddFunc("0 6 * * *", func() {
        log.Println("resetting forecast")
        solcast.ResetForecasts()
    })
    c.Start()
    fmt.Println("starting")

    fs := http.FileServer(http.Dir(config.Config.WebDirPath))
    http.Handle("/web/", http.StripPrefix("/web/", fs))

    http.HandleFunc("/", handler.Web)
    http.HandleFunc("/forecast", handler.RenderForecast)
    http.HandleFunc("/colors", handler.SetColors)
    if err := http.ListenAndServe(":8081", nil); err != nil {
        panic(err)
    }
}
```

- [ ] **Step 2: Build and fix any errors**

```bash
go build ./cmd/frame/...
```
Expected: success (fix any type errors)

- [ ] **Step 3: Commit**

```bash
git add cmd/frame/main.go
git commit -m "feat: wire up palette manager and auto-color"
```

---

## Verification

- [ ] Run all tests: `go test ./...`
- [ ] Build binary: `go build ./cmd/frame/...`
- [ ] Manual test: Set palette via curl and verify HTML changes
