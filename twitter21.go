package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	fmt.Println("...")

	// tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	// 	Count: 20,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.StatusCode)
	// fmt.Println("................................")
	// for i, s := range tweets {
	// 	fmt.Println(i, s.Text)
	// }
	//fmt.Println(tweets)

	fmt.Println("...-----------------------------------------...")

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "Caiado",
		Count: 5,
	})
	if err != nil {
		log.Fatal(err, resp)
	}
	for i, s := range search.Statuses {
		fmt.Println(i, s.Text)
	}
	fmt.Println("📉", search.Metadata.NextResults)

	time.Sleep(2 * time.Second)

	search, resp, err = client.Search.Tweets(&twitter.SearchTweetParams{
		Query:   "cat",
		Count:   5,
		SinceID: search.Metadata.MaxID,
	})

	if err != nil {
		log.Fatal(err, resp)
	}
	for i, s := range search.Statuses {
		fmt.Println(i, s.Text)
	}

}
