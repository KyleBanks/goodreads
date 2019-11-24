package work

type Work struct {
	ID                       int     `xml:"id"`
	BooksCount               int     `xml:"books_count"`
	RatingsCount             int     `xml:"ratings_count"`
	TextReviewsCount         int     `xml:"text_reviews_count"`
	OriginalPublicationYear  int     `xml:"original_publication_year"`
	OriginalPublicationMonth int     `xml:"original_publication_month"`
	OriginalPublicationDay   int     `xml:"original_publication_day"`
	AverageRating            float64 `xml:"average_rating"`
	BestBook                 Book    `xml:"best_book"`
}

type Book struct {
	ID            int    `xml:"id"`
	Title         string `xml:"title"`
	Author        Author `xml:"author"`
	ImageURL      string `xml:"image_url"`
	SmallImageURL string `xml:"small_image_url"`
}

type Author struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}
