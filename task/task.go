package task

import (
	"bytes"
	"errors"
	"fmt"
)

type Task struct {
	Name string
	Id   int
}

func (t Task) String() string {
	return fmt.Sprintf("[%v] %v", t.Id, t.Name)
}

type Tasks []*Task

func (t Tasks) MaxId() int {
	var max int

	for _, task := range t {
		if task.Id > max {
			max = task.Id
		}
	}
	return max
}

func (t Tasks) Get(id int) (*Task, error) {
	var task *Task

	for _, task = range t {
		if task.Id == id {
			return task, nil
		}
	}
	return nil, errors.New("index not found")
}

func (t Tasks) GetByName(name string) (*Task, error) {
	var task *Task

	for _, task = range t {
		if task.Name == name {
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (t *Tasks) Add(name string) {
	*t = append(*t, &Task{name, (*t).MaxId() + 1})
}

func (t *Tasks) Remove(id int) error {
	for i, task := range *t {
		if task.Id == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return nil
		}
	}
	return errors.New("index not found")
}

func (t Tasks) String() string {
	var buf bytes.Buffer

	for _, i := range t {
		buf.WriteString(i.String() + "\n")
	}
	return buf.String()
}
