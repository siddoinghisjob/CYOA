package utils

import (
	"encoding/json"
	"os"
)

func JSONParser(f string) story {
	r, err := os.Open(f)
	if err != nil {
		os.Exit(1)
	}
	d := json.NewDecoder(r)
	var storyArcs story
	err = d.Decode(&storyArcs)
	if err != nil {
		os.Exit(1)
	}
	return storyArcs
}
