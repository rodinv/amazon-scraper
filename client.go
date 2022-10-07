package amazon_scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c Client) GetAppInfo(asin string) (*AppInfo, error) {
	app := &AppInfo{
		ID: asin,
	}

	scraper := colly.NewCollector()

	scraper.OnHTML("#btAsinTitle", func(e *colly.HTMLElement) {
		app.Title = formatRawText(e.Text)
	})

	scraper.OnHTML("#mas-product-description", func(e *colly.HTMLElement) {
		app.Description = e.Text
		app.Description = strings.Replace(app.Description, "Product description", "", -1)
		app.Description = formatRawText(app.Description)
	})

	scraper.OnHTML("#brand", func(e *colly.HTMLElement) {
		app.Developer = formatRawText(e.Text)
	})

	scraper.OnHTML("#js-masrw-main-image", func(e *colly.HTMLElement) {
		app.Icon = e.Attr("src")
	})

	scraper.OnHTML(".masrw-screenshot", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if src != "" {
			app.Screenshots = append(app.Screenshots, src)
		}
	})

	scraper.OnHTML("[data-hook=\"rating-out-of-text\"]", func(e *colly.HTMLElement) {
		app.Rating = e.Text
	})

	err := scraper.Visit(fmt.Sprintf("http://www.amazon.com/dp/%s", asin))
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Not Found"):
			return nil, ErrNotFound
		default:
			return nil, ErrUnknown
		}
	}

	return app, nil
}

func formatRawText(raw string) string {
	return strings.TrimSpace(strings.Replace(raw, "\n", "", -1))
}
