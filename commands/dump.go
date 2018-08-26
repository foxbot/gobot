package commands

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/whats-this/owo.go"
)

var dump = &Command{
	Aliases:     []string{"dump"},
	Description: "Outputs the music queue as a json string",
	Method:      onDump,
}

func onDump(c *Context) Response {
	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil {
		return textResponse("No music is playing on this guild! To play a song use `{{prefix}}play`")
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	q, err := c.Redis.LRange("queues:"+c.Event.GuildID, 0, -1).Result()
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	if t := player.Track(); t != "" {
		q = append([]string{t}, q...)
	}

	b, err := json.Marshal(q)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	r := bytes.NewReader(b)
	nr := owo.NamedReader{
		Filename: "queue.txt",
		Reader:   r,
	}
	resp, err := c.Owo.UploadFile(context.TODO(), nr)
	if err != nil {
		// TODO: post to hastebin of owo is down?
		Errors <- err
		return textResponse("upload err")
	}
	f := resp.Files[0].URL
	return textResponse("Dump created! " + f)
}
