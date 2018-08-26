package commands

import "github.com/foxbot/discordgo"

func join(c *Context) (bool, error) {
	g, err := c.Session.State.Guild(c.Event.GuildID)
	if err != nil {
		return false, err
	}

	var as *discordgo.VoiceState
	var bs *discordgo.VoiceState
	var moveTo string

	for _, vs := range g.VoiceStates {
		if vs.UserID == c.Event.Author.ID {
			as = vs
		} else if vs.UserID == c.Session.State.User.ID {
			bs = vs
		}
	}

	if as == nil {
		_, err := c.Session.ChannelMessageSend(c.Event.ChannelID, "You must be in a voice channel!")
		if err != nil {
			return false, err
		}
		return false, nil
	} else if bs != nil && bs.ChannelID != as.ChannelID {
		// TODO: check permissions
		moveTo = as.ChannelID
	} else if bs == nil {
		moveTo = as.ChannelID
	}

	if len(moveTo) > 0 {
		err = c.Session.ChannelVoiceJoinManual(c.Event.GuildID, moveTo, false, true)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
