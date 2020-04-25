package api

type Response struct {
	Status string `json:"status"`
	News   []News `json:"articles"`
}

type News struct {
	Source  Source `json:"source"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Source struct {
	Name string `json:"name"`
}

type Error struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
