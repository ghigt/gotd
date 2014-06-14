package task

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id       int
	Name     string
	Duration time.Duration
}

func (t Task) String() string {
	return fmt.Sprintf("[%v] %q (%v)", t.Id, t.Name, t.Duration)
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

func (t *Tasks) Add(name string, duration time.Duration) *Task {
	nt := &Task{(*t).MaxId() + 1, name, duration}
	*t = append(*t, nt)
	return nt
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
