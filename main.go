package main

import (
	"cyoa/utils"
	"flag"
	"net/http"
)

func main() {
	filename := flag.String("j", "story.json", "JSON file having the stories")
	flag.Parse()

	jsonData := utils.JSONParser(*filename)
	handler := utils.Handler{Data: jsonData}
	http.ListenAndServe(":8080", handler)
}
