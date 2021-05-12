package helpers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Yash-Garg/nyaa-api-go/models"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func scrapeNyaa(resp *fiber.Ctx, searchUrl string) {
	c := colly.NewCollector()
	c.OnHTML("tbody", torrentScraper)

	c.OnScraped(func(response *colly.Response) {
		if len(torrents) <= 0 {
			_ = resp.SendString("{response: 404, error: No results found}")
		} else {
			_ = resp.Status(200).JSON(torrents)
			// set torrents list to nil after sending JSON
			// otherwise it will append other results
			// to the earlier searched data
			torrents = nil
		}
	})

	c.OnError(func(response *colly.Response, err error) {
		_ = resp.SendString("{ response: 404, error: Not Found}")
	})

	_ = c.Visit(searchUrl)
}

func torrentScraper(element *colly.HTMLElement) {
	element.DOM.Each(func(i int, selection *goquery.Selection) {
		t := models.Torrent{}
		selection.Find("tr").Each(func(i int, selection *goquery.Selection) {
			selection.Find("td:nth-child(2) a").Each(func(i int, selection *goquery.Selection) {
				if !selection.HasClass("comments") {
					t.Title, _ = selection.Attr("title")
					torrentPath, _ := selection.Attr("href")
					t.Link = baseUrl + torrentPath
				}
			})
			t.Category, _ = selection.Find("td:nth-child(1) a").Attr("title")
			filePath, _ := selection.Find("td:nth-child(3) a:nth-child(1)").Attr("href")
			t.File = baseUrl + filePath
			t.Magnet, _ = selection.Find("td:nth-child(3) a:nth-child(2)").Attr("href")
			t.Size = selection.Find("td:nth-child(4)").Text()
			t.Uploaded = selection.Find("td:nth-child(5)").Text()
			t.Seeders, _ = strconv.Atoi(selection.Find("td:nth-child(6)").Text())
			t.Leechers, _ = strconv.Atoi(selection.Find("td:nth-child(7)").Text())
			torrents = append(torrents, t)
		})
	})
}
