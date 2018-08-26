package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dabbotorg/gobot/config"
)

var load = &Command{
	Aliases:     []string{"load", "undump"},
	Description: "Loads in a queue dump",
	Method:      onLoad,
}

func onLoad(c *Context) Response {
	if len(c.Args) == 0 {
		return textResponse("Usage: `{{prefix}}load <dumped playlist url>`")
	}

	cont, err := join(c)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	if !cont {
		return NoResponse{}
	}

	url := c.Args[0]
	if strings.Contains(url, "hastebin.com") && !strings.Contains(url, "raw") {
		name := url[strings.LastIndex(url, "/")+1:]
		url = "https://hastebin.com/raw/" + name
	}

	resp, err := http.Get(url)
	if err != nil {
		Errors <- err
		r := fmt.Sprintf("An error occured! %s", err.Error())
		return textResponse(r)
	}

	var tracks []string
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	err = json.Unmarshal(d, &tracks)
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	if len(tracks) < 1 {
		return textResponse("Error: This playlist contains no tracks.")
	}

	player, err := c.Lavalink.GetPlayer(c.Event.GuildID)
	if player == nil { // this shouldn't happen..
		for i := 0; i < 5; i++ {
			time.Sleep(250 * time.Millisecond)
			player, err = c.Lavalink.GetPlayer(c.Event.GuildID)
			if player != nil {
				break
			}
		}
		if player == nil {
			return textResponse("An error occured: timeout waiting for guild player")
		}
	}
	if err != nil {
		Errors <- err
		return textResponse("err")
	}

	_, err = c.Redis.Del("queues:" + c.Event.GuildID).Result()
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	c.State.QueueFlags[c.Event.GuildID] = config.FlagNone
	err = player.Stop()
	if err != nil {
		Errors <- err
		return textResponse("err")
	}
	err = player.Play(tracks[0])

	if l := len(tracks); l > 1 {
		// nice one go
		z := make([]interface{}, l)
		for i, v := range tracks {
			z[i] = v
		}

		_, err = c.Redis.RPush("queues"+c.Event.GuildID, z[1:]...).Result()
		if err != nil {
			Errors <- err
			return textResponse("err")
		}
	}

	r := fmt.Sprintf("Loaded %d tracks from <%s>", len(tracks), url)
	return textResponse(r)
}
