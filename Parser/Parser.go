package Parser

import (
	"bkcli/Model"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"time"
)

func GetArticles(top int) []Model.Article {
	var url = "https://borakasmer.com"
	var articleResult = make([]Model.Article, 0)
	var client = &http.Client{Timeout: 10 * time.Second}
	var response, err = client.Get(url)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(response.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	doc.Find(".post-title.entry-title a").Each(func(i int, s *goquery.Selection) {
		if i < top {
			articleResult = append(articleResult, Model.Article{
				Title: strconv.Itoa(i+1) + "-) " + s.Text(),
				Url:   s.AttrOr("href", ""),
			})
		} else {
			return
		}
	})
	return articleResult
}
