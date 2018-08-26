package config

const (
	// FlagNone means no flags
	FlagNone = 0
	// FlagRepeat means repeat the current song
	FlagRepeat = 1
	// FlagLoop means loop the entire queue
	FlagLoop = 2
)

// State holds bot state
type State struct {
	QueueFlags map[string]int
}
