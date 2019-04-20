package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jialu98/twitter-app/pkg/twitter"
)

// TweetPageData is the model for the tweet page
type TweetPageData struct {
	User   string          `json:"user"`
	Tweets []twitter.Tweet `json:"tweets"`
}

// Tweets creates the handler to get tweets for the given user
func Tweets(tweetGetter twitter.TweetGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, ok := r.URL.Query()["user"]
		if !ok || len(users[0]) < 1 {
			fmt.Println("user is missing")
			return
		}
		user := users[0]
		tweets, err := tweetGetter.Get(user)
		if err != nil {
			fmt.Printf("fail to get tweet for %v, error: %v", user, err)
			return
		}

		data := TweetPageData{
			User:   user,
			Tweets: tweets,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})
}
