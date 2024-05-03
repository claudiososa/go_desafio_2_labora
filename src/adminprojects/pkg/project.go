package pkg

type Project struct {
	Name string
	Members []string
	Progress int //values 1: start, 2: advanced, 3: finished
	Milestones map[string]Milestone
}

func (p *Project) setMember(name string) {
	p.Members = append(p.Members, name)
}

func (p *Project) NewProject(name string, members []string, progress int, milestones map[string]Milestone) *Project{
	return &Project{
		Name: name,
		Members: members,
		Progress: progress,
		Milestones: milestones,
	}
}





func (p *Project) AddProject( project Project, projects *map[string]Project){
	(*projects)[project.Name] = project
}

func (p *Project) AddMemeberToProject(memberName string, projectName string, projects *map[string]Project){
	//memberToAssign := "Carlos"

	if project,ok :=(*projects)[projectName]; ok {
		project.setMember(memberName)
		(*projects)[projectName] = project
	}
}

