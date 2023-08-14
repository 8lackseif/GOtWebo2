package music

import (
	"botwebo2/lib/embed"
)

//in common

func (p *Player) stop() (embed.Embed, error) {
	message := embed.NewEmbed()
	//call exit function
	exit(p)

	message.SetColor(embed.SuccessColor).SetDescription("disconnected.")

	return *message, nil
}

func (p *Player) song() (embed.Embed, error) {
	message := embed.NewEmbed()
	return *message, nil
}

func (p *Player) skip() (embed.Embed, error) {
	message := embed.NewEmbed()

	if p.looping {
		p.looping = false
	}

	p.stopch <- false

	return *message.SetColor(embed.SuccessColor).SetDescription(p.currentSong.Title + " skipped"), nil
}

func exit(p *Player) {
	//lock list and unlocked when function finished

	//set to zero value
	p.looping = false
	p.emptyList()
	p.isRandom = false
	p.nextPageToken = ""
	p.playlistID = ""
	p.randomSong = ""
	p.randongSongSlug = ""
	p.randomSongImage = ""
	p.isPlaying = false

	//disconnect voice
	p.voiceClient.Disconnect()
	p.voiceClient.Speaking(false)
}
