package main

import (
	"fmt"
	"github.com/gorilla/feeds"
)

func main() {

	//generate RSS feeds data from JSON file
	feedsData := generateRssFeeds()

	/*
	   convert it to RSS.
	   github.com/gorilla/feeds also provides methods for converting to RSS, JSON and Atom
	*/

	rssFeed := (&feeds.Rss{Feed: feedsData}).RssFeed()

	// Printing response to XML
	xmlRssFeeds := rssFeed.FeedXml()

	fmt.Println(xmlRssFeeds)
}
