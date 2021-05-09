package helpers

import (
	"fmt"
	"github.com/Yash-Garg/nyaa-api-go/models"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"

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
	query := strings.ReplaceAll(resp.Query("q"), " ", "+")
	searchUrl := baseUrl + "?q=" + strings.TrimSpace(query)
	c := colly.NewCollector()
	torrents := make([]models.TorrentLinks, 0)

	c.OnHTML("tbody", func(element *colly.HTMLElement) {
		element.DOM.Each(func(i int, selection *goquery.Selection) {
			t := models.TorrentLinks{}
			selection.Find("tr td:nth-child(2) a").Each(func(i int, selection *goquery.Selection) {
				if !selection.HasClass("comments") {
					t.Title, _ = selection.Attr("title")
					torrentLink, _ := selection.Attr("href")
					t.Link = baseUrl + torrentLink
					torrents = append(torrents, t)
				}
			})
		})
	})

	c.OnScraped(func(response *colly.Response) {
		if len(torrents) <= 0 {
			resp.Status(204)
		} else {
			_ = resp.Status(200).JSON(torrents)
		}
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
	})

	err := c.Visit(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
