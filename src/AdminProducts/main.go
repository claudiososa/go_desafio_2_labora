package main

import "fmt"

/*
Implementar una estructura de datos para almacenar la información de cada
producto (nombre, precio y cantidad disponible).
*/

type Product struct {
	name     string
	price    float64
	quantity int
}

func (p Product) String() string {
	return fmt.Sprintf("%-15s $%-9.2f %-10d\n", p.name, p.price, p.quantity)
}

func addProduct(products *map[string]Product, p Product) {
	(*products)[p.name] = p
}

func listProducts(products *map[string]Product, title string) {
	fmt.Println("-----------------------------------------------------")
	fmt.Println(title)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("%-15s %-10s %-10s\n", "Name", "Price", "Quantity")

	for _, product := range *products {
		fmt.Print(product)
	}
	fmt.Println()
}

func main() {

	products := make(map[string]Product)

	// Permitir la adición de nuevos productos con sus respectivas cantidades.
	prod := Product{name: "table", price: 211.20, quantity: 20}
	prod1 := Product{name: "car", price: 33211.20, quantity: 10}
	prod2 := Product{name: "door", price: 351.20, quantity: 12}
	prod3 := Product{name: "computer", price: 4351.20, quantity: 2}

	addProduct(&products, prod)
	addProduct(&products, prod1)
	addProduct(&products, prod2)
	addProduct(&products, prod3)

	listProducts(&products, "Add prduct")

	// Permitir la actualización de la cantidad disponible de un producto existente. -> Product name:  table
	if product, ok := products["table"]; ok {

		product.quantity = 10

		products["table"] = product
	}
	listProducts(&products, "Update product")

	// Permitir la eliminación de productos del inventario. -> quitar product car
	delete(products, "car")
	listProducts(&products, "Delete product")

	// Mostrar el inventario completo, incluyendo el nombre, precio y cantidad disponible de cada producto.
	listProducts(&products, "List products")

}
