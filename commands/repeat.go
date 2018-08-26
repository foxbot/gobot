package commands

import "github.com/dabbotorg/gobot/config"

var repeat = Command{
	Aliases:     []string{"repeat"},
	Description: "Loops the current song",
	Method:      onLoop,
}

func onRepeat(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	flag := c.State.QueueFlags[c.Event.GuildID]
	if flag == config.FlagRepeat {
		c.State.QueueFlags[c.Event.GuildID] = config.FlagNone
		return textResponse("**Disabled** queue repeating!")
	}

	c.State.QueueFlags[c.Event.GuildID] = config.FlagRepeat
	return textResponse("**Enabled** queue repeating!")
}
