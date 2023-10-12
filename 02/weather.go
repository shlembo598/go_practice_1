package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"test02/responses"
)

func main() {
	http.HandleFunc("/", gerWeather)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
	slog.Error("Error: ", http.ListenAndServe(":8080", nil))
}

func gerWeather(w http.ResponseWriter, req *http.Request) {
	if !isValidUrl(req.URL) {
		fmt.Println("Invalid url")
		return
	}
	var pathData string = req.URL.Query()["coords"][0]
	latLng := strings.Split(pathData, ",")
	requestUrl := getRequestUrl(latLng[0], latLng[1])
	result := createRequest(requestUrl)

	fmt.Fprintf(w, "Current temperature = %v", result.CurrentWeather.Temperature)

}

func isValidUrl(url *url.URL) bool {
	var hasCoords = url.Query().Has("coords")
	if !hasCoords {
		return false
	}
	return url.Path == "/"
}

func getRequestUrl(lat string, lon string) string {
	return "https://api.open-meteo.com/v1/forecast?latitude=" + lat + "&longitude=" + lon + "&current_weather=true"
}

func createRequest(url string) responses.WeatherResponse {
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Error: ", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error", err)
	}

	var result responses.WeatherResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return result
}
