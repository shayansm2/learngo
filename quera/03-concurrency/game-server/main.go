package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"unicode"
)

var wg sync.WaitGroup

type Player struct {
	name    string
	channel chan string
	loc     *Map
}

type channelMessage struct {
	name string
	msg  string
}

type Map struct {
	listener          chan channelMessage
	broadcastchannels map[string]chan string
}

type Game struct {
	players map[string]*Player
	maps    map[int]*Map
}

func NewGame(mapIds []int) (*Game, error) {
	g := Game{
		players: make(map[string]*Player),
		maps:    make(map[int]*Map),
	}

	for _, mapId := range mapIds {
		if mapId <= 0 {
			return nil, errors.New("map id is invalid")
		}

		if _, found := g.maps[mapId]; found {
			return nil, errors.New("duplicate map ids")
		}

		g.maps[mapId] = &Map{
			listener:          make(chan channelMessage, 10),
			broadcastchannels: make(map[string]chan string),
		}
	}

	return &g, nil
}

func (g *Game) ConnectPlayer(name string) error {
	if _, err := g.GetPlayer(name); err == nil {
		return errors.New("player already exists")
	}

	g.players[name] = &Player{
		name:    getStandardName(name),
		channel: make(chan string, 10),
	}
	return nil
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	p, err := g.GetPlayer(name)
	if err != nil {
		return err
	}

	mapLoc, err := g.GetMap(mapId)
	if err != nil {
		return err
	}

	if p.loc == mapLoc {
		return errors.New("already in the map")
	}

	p.loc.removeChannel(p.name)
	p.loc = mapLoc
	p.loc.addChannel(p.name, p.channel)

	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	name = getStandardName(name)
	if player, found := g.players[name]; found {
		return player, nil
	}
	return nil, errors.New("player not found")
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	if mapLoc, found := g.maps[mapId]; found {
		return mapLoc, nil
	}
	return nil, errors.New("map not found")
}

func (m *Map) FanOutMessages() {
	for i := 0; i < len(m.listener); i++ {
		msg := <-m.listener
		message := fmt.Sprintf("%s says: %s", msg.name, msg.msg)
		for name, ch := range m.broadcastchannels {
			if msg.name == name {
				continue
			}
			ch <- message
		}
	}

	wg.Done()
}

func (m *Map) addChannel(name string, channel chan string) {
	m.broadcastchannels[name] = channel
}

func (m *Map) removeChannel(name string) {
	if m == nil {
		return
	}

	delete(m.broadcastchannels, name)
}

func (p *Player) GetChannel() <-chan string {
	return p.channel
}

func (p *Player) SendMessage(msg string) error {
	p.loc.listener <- channelMessage{
		name: p.name,
		msg:  msg,
	}

	wg.Add(1)
	go p.loc.FanOutMessages()
	wg.Wait()

	return nil
}

func (p *Player) GetName() string {
	return p.name
}

func getStandardName(name string) string {
	for i, v := range name {
		return string(unicode.ToUpper(v)) + strings.ToLower(name[i+1:])
	}
	return ""
}
