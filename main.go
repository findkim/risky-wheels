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

var (
	zeroDuration time.Duration
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("wheel/public"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/locations", updateLocations)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// PageVariables contains template values
type PageVariables struct {
	NamedLocations   Locations
	UnnamedLocations Locations
}

type Locations []Location

// Location manages metadata for a location
type Location struct {
	Type      string
	Value     string
	Selected  bool
	ImagePath string
}

// newLocationsCookie initializes a new cookie with list of locations
func newLocationsCookie(locations []string, t time.Time, expiry time.Duration) *http.Cookie {
	var expireTime time.Time
	if expiry != zeroDuration {
		expireTime = t.Add(defaultCookieExpiry)
	}

	return &http.Cookie{
		Name:    "locations",
		Value:   strings.Join(locations, ","),
		Path:    "/",
		Expires: expireTime,
	}
}

func newLocations(allLocations []string, selectedLocations []string, locationType string) Locations {
	selectedLocationMap := make(map[string]bool, len(selectedLocations))
	for _, location := range selectedLocations {
		selectedLocationMap[location] = true
	}

	locations := make(Locations, len(allLocations))
	for i, location := range allLocations {
		_, selected := selectedLocationMap[location]
		imgPath, ok := LocationImages[location]
		if ok {
			imgPath = fmt.Sprint(ImagesDir, imgPath)
		}
		locations[i] = Location{
			Type:      locationType,
			Value:     location,
			Selected:  selected,
			ImagePath: imgPath,
		}
	}
	return locations
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
		locationsCookie = newLocationsCookie(locationNames, time.Now(), defaultCookieExpiry)
		http.SetCookie(w, locationsCookie)
	default:
		log.Fatal("doh", err)
	}

	t, err := template.ParseFiles("wheel/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	pageVar := PageVariables{
		NamedLocations:   newLocations(NamedLocations, locationNames, "named"),
		UnnamedLocations: newLocations(UnnamedLocations, locationNames, "unnamed"),
	}
	err = t.Execute(w, pageVar)
	if err != nil {
		log.Print("tempalte executing error: ", err)
	}
}

func updateLocations(w http.ResponseWriter, r *http.Request) {
	// Update cookie with user form
	r.ParseForm()
	newLocationNames := r.Form["location"]
	http.SetCookie(w, newLocationsCookie(newLocationNames, time.Now(), defaultCookieExpiry))
	http.Redirect(w, r, "/", http.StatusFound)
}
