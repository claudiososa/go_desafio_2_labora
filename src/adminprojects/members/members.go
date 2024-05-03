package members

// Definición de la estructura Member
type Member struct {
	Name string
	Age  int
}

// Constructor de Member
func NewMember(name string, age int) *Member {
	return &Member{
		Name: name,
		Age:  age,
	}
}
