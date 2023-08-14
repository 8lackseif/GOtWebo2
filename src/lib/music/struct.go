package music

import (
	"botwebo2/lib/embed"
	"sync"

	"github.com/bwmarrin/discordgo"
)

// player struct
type PlayerFunctions interface {
	//in common
	stop() (embed.Embed, error)
	song() (embed.Embed, error)
	skip() (embed.Embed, error)

	//play
	addSongs(url string, c string, dg *discordgo.Session) (embed.Embed, error)
	list() (embed.Embed, error)
	loop() (embed.Embed, error)
	emptyList() (embed.Embed, error)
	shuffle() (embed.Embed, error)
	addPlayList(id string) (embed.Embed, error)
	addVideo(id string) (embed.Embed, error)
	playerLoop()
	playSong(v Video)
	nextPage()

	//rplay
	random(user string) (embed.Embed, error)
}

type Player struct {
	//common
	guildID   string
	isPlaying bool
	stopch    chan bool
	playingch chan bool

	//voiceclient
	voiceClient *discordgo.VoiceConnection

	//youtube
	playlist      []Video
	currentSong   Video
	searchResults []Video
	looping       bool
	queueMutex    sync.RWMutex

	//for youtube lists +30
	playlistID    string
	nextPageToken string
	//for random
	isRandom        bool
	randomSong      string
	randongSongSlug string
	randomSongImage string
	user            string
}

type Video struct {
	VidID string
	Title string
}

type PkgSong struct {
	data Video
	v    *Player
}

//youtube videolist struct

type YRequest struct {
	NextPageToken string `json:"nextPageToken"`
	Items         []Item `json:"items"`
}

type Item struct {
	ID      string  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Title      string     `json:"title"`
	ResourceID ResourceID `json:"resourceId"`
}

type ResourceID struct {
	VideoID string `json:"videoId"`
}

type VideoResponse struct {
	Formats []Format `json:"formats"`
}
type Format struct {
	Url string `json:"url"`
}

//youtube video struct

//anilist userAnime struct

//animethemes music struct
