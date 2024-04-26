package main

import (
	"fmt"
	"strings"
)

/*
Un usuario necesita un programa para gestionar sus contactos. El programa debe permitir
la adición de nuevos contactos, la búsqueda de contactos por nombre o número de teléfono,
la actualización de la información de un contacto y la eliminación de contactos.
*/

// Implementar una estructura de datos para almacenar la información de cada
// contacto (nombre, número de teléfono, correo electrónico y dirección).
type Contact struct {
	name    string
	phone   string
	address string
}

// Constructor
func NewContact(name string, phone string, address string) *Contact {
	return &Contact{
		name:    name,
		phone:   phone,
		address: address,
	}

}

func (c *Contact) setName(name string) {
	c.name = name
}

func (c *Contact) setPhone(phone string) {
	c.phone = phone
}

func (c *Contact) setAddress(address string) {
	c.address = address
}

func (c Contact) String() string {
	return fmt.Sprintf("%-30s %-20s %-20s", c.name, c.phone, c.address)
}
func addContact(contacts *map[string]Contact, c *Contact) {
	(*contacts)[c.name] = *c
}

func searchContact(contacts map[string]Contact, text string) {

	contactFound := make(map[string]Contact)

	for _, contact := range contacts {
		//Busco por el nombre
		if strings.Contains(contact.name, text) {
			contactFound[contact.name] = contact
		} else {
			//Busco por el telefono
			if strings.Contains(contacts[contact.name].phone, text) {
				contactFound[contact.name] = contact
			}
		}
	}

	listContacts(contactFound, "Contacts found to '"+text+"'")
}

func updateContact(name string, phone string, address string, contacts *map[string]Contact) {
	if contact, ok := (*contacts)[name]; ok {
		contact.setName(name)
		contact.setPhone(phone)
		contact.setAddress(address)
		(*contacts)[name] = contact
	}
}

func listContacts(contacts map[string]Contact, name string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println(name)
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf("%-30s %-20s %-20s\n", "Name", "Phone", "Address")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, contact := range contacts {
		fmt.Println(contact)
	}
	fmt.Println()
}

func deleteContact(titleBook string, contacts *map[string]Contact) {
	delete((*contacts), titleBook)
}

func main() {

	contacts := make(map[string]Contact)

	contact1 := NewContact("John Doe", "1234567890", "123 Main St")
	contact2 := NewContact("Jane Smith", "0987654321", "456 Elm St")
	contact3 := NewContact("Alice Johnson", "5555555555", "789 Oak St")
	contact4 := NewContact("Bob Brown", "1112223333", "101 Pine St")
	contact5 := NewContact("Emily Davis", "4443332222", "202 Maple St")

	// Permitir la adición de nuevos contactos con su información completa.
	addContact(&contacts, contact1)
	addContact(&contacts, contact2)
	addContact(&contacts, contact3)
	addContact(&contacts, contact4)
	addContact(&contacts, contact5)

	listContacts(contacts, "List Books")

	// Permitir la búsqueda de contactos por nombre o número de teléfono.
	textToSearch := "mi"

	searchContact(contacts, textToSearch)

	// Permitir la actualización de la información de un contacto (teléfono, correo electrónico o dirección).

	name := "Bob Brown"
	phone := "7864444443"
	address := "home 222"

	updateContact(name, phone, address, &contacts)

	listContacts(contacts, "Update data")

	// Permitir la eliminación de contactos de la lista.

	nameContactToDelete := "Bob Brown"

	deleteContact(nameContactToDelete, &contacts)

	listContacts(contacts, "Contact deleted : "+nameContactToDelete)

}
