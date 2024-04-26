package main

import "fmt"

/*
Un equipo de desarrollo necesita un programa para gestionar sus tareas pendientes.
El programa debe permitir la creación de nuevas tareas, la asignación de
responsables, la actualización del estado de una tarea y la visualización de
todas las tareas pendientes.
*/

// Implementar una estructura de datos para almacenar la información de cada tarea (nombre, descripción, responsable y estado).
type Task struct {
	name        string
	description string
	responsible string
	status      string
}

func newTask(name, description, responsible string) *Task {
	return &Task{
		name:        name,
		description: description,
		responsible: responsible,
		status:      "pending",
	}
}

func addNewTask(tasks *map[string]Task, task *Task) {
	(*tasks)[task.name] = *task
}

func (t Task) String() string {
	return fmt.Sprintf("%-15s %-18s %-18s %-18s\n", t.name, t.description, t.responsible, t.status)
}

func (t *Task) SetResponsible(responsible string) {
	t.responsible = responsible
}

func (t *Task) SetStatus(status string) {
	t.status = status
}

func listTasks(tasks *map[string]Task, title string) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println(title)
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("%-15s %-18s %-18s %-18s\n", "Name", "Description", "Responsible", "Status")

	for _, task := range *tasks {
		if title == "Only Pending" {
			if task.status == "pending" {
				fmt.Print(task)
			}
		} else {
			fmt.Print(task)
		}
	}
	fmt.Println()
}

func main() {

	tasks := make(map[string]Task)

	// Permitir la creación de nuevas tareas con un estado inicial de "pendiente".
	task := newTask("run", "go to the park", "John")
	task1 := newTask("buy", "go to store", "Clark")

	addNewTask(&tasks, task)
	addNewTask(&tasks, task1)

	listTasks(&tasks, "Tasks")

	// Permitir la asignación de un responsable a cada tarea.
	myTask := "run"

	if task, ok := tasks[myTask]; ok {
		task.SetResponsible("Richard")
		tasks[myTask] = task
	}

	listTasks(&tasks, "Update Responsible")

	// Permitir la actualización del estado de una tarea a "en progreso" o "completada".
	myTask2 := "buy"

	if task, ok := tasks[myTask2]; ok {
		task.SetStatus("complete")
		tasks[myTask2] = task
	}

	listTasks(&tasks, "Update Status")

	// Mostrar todas las tareas pendientes, incluyendo su nombre, descripción, responsable y estado.
	listTasks(&tasks, "Only Pending")

}
