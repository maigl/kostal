package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"text/template"
	"time"

	"maigl/kostal/data"

	"github.com/goburrow/modbus"
	cron "github.com/robfig/cron/v3"
)

type PowerItem struct {
	Label    string
	Unit     string
	Value    string
	Forecast string
}

var modbusAddr = flag.String("modbus_addr", "192.168.0.31:1502", "The addr of the kostal modbus")
var webDirPath = flag.String("web_dir", "/home/pi/kostal/web", "the path to the web dir")
var apiKey = flag.String("api_key", "", "solcast api key")
var propertyID = flag.String("property_id", "95c5-8ddf-1d04-b586", "solcast property id")

func web(w http.ResponseWriter, r *http.Request) {
	power := getPower()

	tmpl, err := template.ParseFiles(*webDirPath + "/frame.html")

	// tmpl, err := template.New("web").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, power); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// comment wiht typo.
// TODO
func renderForecast(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(*webDirPath + "/forecast.html")

	// tmpl, err := template.New("web").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	forecastToday, err := getForcast(today)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tomorrow := today.AddDate(0, 0, 1)
	forecastTomorrow, err := getForcast(tomorrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Today    forecast
		Tomorrow forecast
	}{
		Today:    forecastToday,
		Tomorrow: forecastTomorrow,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var client modbus.Client

func main() {
	flag.Parse()

	c := cron.New()
	c.AddFunc("0 6 * * *", func() {
		log.Println("resetting forecast")
		forecasts = nil
	})
	c.Start()
	fmt.Println("starting")

	fs := http.FileServer(http.Dir(*webDirPath))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.HandleFunc("/", web)
	http.HandleFunc("/forecast", renderForecast)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func getModbusHandler() *modbus.TCPClientHandler {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler(*modbusAddr)
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 71
	// handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return handler
}

func getPower() map[string]PowerItem {
	if client == nil {
		// TODO maybe cache power object and reduce number of modbus calls
		modbusHandler := getModbusHandler()
		client = modbus.NewClient(modbusHandler)
		// defer modbusHandler.Close()
	}

	br := data.Registers["514"]
	err := br.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	battery := fmt.Sprintf("%d", br.Value)

	yr := data.Registers["260"]
	err = yr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	yield := yr.Float32()
	yr = data.Registers["270"]
	err = yr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	yield += yr.Float32()
	yield /= float32(1000)
	if yield < 0 {
		yield = 0
	}
	yieldString := fmt.Sprintf("%1.1f", yield)

	cr := data.Registers["106"]
	err = cr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	consumption := cr.Float32()
	cr = data.Registers["108"]
	err = cr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	consumption += cr.Float32()
	cr = data.Registers["116"]
	err = cr.Read(client)
	if err != nil {
		log.Fatal(err)
	}

	consumption += cr.Float32()
	consumption /= float32(1000)
	if consumption < 0 {
		consumption = 0
	}
	consumptionString := fmt.Sprintf("%1.1f", consumption)

	ir := data.Registers["575"]
	err = ir.Read(client)
	if err != nil {
		log.Fatal(err)
	}
	gridRaw := float32(ir.Uint16()) / 1000.
	// we seem to have an overflow here
	if gridRaw > 60 {
		gridRaw = 0
	}
	grid := gridRaw - consumption
	gridLabel := "to grid"
	if grid <= 0 {
		gridLabel = "from grid"
	}
	gridString := fmt.Sprintf("%1.1f", math.Abs(float64(grid)))

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	forecast, err := getForcast(today)
	if err != nil {
		log.Printf("error getting forecast: %v", err)
	}

	return map[string]PowerItem{
		"battery":     {Label: "battery", Unit: "%", Value: battery, Forecast: forecast.Morning.Icon()},
		"consumption": {Label: "consumption", Unit: "kW", Value: consumptionString, Forecast: forecast.Noon.Icon()},
		"grid":        {Label: gridLabel, Unit: "kW", Value: gridString, Forecast: forecast.Afternoon.Icon()},
		"yield":       {Label: "yield", Unit: "kW", Value: yieldString, Forecast: forecast.Evening.Icon()},
	}
}

type SolCastForecasts struct {
	Forecasts []SolcastForecast `json:"forecasts"`
}

type SolcastForecast struct {
	PvEstimate   float64   `json:"pv_estimate"`
	PvEstimate10 float64   `json:"pv_estimate10"`
	PvEstimate90 float64   `json:"pv_estimate90"`
	PeriodEnd    time.Time `json:"period_end"`
	Period       string    `json:"period"`
}

type forecast struct {
	Morning   Estimate
	Noon      Estimate
	Afternoon Estimate
	Evening   Estimate
}

var forecasts map[time.Time]forecast

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

// for now we only get the forcast from today
// let's read forecast.json in '.'
// find all the forcasts for today
// and aggregate
func getForcast(day time.Time) (forecast, error) {

	if f, ok := forecasts[day]; ok {
		return f, nil
	}

	log.Println("getting forecast for", day)

	var data *SolCastForecasts
	var err error
	if *apiKey == "" {
		log.Println("no api key for solcast")
		data, err = getForecastDataFromFile()
	} else {
		data, err = getForecastDataFromAPI()
	}
	if err != nil {
		return forecast{}, err
	}

	var morning, noon, afternoon, evening Estimate

	var foundValuesForDay bool

	for _, f := range data.Forecasts {
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
		return forecast{}, fmt.Errorf("no values found for day %s", day)
	}

	if forecasts == nil {
		forecasts = make(map[time.Time]forecast)
	}

	forecasts[day] = forecast{
		Morning:   morning,
		Noon:      noon,
		Afternoon: afternoon,
		Evening:   evening,
	}

	return forecasts[day], nil
}

var RateLimitError error = fmt.Errorf("rate limit exceeded")
var PauseSolcastError error = fmt.Errorf("pause solcast")
var PauseSolcastUntil time.Time

func getForecastDataFromAPI() (*SolCastForecasts, error) {

	now := time.Now()
	if PauseSolcastUntil.After(now) {
		return nil, fmt.Errorf("%w until: %v", PauseSolcastError, PauseSolcastUntil)
	}

	url := fmt.Sprintf("https://api.solcast.com.au/rooftop_sites/%s/forecasts?format=json&api_key=%s",
		*propertyID, *apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("erroru getting forecast: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests ||
		resp.StatusCode == 420 { // strangely the api seems returns 420 for rate limit
		PauseSolcastUntil = time.Now().Add(time.Hour * 12)
		return nil, RateLimitError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	var data SolCastForecasts
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func getForecastDataFromFile() (*SolCastForecasts, error) {

	b, err := ioutil.ReadFile("forecasts.json")
	if err != nil {
		return nil, err
	}

	//unmarshal the json
	var forcasts SolCastForecasts
	err = json.Unmarshal(b, &forcasts)
	if err != nil {
		return nil, err
	}
	return &forcasts, nil

}

func PrintAllRegisters() { // nolint
	for _, r := range data.Registers {
		err := r.Read(client)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("r: %+v", r)
	}
}
