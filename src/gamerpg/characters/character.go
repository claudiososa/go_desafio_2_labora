package game

import "fmt"

type Character struct {
	Name  string
	Power int
	Health int
	TypeCharacter string
}

func (c *Character) ShowStatus(health int, typeCharacter string) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("Poder del %s: %d, Vidas del %s: %d\n", c.TypeCharacter, c.Power, typeCharacter,  c.Health)
	fmt.Println("-----------------------------------------------------------")
}

func (c *Character) UpdatedHealth(value int){
	c.Health = c.Health - value
}


func (character *Character) GetName() string {
	return character.Name
}

func (character *Character) GetPower() int {
	return character.Power
}

func (character *Character) GetHealth() int {
	return character.Health
}

func (character *Character) GetTypeCharacter() string {
	return character.TypeCharacter
}

func (character *Character) SetName(Name string) *Character {
	character.Name = Name
	return character
}

func (character *Character) SetPower(Power int) *Character {
	character.Power = Power
	return character
}

func (character *Character) SetHealth(Health int) *Character {
	character.Health = Health
	return character
}

func (character *Character) SetTypeCharacter(TypeCharacter string) *Character {
	character.TypeCharacter = TypeCharacter
	return character
}