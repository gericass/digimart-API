package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"html"
	"strconv"
	"time"
	"fmt"
	"strings"
)

func scanNewInstrument(s *goquery.Selection) *Instrument {
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
			insts = append(insts, scanNewInstrument(s))
		})
	})
	return insts, nil
}

func scanInstrument(s *goquery.Selection) *Instrument {
	inst := &Instrument{}
	dateStr := strings.Replace(s.Find("ul.itemDateInfo").Children().Next().Text(), "登録：", "", -1)
	d, _ := time.Parse("2006-01-02 15:04:05", dateStr)
	inst.RegisterDate = d
	s.Find("p.ttl").Children().First().Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		inst.URL = "https://www.digimart.net" + url
		inst.Name = html.UnescapeString(s.Text())

	})
	inst.Description = s.Find("p.ttl").Next().Text()
	imgUrl, _ := s.Find("div.pic").Children().First().Children().Attr("src")
	inst.Image = "https:" + imgUrl
	s.Find("div.itemState").Each(func(_ int, s *goquery.Selection) {
		s.Find("p.price").First().Children().Remove()
		priceString := strings.Replace(s.Find("p.price").First().Text(), "¥", "", -1)
		price, _ := strconv.Atoi(priceString)
		inst.Price = price
		inst.Condition = s.Find("p.state").Children().First().Text()
	})
	_, exist := s.Find("p.order").Children().First().Attr("alt")
	if exist {
		inst.Status = sold
	} else {
		inst.Status = sell
	}

	return inst
}

func Scrape(keyword string, page int) ([]*Instrument, error) {
	keyword = strings.Replace(keyword, " ", "+", -1)
	url := "https://www.digimart.net/search?category12Id=359&keywordAnd=" + keyword + "&currentPage=" + fmt.Sprint(page)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	var insts = make([]*Instrument, 0)
	doc.Find("div.itemSearchBlock").Each(func(_ int, s *goquery.Selection) {
		insts = append(insts, scanInstrument(s))
	})
	return insts, nil
}
