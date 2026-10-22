package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"strings"
	"sync"
	"time"
)

const baseURL = "https://groupietrackers.01edu.ai/api"

var client = &http.Client{Timeout: 10 * time.Second}

func FetchData[T any](url string) (T, error) {
	var data T
	resp, err := client.Get(url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("unexpected HTTP status %d:",
			resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, err
}

func FetchAllData() ([]models.Artist, error) {
	var (
		artists                []models.Artist
		locations              models.Location
		dates                  models.Date
		relations              models.Relation
		err1, err2, err3, err4 error
		wg                     sync.WaitGroup
	)
	wg.Add(4)
	go func() {
		defer wg.Done()
		artists, err1 = FetchData[[]models.Artist](baseURL + "/artists")
	}()
	go func() {
		defer wg.Done()
		locations, err2 = FetchData[models.Location](baseURL + "/locations")
	}()
	go func() {
		defer wg.Done()
		dates, err3 = FetchData[models.Date](baseURL + "/dates")
	}()
	go func() {
		defer wg.Done()
		relations, err4 = FetchData[models.Relation](baseURL + "/relation")
	}()
	wg.Wait()
	if err := errors.Join(err1, err2, err3, err4); err != nil {
		return nil, err
	}
	locationMap := make(map[int][]string)
	for _, v := range locations.Index {
		formattedLocs := make([]string, len(v.Locations))
		for i, loc := range v.Locations {
			formattedLocs[i] = formatLocation(loc)
		}
		locationMap[v.ID] = formattedLocs
	}
	dateMap := make(map[int][]string)
	for _, v := range dates.Index {
		dateMap[v.ID] = v.Dates
	}
	relationsMap := make(map[int]map[string][]string)
	for _, v := range relations.Index {
		formattedRel := make(map[string][]string, len(v.DatesLocations))
		for loc, d := range v.DatesLocations {
			formattedRel[formatLocation(loc)] = d
		}
		relationsMap[v.ID] = formattedRel
	}
	for i := range artists {
		artists[i].Locations = locationMap[artists[i].ID]
		artists[i].ConcertDates = dateMap[artists[i].ID]
		artists[i].Relations = relationsMap[artists[i].ID]

	}
	return artists, nil
}
func formatLocation(loc string) string {
	parts := strings.Split(loc, "-") // to split country from city
	for i, part := range parts {
		words := strings.Split(part, "_") //to seperate multi-word name like new zealand
		for j, word := range words {
			lower := strings.ToLower(word)
			if lower == "usa" || lower == "uk" || lower == "uae" {
				words[j] = strings.ToUpper(lower)
			} else if len(lower) > 0 {
				words[j] = strings.ToUpper(lower[:1]) + lower[1:]
			}
		}
		parts[i] = strings.Join(words, " ")
	}
	return strings.Join(parts, ", ")
}
