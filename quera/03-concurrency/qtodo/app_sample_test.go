package qtodo_test

import (
	"qtodo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// app
func TestAppInit(t *testing.T) {
	t.Parallel()
	var db qtodo.Database = qtodo.NewDatabase()
	var app qtodo.App = qtodo.NewApp(db)
	assert := assert.New(t)
	assert.NotNil(app)
}

func TestAppCreation(t *testing.T) {
	t.Parallel()
	var db qtodo.Database = qtodo.NewDatabase()
	var app qtodo.App = qtodo.NewApp(db)
	assert := assert.New(t)
	assert.NotNil(app)

	err := app.AddTask("walk", "walking", time.Now().Add(requestTime), func() {}, false)
	assert.Nil(err)
}

func TestAppGetTaskList(t *testing.T) {
	t.Parallel()
	var db qtodo.Database = qtodo.NewDatabase()
	var app qtodo.App = qtodo.NewApp(db)
	assert := assert.New(t)
	assert.NotNil(app)

	err := app.AddTask("walk", "walking", time.Now().Add(requestTime), func() {}, false)
	//fmt.Println("db is", db)
	//fmt.Println("app db is", *app.GetDB())
	assert.Nil(err)

	actual := app.GetTaskList()
	require.Equal(t, 1, len(actual))
	assert.Equal("walk", actual[0].GetName())
}

func TestAppDelTask(t *testing.T) {
	t.Parallel()
	var db qtodo.Database = qtodo.NewDatabase()
	var app qtodo.App = qtodo.NewApp(db)
	assert := assert.New(t)
	assert.NotNil(app)

	err := app.AddTask("walk", "walking", time.Now().Add(requestTime), func() {}, false)
	assert.Nil(err)

	_, err = app.GetTask("walk")
	assert.Nil(err)

	err = app.DelTask("walk")
	assert.Nil(err)

	_, err = app.GetTask("walk")
	assert.Error(err)
}
