package models

type Torrent struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	Uploaded  string `json:"uploaded"`
	Seeders   int    `json:"seeders"`
	Leechers  int    `json:"leechers"`
	Completed int    `json:"completed"`
	Size      string `json:"size"`
	File      string `json:"file"`
	Link      string `json:"link"`
	Magnet    string `json:"magnet"`
}

type File struct {
	Torrent
	Description string   `json:"description"`
	SubmittedBy string   `json:"submittedBy"`
	InfoHash    string   `json:"infoHash"`
	CommentInfo Comments `json:"commentInfo"`
}
