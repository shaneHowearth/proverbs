// Package main -
package main

import (
	"log"
	"os"

	"github.com/shanehowearth/proverbs/publish"
	"github.com/shanehowearth/proverbs/publish/twitter"
)

func main() {
	// Collect twitter credentials
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPISecret := os.Getenv("TWITTER_API_SECRET")
	twitterAccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	// Create the twitter client
	tc, err := twitter.NewTwitterClient(twitterAPIKey, twitterAPISecret, twitterAccessToken, twitterAccessTokenSecret)
	if err != nil {
		log.Fatalf("Unable to create twitter client with error %v", err)
	}
	platforms := []publish.Publisher{}
	platforms = append(platforms, tc)
	// tc.PublishContent("Tēnā koe e hoa")
	// Connect to the data store

}
