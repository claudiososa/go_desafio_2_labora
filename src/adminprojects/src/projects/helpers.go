package projects

import "fmt"

func ChangeProgressToAllProjects(projects *map[string]Project, newProgress int) {
	fmt.Println(newProgress)

	for p := range *projects {
		project := (*projects)[p]
		project.SetProgress(newProgress)
		(*projects)[p] = project
	}

}

func QuantityProjectsByProgressLevel(projects *map[string]Project, progressLevel int) int {

	quantity := 0

	for p := range *projects {
		if (*projects)[p].Progress == progressLevel {
			quantity++
		}
	}

	return quantity
}
