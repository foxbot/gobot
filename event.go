package main

import (
	"github.com/dabbotorg/gobot/config"
	"github.com/foxbot/gavalink"
)

var handler = lavalinkHandler{}

type lavalinkHandler struct{}

// OnTrackEnd is raised when a track ends
func (l lavalinkHandler) OnTrackEnd(player *gavalink.Player, track string, reason string) error {
	// track ended because we invoked Play
	if reason == "REPLACED" {
		return nil
	}

	flag := state.QueueFlags[player.GuildID()]

	var next string

	// repeat this track
	if flag == config.FlagRepeat {
		next = track
	} else {
		var err error
		next, err = rdis.LPop("queues:" + player.GuildID()).Result()
		if err != nil {
			errors <- err
			return err
		}
	}

	// if no more songs in queue, and queue loop disabled, queue is done
	if next == "" && flag != config.FlagLoop {
		err := player.Destroy()
		if err != nil {
			errors <- err
			return err
		}
		return nil
	} else if next == "" && flag == config.FlagLoop { // no songs in queue, but loop is on
		next = track
	} else if flag == config.FlagLoop { // more songs in queue, and loop is on
		_, err := rdis.RPush("queues:"+player.GuildID(), track).Result()
		if err != nil {
			errors <- err
			// don't break out here, it would impede on user experience
		}
	}

	err := player.Play(next)
	if err != nil {
		errors <- err
	}
	return err
}

// TODO: do we need to do anything for these situations?
// OnTrackException is raised when a track throws an exception
func (l lavalinkHandler) OnTrackException(player *gavalink.Player, track string, reason string) error {
	return nil
}

// OnTrackStuck is raised when a track gets stuck
func (l lavalinkHandler) OnTrackStuck(player *gavalink.Player, track string, threshold int) error {
	return nil
}
