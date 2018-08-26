package commands

const aboutText = `
**dabBot - the music bot that makes you dab**
Command prefix:` + "`!!!`" + `
Invite me to your server: <https://dabbot.org/invite>
Support server: https://discord.gg/9gwuZsv
Github: <https://github.com/ducc/JavaMusicBot>
`

var about = Command{
	Aliases: []string{"about", "info", "support"},
	Method:  onAbout,
}

func onAbout(c *Context) Response {
	return textResponse(aboutText)
}
