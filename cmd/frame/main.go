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
			source := config.Config.Palette.AutoColorSrc
			if source == "" {
				source = "colormind"
			}
			go func() {
				ticker := time.NewTicker(duration)
				for range ticker.C {
					colors, err := handler.FetchPalette(source)
					if err != nil {
						log.Printf("auto-color fetch failed: %v", err)
						continue
					}
					handler.GlobalPaletteManager.SetPalette(colors)
					log.Printf("auto-color: fetched new palette from %s", source)
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
