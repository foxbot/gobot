package commands

var invite = &Command{
	Aliases:     []string{"invite", "addbot"},
	Description: "Adds the bot to your server",
	Method:      onInvite,
}

func onInvite(c *Context) Response {
	return textResponse("Invite dabBot: " + c.Config.Invite)
}
