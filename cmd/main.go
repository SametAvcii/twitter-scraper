package main

import (
	"context"
	"fmt"

	"github.com/SametAvcii/twitter-scraper/scrapper"
)

func main() {
	scrapper.InitScrapper("Kriptos1752931", "Uj3tLAfowEtQM8m")
	scp := scrapper.Client()
	var err error
	count := 0
	for tweet := range scp.GetTweets(context.Background(), "Samet_Avcii", 39) {
		fmt.Println("")
		fmt.Println("count:", count)
		count++
		if tweet.Error != nil {
			fmt.Println("err:", tweet.Error.Error())
			continue
		}

	}
	if count != 39 {
		ids, err := scp.FetchTweetIDS("Samet_Avcii", 39, "")
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		fmt.Println("len ids:", len(ids))
	}

	//res, err := scp.GetTweet("1116785280797880320")

	if err != nil {
		fmt.Println("err:", err.Error())
	}
	//fmt.Println("res:", &res)
}
