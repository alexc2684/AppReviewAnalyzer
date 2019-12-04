package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	// "projects/app_review_analyzer/models"
)

// GetReviewsByID - http get request to get all reviews
func GetReviewsByID(id int) { //[]models.Review {
	url := createURLFromID(id)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func createURLFromID(id int) string {
	//TODO: pages
	url := "https://itunes.apple.com/rss/customerreviews/id=" + strconv.Itoa(id) + "/page=2/sortby=mostrecent/json"
	return url
}
