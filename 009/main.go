package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imageUrl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://jiji.co.ke"),
	)

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		fmt.Println(h)
	})

	c.Visit("https://jiji.co.ke/houses-apartments-for-rent")
}
