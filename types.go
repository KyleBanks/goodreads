package goodreads

type User struct {
	ID            string `xml:"id"`
	Name          string `xml:"name"`
	Link          string `xml:"link"`
	ImageURL      string `xml:"image_url"`
	SmallImageURL string `xml:"small_image_url"`
	About         string `xml:"about"`
	Gender        string `xml:"gender"`
	Location      string `xml:"location"`
	Website       string `xml:"website"`
	Joined        string `xml:"joined"`
	LastActive    string `xml:"last_active"`
	FriendsCount  int    `xml:"friends_count"`
	GroupsCount   int    `xml:"groups_count"`
	ReviewCount   int    `xml:"reviews_count"`

	UserShelves []UserShelf `xml:"user_shelves>user_shelf"`
}

type UserShelf struct {
	ID            string `xml:"id"`
	Name          string `xml:"name"`
	BookCount     string `xml:"book_count"`
	ExclusiveFlag bool   `xml:"exclusive_flag"`
	Description   string `xml:"description"`
}
