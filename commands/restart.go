package commands

import (
	"fmt"

	"github.com/foxbot/gavalink"
)

var restart = &Command{
	Aliases:     []string{"restart"},
	Description: "Plays a song from the beginning",
	Method:      onRestart,
}

func onRestart(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	err = player.Seek(0)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	t, err := gavalink.DecodeString(player.Track())
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	// TODO: duration
	r := fmt.Sprintf("Restarted **%s** by **%s** `[%d]`", t.Title, t.Author, t.Length)
	return textResponse(r)
}
