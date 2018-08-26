package commands

import (
	"fmt"
	"strconv"
)

var volume = Command{
	Aliases:     []string{"volume", "v"},
	Description: "Changes the music volume",
	Method:      onVolume,
}

const volPatronOnly = `**The volume command is dabBot premium only!**
Donate for the ` + "`Volume Control`" + `tier on Patreon at https://patreon.com/dabbot to gain access.`

func onVolume(c *Context) Response {
	if !c.Config.Patreon {
		return textResponse(volPatronOnly)
	}

	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	// TODO: check for patron status
	g, err := c.Session.Guild(c.Event.GuildID)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	var state bool
	for _, vs := range g.VoiceStates {
		if vs.UserID == c.Event.Author.ID {
			state = true
		}
	}

	if !state {
		return textResponse("You must be in a voice channel!")
	}

	if len(c.Args) == 0 {
		r := fmt.Sprintf("Current volume: **%d**", player.GetVolume())
		return textResponse(r)
	}

	v, err := strconv.Atoi(c.Args[0])
	if err != nil {
		return textResponse("Invalid volume. Bounds: `1 - 150`")
	}

	if v > 150 {
		v = 150
	}
	err = player.Volume(v)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	r := fmt.Sprintf("Set volume to **%d**", v)
	return textResponse(r)
}
