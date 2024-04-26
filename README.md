# Task Management System ( src/AdminTasks )

This Go project implements a simple task management system to manage pending tasks for a development team. The program allows for creating new tasks, assigning responsibilities, updating task statuses, and viewing all pending tasks.

## Usage

### Create New Task
To create a new task, use the `newTask` function. Provide the name, description, and responsible person for the task. The task will be initialized with a status of "pending".

### Assign Responsible
To assign a responsible person to a task, use the `SetResponsible` method of the `Task` struct. Provide the name of the responsible person.

### Update Task Status
To update the status of a task to "in progress" or "complete", use the `SetStatus` method of the `Task` struct. Provide the desired status.

### View Pending Tasks
To view all pending tasks, use the `listTasks` function with the title "Only Pending".


# Product Inventory Management System ( src/AdminProducts )

This Go project implements a simple inventory management system to manage products. The program allows for adding new products, updating the quantity of existing products, and deleting products from the inventory.

## Usage

### Add Product
To add a new product to the inventory, use the `addProduct` function. Provide the name, price, and quantity of the product.

### Update Product Quantity
To update the quantity of an existing product, modify the `quantity` field of the respective product directly.

### Delete Product
To delete a product from the inventory, use the `delete` function. Provide the name of the product to be deleted.


# Contact Management System ( src/AdminContacts )

This Go project implements a simple contact management system to manage a user's contacts. The program allows for adding new contacts, searching for contacts by name or phone number, updating contact information, and deleting contacts.

## Usage

### Add Contact
To add a new contact to the list, use the `addContact` function. Provide the name, phone number, and address of the contact.

### Search Contact
To search for contacts by name or phone number, use the `searchContact` function. Provide a text string to search for in either the name or phone number fields.

### Update Contact
To update the information of a contact (phone number, email, or address), use the `updateContact` function. Provide the name of the contact to be updated along with the new phone number, address, or other information.

### Delete Contact
To delete a contact from the list, use the `deleteContact` function. Provide the name of the contact to be deleted.


# Library Book Management System ( src/AdminBooks )

This Go project implements a simple library management system to manage a collection of books. The program allows for adding new books, searching for books by title or author, updating the status of a book (available or borrowed), and deleting books from the collection.

## Usage

### Add Book
To add a new book to the collection, use the `addBook` function. Provide the title, author, genre, and status of the book.

### Search Book
To search for books by title or author, use the `searchBook` function. Provide a text string to search for in either the title or author fields.

### Update Status
To update the status of a book to either "available" or "borrowed", use the `updateStatus` function. Provide the new status and the title of the book to be updated.

### Delete Book
To delete a book from the collection, use the `deleteBook` function. Provide the title of the book to be deleted.

