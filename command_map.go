package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func commandMap(cfg *config) error {
	var id int
	var urls []string
	if cfg.Next == nil {
		id = 0
		for i := 0; i < 20; i++ {
			fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", id+1+i)
			cfg.Next = append(cfg.Next, fullURL)
		}
		urls = cfg.Next
	} else {
		if cfg.Previous == nil {
			cfg.Previous = make([]string, len(cfg.Next))
		}
		copy(cfg.Previous, cfg.Next)
		lastURL := cfg.Previous[len(cfg.Previous)-1]
		urlParsed, err := url.Parse(lastURL)
		if err != nil {
			return err
		}
		urlPath := strings.Split(urlParsed.Path, "/")
		idString := urlPath[len(urlPath)-2]
		id, err = strconv.Atoi(idString)
		if err != nil {
			return err
		}
		for i := 0; i < 20; i++ {
			fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", id+1+i)
			cfg.Next[i] = fullURL
		}
		urls = cfg.Next
	}

	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, data)
		}
		var locationArea localArea
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return err
		}
		fmt.Println(locationArea.Name)
	}

	return nil
}

type localArea struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
