package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/garyburd/redigo/redis"
	"time"
)

type Item struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"pub_date"`
	SimilarUrls []string  `json:"similar_items"`
}

type Items []Item

func (_ Items) get(url string) (items Items) {
	doc, _ := goquery.NewDocument(url)

	doc.Find("item").Each(func(_ int, s *goquery.Selection) {
		title := s.Find("title").Text()
		link := s.Find("link").Text()
		description := s.Find("description").Text()
		pub_date, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", s.Find("pubDate").Text())
		key := fmt.Sprintf("%s/similar_news", link)
		// id := 4
		// rc.Do("ZINCRBY", key, 1, id)

		similar_urls, err := redis.Strings(rc.Do("ZREVRANGE", key, 0, -1))
		if err != nil {
			panic(err)
		}

		items = append(items, Item{Title: title, Link: link, Description: description, PubDate: pub_date, SimilarUrls: similar_urls})
	})
	return
}
