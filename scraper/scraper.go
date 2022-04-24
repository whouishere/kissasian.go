package scraper

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const episodeListUrl = "https://kissasiandb.com/tvshows/soredemo-ai-wo-chikaimasu-ka-2021/"
const episodePageBaseUrl = "https://kissasiandb.com/soredemo-ai-wo-chikaimasu-ka-2021-episode-"
const showName = "Soredemo Ai wo Chikaimasu ka?"
const listElement = "#all-episodes .list li"

var newEpisodeReleased bool
var lastEpisode int
var newEpisode int

func ConnectToEpisodeList() {
	res, err := http.Get(episodeListUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("%d error: %s", res.StatusCode, res.Status)
	}
	fmt.Println("Connected to kissasiandb.com")

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	lastEpisode = GetWatchedEpisode()

	fmt.Printf("Fetching '%s' episode list...\n", showName)
	doc.Find(listElement).Each(checkForEpisodeList)

	// this doesn't consider more than one new episodes. might need to fix that later
	if newEpisodeReleased {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("New episode released (", newEpisode, ")! Do you want to mark it as watched? (Y/N) ")
		input.Scan()
		answer := input.Text()

		if strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes") {
			UpdateEpisode(newEpisode)
			lastEpisode = newEpisode
		} else if strings.EqualFold(answer, "n") || strings.EqualFold(answer, "no") {
			return
		}
	}
}

func checkForEpisodeList(i int, s *goquery.Selection) {
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

	if episodeInt > lastEpisode {
		newEpisodeReleased = true
		newEpisode = episodeInt
	} else {
		newEpisodeReleased = false
	}
}

func CheckForEpisodePage(episode int) {
	fullUrl := episodePageBaseUrl + fmt.Sprint(episode) + "/"
	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Println("Episode", episode, "was released.")
	} else {
		fmt.Println("Episode", episode, "not yet released.")
	}
}
