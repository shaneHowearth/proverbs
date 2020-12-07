// Package main -
package main

import (
	"log"
	"os"
	"strings"

	"github.com/shanehowearth/proverbs/publish"
	"github.com/shanehowearth/proverbs/publish/twitter"
	"github.com/shanehowearth/proverbs/storage"
	postgresstore "github.com/shanehowearth/proverbs/storage/postgres"
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
	// Connect to the data store
	pgHost := os.Getenv("POSTGRES_HOST")
	pgDB := os.Getenv("POSTGRES_DATABASE")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPW := os.Getenv("POSTGRES_PASSWORD")
	db, err := postgresstore.NewPGStore(pgHost, pgDB, pgUser, pgPW)
	if err != nil {
		log.Fatalf("Unable to create postgres connection with error %v", err)
	}

	// Retrieve content
	data, translation, explanation, err := storage.GetContent(db)
	if err != nil {
		log.Fatalf("Unable to retrieve content with error %v", err)
	}

	content := strings.Join([]string{data, translation, explanation}, "\n\n")
	// Publish content
	for idx := range platforms {
		err := platforms[idx].PublishContent(content)
		if err != nil {
			// not the end of the world so don't die, but someone somewhere
			// should have a look at it
			log.Printf("ERROR: Getting content for %v returned error %v", platforms[idx], err)
		}
	}

}
