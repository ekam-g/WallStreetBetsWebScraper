package main

import (
	"WallStreetBetsWebScraper/Reddit"
	"fmt"
)

func main() {
	r, err := Reddit.Scrape{}.Init()
	if err != nil {
		fmt.Println(r)
		panic(err)
	}
}
