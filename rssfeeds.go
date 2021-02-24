package main

import (
	"encoding/json"
	"github.com/gorilla/feeds"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type FeedsData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func generateRssFeeds() *feeds.Feed {

	//Initialize RSS Feeds
	feed := &feeds.Feed{
		Title:       "RSS Feeds Example",
		Link:        &feeds.Link{Href: "/rssFeedsLink"},
		Description: "This is RSS Feeds Example In Go language",
		Author:      &feeds.Author{Name: "author"},
		Created:     time.Now(),
	}

	//Get JSON data from file, here you can retrieve data from database table also.
	feedsData := getJSONFeedsDataFromFile()
	var feedItems []*feeds.Item

	// Add items to RSS feeds
	for i := 0; i < len(feedsData); i++ {

		item := feedsData[i]

		feedItems = append(feedItems,
			&feeds.Item{
				Id:          item.ID,
				Title:       item.Title,
				Link:        &feeds.Link{Href: "/rssFeedsLink"},
				Description: item.Description,
				Created:     time.Now(),
			})
	}

	//Append Items
	feed.Items = feedItems

	return feed
}

func getJSONFeedsDataFromFile() []FeedsData {

	//open JSON file
	jsonFile, err := os.Open("feeds.json")
	if err != nil {
		log.Fatal(err)
	}

	//defer the closing of jsonFile, we can parse it later on
	defer jsonFile.Close()

	//read opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	var data []FeedsData

	// Unmarshal byteArray into FeedsData
	err = json.Unmarshal(byteValue, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
