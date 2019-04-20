package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/jialu98/twitter-app/pkg/twitter"
	"net/http"
	"net/http/httptest"
	"testing"
)

type tweetGetterMock struct {
	tweetsMap  map[string][]twitter.Tweet
	getInvoked bool
}

func (mock *tweetGetterMock) Get(user string) ([]twitter.Tweet, error) {
	mock.getInvoked = true
	tweets, ok := mock.tweetsMap[user]
	if !ok {
		return []twitter.Tweet{}, fmt.Errorf("no tweets for user %v", user)
	}
	return tweets, nil
}
func TestTweets(t *testing.T) {
	user := "user1"
	tweets := []twitter.Tweet{
		twitter.Tweet{Text: "text1", Timestamp: "ts1"},
		twitter.Tweet{Text: "text2", Timestamp: "ts2"},
	}
	tweetsMap := map[string][]twitter.Tweet{"user1": tweets}
	mock := &tweetGetterMock{tweetsMap: tweetsMap}

	handler := Tweets(mock)
	// Invoke the handler.
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "?user="+user, nil)
	handler.ServeHTTP(w, r)

	// Validate mock.
	if !mock.getInvoked {
		t.Fatal("expected User() to be invoked")
	}

	// Check http status
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we want.
	want := TweetPageData{
		User:   user,
		Tweets: tweets,
	}
	got := TweetPageData{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if !cmp.Equal(got, want) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, want)
	}
}
