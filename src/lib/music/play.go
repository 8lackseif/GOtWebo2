package music

import (
	"botwebo2/lib/embed"
	"botwebo2/lib/myRequests"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const (
	youtubeListEndpoint     = "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&key=%s&playlistId=%s&maxResults=30"
	youtubeVideoEndpoint    = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key=%s&id=%s"
	youtubeNextPageEndpoint = "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&key=%s&playlistId=%s&maxResults=30&pageToken=%s"
)

//play

func (p *Player) addSongs(url string, c string, dg *discordgo.Session) (embed.Embed, error) {
	message := embed.NewEmbed()
	var err error
	p.voiceClient, err = dg.ChannelVoiceJoin(p.guildID, c, false, true)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("ERR: Error when joining voice channel"), err
	}

	err = p.voiceClient.Speaking(false)
	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("ERR: Error when joining voice channel"), err
	}

	//regex to detect if its a list or a video
	ok, err := regexp.MatchString("(youtube\\.com|youtu\\.be)(\\/playlist\\?list=)([a-zA-Z0-9\\-\\_]+)", url)

	if ok {
		return p.addPlayList(getYoutubeID(url))

	} else {
		ok, err = regexp.MatchString("(youtu\\.be\\/|v\\/|u\\/\\w\\/|embed\\/|watch\\?v=|\\&v=)([^#\\&\\?]*)", url)

		if ok {
			return p.addVideo(getYoutubeID(url))
		}
	}

	return *message, nil
}

func (p *Player) list() (embed.Embed, error) {
	//lock list and unlocked when function finished for read
	p.queueMutex.RLock()
	defer p.queueMutex.RUnlock()

	message := embed.NewEmbed()

	message.AddField("Now playing:", p.currentSong.Title)

	for i, video := range p.playlist {
		message.AddField("", fmt.Sprintf("%d", i+1)+") "+video.Title)
	}

	return *message, nil
}

func (p *Player) loop() (embed.Embed, error) {
	message := embed.NewEmbed()

	//make looping the opposite
	p.looping = !p.looping

	message.SetColor(embed.SuccessColor).SetDescription("looping: " + strconv.FormatBool(p.looping))

	return *message, nil
}

func (p *Player) emptyList() (embed.Embed, error) {
	//lock list and unlocked when function finished
	p.queueMutex.Lock()
	defer p.queueMutex.Unlock()

	message := embed.NewEmbed()

	//set to zero value
	p.playlist = []Video{}

	message.SetColor(embed.SuccessColor).SetDescription("Playlist cleared.")

	return *message, nil

}

func (p *Player) shuffle() (embed.Embed, error) {

	//lock list and unlocked when function finished
	p.queueMutex.Lock()
	defer p.queueMutex.Unlock()

	message := embed.NewEmbed()

	//shuffle the list
	rand.Shuffle(len(p.playlist), func(i, j int) { p.playlist[i], p.playlist[j] = p.playlist[j], p.playlist[i] })

	message.SetColor(embed.SuccessColor).SetDescription("list shuffled.")

	return *message, nil
}

func (p *Player) playSong(v Video) {
	//download audio
	cmd := exec.Command("yt-dlp", "-o", "serverAudio/"+p.guildID, "-x", "--audio-format", "opus", "--force-overwrites", "https://www.youtube.com/watch?v="+v.VidID)

	//get output
	out, err := cmd.Output()
	if err != nil {
		log.Println("error downloading song")
		return
	}

	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))

	//play song
	p.stopch = make(chan bool)
	dgvoice.PlayAudioFile(p.voiceClient, "serverAudio/"+p.guildID+".opus", p.stopch)
}

//youtube

func (p *Player) addPlayList(id string) (embed.Embed, error) {
	message := embed.NewEmbed()
	var res YRequest

	//request
	r, err := myRequests.GetByteResponse(fmt.Sprintf(youtubeListEndpoint, os.Getenv("YOUTUBE_KEY"), id))
	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error request"), err
	}

	//get the json into struct
	err = json.Unmarshal(r, &res)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error parsing json"), err
	}

	if p.nextPageToken != "" {
		return *message.SetColor(embed.ErrorColor).SetDescription("Cannot add more than 1 playlist at the same time, wait until finishing extract current youtube playlist"), err
	}

	//if playlist have more than 30 songs
	if res.NextPageToken != "" {
		p.nextPageToken = res.NextPageToken
		p.playlistID = id
	}

	//add songs to playlist
	for _, v := range res.Items {
		println(v.Snippet.ResourceID.VideoID)
		p.queueMutex.Lock()
		p.playlist = append(p.playlist, Video{v.Snippet.ResourceID.VideoID, v.Snippet.Title})
		p.queueMutex.Unlock()
	}

	//call the player
	p.playerLoop()

	//send message succes back
	message.SetColor(embed.SuccessColor).SetDescription("added youtube playlist to playlist")

	return *message, nil

}

func (p *Player) nextPage() {
	var res YRequest
	//request
	r, err := myRequests.GetByteResponse(fmt.Sprintf(youtubeNextPageEndpoint, os.Getenv("YOUTUBE_KEY"), p.playlistID, p.nextPageToken))
	if err != nil {
		log.Println("error getting next youtube list page")
		return
	}

	//get the json into struct
	err = json.Unmarshal(r, &res)

	if err != nil {
		log.Println("error parsing youtube list next page json to struct")
		return
	}

	//if playlist have more than 30 songs
	p.nextPageToken = res.NextPageToken

	//reset when finish playlist
	if res.NextPageToken == "" {
		p.playlistID = ""
	}

	//add songs to playlist
	for _, v := range res.Items {
		p.queueMutex.Lock()
		p.playlist = append(p.playlist, Video{v.Snippet.ResourceID.VideoID, v.Snippet.Title})
		p.queueMutex.Unlock()
	}

}

func (p *Player) addVideo(id string) (embed.Embed, error) {
	message := embed.NewEmbed()
	var res YRequest

	//get the video
	r, err := myRequests.GetByteResponse(fmt.Sprintf(youtubeVideoEndpoint, os.Getenv("YOUTUBE_KEY"), id))

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error request"), err
	}

	//get the json into struct
	err = json.Unmarshal(r, &res)

	if err != nil {
		return *message.SetColor(embed.ErrorColor).SetDescription("error parsing json"), err
	}

	//add to playlist
	p.queueMutex.Lock()
	p.playlist = append(p.playlist, Video{res.Items[0].ID, res.Items[0].Snippet.Title})
	p.queueMutex.Unlock()

	//call the player to loop the songs
	p.playerLoop()

	//send message succes back
	message.SetColor(embed.SuccessColor).SetDescription("added: " + res.Items[0].Snippet.Title)

	return *message, nil
}

func (p *Player) playerLoop() {

	if p.isPlaying {
		// the bot is playing
		return
	}

	p.isPlaying = true

	go func() {
		for {
			//if playlist is empty check if has nextpagetoken
			if len(p.playlist) == 0 {
				if p.nextPageToken != "" {
					p.nextPage()
				} else {
					exit(p)
					return
				}
			}

			//get first video
			p.queueMutex.Lock()
			p.currentSong = Video{
				p.playlist[0].VidID,
				p.playlist[0].Title,
			}

			if len(p.playlist) != 0 {
				p.playlist = p.playlist[1:]
			}

			p.queueMutex.Unlock()

			//send song to player
			p.playSong(p.currentSong)

			//if looping true append the song to last
			if p.looping {
				p.queueMutex.Lock()
				p.playlist = append(p.playlist, p.currentSong)
				p.queueMutex.Unlock()
			}
		}
	}()
}

func getYoutubeID(url string) string {
	re := regexp.MustCompile("(https://www\\.youtube\\.com/watch\\?v=|https://youtu\\.be/|https\\://www.youtube\\.com/playlist\\?list=)")

	id := re.ReplaceAllString(url, "")
	re = regexp.MustCompile("(\\?|&)")
	ind := re.FindStringIndex(id)

	if ind != nil {
		id = id[0:re.FindStringIndex(id)[0]]
	}

	return id
}
