package commands

var invite = Command{
	Aliases: []string{"invite", "addbot"},
	Method:  onInvite,
}

func onInvite(c *Context) Response {
	return textResponse("Invite dabBot: " + c.Config.Invite)
}
