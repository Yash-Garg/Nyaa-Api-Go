package helpers

import (
	"fmt"
	"strings"

	"github.com/Yash-Garg/nyaa-api-go/models"
	"github.com/gofiber/fiber/v2"
)

var baseUrl = CheckNyaaUrl()
var torrents = make([]models.Torrent, 0)

func GetCategoryTorrents(resp *fiber.Ctx) error {
	var searchUrl string

	category := GetCategoryID(resp.Params("category"), resp.Params("sub_category"))
	query, sortParam, sortOrder, pageNum, filter := GetSearchParameters(resp)

	searchUrl = fmt.Sprintf("%s?q=%s&c=%s&p=%d&s=%s&o=%s&f=%d", baseUrl, strings.TrimSpace(query), category, pageNum, sortParam, sortOrder, filter)

	scrapeNyaa(resp, searchUrl)
	return nil
}

func GetUserUploads(resp *fiber.Ctx) error {
	username := resp.Params("username")
	query, sortParam, sortOrder, pageNum, filter := GetSearchParameters(resp)

	searchUrl := fmt.Sprintf("%s/user/%s?q=%s&p=%d&s=%s&o=%s&f=%d", baseUrl, username, strings.TrimSpace(query), pageNum, sortParam, sortOrder, filter)

	scrapeNyaa(resp, searchUrl)
	return nil
}

func GetInfoFromID(resp *fiber.Ctx) error {
	id := resp.Params("id")
	searchUrl := baseUrl + "/view/" + id

	fileInfoScraper(resp, searchUrl)
	return nil
}
