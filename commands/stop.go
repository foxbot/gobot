package commands

var stop = Command{
	Aliases:     []string{"stop", "leave", "clear"},
	Description: "Stops playing and leaves the voice channel",
	Method:      onStop,
}

func onStop(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("Not currently playing music.")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	err = player.Destroy()
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	c.Redis.Del("queues:" + c.Event.GuildID)

	return textResponse("Stopped playing music & left the voice channel.")
}
