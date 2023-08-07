package sauces

import (
	"botwebo2/lib/embed"
	"botwebo2/lib/myRequests"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var (
	apiCall = "https://saucenao.com/search.php?output_type=2&api_key=%s&url=%s"
)

// struct for json response
type Sauces struct {
	Header  Header   `json:"header"`
	Results []Result `json:"results"`
}

type Header struct {
	Status int `json:"status"`
}

type Result struct {
	Header ResultHeader `json:"header"`
	Data   ResultData   `json:"data"`
}

type ResultHeader struct {
	Similarity string `json:"similarity"`
}

type ResultData struct {
	//common
	ExtUrls    []string `json:"ext_urls"`
	Title      string   `json:"title"`
	AuthorName string   `json:"author_name"`
	AuthorUrl  string   `json:"author_url"`

	//pixiv

	PixivId    int    `json:"pixiv_id"`
	MemberName string `json:"member_name"`
	MemberId   int    `json:"member_id"`

	//danbooru
	DanbooruId int    `json:"danbooru_id"`
	GelbooruId int    `json:"gelbooru_id"`
	Material   string `json:"material"`
	Characters string `json:"characters"`
	Source     string `json:"source"`

	//e621
	E621ID int `json:"e621_id"`

	//skeb
	Path        string `json:"path"`
	CreatorName string `json:"creator_name"`

	//deviantart
	DaId string `json:"da_id"`
}

func GetSauce(url string) (embed.Embed, error) {

	message := embed.NewEmbed()

	response, err := myRequests.GetByteResponse(fmt.Sprintf(apiCall, os.Getenv("SAUCENAO_KEY"), url))

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error request"), err
	}

	var sauces Sauces

	err = json.Unmarshal(response, &sauces)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error json decode"), err
	}

	if sauces.Header.Status == -3 {
		return *message.SetColor(embed.ErrorColor).SetDescription("that is not an image"), err
	} else if sauces.Header.Status == 0 {

		//filter and resume content
		for ind, result := range sauces.Results {

			s, err := strconv.ParseFloat(result.Header.Similarity, 64)

			if err != nil {
				return *message.SetColor(embed.ErrorColor).SetDescription("cannot convert similarity to float"), err
			}

			if s > 65 {
				msg := ""
				var temp []string

				temp = append(temp, result.Data.ExtUrls[0])

				result.Data.ExtUrls = temp

				var tempMap map[string]interface{}

				data, _ := json.Marshal(result.Data)

				json.Unmarshal(data, &tempMap)

				for key, content := range tempMap {

					v, ok := content.(float64)
					if v == 0 && ok {

					} else if ok {
						msg += key + ": " + fmt.Sprintf("%.0f", v) + "\n"
					} else if content != "" {
						msg += key + ": " + fmt.Sprintf("%v", content) + "\n"
					}

				}

				message.AddField(strconv.Itoa(ind+1)+")", msg)

			}
		}

	}
	message.SetColor(embed.SuccessColor)
	return *message, err
}
