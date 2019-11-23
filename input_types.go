package goodreads

type SearchField string

const (
	AuthorField SearchField = "author"
	TitleField  SearchField = "title"
	AllFields   SearchField = "all"
)
