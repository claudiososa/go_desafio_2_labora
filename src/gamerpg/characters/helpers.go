package game

import (
	"time"
	"fmt"
	"math/rand"
)

func TurnRight() {
	fmt.Println("Giras a la derecha y te encuentras con un río.")
	time.Sleep(2 * time.Second)
}

func TurnLeft() {
	fmt.Println("Giras a la izquierda y encuentras una cueva oscura.")
	time.Sleep(2 * time.Second)
}

func Battle(villain *Character) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("Comienza la batalla con %s!\n", villain.Name)
	fmt.Println("-------------------------------------------------------------")

	hero := Character{"Hero", rand.Intn(10) + 1, rand.Intn(5) + 1, "hero"}

	for villain.Health > 0 && hero.Health > 0 {

		villain.ShowStatus(hero.Health, villain.TypeCharacter)
		hero.ShowStatus(villain.Health, hero.TypeCharacter)

		// Heroe ataca
		playerAttack := rand.Intn(10) + 1
		fmt.Println("Heroe ataca al Villano!")
		time.Sleep(2 * time.Second)
		if playerAttack > villain.Power {
			fmt.Println("Heroe ha infligido daño al villano!")
			villain.UpdatedHealth(hero.Power)
		} else {
			fmt.Println("El villano esquiva tu ataque!")
		}

		if villain.Health <= 0 {
			fmt.Printf("Has derrotado al villano %s!\n", villain.Name)
			break
		}

		villain.ShowStatus(hero.Health, villain.TypeCharacter)
		hero.ShowStatus(villain.Health, hero.TypeCharacter)

		// Villano ataca
		fmt.Println("El villano ataca al Hero")
		time.Sleep(2 * time.Second)
		villainAttack := rand.Intn(10) + 1
		if villainAttack > hero.Power {
			fmt.Println("El villano te ha infligido daño!")
			hero.UpdatedHealth(villain.Power)
		} else {
			fmt.Println("Has esquivado el ataque del villano!")
		}

		if hero.Health <= 0 {
			fmt.Println("¡Has sido derrotado! Fin del juego.")
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func ShowIntroduction() {
	fmt.Println("¡Bienvenido al mundo del RPG!")
	fmt.Println("Te encuentras en un bosque oscuro y misterioso...")
	fmt.Println("Tu misión es encontrar el tesoro perdido.")
}

func SelectOption() string {
	var option string
	fmt.Println("\n¿Qué deseas hacer?")
	fmt.Println("1. Avanzar hacia adelante")
	fmt.Println("2. Girar a la derecha")
	fmt.Println("3. Girar a la izquierda")
	fmt.Println("4. Salir del juego")
	fmt.Print("Selecciona una opción: ")
	fmt.Scanln(&option)
	return option
}