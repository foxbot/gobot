package commands

import (
	"github.com/dabbotorg/gobot/config"
	"github.com/foxbot/discordgo"
	"github.com/foxbot/gavalink"
	"github.com/go-redis/redis"
)

// Errors pumps out internal commands errors
var Errors = make(chan error)

// Commands returns the bot's commands
func Commands() []*Command {
	return []*Command{
		// meta
		about,
		invite,
		shard,

		// playstate
		nowPlaying,
		pause,
		stop,
		volume,
		restart,
		skip,
		loop,
	}
}

// Command is a command
type Command struct {
	Aliases     []string
	Description string
	Method      Executor
}

// Executor defines a command action
type Executor func(c *Context) Response

// Context is a command context
type Context struct {
	Args     []string
	Config   *config.Config
	Event    *discordgo.MessageCreate
	Lavalink *gavalink.Lavalink
	Redis    *redis.Client
	Session  *discordgo.Session
	State    *config.State
}

// Response is an interface for a command response
type Response interface {
	Act(c *Context) error
}

// NoResponse is a response which does nothing
type NoResponse struct{}

// Act for NoResponse does nothing
func (r NoResponse) Act(c *Context) error {
	return nil
}

// TextResponse is a response which sends a message to the channel
type TextResponse struct {
	text string
}

// Act for TextResponse sends a message to the channel
func (r TextResponse) Act(c *Context) error {
	_, err := c.Session.ChannelMessageSend(c.Event.ChannelID, r.text)
	return err
}
func textResponse(c string) TextResponse {
	return TextResponse{
		text: c,
	}
}
