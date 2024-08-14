package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	var err error
	s := NewSurvey()

	err = s.AddFlight("my-flight")
	assert.NoError(t, err)
	err = s.AddFlight("my-flight")
	assert.Error(t, err)

	err = s.AddTicket("my-flight", "p1")
	assert.NoError(t, err)
	err = s.AddTicket("my-flight", "p2")
	assert.NoError(t, err)

	err = s.AddComment("my-flight", "p1", Comment{Score: 9, Text: "good"})
	assert.NoError(t, err)

	err = s.AddComment("other flight", "p2", Comment{Score: 8, Text: "ok"})
	assert.Error(t, err)

	assert.EqualValues(t, 9, s.GetAllCommentsAverage()["my-flight"])

	comments, err := s.GetComments("my-flight")
	assert.NoError(t, err)
	assert.Equal(t, []string{"good"}, comments)
	assert.Equal(t, []string{"good"}, s.GetAllComments()["my-flight"])
}

func TestFLightAddConcurrent(t *testing.T) {
	s := NewSurvey()
	wg := sync.WaitGroup{}
	var err1, err2 error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err1 = s.AddFlight("my-flight")
	}()
	go func() {
		defer wg.Done()
		err2 = s.AddFlight("my-flight")
	}()
	wg.Wait()
	assert.True(t, err1 == nil || err2 == nil)
	assert.False(t, err1 == nil && err2 == nil)
}

func TestTicketAddConcurrent(t *testing.T) {
	s := NewSurvey()
	s.AddFlight("my-flight")
	wg := sync.WaitGroup{}
	var err1, err2 error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err1 = s.AddTicket("my-flight", "p1")
	}()
	go func() {
		defer wg.Done()
		err2 = s.AddTicket("my-flight", "p1")
	}()
	wg.Wait()
	assert.True(t, err1 == nil || err2 == nil)
	assert.False(t, err1 == nil && err2 == nil)
}
