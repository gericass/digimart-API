package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"html"
	"strconv"
	"time"
)

func scanInstrument(s *goquery.Selection) *Instrument {
	inst := &Instrument{}
	inst.Status = sell
	img, _ := s.Find("div.pic").Children().Children().First().Attr("src")
	inst.Image = "https:" + img
	s.Find("p.ttl").Children().First().Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		inst.URL = "https://www.digimart.net" + url
		inst.Name = html.UnescapeString(s.Text())
	})
	s.Find("div.itemState").Each(func(_ int, s *goquery.Selection) {
		s.Find("p.price").Find("span").Remove()
		i, _ := strconv.Atoi(s.Find("p.price").Text()[2:])
		inst.Price = i
		inst.Condition = s.Find("p.state").Children().First().Text()
	})
	inst.RegisterDate = time.Now()

	return inst
}

func NewArrival() ([]*Instrument, error) {
	doc, err := goquery.NewDocument("https://www.digimart.net")
	if err != nil {
		return nil, err
	}
	var insts = make([]*Instrument, 0)
	doc.Find("div.NewProductBlock").Each(func(_ int, s *goquery.Selection) {
		s.Find("li.ProductBox").Each(func(_ int, s *goquery.Selection) {
			insts = append(insts, scanInstrument(s))
		})
	})
	return insts, nil
}

func Scrape(keyword string) ([]*Instrument, error) {
	return nil, nil
}
