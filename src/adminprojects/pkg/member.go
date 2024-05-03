package pkg

// Definición de la estructura Member
type Member struct {
	name string
	age  int
}

// Constructor de Member
func (m *Member) NewMember(name string, age int) *Member {
	return &Member{
		name: name,
		age:  age,
	}
}



// Generación de informes de estado del proyecto, que muestren el progreso, los miembros del equipo asignados y cualquier problema o desafío que surja durante el desarrollo del proyecto.
// Utiliza tu creatividad para definir otros requisitos y funcionalidades que consideres importantes para la gestión eficaz de proyectos creativos.

func (m *Member) AddMember(member Member, members *map[string]Member) {
	(*members)[member.name] = member
}