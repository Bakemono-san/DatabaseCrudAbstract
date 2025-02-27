package databasecrudabstract

import "fmt"

// generic type
type JpaRepository struct {
	entity interface{}
}

// method to create a new record in the database
func (r *JpaRepository) Create() {
	fmt.Println("Create ", r.entity)
}

// method to update an existing record in the database
func (r *JpaRepository) Update() {
	fmt.Println("Update ", r.entity)
}

// method to delete a record from the database
func (r *JpaRepository) Delete() {
	fmt.Println("Delete ", r.entity)
}

// method to find all records in the database
func (r *JpaRepository) FindAll() {
	fmt.Println("Find All ", r.entity)
}

// method to find a record by ID in the database
func (r *JpaRepository) FindById(id int) {
	fmt.Printf("Find User by ID: %d\n", id)
}
