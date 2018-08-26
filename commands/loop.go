package commands

import (
	"github.com/dabbotorg/gobot/config"
)

var loop = Command{
	Aliases:     []string{"loop"},
	Description: "Repeats the whole song queue",
	Method:      onLoop,
}

func onLoop(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	flag := c.State.QueueFlags[c.Event.GuildID]
	if flag == config.FlagLoop {
		c.State.QueueFlags[c.Event.GuildID] = config.FlagNone
		return textResponse("**Disabled** queue looping!")
	}

	c.State.QueueFlags[c.Event.GuildID] = config.FlagLoop
	return textResponse("**Enabled** queue looping!")
}
