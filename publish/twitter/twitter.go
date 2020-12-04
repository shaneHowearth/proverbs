// Package twitter -
package twitter

import (
	"fmt"
	"log"
	"net/http"

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

func (t *twitterClient) PublishContent(s string) error {
	_, resp, err := t.Statuses.Update(s, nil)
	if resp.StatusCode != http.StatusOK {
		log.Printf("http return status was %d, with %s", resp.StatusCode, resp.Status)
		log.Printf("accompanied with error: %v", err)
		err = fmt.Errorf("bad status code returned with error %w", err)
	}
	return err

}
