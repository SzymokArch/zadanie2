package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const AUTHOR = "Marek Zając"

// Dane pogodowe
var weatherData = map[string]map[string]interface{}{
	"Warszawa":     {"temp": 18, "desc": "Słonecznie", "humidity": 50, "wind": 3},
	"Krakow":       {"temp": 16, "desc": "Pochmurno", "humidity": 60, "wind": 2},
	"Gdansk":       {"temp": 14, "desc": "Deszcz", "humidity": 70, "wind": 5},
	"Berlin":       {"temp": 17, "desc": "Słonecznie", "humidity": 55, "wind": 3},
	"Munich":       {"temp": 15, "desc": "Zachmurzenie umiarkowane", "humidity": 65, "wind": 4},
	"Hamburg":      {"temp": 13, "desc": "Deszcz", "humidity": 75, "wind": 6},
	"Nowy York":    {"temp": 20, "desc": "Ciepło", "humidity": 55, "wind": 5},
	"Los Angeles":  {"temp": 25, "desc": "Gorąco", "humidity": 40, "wind": 2},
	"Chicago":      {"temp": 10, "desc": "Wietrznie", "humidity": 65, "wind": 7},
}

// Lokalizacje
var locations = map[string][]string{
	"Polska": {"Warszawa", "Krakow", "Gdansk"},
	"Niemcy": {"Berlin", "Munich", "Hamburg"},
	"USA":    {"Nowy York", "Los Angeles", "Chicago"},
}

// Szablon HTML
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	selectedCountry := "Polska"
	selectedCity := "Warszawa"
	var weather map[string]interface{}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Nie można przetworzyć formularza", http.StatusBadRequest)
			return
		}
		selectedCountry = r.FormValue("country")
		selectedCity = r.FormValue("city")
		weather = weatherData[selectedCity]
	}

	data := map[string]interface{}{
		"Locations":       locations,
		"SelectedCountry": selectedCountry,
		"SelectedCity":    selectedCity,
		"Weather":         weather,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "5000"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Nieprawidłowy port: %v", err)
	}

	log.Printf("Data uruchomienia: %s, Autor: %s, Nasłuchuje na porcie: %d", time.Now().Format("2006-01-02 15:04:05"), AUTHOR, port)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

