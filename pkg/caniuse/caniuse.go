package caniuse

import (
	"encoding/json"
	"io/ioutil"
)

type agent struct {
	Browser     string             `json:"browser"`
	Abbr        string             `json:"abbr"`
	Prefix      string             `json:"prefix"`
	Type        string             `json:"type"`
	UsageGlobal map[string]float64 `json:"usage_global"`
	Versions    []string           `json:"versions"`
}

// DB contains all information of original json.
type DB struct {
	Agents   map[string]agent    `json:"agents"`
	Cats     map[string][]string `json:"cats"`
	Data     map[string]datum    `json:"data"`
	Eras     map[string]string   `json:"eras"`
	Statuses map[string]string   `json:"statuses"`
	Updated  int32               `json:"updated"`
}

type datum struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Spec        string `json:"spec"`
	Status      string `json:"status"`
	Links       []struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"links"`
	Categories []string                       `json:"categories"`
	Stats      map[string](map[string]string) `json:"stats"`
	Notes      string                         `json:"notes"`
	NotesByNum map[string]string              `json:"notes_by_num"`
	UsagePercY float64                        `json:"usage_perc_y"`
	UsagePercA float64                        `json:"usage_perc_a"`
	Ucprefix   bool                           `json:"ucprefix"`
	Parent     string                         `json:"parent"`
	Keywords   string                         `json:"keywords"`
	IeID       string                         `json:"ie_id"`
	ChromeID   string                         `json:"chrome_id"`
	FirefoxID  string                         `json:"firefox_id"`
	WebkitID   string                         `json:"webkit_id"`
}

// GetAll return all information of original json
func GetAll() DB {
	jsonPath := getDataPath()

	fullData, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}

	var db DB
	e := json.Unmarshal(fullData, &db)

	if e != nil {
		panic(e)
	}

	return db
}
