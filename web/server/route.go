package server

import (
	"net/http"

	"github.com/jialu98/twitter-app/pkg/twitter"
	"github.com/jialu98/twitter-app/web/server/handler"
)

// Router setts up url to hanlder route
func Router(tweetGetter twitter.TweetGetter) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/static/", handler.Static())
	router.Handle("/tweets/", handler.Tweets(tweetGetter))
	return router
}
