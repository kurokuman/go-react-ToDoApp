package databases

type Todo struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"contents"`
}

func (todo *Todo) Validate() string {
	if todo.Title == "" {
		return "todo title is none"
	}
	if todo.Content == "" {
		return "todo content is none"
	}
	return ""
}
