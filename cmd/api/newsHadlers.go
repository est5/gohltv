package main

import (
	"encoding/json"
	"github.com/est5/gohltv/pkg/models"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) GetNews(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	var articles []models.NewsArticle

	prefix := "https://www.hltv.org"
	c.OnHTML("div.standard-box.standard-list", func(e *colly.HTMLElement) {
		e.ForEach("a.newsline.article", func(_ int, element *colly.HTMLElement) {
			link := prefix + element.Attr("href")
			text := element.ChildText("div.newstext")
			data := strings.Split(element.ChildText("div.newstc"), "\n")
			date := data[0]
			comments := strings.Split(strings.TrimSpace(data[1]), " ")
			commentCount, _ := strconv.Atoi(comments[0])
			article := models.NewsArticle{
				Link:          link,
				Text:          text,
				CommentsCount: commentCount,
				Date:          date,
			}
			articles = append(articles, article)
		})
	})

	url := "https://www.hltv.org/news/archive/" + year + "/" + month
	err := c.Visit(url)
	// u can visit smth like /news/archive/2025/january
	if err != nil {
		app.log.Error(err)
		err := c.Visit("https://www.hltv.org/news/archive/")
		if err != nil {
			app.log.Fatal(err)
		}
	}

	js, err := json.MarshalIndent(articles, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		app.log.Fatal(err)
	}

}
