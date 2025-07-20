package domain

import (
	"net/url"

	dto "infrastructure/http/feed/dto"
)

// Feed defines the methods to interact with the RSS Feed URL
type Feed interface {
	GetUrl() string
}

type feed struct {
	url *url.URL
}

// NewFeed returns an instance of a Feed.
//
// The function with return an error if the Url is not valid
func NewFeed(feedDto dto.FeedDto) (Feed, error) {
	url, err := url.ParseRequestURI(feedDto.Url)
	if err != nil {
		return nil, err
	}

	return &feed{url}, nil
}

func (f *feed) GetUrl() string {
	return f.url.String()
}
