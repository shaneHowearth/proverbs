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

func (t *twitterClient) chunkContent(content string) []string {
	// break the content up into chunks that are 140 or less characters
	// if the 140 character isn't whitespace, then look at the 139th, and so on
	// until a space is found. The next chunk then starts at that position.
	chunks := []string{}
	bottom := 0
	top := 140
	l := utf8.RuneCountInString(content)
	for bottom < l {
		if top >= l {
			top = l
			chunks = append(chunks, content[bottom:top])
			break
		}
		chunk := content[bottom:top]
		for i := utf8.RuneCountInString(chunk) - 1; i > 0; i-- {
			r := []rune(chunk)[i]
			if r == ' ' || r == '\n' || r == '\t' {
				top = bottom + i
				break
			}
		}
		chunks = append(chunks, content[bottom:top])
		bottom = top + 1
		top += 140
	}

	return chunks
}

// PublishContent -
func (t *twitterClient) PublishContent(content string) error {
	params := &twitter.StatusUpdateParams{}
	for _, snippet := range t.chunkContent(content) {
		tweet, resp, err := t.Statuses.Update(string(snippet), params)
		if resp.StatusCode != http.StatusOK {
			log.Printf("http return status was %d, with %s", resp.StatusCode, resp.Status)
			log.Printf("accompanied with error: %v", err)
			err = fmt.Errorf("bad status code returned with error %w", err)
			return err
		}
		params.InReplyToStatusID = tweet.ID
	}
	return nil

}
