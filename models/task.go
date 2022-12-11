package models

type Task struct {
	Name        string
	Description string
	Completed   bool
}

func (t *Task) UpdateName(name string) {
	t.Name = name
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}

func (t *Task) MarkAsComplete() {
	t.Completed = true
}
