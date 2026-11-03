package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const baseURL = "https://api.monkeytype.com"

var ErrUserNotFound = errors.New("user not found")

type PB struct {
	WPM         float64 `json:"wpm"`
	Language    string  `json:"language"`
	Punctuation bool    `json:"punctuation"`
	Numbers     bool    `json:"numbers"`
	Difficulty  string  `json:"difficulty"`
	LazyMode    bool    `json:"lazyMode"`
}

type ProfileResponse struct {
	Data struct {
		Name          string `json:"name"`
		PersonalBests struct {
			Time  map[string][]PB `json:"time"`
			Words map[string][]PB `json:"words"`
		} `json:"personalBests"`
	} `json:"data"`
}

type Records struct {
	Name                                string
	Time15, Time30, Time60, Time120     float64
	Words10, Words25, Words50, Words100 float64
}

func bestNormal(pbs []PB) float64 {
	best := 0.0
	for _, pb := range pbs {
		if pb.Language == "english" && !pb.Punctuation && !pb.Numbers &&
			pb.Difficulty == "normal" && !pb.LazyMode && pb.WPM > best {
			best = pb.WPM
		}
	}
	return best
}

func FetchRecords(name string) (Records, error) {
	resp, err := http.Get(baseURL + "/users/" + url.PathEscape(name) + "/profile")
	if err != nil {
		return Records{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return Records{}, ErrUserNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return Records{}, fmt.Errorf("status %d", resp.StatusCode)
	}

	var p ProfileResponse
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return Records{}, err
	}

	t, w := p.Data.PersonalBests.Time, p.Data.PersonalBests.Words
	return Records{
		Name:     p.Data.Name,
		Time15:   bestNormal(t["15"]),
		Time30:   bestNormal(t["30"]),
		Time60:   bestNormal(t["60"]),
		Time120:  bestNormal(t["120"]),
		Words10:  bestNormal(w["10"]),
		Words25:  bestNormal(w["25"]),
		Words50:  bestNormal(w["50"]),
		Words100: bestNormal(w["100"]),
	}, nil
}
