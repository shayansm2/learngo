package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleGameCreation(t *testing.T) {
	g, err := NewGame([]int{})
	assert.Nil(t, err)
	assert.NotNil(t, g)
}

func TestSampleAddPlayer(t *testing.T) {
	g, err := NewGame([]int{1, 2, 3})
	assert.Nil(t, err)

	err = g.ConnectPlayer("Cyn")
	assert.Nil(t, err)
}

func TestSampleGetPlayer(t *testing.T) {
	g, err := NewGame([]int{1, 2, 3})
	assert.Nil(t, err)

	err = g.ConnectPlayer("Cyn")
	assert.Nil(t, err)

	p, err := g.GetPlayer("CyN")
	assert.Nil(t, err)
	assert.NotNil(t, p)
}

func TestShayan(t *testing.T) {
	g, err := NewGame([]int{1})
	assert.Nil(t, err)

	g.ConnectPlayer("Shayan")
	err = g.SwitchPlayerMap("shayan", 1)
	assert.Nil(t, err)

	g.ConnectPlayer("Aidin")
	g.SwitchPlayerMap("aidin", 1)

	g.ConnectPlayer("mahsa")
	g.SwitchPlayerMap("mahsa", 1)

	s, _ := g.GetPlayer("shayan")
	s.SendMessage("hello")
	// s.loc.FanOutMessages()
}
