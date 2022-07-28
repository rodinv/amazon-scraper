package amazon_scraper

import "github.com/pkg/errors"

type AppInfo struct {
	ID          string
	Description string
	Developer   string
	Title       string

	Icon        string
	Screenshots []string
}

var ErrNotFound = errors.New("app is not found")
var ErrUnknown = errors.New("unknown error")
