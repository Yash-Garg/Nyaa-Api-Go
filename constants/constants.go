package constants

var NyaaBaseUrl = "https://nyaa.si"
var NyaaAltUrl = "https://nyaa.kryptk.ml"

var NyaaEndpoints = map[string]map[string]string{
	"anime": {
		"all":     "1_0",
		"amv":     "1_1",
		"eng":     "1_2",
		"non-eng": "1_3",
		"raw":     "1_4",
	},
	"audio": {
		"all":      "2_0",
		"lossless": "2_1",
		"lossy":    "2_2",
	},
	"manga": {
		"all":     "3_0",
		"eng":     "3_1",
		"non-eng": "3_2",
		"raw":     "3_3",
	},
	"live_action": {
		"all":     "4_0",
		"eng":     "4_1",
		"promo":   "4_2",
		"non-eng": "4_3",
		"raw":     "4_4",
	},
	"pictures": {
		"all":      "5_0",
		"graphics": "5_1",
		"photos":   "5_2",
	},
	"software": {
		"all":          "6_0",
		"applications": "6_1",
		"games":        "6_2",
	},
}
