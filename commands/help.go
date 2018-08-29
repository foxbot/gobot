package commands

import (
	"fmt"
	"sort"
	"strings"
)

var help = &Command{
	Aliases:     []string{"help", "commands", "h", "music"},
	Description: "Shows command help",
	Method:      onHelp,
}

const helpFmt = `**Commands:**\n%s\n\n**Quick start:** Use ` + "`{{prefix}}play <link>`" + `to start playing a song, 
use the same command to add another song, ` + "`{{prefix}}skip`" + `to go to the next song and
` + "`{{prefix}}stop`" + `to stop playing and leave.`

func onHelp(c *Context) Response {
	parts := make([]string, len(commands))
	for i, c := range commands {
		if c.Description == "" {
			parts[i] = ""
		}
		f := fmt.Sprintf("`%s` %s\n", c.Aliases[0], c.Description)
		parts[i] = f
	}
	sort.SliceStable(parts, func(i, j int) bool {
		return parts[i] < parts[j]
	})

	var b strings.Builder
	for _, f := range parts {
		b.WriteString(f)
	}

	r := fmt.Sprintf(helpFmt, b.String())
	return textResponse(r)
}
