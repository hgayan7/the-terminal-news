package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func FetchNews() {
	req, err := http.NewRequest("GET", "http://newsapi.org/v2/everything", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	query := req.URL.Query()
	query.Add("apiKey", os.Getenv("apiKey"))
	query.Add("q", "swift OR iOS")
	req.URL.RawQuery = query.Encode()

	res, err := http.Get(req.URL.String())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var responseObject Response
	json.Unmarshal(body, &responseObject)

	fmt.Println(len(responseObject.News))

}
