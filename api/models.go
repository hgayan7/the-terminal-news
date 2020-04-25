package api

type Response struct {
	Status string `json:"status"`
	News   []News `json:"articles"`
}

type News struct {
	Source      Source `json:"source"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Source struct {
	Name string `json:"name"`
}
