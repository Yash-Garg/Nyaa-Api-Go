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
	query, sortParam, sortOrder, pageNum := GetSearchParameters(resp)

	if category != "" {
		searchUrl = fmt.Sprintf("%s?q=%s&c=%s&p=%d&s=%s&o=%s", baseUrl, strings.TrimSpace(query), category, pageNum, sortParam, sortOrder)
	} else {
		return resp.SendString("{response: 400, error: Invalid parameters}")
	}

	scrapeNyaa(resp, searchUrl)
	return nil
}

func GetUserUploads(resp *fiber.Ctx) error {
	username := resp.Params("username")
	query, sortParam, sortOrder, pageNum := GetSearchParameters(resp)

	searchUrl := fmt.Sprintf("%s/user/%s?q=%s&p=%d&s=%s&o=%s", baseUrl, username, strings.TrimSpace(query), pageNum, sortParam, sortOrder)

	scrapeNyaa(resp, searchUrl)
	return nil
}

func GetInfoFromID(resp *fiber.Ctx) error {
	id := resp.Params("id")
	// TODO: Implement file scraper function
	_ = resp.SendString("Requested File ID - " + id)
	return nil
}
