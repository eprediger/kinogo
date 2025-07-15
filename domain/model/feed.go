package domain

import (
	"net/url"

	dto "infrastructure/http/feed/dto"
)

type Feed interface {
	GetUrl() string
}

type feed struct {
	url *url.URL
}

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
