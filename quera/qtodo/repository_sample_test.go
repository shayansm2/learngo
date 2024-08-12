package qtodo_test

import (
	"testing"
	"time"

	"qtodo"

	"github.com/stretchr/testify/assert"
)

// repository

func TestDBCreation(t *testing.T) {
	t.Parallel()
	var db qtodo.Database = qtodo.NewDatabase()
	assert.NotNil(t, db)
}

func TestGetTaskList(t *testing.T) {
	t.Parallel()
	var db qtodo.Database = qtodo.NewDatabase()
	assert := assert.New(t)
	name, description, alarmTime, action := "walk", "walk somewhere", time.Now().Add(requestTime), func() {}
	newTask, err := qtodo.NewTask(action, alarmTime, name, description)
	assert.Nil(err)

	name, description, alarmTime, action = "study", "do some studying", time.Now().Add(requestTime*2), func() {}
	newTask2, err := qtodo.NewTask(action, alarmTime, name, description)
	assert.Nil(err)

	err = db.SaveTask(newTask)
	assert.Nil(err)
	err = db.SaveTask(newTask2)
	assert.Nil(err)

	assert.Len(db.GetTaskList(), 2)
}
