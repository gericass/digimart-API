package scraper

import (
	"testing"
	"fmt"
)

func TestNewArrival(t *testing.T) {
	res, err := NewArrival()
	if err != nil {
		t.Errorf("%v occured", err)
	}
	for _, v := range res {
		fmt.Printf("result: %v\n", v)
	}
}

func TestScrape(t *testing.T) {
	res, err := Scrape("T's guitars",1)
	if err != nil {
		t.Errorf("%v occured", err)
	}
	for _, v := range res {
		fmt.Printf("result: %v\n", v)
	}
}