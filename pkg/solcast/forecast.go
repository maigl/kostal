package solcast

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/maigl/kostal/pkg/config"
)

// forecasts is the struct to parse the json from solcast
type forecasts struct {
	PeriodEstimates []struct {
		PvEstimate   float64   `json:"pv_estimate"`
		PvEstimate10 float64   `json:"pv_estimate10"`
		PvEstimate90 float64   `json:"pv_estimate90"`
		PeriodEnd    time.Time `json:"period_end"`
		Period       string    `json:"period"`
	} `json:"forecasts"`
}

// Forecast holds our projected Estimates for the day
type Forecast struct {
	Morning   Estimate
	Noon      Estimate
	Afternoon Estimate
	Evening   Estimate
}

type Estimate struct {
	Name      string
	Estimates []float64
}

func (e Estimate) Avg() float64 {
	if len(e.Estimates) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range e.Estimates {
		sum += v
	}
	return sum / float64(len(e.Estimates))
}

func (e *Estimate) Add(v float64) {
	if e.Estimates == nil {
		e.Estimates = []float64{}
	}
	e.Estimates = append(e.Estimates, v)
}

// no pointer to be able to use it in template
func (e Estimate) Sum() float64 {
	sum := 0.0
	for _, v := range e.Estimates {
		sum += v
	}
	return sum
}

// as we have estimates for 30 min we need
// to divide by 2 to get kWh
func (e Estimate) Total() float64 {
	return e.Sum() / 2
}

func (e Estimate) String() string {
	return fmt.Sprintf("Ø %1.1f\nΣ %1.1f", e.Avg(), e.Total())
}

func (e Estimate) Icon() string {

	avg := e.Avg()

	// don't show anything if we don't have estimates
	// this might happen if the solcast api is down
	// or we ran in rate limits
	if e.Estimates == nil {
		return ""
	}

	if avg < 0.7 {
		return "cloud"
	}
	if avg < 1.5 {
		return "sun-cloud"
	}
	if avg < 3.0 {
		return "sun"
	}
	return "full-sun"
}

// cache
var forecastsCache map[time.Time]Forecast

func ResetForecasts() {
	fmt.Printf("resetting forecasts cache, we needed %d calls\n", apiCallCount)
	forecastsCache = nil
	// don't forget to reset the api call count
	apiCallCount = 0
}

// for now we only get the forecast from today
// let's read forecast.json in '.'
// find all the forecasts for today
// and aggregate
func GetForecast(day time.Time) (Forecast, error) {

	if f, ok := forecastsCache[day]; ok {
		return f, nil
	}

	log.Println("getting forecast for", day)

	var data *forecasts
	var err error
	if config.Config.SolcastApiKey == "" || config.Config.SolcastPropertyID == "" {
		log.Println("no api key or property id for solcast")
		data, err = getForecastDataFromFile()
	} else {
		data, err = getForecastDataFromAPI()
	}
	if err != nil {
		return Forecast{}, err
	}

	var morning, noon, afternoon, evening Estimate

	var foundValuesForDay bool

	for _, f := range data.PeriodEstimates {
		//check if the forecast is for today
		if f.PeriodEnd.Day() != day.Day() {
			continue
		}

		foundValuesForDay = true

		if f.PvEstimate < 0.01 {
			continue
		}

		// if the forecast is for today morning
		if f.PeriodEnd.Before(day.Add(time.Hour * 10)) {
			morning.Add(f.PvEstimate)
		} else if f.PeriodEnd.Before(day.Add(time.Hour * 13)) {
			noon.Add(f.PvEstimate)
		} else if f.PeriodEnd.Before(day.Add(time.Hour * 16)) {
			afternoon.Add(f.PvEstimate)
		} else if f.PeriodEnd.Before(day.Add(time.Hour * 22)) {
			evening.Add(f.PvEstimate)
		}
	}

	if !foundValuesForDay {
		return Forecast{}, fmt.Errorf("no values found for day %s", day)
	}

	if forecastsCache == nil {
		forecastsCache = make(map[time.Time]Forecast)
	}

	forecastsCache[day] = Forecast{
		Morning:   morning,
		Noon:      noon,
		Afternoon: afternoon,
		Evening:   evening,
	}

	return forecastsCache[day], nil
}

var RateLimitError error = fmt.Errorf("rate limit exceeded")
var PauseSolcastError error = fmt.Errorf("pause solcast")
var PauseSolcastUntil time.Time

func getForecastDataFromAPI() (*forecasts, error) {

	now := time.Now()
	if PauseSolcastUntil.After(now) {
		return nil, fmt.Errorf("%w until: %v", PauseSolcastError, PauseSolcastUntil)
	}

	url := fmt.Sprintf("https://api.solcast.com.au/rooftop_sites/%s/forecasts?format=json&api_key=%s",
		config.Config.SolcastPropertyID, config.Config.SolcastApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error getting forecast: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests ||
		resp.StatusCode == 420 { // strangely the api seems returns 420 for rate limit
		PauseSolcastUntil = time.Now().Add(time.Hour * 12)
		return nil, RateLimitError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	var data forecasts
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func getForecastDataFromFile() (*forecasts, error) {

	b, err := os.ReadFile("forecasts.json")
	if err != nil {
		return nil, err
	}

	//unmarshal the json
	var forecasts forecasts
	err = json.Unmarshal(b, &forecasts)
	if err != nil {
		return nil, err
	}
	return &forecasts, nil

}
