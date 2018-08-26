package commands

var pause = Command{
	Aliases:     []string{"pause", "resume"},
	Description: "Pauses & resumes the song",
	Method:      onPause,
}

func onPause(ctx *Context) Response {
	player, err := ctx.Lavalink.GetPlayer(ctx.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	action := !player.Paused()
	err = player.Pause(action)

	// this should never happen, but report it if it does
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	if action {
		return textResponse("Paused music playback! Use `{{prefix}}resume` to resume.")
	}
	return textResponse("Resumed music playback!")
}
