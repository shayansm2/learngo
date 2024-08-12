package qtodo

import (
	"errors"
)

type Database interface {
	GetTaskList() []Task
	GetTask(string) (Task, error)
	SaveTask(Task) error
	DelTask(string) error
}

// type JSONFileDB string

type InMemoryDB map[string]Task

func (hashmap InMemoryDB) GetTaskList() []Task {
	result := make([]Task, 0)
	for _, val := range hashmap {
		result = append(result, val)
	}
	return result
}

func (hashmap InMemoryDB) GetTask(name string) (Task, error) {
	val, found := hashmap[name]
	if !found {
		return nil, errors.New("finding task with the given name")
	}
	return val, nil
}

func (hashmap InMemoryDB) SaveTask(task Task) error {
	if task.GetName() == "" {
		return errors.New("task with no name")
	}
	hashmap[task.GetName()] = task
	return nil
}

func (hashmap InMemoryDB) DelTask(name string) error {
	if _, found := hashmap[name]; !found {
		return errors.New("cannot find and delete the task")
	}
	delete(hashmap, name)
	return nil
}

func NewDatabase() InMemoryDB {
	return make(InMemoryDB)
}
