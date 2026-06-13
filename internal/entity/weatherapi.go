package entity

type Location struct {
	City            string  `json:"name"`
	Region          string  `json:"region"`
	Country         string  `json:"country"`
	Latitude        float64 `json:"lat"`
	Longitude       float64 `json:"lon"`
	TimezoneId      string  `json:"tz_id"`
	Localtime_epoch int64   `json:"localtime_epoch"`
	Localtime       string  `json:"localtime"`
}

type Current struct {
	Temp_c float64 `json:"temp_c"`
	Temp_f float64 `json:"temp_f"`
}

type WeatherAPI struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func NewWeatherAPI() *WeatherAPI {
	return &WeatherAPI{}
}
