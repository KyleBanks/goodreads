package goodreads

type Author struct {
	ID               string  `xml:"id"`
	Name             string  `xml:"name"`
	ImageURL         string  `xml:"image_url"`
	SmallImageURL    string  `xml:"small_image_url"`
	Link             string  `xml:"link"`
	AverageRating    float32 `xml:"average_rating"`
	RatingsCount     int     `xml:"ratings_count"`
	TextReviewsCount int     `xml:"text_reviews_count"`
}

type Book struct {
	ID                 string   `xml:"id"`
	ISBN               string   `xml:"isbn"`
	ISBN13             string   `xml:"isbn13"`
	TextReviewsCount   int      `xml:"text_reviews_count"`
	URI                string   `xml:"uri"`
	Title              string   `xml:"title"`
	TitleWithoutSeries string   `xml:"title_without_series"`
	ImageURL           string   `xml:"image_url"`
	SmallImageURL      string   `xml:"small_image_url"`
	LargeImageURL      string   `xml:"large_image_url"`
	Link               string   `xml:"link"`
	NumPages           int      `xml:"num_pages"`
	Format             string   `xml:"format"`
	EditionInformation string   `xml:"edition_information"`
	Publisher          string   `xml:"publisher"`
	PublicationDay     int      `xml:"publication_day"`
	PublicationYear    int      `xml:"publication_year"`
	PublicationMonth   int      `xml:"publication_month"`
	AverageRating      float32  `xml:"average_rating"`
	RatingsCount       int      `xml:"ratings_count"`
	Description        string   `xml:"description"`
	Authors            []Author `xml:"authors>author"`
}

type Review struct {
	ID          string `xml:"id"`
	Book        Book   `xml:"book"`
	Rating      int    `xml:"rating"`
	StartedAt   string `xml:"started_at"`
	ReadAt      string `xml:"read_at"`
	DateAdded   string `xml:"date_added"`
	DateUpdated string `xml:"date_updated"`
	ReadCount   int    `xml:"read_count"`
	Body        string `xml:"body"`
}

type User struct {
	ID            string      `xml:"id"`
	Name          string      `xml:"name"`
	Link          string      `xml:"link"`
	ImageURL      string      `xml:"image_url"`
	SmallImageURL string      `xml:"small_image_url"`
	About         string      `xml:"about"`
	Gender        string      `xml:"gender"`
	Location      string      `xml:"location"`
	Website       string      `xml:"website"`
	Joined        string      `xml:"joined"`
	LastActive    string      `xml:"last_active"`
	FriendsCount  int         `xml:"friends_count"`
	GroupsCount   int         `xml:"groups_count"`
	ReviewCount   int         `xml:"reviews_count"`
	UserShelves   []UserShelf `xml:"user_shelves>user_shelf"`
}

type UserShelf struct {
	ID            string `xml:"id"`
	Name          string `xml:"name"`
	BookCount     string `xml:"book_count"`
	ExclusiveFlag bool   `xml:"exclusive_flag"`
	Description   string `xml:"description"`
}
