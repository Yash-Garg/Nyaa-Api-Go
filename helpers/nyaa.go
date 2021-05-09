package helpers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"net/http"

	"strings"

	"github.com/Yash-Garg/nyaa-api-go/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckNyaaUrl() string {
	_, err := http.Get(utils.NyaaBaseUrl)
	if err != nil {
		return utils.NyaaAltUrl
	} else {
		return utils.NyaaBaseUrl
	}
}

func GetNyaa(resp *fiber.Ctx) error {
	baseUrl := CheckNyaaUrl()
	var query string = strings.ReplaceAll(resp.Query("q"), " ", "+")
	var searchUrl string = baseUrl + "?q=" + strings.TrimSpace(query)
	c := colly.NewCollector()
	// torrents := make([]models.Torrent, 0)

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.DOM.Each(func(i int, selection *goquery.Selection) {
			// t := models.Torrent{}
			fmt.Print(selection.Find("td:nth-child(2) a").Text())
		})
	})

	err := c.Visit(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
