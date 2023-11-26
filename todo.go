package go_todo_app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task      string
	Done      bool
	Created   time.Time
	Completed time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:      task,
		Done:      false,
		Created:   time.Now(),
		Completed: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].Completed = time.Now()
	ls[index-1].Done = true
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Read(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Save(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Todos"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Created"},
			{Align: simpletable.AlignCenter, Text: "Completed"},
		},
	}
	var cells [][]*simpletable.Cell
	for index, item := range *t {
		index++
		task := blue(item.Task)
		done := blue("no")
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", index)},
			{Text: task},
			{Text: done},
			{Text: item.Created.Format(time.RFC850)},
			{Text: item.Completed.Format(time.RFC850)},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("you still have %d todos left", t.PendingTasks()))},
	}}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *Todos) PendingTasks() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return total
}
