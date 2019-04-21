# twitter-app

1. Project setup
   - checkout project
   ```go get github.com/jialu98/twitter-app```
   - run dep ensure (make sure dep is installed)
   ```cd $GOPATH/src/github.com/jialu98/twitter-app
   $GOPATH/src/github.com/jialu98/twitter-app> dep ensure
   ```
   - build project
   ```
   cd $GOPATH/src/github.com/jialu98/twitter-app/cmd/app
   $GOPATH/src/github.com/jialu98/twitter-app/cmd/app> go build
   ```
   The executable app.exe(on windows) is created.
   - run app
   Set up follow environment variable:
      ```
      CONSUMER_KEY
      CONSUMER_SECRET
      ACCESS_TOKEN
      ACCESS_SECRET
      ```
      For example:
      ```
        set CONSUMER_KEY=abcd
        set CONSUMER_SECRET=abcd
        set ACCESS_TOKEN=abcd
        set ACCESS_SECRET=abcd
      ```  
      Note: The values of these environment variables are provided by your twitter developer account.
    Run app.exe (on windows), and the app runs on localhost and port 8080 by default.
       ```$GOPATH/src/github.com/jialu98/twitter-app/cmd/app> app.exe```

    - test
    In browser, go to http://localhost:8080/static/
    Test with user 'BillGates', and click on the "get tweets" button, the latest 10 tweets will be displayed.

2. Design
  - Project layout.
  ```
     twitter-app
              |--cmd
              |--pkg
              |--web
   ```
     The cmd folder contains the main function. The pkg folder is for shared library. The service to get the tweets is in this folder. The web folder contains web related code such as web server, handler, route etc.
  - The TweetService.
  
      The logic to get tweets from twitter for a given user is implemented in twitter-app/pkg/twitter/twitter.go. The twitter library github.com/dghubble/go-twitter is used. The TweetService implements a single method interface TweetGetter. This approach makes it easy to change the implementation, for example, using a different twitter library, without impact on the consumer of the service.
  - The tweet web handler.

    The logic to serve the request to get the tweets for given user is implemented in this handler. The TweetService is injected into this handler by using the single method interface TweetGetter, which makes it easy to mock the TweetService in test. This handler serves the json message instead of html using server site template, which separates the backend from the front end. We can choose appropriate front end technology to implement single page app, and offers better performance during the high load.
    
    
