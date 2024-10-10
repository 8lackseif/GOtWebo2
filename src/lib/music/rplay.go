package music

import "botwebo2/lib/embed"

const (
	graphql = "https://graphql.anilist.co"

	userQuery = `query($name: String) {
        User (name: $name) {
            id,
			name
        }
    }`
)

//Rplay

func (p *Player) random(user string) (embed.Embed, error) {
	message := embed.NewEmbed()
	return *message, nil
}
