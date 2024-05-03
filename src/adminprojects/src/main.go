package main

import (
	"fmt"
	"strings"
	"adminprojects/pkg"
)

/*
Imagina que eres parte de un equipo de desarrollo de software en una
empresa de tecnología creativa. Se te ha asignado la tarea de crear
un sistema para gestionar proyectos creativos.
El sistema debe permitir la creación de nuevos proyectos, asignación
de miembros del equipo a proyectos, seguimiento del progreso del
proyecto y generación de informes de estado del proyecto.
Utiliza tu imaginación para definir cómo serán los proyectos creativos
y cómo se gestionarán en este sistema.
*/

func reportProjects(projects map[string]pkg.Project, members map[string]pkg.Member){
	fmt.Println("--------------------------------------------------------------------")
	fmt.Printf("%-20s %-30s %-20s\n","Project name", "Members", "Progress")
	fmt.Println("--------------------------------------------------------------------")

	
	for p := range projects{
		project := projects[p]

		members := strings.Join(project.Members, ",")

		// //values 1: start, 2: advanced, 3: finished
		var progress string

		switch project.Progress {
		case 1:
			progress = "Started"
		case 2:
			progress = "Adavanced"
		case 3:
			progress = "Finished"
		}

		fmt.Printf("%-20s %-30s %-20s\n",project.Name, members, progress)
	}
	fmt.Println("--------------------------------------------------------------------")
}


func main() {

	members := make(map[string]pkg.Member)

	var myMember pkg.Member


	//Creo y agrego nuevos Members al map members
	memb := myMember.NewMember("Juan", 22)
	memb1 := myMember.NewMember("Maria", 42)
	memb2 := myMember.NewMember("Carlos", 32)

	myMember.AddMember(*memb, &members)
	myMember.AddMember(*memb1, &members)
	myMember.AddMember(*memb2, &members)

	projects := make(map[string]pkg.Project)

	// Seguimiento del progreso del proyecto, incluyendo hitos y fechas límite.

	// Creo una variable de tipo Milestone

	var myMilestone pkg.Milestone

	// Obtengo una map con un listado de hitos creado para asignarlo a un proyecto.
	milestone1 := myMilestone.CreateMilestone(1)

	milestone2 := myMilestone.CreateMilestone(2)

	// Implementar un sistema para la creación de nuevos proyectos creativos.
	listMembers := []string{}

	var myProject pkg.Project

	project1 := myProject.NewProject("project1", listMembers, 1, milestone1)
	project2 := myProject.NewProject("project2", listMembers, 1, milestone2)
	project3 := myProject.NewProject("project3", listMembers, 2, milestone1)
	project4 := myProject.NewProject("project4", listMembers, 3, milestone2)

	myProject.AddProject(*project1, &projects)
	myProject.AddProject(*project2, &projects)
	myProject.AddProject(*project3, &projects)
	myProject.AddProject(*project4, &projects)


	// Permitir la asignación de miembros del equipo a proyectos.

	myProject.AddMemeberToProject("Carlos", "project1", &projects)
	myProject.AddMemeberToProject("Maria", "project1", &projects)
	myProject.AddMemeberToProject("Juan", "project2", &projects)
	myProject.AddMemeberToProject("Maria", "project3", &projects)
	myProject.AddMemeberToProject("Carlos", "project4", &projects)

	// Generación de informes de estado del proyecto, que muestren el progreso, los miembros 

	reportProjects(projects, members)
}
