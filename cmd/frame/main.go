package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maigl/kostal/pkg/config"
	"github.com/maigl/kostal/pkg/handler"
	"github.com/maigl/kostal/pkg/solcast"

	cron "github.com/robfig/cron/v3"
)


func main() {

	config.Init()

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
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
