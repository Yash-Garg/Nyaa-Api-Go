package helpers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Yash-Garg/nyaa-api-go/constants"
	"github.com/gofiber/fiber/v2"
)

func AliveHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Nyaa API v1 // Alive")
}

func CheckNyaaUrl() string {
	_, err := http.Get(constants.NyaaBaseUrl)
	if err != nil {
		return constants.NyaaAltUrl
	} else {
		return constants.NyaaBaseUrl
	}
}

func GetSearchParameters(resp *fiber.Ctx) (string, string, string, int, int) {
	q := strings.ReplaceAll(resp.Query("q"), " ", "+")
	p, _ := strconv.Atoi(resp.Query("p"))
	o := resp.Query("o")
	s := resp.Query("s")
	f, _ := strconv.Atoi(resp.Query("f"))
	if s == "date" {
		s = "id"
	}
	return q, s, o, p, f
}

func GetCategoryID(c string, s string) string {
	if len(s) == 0 {
		return constants.NyaaEndpoints[c]["all"]
	} else if c == "all" {
		return ""
	} else {
		return constants.NyaaEndpoints[c][s]
	}
}
