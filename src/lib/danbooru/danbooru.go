package danbooru

import (
	"botwebo2/lib/embed"
	"botwebo2/lib/myRequests"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type Tag struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	PostCount    int       `json:"post_count"`
	Category     int16     `json:"category"`
	Created      time.Time `json:"created_at"`
	Updated      time.Time `json:"updated_at"`
	IsDeprecated bool      `json:"is_deprecated"`
	Words        []string  `json:"words"`
}

func SendDanbooruImage(tag string) (embed.Embed, error) {
	message := embed.NewEmbed()

	response, err := myRequests.GetByteResponse("https://danbooru.donmai.us/posts/random.json?tags=" + tag)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error getting request"), err
	}

	sec := map[string]interface{}{}

	err = json.Unmarshal(response, &sec)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error json decode"), err
	}

	image := ""

	if sec["file_url"] != nil {
		image = fmt.Sprintf("%v", sec["file_url"])
	}

	return *message.SetColor(embed.SuccessColor).SetImage(image), err

}

func GetSimilarTags(tag string) (embed.Embed, error) {
	message := embed.NewEmbed()

	response1, err := myRequests.GetByteResponse("https://danbooru.donmai.us/tags.json" +
		"?limit=5" +
		"&search[hide_empty]=true" +
		"&search[order]=count" +
		"&search[fuzzy_name_matches]=" + tag)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error getting request"), err
	}

	var tags []Tag

	var temp []Tag

	err = json.Unmarshal(response1, &temp)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error json decode"), err
	}

	for _, tag := range temp {
		tags = append(tags, tag)
	}

	response2, err := myRequests.GetByteResponse("https://danbooru.donmai.us/tags.json" +
		"?limit=5" +
		"&search[hide_empty]=true" +
		"&search[order]=count" +
		"&search[name_or_alias_matches]=*" + tag + "*")

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error getting request"), err
	}

	err = json.Unmarshal(response2, &temp)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error json decode"), err
	}

	for _, tag := range temp {
		tags = append(tags, tag)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].PostCount > tags[j].PostCount
	})

	message.SetTitle("Maybe you mean: ").SetColor(embed.SuccessColor)

	for _, tag := range tags {
		message.AddField("", tag.Name)
	}

	return *message, nil
}

func GetRandomImage() {

}
