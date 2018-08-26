package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dabbotorg/gobot/config"
	"github.com/foxbot/gavalink"

	"github.com/foxbot/discordgo"
)

var (
	low        int
	high       int
	total      int
	prometheus int
)

var conf config.Config
var errors chan error
var logger *log.Logger
var lavalink *gavalink.Lavalink

func init() {
	errors = make(chan error)
	logger = log.New(os.Stdout, "(main) ", log.Ldate|log.Ltime)

	flag.IntVar(&low, "low", 0, "low shard ID")
	flag.IntVar(&high, "high", 1, "upper shard ID")
	flag.IntVar(&total, "total", 1, "total shards")
	flag.IntVar(&prometheus, "prom", 0, "prometheus port")
}

func main() {
	logger.Println("init...")
	conf = config.LoadConfig()

	clients, err := makeClients()
	if err != nil {
		panic(err)
	}

	t := strconv.Itoa(total)
	lavalink = gavalink.NewLavalink(t, conf.UserID)
	lavalink.AddNodes(conf.Nodes...)

	for _, s := range clients {
		err = s.Open()
		if err != nil {
			panic(err)
		}
	}
	logger.Println("startup OK!")

	go pumpErrors()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	for _, s := range clients {
		s.Close()
	}
}

func makeClients() ([]*discordgo.Session, error) {
	t := fmt.Sprintf("Bot %s", conf.Token)
	s := high - low

	shards := make([]*discordgo.Session, s)

	for i := 0; i < s; i++ {
		id := low + i
		s, err := discordgo.New(t)
		if err != nil {
			return nil, err
		}

		s.ShardCount = total
		s.ShardID = id
		s.LogLevel = discordgo.LogInformational

		s.AddHandler(onReady)
		s.AddHandler(onMessage)
		s.AddHandler(onVoiceServerUpdate)

		shards[i] = s
	}

	return shards, nil
}

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	mentionPrefix = e.User.Mention()

	logger.Println("shard ready", s.ShardID)
}

func onVoiceServerUpdate(s *discordgo.Session, e *discordgo.VoiceServerUpdate) {
	v := gavalink.VoiceServerUpdate{
		Endpoint: e.Endpoint,
		GuildID:  e.GuildID,
		Token:    e.Token,
	}

	if p, _ := lavalink.GetPlayer(e.GuildID); p != nil {
		err := p.Forward(s.State.SessionID, v)
		if err != nil {
			errors <- err
			return
		}
		return
	}

	node, err := lavalink.BestNode()
	if err != nil {
		errors <- err
		return
	}

	h := gavalink.DummyEventHandler{}
	_, err = node.CreatePlayer(e.GuildID, s.State.SessionID, v, h)
	if err != nil {
		errors <- err
	}
}

func pumpErrors() {
	l := log.New(os.Stdout, "(error) ", log.Ldate|log.Ltime)
	for {
		e, ok := <-errors
		if !ok {
			l.Println("errors channel closed, goodbye?")
			return
		}
		l.Println(e)
		// TODO: push to raven
	}
}
