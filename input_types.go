package goodreads

// SearchField defines the field types within which you can search.
// Defaults to AllFields.
type SearchField string

const (
	// AuthorField lets you specify that you want to search author names only.
	AuthorField SearchField = "author"

	// TitleField lets you specify that you want to only search through book titles.
	TitleField SearchField = "title"

	// AllFields (the default) lets you search over everything.
	AllFields SearchField = "all"
)
