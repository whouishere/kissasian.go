package scraper

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/whouishere/kissasiandb/status"
)

var NewEpisodeReleased bool
var LastEpisode int
var NewEpisode int

type ShowInterface interface {
	ConnectToEpisodeList()
	CheckForEpisodePage()
}

func (show Show) ConnectToEpisodeList() {
	res, err := http.Get(show.episodeListUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("%d error: %s", res.StatusCode, res.Status)
	}
	fmt.Println("Connected to ", show.indexerUrl)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	show.lastEpisode = status.GetWatchedEpisode()

	fmt.Printf("Fetching '%s' episode list...\n", show.name)
	doc.Find(show.episodeListElement).Each(show.getEpisode)

	// this doesn't consider more than one new episodes. might need to fix that later
	if show.newEpisodeReleased {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("New episode released (", show.newEpisode, ")! Do you want to mark it as watched? (Y/N) ")
		input.Scan()
		answer := input.Text()

		if strings.EqualFold(answer, "y") || strings.EqualFold(answer, "yes") {
			status.UpdateEpisode(show.newEpisode)
			show.lastEpisode = show.newEpisode
		} else if strings.EqualFold(answer, "n") || strings.EqualFold(answer, "no") {
			return
		}
	}
}

func (show Show) CheckForEpisodePage(episode int) {
	fullUrl := show.episodePageBaseUrl + fmt.Sprint(episode) + "/"
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
