package domain

import "github.com/mmcdole/gofeed"

type item struct {
	Title     string `json:"title"`
	Published string `json:"published"`
}

type Source struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	Items       []item `json:"items"`
}

// NewSource creates from the RSS Feed with the following fields:
//   - feed.Title
//   - feed.Description
//   - feed.Image.URL
//   - feed.Link
//
// And for every item:
//   - Title
//   - Published
func NewSource(feed *gofeed.Feed) *Source {
	items := []item{}

	for _, it := range feed.Items {
		items = append(items, item{
			Title:     it.Title,
			Published: it.Published,
		})
	}
	return &Source{
		Title:       feed.Title,
		Description: feed.Description,
		Image:       feed.Image.URL,
		Link:        feed.FeedLink,
		Items:       items,
	}
}
