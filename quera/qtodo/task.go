package qtodo

import (
	"errors"
	"time"
)

type Task interface {
	DoAction()
	GetAlarmTime() time.Time
	GetAction() func()
	GetName() string
	GetDescription() string
}

type ToDoTask struct {
	action      func()
	name        string
	description string
	alarmTime   time.Time
}

func (this *ToDoTask) DoAction() {
	this.action()
}

func (this *ToDoTask) GetAlarmTime() time.Time {
	return this.alarmTime
}

func (this *ToDoTask) GetAction() func() {
	return this.action
}

func (this *ToDoTask) GetName() string {
	return this.name
}

func (this *ToDoTask) GetDescription() string {
	return this.description
}

func NewTask(
	action func(),
	alarmTime time.Time,
	name string,
	description string,
) (*ToDoTask, error) {
	if name == "" {
		return nil, errors.New("empty name")
	}

	if description == "" {
		return nil, errors.New("empty description")
	}

	if alarmTime.Before(time.Now()) {
		return nil, errors.New("alarm time is in the past")
	}

	return &ToDoTask{
		action:      action,
		alarmTime:   alarmTime,
		name:        name,
		description: description,
	}, nil
}
