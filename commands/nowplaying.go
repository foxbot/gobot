package commands

import (
	"fmt"

	"github.com/foxbot/gavalink"
)

var nowPlaying = Command{
	Aliases:     []string{"nowplaying", "current", "now", "np"},
	Description: "Shows the current song",
	Method:      onNowPlaying,
}

const npf = "Currently playing **%s** by **%s** `[%d/%d]`\nSong URL: %s"

func onNowPlaying(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	t, err := gavalink.DecodeString(player.Track())
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	// TODO: prettify durations
	r := fmt.Sprintf(npf, t.Title, t.Author, t.Position, t.Length, t.URI)
	return textResponse(r)
}
