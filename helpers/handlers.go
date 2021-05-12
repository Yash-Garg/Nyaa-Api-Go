package helpers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Yash-Garg/nyaa-api-go/utils"
	"github.com/gofiber/fiber/v2"
)

func AliveHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Nyaa API v1 // Alive")
}

func CheckNyaaUrl() string {
	_, err := http.Get(utils.NyaaBaseUrl)
	if err != nil {
		return utils.NyaaAltUrl
	} else {
		return utils.NyaaBaseUrl
	}
}

func GetSearchParameters(resp *fiber.Ctx) (string, string, string, int) {
	q := strings.ReplaceAll(resp.Query("q"), " ", "+")
	p, _ := strconv.Atoi(resp.Query("p"))
	o := resp.Query("o")
	s := resp.Query("s")
	if s == "date" {
		s = "id"
	}
	return q, s, o, p
}

func GetCategoryID(c string, s string) string {
	if len(s) == 0 {
		return utils.NyaaEndpoints[c]["all"]
	} else {
		return utils.NyaaEndpoints[c][s]
	}
}
