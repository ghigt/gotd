package task

import (
	"bytes"
	"fmt"
)

type Task struct {
	Name string
	Id   int
}

func (t *Task) String() string {
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

func (t Tasks) Add(name string) Tasks {
	t = append(t, &Task{name, t.MaxId() + 1})
	return t
}

func (t Tasks) Remove(id int) Tasks {
	var i int
	var task *Task

	for i, task = range t {
		if task.Id == id {
			break
		}
	}
	t = append(t[:i], t[i+1:]...)
	return t
}

func (t Tasks) String() string {
	var buf bytes.Buffer

	for _, i := range t {
		buf.WriteString(i.String() + "\n")
	}
	return buf.String()
}
