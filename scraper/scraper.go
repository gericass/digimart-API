package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"html"
	"strconv"
	"time"
	"fmt"
	"strings"
	"net/http"
)

type scraper interface {
	scanName(*goquery.Selection) error
	scanDescription(s *goquery.Selection) error
	scanPrice(s *goquery.Selection) error
	scanCondition(s *goquery.Selection) error
	scanStatus(s *goquery.Selection) error
	scanURL(s *goquery.Selection) error
	scanImage(s *goquery.Selection) error
	scanRegisterDate(s *goquery.Selection) error
	toInstrument() *Instrument
}

func (i *NewArrivalInstrument) scanName(s *goquery.Selection) error {
	s.Find("p.ttl").Children().First().Each(func(_ int, s *goquery.Selection) {
		i.Name = html.UnescapeString(s.Text())
	})
	return nil
}

func (i *NewArrivalInstrument) scanDescription(s *goquery.Selection) error {

	return nil
}

func (i *NewArrivalInstrument) scanPrice(s *goquery.Selection) error {
	s.Find("div.itemState").Each(func(_ int, s *goquery.Selection) {
		s.Find("p.price").Find("span").Remove()
		p, _ := strconv.Atoi(s.Find("p.price").Text()[2:])
		i.Price = p

	})
	return nil
}

func (i *NewArrivalInstrument) scanCondition(s *goquery.Selection) error {
	s.Find("div.itemState").Each(func(_ int, s *goquery.Selection) {
		i.Condition = s.Find("p.state").Children().First().Text()
	})
	return nil
}

func (i *NewArrivalInstrument) scanStatus(s *goquery.Selection) error {
	i.Status = sell
	return nil
}

func (i *NewArrivalInstrument) scanURL(s *goquery.Selection) error {
	s.Find("p.ttl").Children().First().Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		i.URL = "https://www.digimart.net" + url
	})

	return nil
}

func (i *NewArrivalInstrument) scanImage(s *goquery.Selection) error {
	img, _ := s.Find("div.pic").Children().Children().First().Attr("src")
	i.Image = "https:" + img
	return nil
}

func (i *NewArrivalInstrument) scanRegisterDate(s *goquery.Selection) error {
	i.RegisterDate = time.Now()
	return nil
}

func (i *NewArrivalInstrument) toInstrument() *Instrument {
	inst := &Instrument{}
	inst.RegisterDate = i.RegisterDate
	inst.URL = i.URL
	inst.Status = i.Status
	inst.Condition = i.Condition
	inst.Price = i.Price
	inst.Category = i.Category
	inst.Description = i.Description
	inst.Image = i.Image
	inst.Name = i.Name
	return inst
}

func NewArrival() ([]*Instrument, error) {
	req, err := http.NewRequest("GET", "https://www.digimart.net", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:43.0) Gecko/20100101 Firefox/43.0 ")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	var insts = make([]*Instrument, 0)
	doc.Find("div.NewProductBlock").Each(func(_ int, s *goquery.Selection) {
		s.Find("li.ProductBox").Each(func(_ int, s *goquery.Selection) {
			insts = append(insts, scanInstrument(s,NewArrivalInstrument{}))
		})
	})
	return insts, nil
}

func (i *SearchInstrument) scanName(s *goquery.Selection) error {
	s.Find("p.ttl").Children().First().Each(func(_ int, s *goquery.Selection) {
		i.Name = html.UnescapeString(s.Text())

	})
	return nil
}

func (i *SearchInstrument) scanDescription(s *goquery.Selection) error {
	i.Description = s.Find("p.ttl").Next().Text()
	return nil
}

func (i *SearchInstrument) scanPrice(s *goquery.Selection) error {
	s.Find("div.itemState").Each(func(_ int, s *goquery.Selection) {
		s.Find("p.price").First().Children().Remove()
		priceString := strings.Replace(s.Find("p.price").First().Text(), "¥", "", -1)
		price, _ := strconv.Atoi(priceString)
		i.Price = price
	})
	return nil
}

func (i *SearchInstrument) scanCondition(s *goquery.Selection) error {
	s.Find("div.itemState").Each(func(_ int, s *goquery.Selection) {
		i.Condition = s.Find("p.state").Children().First().Text()
	})
	return nil
}

func (i *Instrument) scanStatus(s *goquery.Selection) error {
	_, exist := s.Find("p.order").Children().First().Attr("alt")
	if exist {
		i.Status = sold
	} else {
		i.Status = sell
	}
	return nil
}

func (i *SearchInstrument) scanURL(s *goquery.Selection) error {
	s.Find("p.ttl").Children().First().Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		i.URL = "https://www.digimart.net" + url
	})
	return nil
}

func (i *Instrument) scanImage(s *goquery.Selection) error {
	imgUrl, _ := s.Find("div.pic").Children().First().Children().Attr("src")
	i.Image = "https:" + imgUrl
	return nil
}

func (i *SearchInstrument) scanRegisterDate(s *goquery.Selection) error {
	dateStr := strings.Replace(s.Find("ul.itemDateInfo").Children().Next().Text(), "登録：", "", -1)
	d, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return err
	}
	i.RegisterDate = d
	return nil
}

func (i *SearchInstrument) toInstrument() *Instrument {
	inst := &Instrument{}
	inst.RegisterDate = i.RegisterDate
	inst.URL = i.URL
	inst.Status = i.Status
	inst.Condition = i.Condition
	inst.Price = i.Price
	inst.Category = i.Category
	inst.Description = i.Description
	inst.Image = i.Image
	inst.Name = i.Name
	return inst
}

func scanInstrument(s *goquery.Selection, instType interface{}) *Instrument {
	var inst scraper
	switch instType.(type) {
	case SearchInstrument:
		inst = &SearchInstrument{}
	case NewArrivalInstrument:
		inst = &NewArrivalInstrument{}
	}
	inst.scanRegisterDate(s)
	inst.scanURL(s)
	inst.scanName(s)
	inst.scanDescription(s)
	inst.scanImage(s)
	inst.scanPrice(s)
	inst.scanCondition(s)
	inst.scanStatus(s)
	return inst.toInstrument()
}

func Scrape(keyword string, page int) ([]*Instrument, error) {
	keyword = strings.Replace(keyword, " ", "+", -1)
	url := "https://www.digimart.net/search?category12Id=359&keywordAnd=" + keyword + "&currentPage=" + fmt.Sprint(page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:43.0) Gecko/20100101 Firefox/43.0 ")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	var insts = make([]*Instrument, 0)
	doc.Find("div.itemSearchBlock").Each(func(_ int, s *goquery.Selection) {
		insts = append(insts, scanInstrument(s,SearchInstrument{}))
	})

	return insts, nil
}
