package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	// cookie is set to expire after a week
	defaultCookieExpiry = time.Hour * 24 * 7
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("wheel/public"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/locations", updateLocations)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type PageVariables struct {
	NamedLocations   []LocationCheckBox
	UnnamedLocations []LocationCheckBox
	Wheel            Wheel
}

type LocationCheckBox struct {
	Name      string
	Value     string
	Selected  bool
	ImagePath string
}

type Wheel struct {
	LocationWeights map[string]int
}

func (w Wheel) Data() string {
	var s string
	for location, weight := range w.LocationWeights {
		if s != "" {
			s = fmt.Sprintf("%s;%s,%d", s, location, weight)
		} else {
			s = fmt.Sprintf("%s,%d", location, weight)
		}
	}
	return s
}

func newLocationsCookie(locations []string, t time.Time) *http.Cookie {
	return &http.Cookie{
		Name:    "locations",
		Value:   strings.Join(locations, ","),
		Path:    "/",
		Expires: t.Add(defaultCookieExpiry),
	}
}

func newLocationCheckBoxes(locations []string, selectedLocations []string) []LocationCheckBox {
	selectedLocationMap := make(map[string]bool, len(selectedLocations))
	for _, location := range selectedLocations {
		selectedLocationMap[location] = true
	}

	checkBoxes := make([]LocationCheckBox, len(locations))
	for i, location := range locations {
		_, selected := selectedLocationMap[location]
		checkBoxes[i] = LocationCheckBox{
			Name:      "location",
			Value:     location,
			Selected:  selected,
			ImagePath: fmt.Sprint(ImagesDir, LocationImages[location]),
		}
	}
	return checkBoxes
}

func newWheelData(locations []string) Wheel {
	locationWeights := make(map[string]int, len(locations))
	for _, location := range locations {
		// TODO hard coded
		locationWeights[location] = 1
	}
	return Wheel{locationWeights}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Load default or from cookie if exists
	var locationNames []string
	locationsCookie, err := r.Cookie("locations")
	switch err {
	case nil:
		locationNames = strings.Split(locationsCookie.Value, ",")
	case http.ErrNoCookie:
		locationNames = NamedLocations
		locationsCookie = newLocationsCookie(locationNames, time.Now())
		http.SetCookie(w, locationsCookie)
	default:
		log.Fatal("doh", err)
	}

	t, err := template.ParseFiles("wheel/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	pageVar := PageVariables{
		NamedLocations:   newLocationCheckBoxes(NamedLocations, locationNames),
		UnnamedLocations: newLocationCheckBoxes(UnnamedLocations, locationNames),
		Wheel:            newWheelData(locationNames),
	}
	fmt.Println(pageVar.Wheel.Data())
	err = t.Execute(w, pageVar)
	if err != nil {
		log.Print("tempalte executing error: ", err)
	}
}

func updateLocations(w http.ResponseWriter, r *http.Request) {
	// Update cookie with user form
	r.ParseForm()
	newLocationNames := r.Form["location"]
	http.SetCookie(w, newLocationsCookie(newLocationNames, time.Now()))
	http.Redirect(w, r, "/", http.StatusFound)
}
