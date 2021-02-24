package main

import (
	"fmt"
	"log"
)

func main() {

	//generate RSS feeds data from JSON file
	feedsData := generateRssFeeds()

	/*
	   convert it to RSS.
	   github.com/gorilla/feeds also provides methods for converting to JSON and Atom
	*/

	rssFeeds, err := feedsData.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rssFeeds)
}
