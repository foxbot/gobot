package main

import (
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

	next, err := rdis.LPop("queues:" + player.GuildID()).Result()
	if err != nil {
		errors <- err
		return err
	}
	if next == "" {
		return nil
	}

	err = player.Play(next)
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
