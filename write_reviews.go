package main

import (
	"flag"
	"fmt"
	"projects/app_review_analyzer/scraper"
)

func main() {
	id := getIdFromFlags()
	scraper.GetReviewsByID(id)
}

func getIdFromFlags() int {
	idPtr := flag.Int("id", 0, "integer for appId")
	flag.Parse()
	return *idPtr
}
