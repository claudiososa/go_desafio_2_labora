package pkg


type Milestone struct {
	Description string
	Date string
	Status string //done or todo

}

func (m *Milestone) CreateMilestone(num int) map[string]Milestone {
	if num ==  1 {
		return map[string]Milestone{
			"mileston1": Milestone{ Description: "mileston1", Date: "10/10/2023", Status: "done"},
			"mileston2": Milestone{ Description: "mileston1", Date: "10/10/2023", Status: "todo"},
		}
	} else {
		return map[string]Milestone{
			"mileston1": Milestone{ Description: "mileston1", Date: "09/09/2023", Status: "done"},
			"mileston2":  Milestone{ Description: "mileston1", Date: "18/08/2023", Status: "todo"},
		}
	}

}