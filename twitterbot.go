// twitterbot
package main

import (
	"fmt"
	"log"
	"strconv"

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

	//user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	//fmt.Printf("ScreenName: %v \n", user.ScreenName)

	_, _, err := client.Accounts.VerifyCredentials(nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getUser(client *twitter.Client) (*twitter.User, error) {
	user, _, err := client.Accounts.VerifyCredentials(nil)
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

func SendMessage(client *twitter.Client, sMsg string) {

	_, _, err := client.Statuses.Update(sMsg, nil)
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

func main() {

	client, err := getClient()
	if err != nil {
		log.Print(err)
	}

	user, err := getUser(client)
	if err != nil {
		log.Print(err)
	}

	fmt.Printf("%s\n", user.ScreenName)
	fmt.Printf("%s\n", user.IDStr)
	fmt.Printf("%v\n", user.ID)
	fmt.Printf("%s\n", user.Name)
	fmt.Printf("%s\n", user.Location)
	fmt.Printf("%d\n", user.FriendsCount)

	//SendMessage(client, "test")

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

	tweets, _ := searchKeyword(client, "최신 영화")
	for _, tweet := range tweets {
		sMsg := fmt.Sprintf("%d/%s - %s\n", tweet.ID, tweet.User.ScreenName, tweet.Text)
		fmt.Printf(sMsg)
	}

}
