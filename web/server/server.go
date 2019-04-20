package server

import (
	"fmt"
	"github.com/jialu98/twitter-app/pkg/twitter"
	"net/http"
)

// Start starts web server at given listen address
func Start(listenAddr string, tweetGetter twitter.TweetGetter) {
	fmt.Printf("Start server at %v", listenAddr)
	server := &http.Server{
		Addr:    listenAddr,
		Handler: Router(tweetGetter),
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen on %v: %v\n", listenAddr, err)
	}
}
