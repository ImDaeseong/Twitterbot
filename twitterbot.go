// twitterbot
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var api_key = "pnDHE9x8v0UIUUann6IMQ5EHP"
var api_secret_key = "n2o1gLnW9k13GS8bXjwPWtBxw68R4SBP8LGD3qk8EApbeyW3Ry"
var api_access_token = "914299970974490624-dOPMMTDeq9IzbeSydEtpUNWnba438ar"
var api__secret_access_token = "g52j8Mdr5oRq0E8t1KvSyIuHLBkl5R65Q3qklB1cMHHtu"

func getClient() (*twitter.Client, error) {
	config := oauth1.NewConfig(api_key, api_secret_key)
	token := oauth1.NewToken(api_access_token, api__secret_access_token)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	_, _, err := client.Accounts.VerifyCredentials(nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getUser(client *twitter.Client) (*twitter.User, error) {

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	//user, _, err := client.Accounts.VerifyCredentials(nil)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func followerList(client *twitter.Client) []string {

	var followerList []string

	followers, _, err := client.Followers.IDs(&twitter.FollowerIDParams{})
	if err != nil {
		return nil
	}

	for _, follower := range followers.IDs {
		followerList = append(followerList, strconv.FormatInt(follower, 10))
	}
	return followerList
}

func friendList(client *twitter.Client) []string {

	var friendList []string

	friends, _, err := client.Friends.IDs(&twitter.FriendIDParams{})
	if err != nil {
		return nil
	}

	for _, friend := range friends.IDs {
		friendList = append(friendList, strconv.FormatInt(friend, 10))
	}
	return friendList
}

func hometimeline(client *twitter.Client) []twitter.Tweet {

	tweets, _, _ := client.Timelines.HomeTimeline(nil)
	return (tweets)
}

func sendLink(client *twitter.Client, sText string, sLink string) {

	sMsg := fmt.Sprintf("%s\n%s", sText, sLink)
	_, _, err := client.Statuses.Update(sMsg, nil)
	if err != nil {
		log.Print(err)
	}
}

func SendMessage(client *twitter.Client, sMsg string) {

	_, _, err := client.Statuses.Update(sMsg, nil)
	if err != nil {
		log.Print(err)
	}
}

func SendMessageID(client *twitter.Client, sMsg string, tweetid int64) {

	replyTo := new(twitter.StatusUpdateParams)
	replyTo.InReplyToStatusID = tweetid

	_, _, err := client.Statuses.Update(sMsg, replyTo)
	if err != nil {
		log.Print(err)
	}
}

func searchKeyword(client *twitter.Client, skey string) ([]twitter.Tweet, error) {

	searchParams := &twitter.SearchTweetParams{
		Query:      skey,
		Count:      1000,
		ResultType: "recent",
	}
	search, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		log.Print(err)
	}

	return search.Statuses, nil
}

func searchretweet(client *twitter.Client, skey string) {

	searchParams := &twitter.SearchTweetParams{
		Query:      skey,
		Count:      1000,
		ResultType: "recent",
	}
	search, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		log.Print(err)
	}

	for _, tweet := range search.Statuses {
		//fmt.Printf("%v\n", tweet.ID)
		//fmt.Printf("%s\n", tweet.User.ScreenName)
		//fmt.Printf("%v\n", tweet.Text)
		//fmt.Printf("%v\n", tweet.FullText)

		sMsg := fmt.Sprintf("[%v] - %s\n", tweet.ID, tweet.Text)
		fmt.Printf(sMsg)

		client.Statuses.Retweet(tweet.ID, &twitter.StatusRetweetParams{})
		time.Sleep(100 * time.Millisecond)
	}
}

func checkfavorite(client *twitter.Client, skey string) {

	searchParams := &twitter.SearchTweetParams{
		Query:      skey,
		Count:      10,
		ResultType: "recent",
	}
	search, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		log.Print(err)
	}

	for _, tweet := range search.Statuses {

		//sMsg := fmt.Sprintf("[%s] - %s\n", tweet.User.ScreenName, tweet.Text)
		//fmt.Printf(sMsg)

		//sMsg := fmt.Sprintf("%v - %v\n", tweet.Entities.Media, tweet.Entities.Urls)
		//fmt.Printf(sMsg)

		if tweet.Favorited == false {
		}

		if tweet.User.Following == false {
		}
	}
}

func main() {

	client, err := getClient()
	if err != nil {
		log.Print(err)
	}

	user, err := getUser(client)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("user.FriendsCount:%d\n", user.FriendsCount)
	fmt.Printf("user.Description:%s\n", user.Description)
	fmt.Printf("user.Location:%s\n", user.Location)
	fmt.Printf("user.ID:%v\n", user.ID)
	fmt.Printf("user.IDStr:%s\n", user.IDStr)
	fmt.Printf("user.ScreenName:%s\n", user.ScreenName)
	fmt.Printf("user.Name:%s\n", user.Name)
	fmt.Printf("user.FollowersCount:%d\n", user.FollowersCount)
	fmt.Printf("user.StatusesCount:%d\n", user.StatusesCount)
	fmt.Printf("user.Timezone:%s\n", user.Timezone)

	/*
		tweets := hometimeline(client)
		for _, tweet := range tweets {
			sMsg := fmt.Sprintf("[%d / %s] - %s\n", tweet.ID, tweet.User.ScreenName, tweet.Text)
			fmt.Printf(sMsg)
		}
	*/

	/*
		data := followerList(client)
		for _, item := range data {
			fmt.Println(item)
		}

		data := friendList(client)
		for _, item := range data {
			fmt.Println(item)
		}
	*/

	/*
		tweets, _ := searchKeyword(client, "최신 영화")
		for _, tweet := range tweets {
			sMsg := fmt.Sprintf("%d/%s - %s\n", tweet.ID, tweet.User.ScreenName, tweet.Text)
			fmt.Printf(sMsg)
		}
	*/

	//searchretweet(client, "#python")
	//checkfavorite(client, "#python")

	//SendMessage(client, "공부중")
	//SendMessageID(client, "공부중", user.ID)
	//sendLink(client, "나의 이미지", "https://avatars0.githubusercontent.com/u/10001221?s=460&v=4")
}
