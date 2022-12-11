package models

import "fmt"

type Project struct {
	Name  string
	Tasks []Task
}

func (p *Project) MarkTaskAsCompleted(id int) {
	p.Tasks[id-1].MarkAsComplete()
}

func (p *Project) AddTask(task Task) {
	p.Tasks = append(p.Tasks, task)
}

func (p *Project) RemoveTask(index int) {
	p.Tasks = append(p.Tasks[:index], p.Tasks[index+1:]...)
}

func (p *Project) PrintTasks() {
	if len(p.Tasks) == 0 {
		fmt.Println("Aún no tienes tareas")
		return
	}

	for index, task := range p.Tasks {
		fmt.Println("ID:", index+1)
		fmt.Println("Tarea:", task.Name)
		fmt.Println("Descripción:", task.Description)
		fmt.Println("Completado:", task.Completed)
		fmt.Println()
	}
}
