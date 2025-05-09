package handler

import (
	"net/http"
	"text/template"
	"time"

	"github.com/maigl/kostal/pkg/config"
	"github.com/maigl/kostal/pkg/kostalModbus"
	"github.com/maigl/kostal/pkg/solcast"
)

func Web(w http.ResponseWriter, r *http.Request) {
	power := kostalModbus.GetPower()

	tmpl, err := template.ParseFiles(config.Config.WebDirPath + "/frame.html")

	// tmpl, err := template.New("web").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, power); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderForecast(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(config.Config.WebDirPath + "/forecast.html")

	// tmpl, err := template.New("web").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, config.Config.Location)
	forecastToday, err := solcast.GetForcast(today)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tomorrow := today.AddDate(0, 0, 1)
	forecastTomorrow, err := solcast.GetForcast(tomorrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Today    solcast.Forecast
		Tomorrow solcast.Forecast
	}{
		Today:    forecastToday,
		Tomorrow: forecastTomorrow,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
