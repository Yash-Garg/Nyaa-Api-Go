package helpers

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Yash-Garg/nyaa-api-go/models"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

func scrapeNyaa(resp *fiber.Ctx, searchUrl string) {
	c := colly.NewCollector()
	c.OnHTML("tbody", torrentScraper)

	c.OnScraped(func(response *colly.Response) {
		if len(torrents) <= 0 {
			_ = resp.Redirect("/404")
		} else {
			_ = resp.Status(200).JSON(torrents)
			// set torrents list to nil after sending JSON
			// otherwise it will append other results
			// to the earlier searched data
			torrents = nil
		}
	})

	c.OnError(func(response *colly.Response, err error) {
		_ = resp.Redirect("/404")
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
					t.ID, _ = strconv.Atoi(strings.Split(torrentPath, "/")[2])
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

func fileInfoScraper(resp *fiber.Ctx, searchUrl string) {
	c := colly.NewCollector()
	t := models.File{}

	c.OnHTML("div.container", func(element *colly.HTMLElement) {
		t.Title = strings.TrimSpace(element.DOM.Find("h3.panel-title").First().Text())
		t.File = baseUrl + element.ChildAttr("div.panel-footer a", "href")
		t.Link = searchUrl
		t.ID, _ = strconv.Atoi(strings.Split(searchUrl, "/")[4])
		t.Magnet = element.ChildAttr("div.panel-footer a:nth-child(2)", "href")
		t.Description = element.ChildText("div.panel-body#torrent-description")
		t.Size = element.ChildText("div.panel-body div.row:nth-child(4) .col-md-5:nth-child(2)")
		t.Category = element.ChildText("div.panel-body div.row:nth-child(1) .col-md-5:nth-child(2)")
		t.Uploaded = element.ChildText("div.panel-body div.row:nth-child(1) .col-md-5:nth-child(4)")
		t.InfoHash = element.ChildText("div.panel-body div.row:nth-child(5) .col-md-5:nth-child(2)")
		t.SubmittedBy = element.ChildText("div.panel-body div.row:nth-child(2) .col-md-5:nth-child(2)")
		t.Seeders, _ = strconv.Atoi(element.ChildText("div.panel-body div.row:nth-child(2) .col-md-5:nth-child(4)"))
		t.Leechers, _ = strconv.Atoi(element.ChildText("div.panel-body div.row:nth-child(3) .col-md-5:nth-child(4)"))
	})

	c.OnError(func(response *colly.Response, err error) {
		_ = resp.Redirect("/404")
	})

	c.OnScraped(func(response *colly.Response) {
		_ = resp.Status(200).JSON(t)
	})

	_ = c.Visit(searchUrl)
}
