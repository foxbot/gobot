package config

import (
	"github.com/BurntSushi/toml"
	"github.com/foxbot/gavalink"
)

// Config configures the bot
type Config struct {
	Token   string
	UserID  string
	Prefix  string
	Invite  string
	Patreon bool
	Redis   string
	Nodes   []gavalink.NodeConfig
}

// LoadConfig loads the config
func LoadConfig() Config {
	var conf Config
	_, err := toml.DecodeFile("./config.toml", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
