package databases

type Todo struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"contents"`
}
