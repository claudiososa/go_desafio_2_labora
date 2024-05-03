package projects

import "testing"

func TestQuantityProjectsByProgressLevel(t *testing.T) {

	projects := map[string]Project{
		"proyecto1": {Name: "Proyecto 1", Progress: 1},
		"proyecto2": {Name: "Proyecto 2", Progress: 2},
		"proyecto3": {Name: "Proyecto 3", Progress: 1},
		"proyecto4": {Name: "Proyecto 4", Progress: 3},
	}

	quantity := QuantityProjectsByProgressLevel(&projects, 1)

	expected := 2

	if quantity != expected {
		t.Errorf("Quantity is incorrect, got %d expected %d", expected, quantity)
	}

}
