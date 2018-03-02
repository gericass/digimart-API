package scraper

import "github.com/PuerkitoBio/goquery"

func NewArrival() ([]*Instrument, error) {
	doc, err := goquery.NewDocument("https://www.digimart.net")
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func Scrape(keyword string) ([]*Instrument, error) {
	return nil, nil
}
