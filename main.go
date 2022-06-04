package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/whouishere/kissasian.go/scraper"
)

var Version string = "DEV"

func main() {
	fmt.Printf("Version %s\n\n", Version)

	testShow := scraper.NewShow("Soredemo Ai wo Chikaimasu ka?",
		"https://kissasiandb.com",
		"https://kissasiandb.com/tvshows/soredemo-ai-wo-chikaimasu-ka-2021/",
		"https://kissasiandb.com/soredemo-ai-wo-chikaimasu-ka-2021-episode-",
		"#all-episodes .list li",
		func(i int, s *goquery.Selection) {
			episode := s.Find("h3").Text()
			episodeUrl, exists := s.Find("h3 a").Attr("href")

			if exists {
				fmt.Printf("%s -> %s\n\n", episode, episodeUrl)
			} else {
				fmt.Printf("%s -> URL not found.\n\n", episode)
			}

			episodeInt, err := strconv.Atoi(string([]rune(episodeUrl)[len(episodeUrl)-2]))
			if err != nil {
				log.Fatal(err)
			}

			if episodeInt > scraper.LastEpisode {
				scraper.NewEpisodeReleased = true
				scraper.NewEpisode = episodeInt
			} else {
				scraper.NewEpisodeReleased = false
			}
		})

	testShow.ConnectToEpisodeList()

	fmt.Print("\nPress any key to exit.")
	bufio.NewScanner(os.Stdin).Scan()
}
