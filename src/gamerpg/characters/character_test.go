package game


import (
    "testing"
)

func TestGetName(t *testing.T) {

    character := &Character{Name: "Hero"}

    name := character.GetName()

    if name != "Hero" {
        t.Errorf("GetName() no produjo el resultado esperado. Obtenido: %s, Esperado: %s", name, "Hero")
    }
}

func TestGetPower(t *testing.T) {

    character := &Character{Power: 10}

    power := character.GetPower()

    if power != 10 {
        t.Errorf("GetPower() no produjo el resultado esperado. Obtenido: %d, Esperado: %d", power, 10)
    }
}

func TestGetHealth(t *testing.T) {

	character := &Character{Health: 5}

    health := character.GetHealth()

    if health != 5 {
        t.Errorf("GetHealth() no produjo el resultado esperado. Obtenido: %d, Esperado: %d", health, 5)
    }
}

func TestGetTypeCharacter(t *testing.T) {

	character := &Character{TypeCharacter: "hero"}

    characterType := character.GetTypeCharacter()

    if characterType != "hero" {
        t.Errorf("GetTypeCharacter() no produjo el resultado esperado. Obtenido: %s, Esperado: %s", characterType, "hero")
    }
}
