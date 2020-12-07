// Package twitter -
package twitter

import (
	"fmt"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type twitterClient struct {
	*twitter.Client
}

// NewTwitterClient -
// nolint: golint
func NewTwitterClient(consumerKey, consumerSecret, accessToken, accessSecret string) (*twitterClient, error) {

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return nil, fmt.Errorf("missing consumerKey, consumerSecret, accessToken, or accessSecret cannot continue")
	}
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return &twitterClient{twitter.NewClient(httpClient)}, nil
}

func (t *twitterClient) PublishContent(content string) error {
	bottom := 0
	top := 280
	l := utf8.RuneCountInString(content)
	params := &twitter.StatusUpdateParams{}
	for bottom < l {
		snippet := []rune(content)[bottom:top]
		for {
			if snippet[top] != ' ' {
				top--
			} else {
				break
			}
		}
		tweet, resp, err := t.Statuses.Update(string(snippet[bottom:top]), params)
		if resp.StatusCode != http.StatusOK {
			log.Printf("http return status was %d, with %s", resp.StatusCode, resp.Status)
			log.Printf("accompanied with error: %v", err)
			err = fmt.Errorf("bad status code returned with error %w", err)
			return err
		}
		params.InReplyToStatusID = tweet.ID
		bottom = top
		top += 280
	}
	return nil

}
