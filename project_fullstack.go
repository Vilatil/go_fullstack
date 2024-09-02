package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

type WeatherData struct {
    Location struct {
        Name string `json:"name"`
		Country string `json:"country"`
		Localtime string `json:"localtime"`
    } `json:"location"`
	Condition struct{
		Weather string `json:"text"`
	}`json:"condition"`
}



func getWeather() (WeatherData, error){
	var weather WeatherData
    APIKEY := os.Getenv("APIKEY")
	if APIKEY == "" {
		log.Fatal("APIKEY не задан. Проверьте файл .env")
	}
    url := "https://api.weatherapi.com/v1/current.json?key=" + APIKEY + "&q=Odesa"

	response, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&weather)
	return weather, err
}

func main() {
	godotenv.Load()
	weather, err := getWeather()
	log.Print(weather)
	if err != nil {
		log.Println("Ошибка получения погоды:", err)
	}
	component := hello(weather)

	http.Handle("/",templ.Handler(component))

	log.Fatal(http.ListenAndServe(":8080",nil))
}