package utils

var NyaaBaseUrl string = "https://nyaa.si"
var NyaaAltUrl string = "https://rss.webhook08.cf"

var AnimeSubCategoryMap = map[string]string{
	"all":     "1_0",
	"amv":     "1_1",
	"eng":     "1_2",
	"non-eng": "1_3",
	"raw":     "1_4",
}

var MangaSubCategoryMap = map[string]string{
	"all":     "1_0",
	"eng":     "1_2",
	"non-eng": "1_3",
	"raw":     "1_4",
}

var LiveActionSubCategoryMap = map[string]string{
	"all":     "4_0",
	"eng":     "4_1",
	"promo":   "4_2",
	"non-eng": "4_3",
	"raw":     "4_4",
}

var AudioSubCategoryMap = map[string]string{
	"all":      "2_0",
	"lossless": "2_1",
	"lossy":    "2_2",
}

var PicturesSubCategoryMap = map[string]string{
	"all":      "5_0",
	"graphics": "5_1",
	"photos":   "5_2",
}

var SoftwareSubCategoryMap = map[string]string{
	"all":          "6_0",
	"applications": "6_1",
	"games":        "6_2",
}
