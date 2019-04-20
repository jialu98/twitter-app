package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/pkg/errors"
)

// Tweet represents a tweet text and timestamp
type Tweet struct {
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

// TweetGetter is a interface providing Get method
type TweetGetter interface {
	Get(user string) ([]Tweet, error)
}

// TweetService provide tweet service
type TweetService struct {
	tc    *twitter.Client
	count int
}

// NewTweetService returns a new Client.
func NewTweetService(consumerKey, consumerSecret, accessToken, accessSecret string, count int) *TweetService {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	tc := twitter.NewClient(httpClient)
	return &TweetService{tc, count}
}

// Get returns the tweets by given user handle and count
func (service *TweetService) Get(user string) ([]Tweet, error) {
	var result []Tweet
	userTimelineParams := &twitter.UserTimelineParams{
		ScreenName: user,
		Count:      service.count,
	}
	tweets, resp, err := service.tc.Timelines.UserTimeline(userTimelineParams)
	if err != nil {
		return result, errors.Wrapf(err, "response StatusCode is %v", resp.StatusCode)
	}
	if resp.StatusCode > 300 {
		return result, errors.Errorf("response StatusCode is %v", resp.StatusCode)
	}
	for _, tweet := range tweets {
		result = append(result, Tweet{Text: tweet.Text, Timestamp: tweet.CreatedAt})
	}
	return result, nil
}
