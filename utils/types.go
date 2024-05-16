package utils

type chapter struct {
	Title   string    `json:"title"`
	Desc    []string  `json:"story"`
	Options []options `json:"options"`
}

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type story map[string]chapter

type Handler struct {
	Data story
}
