package models

type Torrent struct {
	Title       string `json:"title"`
	Uploaded    string `json:"uploaded"`
	Seeders     string `json:"seeders"`
	Leechers    string `json:"leechers"`
	Size        string `json:"size"`
	File        string `json:"file"`
	Information string `json:"information"`
	Magnet      string `json:"magnet"`
}
