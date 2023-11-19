package scrapper

import (
	"encoding/json"
	"fmt"
	"os"

	twitterscraper "github.com/SametAvcii/twitter-scraper"
)

var scrapper *twitterscraper.Scraper

func InitScrapper(username, pass string) {
	scraper := twitterscraper.New()
	cookies := scraper.GetCookies()

	f, _ := os.Open("cookies.json")
	json.NewDecoder(f).Decode(&cookies)
	scraper.SetCookies(cookies)

	isLogin := scraper.IsLoggedIn()
	fmt.Println("is login:", isLogin)
	if !isLogin {
		fmt.Println("Cookied success")
		err := scraper.Login(username, pass)
		if err != nil {
			fmt.Println("login is failed", err.Error())
		}
		cookies = scraper.GetCookies()
		js, _ := json.Marshal(cookies)
		// save to file
		f, _ = os.Create("cookies.json")
		f.Write(js)
	}
	scrapper = scraper
}

func Client() *twitterscraper.Scraper {
	return scrapper
}
