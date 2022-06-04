package scraper

import (
	"github.com/PuerkitoBio/goquery"
)

type Show struct {
	name               string
	indexerUrl         string
	episodeListUrl     string
	episodePageBaseUrl string
	episodeListElement string
	getEpisode         func(int, *goquery.Selection)
	newEpisodeReleased bool
	lastEpisode        int
	newEpisode         int
}

func NewShow(name,
	indexerUrl,
	episodeListUrl,
	episodePageBaseUrl,
	episodeListElement string,
	episodeScraper func(int, *goquery.Selection)) Show {

	return Show{
		name:               name,
		indexerUrl:         indexerUrl,
		episodeListUrl:     episodeListUrl,
		episodePageBaseUrl: episodePageBaseUrl,
		episodeListElement: episodeListElement,
		getEpisode:         episodeScraper,
	}
}
