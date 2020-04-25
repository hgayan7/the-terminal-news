package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//FetchNews from News API
func FetchNews(language string, topics []string) {

	req, err := http.NewRequest("GET", "http://newsapi.org/v2/everything", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	query := req.URL.Query()
	query.Add("apiKey", os.Getenv("apiKey"))
	if len(topics) > 0 {
		queryString := ""
		for _, topic := range topics {
			if topic != topics[len(topics)-1] {
				queryString += topic + " OR "
			} else {
				queryString += topic
			}
		}
		query.Add("q", queryString)
	} else {
		query.Add("q", "All")
	}
	query.Add("language", language)
	query.Add("sortBy", "publishedAt")

	query.Add("excludeDomains", "stackoverflow.com")

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
	jsonError := json.Unmarshal(body, &responseObject)
	if jsonError != nil {
		var errorObject Error
		json.Unmarshal(body, &errorObject)
		showError(errorObject)
		os.Exit(1)
	}
	if len(responseObject.News) == 0 {
		fmt.Println("No News Found. Try Using different flags.")
	} else {
		showNews(responseObject)
	}
}

func showNews(responseObject Response) {
	for index, news := range responseObject.News {
		fmt.Println(index+1, ".", news.Title)
		fmt.Println("Source :", news.Source.Name)
		fmt.Println(news.Content, "\n\n")
	}
}

func showError(errorObject Error) {
	fmt.Println(errorObject.Status)
	fmt.Println("Error Code : ", errorObject.Code)
	fmt.Println("Error Message : ", errorObject.Message)
}
