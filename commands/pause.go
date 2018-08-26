package commands

var pause = Command{
	Aliases: []string{"pause", "resume"},
	Method:  onPause,
}

func onPause(ctx *Context) Response {
	player, err := ctx.Lavalink.GetPlayer(ctx.Event.GuildID)
	if err != nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}

	action := !player.Paused()
	err = player.Pause(action)

	// this should never happen, but report it if it does
	if err != nil {
		Errors <- err
	}

	if action {
		return textResponse("Paused music playback! Use `{{prefix}}resume` to resume.")
	} else {
		return textResponse("Resumed music playback!")
	}
}
