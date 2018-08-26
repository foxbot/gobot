package commands

var skip = Command{
	Aliases:     []string{"skip", "s", "next"},
	Description: "Plays the next song in the queue",
	Method:      onSkip,
}

func onSkip(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	next, err := c.Redis.LPop("queues:" + c.Event.GuildID).Result()
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	if next != "" {
		player.Play(next)
	}
	return NoResponse{}
}
