package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeProduct(url string) (float64, string, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64)"),
	)

	var price float64
	var title string
	var err error

	c.OnHTML("h1.ui-pdp-title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".andes-money-amount__fraction", func(e *colly.HTMLElement) {

		rawPrice := strings.ReplaceAll(e.Text, ".", "")
		rawPrice = strings.ReplaceAll(rawPrice, ",", ".")

		p, convErr := strconv.ParseFloat(rawPrice, 64)
		if convErr == nil {
			price = p
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		err = e
	})

	c.Visit(url)

	if price == 0 {
		return 0, "", fmt.Errorf("preço não encontrado ou site desconhecido")
	}

	return price, title, err
}
