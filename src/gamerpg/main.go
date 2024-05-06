package main

import (
	"fmt"
	"math/rand"
	"time"
	c "game/characters"
)

func advance() {
	fmt.Println("Avanzas hacia adelante...")
	time.Sleep(2 * time.Second)
	encounter := rand.Intn(10) + 1
	if encounter == 1 {
		fmt.Println("¡Te encuentras con un villano!")
		villain := c.Character{"Malvado", rand.Intn(10) + 1, rand.Intn(5) + 1, "Villain"}
		c.Battle(&villain)
	} else {
		fmt.Println("Continúas avanzando y llegas a una zona clara.")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	c.ShowIntroduction()
	continuar := true
	for continuar {
		option := c.SelectOption()
		switch option {
		case "1":
			advance()
		case "2":
			c.TurnRight()
		case "3":
			c.TurnLeft()
		case "4":
			fmt.Println("Gracias por jugar. Hasta luego...!")
			continuar = false
		default:
			fmt.Println("Opción inválida. Por favor, selecciona una opción válida.")
		}
	}
}
