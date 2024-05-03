package projects

import "testing"

func TestQuantityProjectsByProgressLevel(t *testing.T) {

	projects := make(map[string]Project)

	// Agregar algunos proyectos al mapa
	projects["proyecto1"] = Project{Name: "Proyecto 1", Progress: 1}
	projects["proyecto2"] = Project{Name: "Proyecto 2", Progress: 2}
	projects["proyecto3"] = Project{Name: "Proyecto 3", Progress: 1}
	projects["proyecto4"] = Project{Name: "Proyecto 4", Progress: 3}

	quantity := QuantityProjectsByProgressLevel(&projects, 1)

	expected := 2

	if quantity != expected {
		t.Errorf("Quantity is incorrect, got %d expected %d", expected, quantity)
	}

}
