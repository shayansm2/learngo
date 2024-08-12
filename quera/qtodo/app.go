package qtodo

import (
	"errors"
	"time"
)

type App interface {
	StartTask(string) error
	StopTask(string)
	AddTask(string, string, time.Time, func(), bool) error
	DelTask(string) error
	GetTaskList() []Task
	GetActiveTaskList() []Task
	GetTask(string) (Task, error)
}

// task runner
// runs these in goroutine and has 1 channel for communication and interruptions
// waits for interruptions
// waits for specific time to run a job
// runs a job
// should keep track of the distribution channels like map[name]chan
// also need another channel for the app for pausing all removing finished task

type ToDoApp struct {
	db                  Database
	finishChannel       chan string
	runningTaskChannels map[string]chan bool
	tempTasks           map[string]bool
}

func (app ToDoApp) StartTask(name string) error {
	task, err := app.db.GetTask(name)
	if err != nil {
		return errors.Join(err, errors.New("no task to run"))
	}

	if _, found := app.runningTaskChannels[task.GetName()]; found {
		return errors.New("task already started")
	}

	interruptChan := make(chan bool)
	app.runningTaskChannels[task.GetName()] = interruptChan
	go app.runTask(task, interruptChan)
	return nil
}

func (app ToDoApp) runTask(task Task, interruptChan chan bool) {
	select {
	case <-time.After(task.GetAlarmTime().Sub(time.Now())):
		task.DoAction()
		app.finishChannel <- task.GetName()
	case <-interruptChan:
		delete(app.runningTaskChannels, task.GetName())
		// should I send to finishChannel? I think NO
	}
}

func (app ToDoApp) StopTask(name string) {
	if _, err := app.db.GetTask(name); err != nil {
		return
	}

	interruptChan, found := app.runningTaskChannels[name]
	if !found {
		return
	}

	interruptChan <- true
}

func (app ToDoApp) AddTask(
	name string,
	description string,
	alarmTime time.Time,
	action func(),
	isTemp bool,
) error {
	task, err := NewTask(action, alarmTime, name, description)
	if err != nil {
		return err
	}

	err = app.db.SaveTask(task)
	if err != nil {
		return err
	}

	app.tempTasks[name] = isTemp
	return nil
}

func (app ToDoApp) DelTask(name string) error {
	err := app.db.DelTask(name)
	if err != nil {
		return err
	}
	delete(app.tempTasks, name)
	delete(app.runningTaskChannels, name)
	return nil
}

func (app ToDoApp) GetTaskList() []Task {
	return app.db.GetTaskList()
}

func (app ToDoApp) GetActiveTaskList() []Task {
	result := make([]Task, 0)
	for name := range app.runningTaskChannels {
		task, _ := app.db.GetTask(name)
		result = append(result, task)
	}
	return result
}

func (app ToDoApp) GetTask(name string) (Task, error) {
	return app.db.GetTask(name)
}

func (app ToDoApp) deleteTempTasks() {
	for name := range app.finishChannel {
		if isTemp, found := app.tempTasks[name]; found && isTemp {
			app.db.DelTask(name)
		}
	}
}

func NewApp(db Database) ToDoApp {
	app := ToDoApp{
		db:                  db,
		finishChannel:       make(chan string),
		runningTaskChannels: make(map[string]chan bool),
		tempTasks:           make(map[string]bool),
	}

	go app.deleteTempTasks()

	return app
}
