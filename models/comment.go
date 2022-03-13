package models

type Comments struct {
	Count    int       `json:"count"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Image     string `json:"image"`
	Timestamp string `json:"timestamp"`
}
