package main

import (
	"flag"
	"github.com/jialu98/twitter-app/pkg/twitter"
	"github.com/jialu98/twitter-app/web/server"
	"os"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Consumer key/secret and Access token/secret required")
	}
	listenAddr := flag.String("listen-addr", ":8080", "server listen address")
	count := flag.Int("number-of-tweet", 10, "number of tweets to display")
	flag.Parse()
	tweetService := twitter.NewTweetService(consumerKey, consumerSecret, accessToken, accessSecret, *count)
	server.Start(*listenAddr, tweetService)
}
