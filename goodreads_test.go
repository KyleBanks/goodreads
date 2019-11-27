package goodreads

import (
	"fmt"
	"github.com/KyleBanks/goodreads/responses"
	"github.com/KyleBanks/goodreads/responses/work"
	"github.com/enribd/goodreads/responses/series"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testAPIKey = "test-api-key"

func TestNewClient(t *testing.T) {
	c := NewClient("api-key")
	assert.NotNil(t, c)
	assert.Equal(t, "api-key", c.APIKey)
	assert.Equal(t, defaultAPIClient, c.httpClient)
}

func TestClient_AuthorBooks(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/author/list/12345?key=%s&page=1", testAPIKey),
		response:  `<response><author><id>AuthorID</id><name>AuthorName</name></author></response>`,
	})
	defer done()

	a, err := c.AuthorBooks("12345", 1)
	assert.Nil(t, err)
	assert.Equal(t, responses.Author{
		ID:   "AuthorID",
		Name: "AuthorName",
	}, *a)
}

func TestClient_AuthorShow(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/author/show/12345?key=%s", testAPIKey),
		response:  `<response><author><id>AuthorID</id><name>AuthorName</name></author></response>`,
	})
	defer done()

	a, err := c.AuthorShow("12345")
	assert.Nil(t, err)
	assert.Equal(t, responses.Author{
		ID:   "AuthorID",
		Name: "AuthorName",
	}, *a)
}

func TestClient_BookReviewCounts(t *testing.T) {
	isbn := "9781400078776"
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/book/review_counts.json?isbns=%s&key=%s", isbn, testAPIKey),
		response: `{
			"books": [{
				"average_rating": "3.82",
				"id": 15,
				"isbn": "1400078776",
				"isbn13": "9781400078776",
				"ratings_count": 1,
				"reviews_count": 2,
				"text_reviews_count": 3,
				"work_ratings_count": 4,
				"work_reviews_count": 5,
				"work_text_reviews_count": 6
			}]
		}`,
	})
	defer done()

	counts, err := c.BookReviewCounts([]string{isbn})
	assert.Nil(t, err)
	assert.Equal(t, []responses.ReviewCounts{
		{
			ID:                   15,
			ISBN:                 "1400078776",
			ISBN13:               "9781400078776",
			RatingsCount:         1,
			ReviewsCount:         2,
			TextReviewsCount:     3,
			WorkRatingsCount:     4,
			WorkReviewsCount:     5,
			WorkTextReviewsCount: 6,
			AverageRating:        "3.82",
		},
	}, counts)
}

func TestClient_ReviewList(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/review/list/user-id.xml?key=%s&order=d&page=1&per_page=200&search=search&shelf=read&sort=date_read&v=2", testAPIKey),
		response: `<response>
			<reviews>
				<review><id>review1</id><rating>1</rating></review>
				<review><id>review2</id><rating>2</rating></review>
				<review><id>review3</id><rating>3</rating></review>
			</reviews>
		</response>`,
	})
	defer done()

	r, err := c.ReviewList("user-id", "read", "date_read", "search", "d", 1, 200)
	assert.Nil(t, err)
	assert.Equal(t, []responses.Review{
		{ID: "review1", Rating: 1},
		{ID: "review2", Rating: 2},
		{ID: "review3", Rating: 3},
	}, r)
}

func TestClient_SeriesShow(t *testing.T) {
	seriesID := "12345"
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/series/show/%s?format=xml&key=%s", seriesID, testAPIKey),
		response: `<response>
        <series>
           <id>12345</id>
           <title>Test Series</title>
           <description>Test Description</description>
           <note>Test Note</note>
           <series_works_count>2</series_works_count>
           <primary_work_count>2</primary_work_count>
           <numbered>true</numbered>
           <series_works>
              <series_work>
                 <id>1</id>
                 <user_position>1</user_position>
                 <work>
                    <id>11</id>
                    <uri>uri://fake-uri</uri>
                    <best_book>
                       <id>111</id>
                       <title>Test Book 1</title>
                       <author>
                          <id>9999</id>
                          <name>Test Author</name>
                       </author>
                       <image_url>https://image-provider.com/image.jpg</image_url>
                    </best_book>
                    <books_count>2</books_count>
                    <original_publication_day />
                    <original_publication_month />
                    <original_publication_year>1957</original_publication_year>
                    <original_title>Test Book 1</original_title>
                    <ratings_count>817</ratings_count>
                    <ratings_sum>3311</ratings_sum>
                    <reviews_count>1638</reviews_count>
                    <text_reviews_count>33</text_reviews_count>
                    <average_rating />
                 </work>
              </series_work>
           </series_works>
           <series_works>
              <series_work>
                 <id>2</id>
                 <user_position>2</user_position>
                 <work>
                    <id>22</id>
                    <uri>uri://fake-uri</uri>
                    <best_book>
                       <id>222</id>
                       <title>Test Book 2</title>
                       <author>
                          <id>9999</id>
                          <name>Test Author</name>
                       </author>
                       <image_url>https://image-provider.com/image2.jpg</image_url>
                    </best_book>
                    <books_count>2</books_count>
                    <original_publication_day>10</original_publication_day>
                    <original_publication_month>9</original_publication_month>
                    <original_publication_year>1989</original_publication_year>
                    <original_title>Test Book 2</original_title>
                    <ratings_count>933</ratings_count>
                    <ratings_sum>411</ratings_sum>
                    <reviews_count>842</reviews_count>
                    <text_reviews_count>53</text_reviews_count>
                    <average_rating>3.59</average_rating>
                 </work>
              </series_work>
        </series>
	</response>`})
	defer done()
	series, err := c.SeriesShow("12345")
	assert.Nil(t, err)
	assert.Equal(t, series.Series{
		{
			ID:               12345,
			Title:            "Test Series",
			Description:      "Test Description",
			Note:             "Test Note",
			SeriesWorksCount: 2,
			PrimaryWorkCount: 2,
			Numbered:         true,
			SeriesWorks: {
				{
					ID:           1,
					UserPosition: "1",
					Work: work.Work{
						ID:                       11,
						BooksCount:               2,
						RatingsCount:             817,
						TextReviewsCount:         33,
						OriginalPublicationYear:  1989,
						OriginalPublicationMonth: 9,
						OriginalPublicationDay:   10,
						AverageRating:            3.59,
						BestBook: work.Book{
							ID:    111,
							Title: "Test Book 1",
							Author: work.Author{
								ID:   9999,
								Name: "Test Author",
							},
							ImageURL: "https://image-provider.com/image.jpg",
						},
					},
				},
				{
					ID:           2,
					UserPosition: "2",
					Work: work.Work{
						ID:                       22,
						BooksCount:               2,
						RatingsCount:             933,
						TextReviewsCount:         53,
						OriginalPublicationYear:  1989,
						OriginalPublicationMonth: 9,
						OriginalPublicationDay:   10,
						AverageRating:            3.59,
						BestBook: work.Book{
							ID:    2,
							Title: "Test Book 2",
							Author: work.Author{
								ID:   9999,
								Name: "Test Author",
							},
							ImageURL: "https://image-provider.com/image2.jpg",
						},
					},
				},
			},
		},
	}, *series)
}

func TestClient_SearchBooks(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/search/index.xml?key=%s&page=1&q=hello&search%%5Bfield%%5D=all", testAPIKey),
		response: `<response>
		<search>
		  <results>
			<work>
			  <id type="integer">1</id>
			  <books_count type="integer">2</books_count>
			  <ratings_count type="integer">3</ratings_count>
			  <text_reviews_count type="integer">4</text_reviews_count>
			  <original_publication_year type="integer">2019</original_publication_year>
			  <original_publication_month type="integer">8</original_publication_month>
			  <original_publication_day type="integer">6</original_publication_day>
			  <average_rating>3.59</average_rating>
			  <best_book type="Book">
				<id type="integer">1</id>
				<title>book1</title>
				<author>
				  <id type="integer">1</id>
				  <name>Author 1</name>
				</author>
				<image_url>https://image1.jpg</image_url>
				<small_image_url>https://small_image1.jpg</small_image_url>
			  </best_book>
			</work>
			<work>
			  <id type="integer">5</id>
			  <books_count type="integer">6</books_count>
			  <ratings_count type="integer">7</ratings_count>
			  <text_reviews_count type="integer">8</text_reviews_count>
			  <original_publication_year type="integer">2018</original_publication_year>
			  <original_publication_month type="integer" nil="true" />
			  <original_publication_day type="integer" nil="true" />
			  <average_rating>3.68</average_rating>
			  <best_book type="Book">
				<id type="integer">2</id>
				<title>Hello: The Sequel</title>
				<author>
				  <id type="integer">2</id>
				  <name>Author 2</name>
				</author>
				<image_url>https://image2.jpg</image_url>
				<small_image_url>https://small_image2.jpg</small_image_url>
			  </best_book>
			</work>
		  </results>
		</search>
	</response>`})
	defer done()
	books, err := c.SearchBooks("hello", 1, AllFields)
	assert.Nil(t, err)
	assert.Equal(t, []work.Work{
		{
			ID:                       1,
			BooksCount:               2,
			RatingsCount:             3,
			TextReviewsCount:         4,
			OriginalPublicationYear:  2019,
			OriginalPublicationMonth: 8,
			OriginalPublicationDay:   6,
			AverageRating:            3.59,
			BestBook: work.Book{
				ID:    1,
				Title: "book1",
				Author: work.Author{
					ID:   1,
					Name: "Author 1",
				},
				ImageURL:      "https://image1.jpg",
				SmallImageURL: "https://small_image1.jpg",
			},
		},
		{
			ID:                       5,
			BooksCount:               6,
			RatingsCount:             7,
			TextReviewsCount:         8,
			OriginalPublicationYear:  2018,
			OriginalPublicationMonth: 0,
			OriginalPublicationDay:   0,
			AverageRating:            3.68,
			BestBook: work.Book{
				ID:    2,
				Title: "Hello: The Sequel",
				Author: work.Author{
					ID:   2,
					Name: "Author 2",
				},
				ImageURL:      "https://image2.jpg",
				SmallImageURL: "https://small_image2.jpg",
			},
		},
	}, books)
}

func TestClient_ShelvesList(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/shelf/list.xml?key=%s&user_id=user-id", testAPIKey),
		response: `<response>
			<shelves>
				<user_shelf><id>shelf1</id><name>Shelf 1</name></user_shelf>
				<user_shelf><id>shelf2</id><name>Shelf 2</name></user_shelf>
				<user_shelf><id>shelf3</id><name>Shelf 3</name></user_shelf>
			</shelves>
		</response>`,
	})
	defer done()

	s, err := c.ShelvesList("user-id")
	assert.Nil(t, err)
	assert.Equal(t, []responses.UserShelf{
		{ID: "shelf1", Name: "Shelf 1"},
		{ID: "shelf2", Name: "Shelf 2"},
		{ID: "shelf3", Name: "Shelf 3"},
	}, s)
}

func TestClient_UserShow(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/user/show/user-id.xml?key=%s", testAPIKey),
		response: `<response>
			<user>
				<id>user-id</id>
				<name>User Name</name>
			</user>
		</response>`,
	})
	defer done()

	u, err := c.UserShow("user-id")
	assert.Nil(t, err)
	assert.Equal(t, responses.User{
		ID:   "user-id",
		Name: "User Name",
	}, *u)
}

type decodeTestCase struct {
	expectURL string
	response  string
}

func newTestClient(t *testing.T, tc decodeTestCase) (*Client, func()) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, tc.expectURL, r.URL.String())
		_, _ = w.Write([]byte(tc.response))
	}))

	return &Client{
		APIKey: testAPIKey,
		httpClient: &httpClient{
			Client:  http.DefaultClient,
			APIRoot: s.URL,
			Verbose: true,
		},
	}, s.Close
}
