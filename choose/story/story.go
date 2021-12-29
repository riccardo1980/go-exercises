package story

import (
	"encoding/json"
	"log"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Story map[string]Arc

func FromFile(s string) (Story, error) {
	log.Printf("Loading from %s", s)
	bytes, err := os.ReadFile(s)
	if err != nil {
		log.Fatal("Can't open s")
	}
	st, err := load([]byte(bytes))
	if err != nil {
		log.Fatal("Can't load from s")
	}
	log.Printf("Arcs: %d", len(st))
	return st, err
}

func load(byteValue []byte) (Story, error) {

	var story Story
	err := json.Unmarshal(byteValue, &story)

	return story, err

}
