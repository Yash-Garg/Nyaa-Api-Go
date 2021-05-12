package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/Yash-Garg/nyaa-api-go/models"
	"github.com/Yash-Garg/nyaa-api-go/utils"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

func GetCategoryID(c string, s string) string {
	if len(s) == 0 {
		return utils.NyaaEndpoints[c]["all"]
	} else {
		return utils.NyaaEndpoints[c][s]
	}
}

func GetNyaa(resp *fiber.Ctx) error {
	var searchUrl string
	baseUrl := CheckNyaaUrl()

	searchQuery := strings.ReplaceAll(resp.Query("q"), " ", "+")
	pageNum, _ := strconv.Atoi(resp.Query("p"))
	sortOrder := resp.Query("o")
	category := GetCategoryID(resp.Params("category"), resp.Params("sub_category"))
	sortParam := resp.Query("s")

	if sortParam == "date" {
		sortParam = "id"
	}

	if category != "" {
		searchUrl = fmt.Sprintf("%s?q=%s&c=%s&p=%d&s=%s&o=%s", baseUrl, strings.TrimSpace(searchQuery), category, pageNum, sortParam, sortOrder)
	} else {
		return resp.SendString("{response: 400, error: Invalid parameters}")
	}

	c := colly.NewCollector()
	torrents := make([]models.Torrent, 0)

	c.OnHTML("tbody", func(element *colly.HTMLElement) {
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
	})

	c.OnScraped(func(response *colly.Response) {
		if len(torrents) <= 0 {
			_ = resp.SendString("{response: 404, error: No results found}")
		} else {
			_ = resp.Status(200).JSON(torrents)
		}
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Printf("{ response : %d, error: %s}", response.StatusCode, err)
	})

	_ = c.Visit(searchUrl)

	return nil
}
