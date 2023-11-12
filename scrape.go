package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string
	Link  string
}

func main() {
	articles := make([]Article, 0)

	c := colly.NewCollector(
		colly.AllowedDomains("technewssite.com"),
	)

	c.OnHTML(".article", func(e *colly.HTMLElement) {
		title := e.ChildText("h2")
		link := e.ChildAttr("a", "href")
		articles = append(articles, Article{
			Title: title,
			Link:  link,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	http.HandleFunc("/scrape", func(w http.ResponseWriter, r *http.Request) {
		_ = c.Visit("https://technewssite.com/latest")
		js, err := json.Marshal(articles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
