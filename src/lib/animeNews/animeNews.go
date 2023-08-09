package animeNews

import (
	"botwebo2/lib/myRequests"
	"encoding/xml"
	"time"
)

const (
	ANIMENEWSNETWORK = "https://www.animenewsnetwork.com/news/atom.xml"
)

var (
	lastMangaNews   = ""
	lastAnimeNews   = ""
	lastTimeChecked = time.Time{}
)

type Feed struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Link       Link       `xml:"link"`
	Published  time.Time  `xml:"published"`
	Categories []Category `xml:"category"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

type Link struct {
	Href string `xml:"href,attr"`
}

func CheckNews() ([]string, []string, error) {

	var animeNews, mangaNews []string

	//get request
	response, err := myRequests.GetByteResponse(ANIMENEWSNETWORK)

	if err != nil {
		return nil, nil, err
	}

	//unmarshal to struct
	var news Feed
	err = xml.Unmarshal(response, &news)

	if err != nil {
		return nil, nil, err
	}

	//check if is first time after bot boot
	if lastTimeChecked.IsZero() {
		lastTimeChecked = news.Entries[0].Published
		return nil, nil, nil
	}

	//parse news
	for _, new := range news.Entries {
		if new.Published.After(lastTimeChecked) {
			for _, category := range new.Categories {
				switch category.Term {
				case "Anime", "Anime.", "Animation":
					animeNews = append(animeNews, new.Link.Href)
				case "Manga":
					mangaNews = append(mangaNews, new.Link.Href)
				}
			}
		}
	}

	//update last new date
	lastTimeChecked = news.Entries[0].Published

	return animeNews, mangaNews, nil
}
