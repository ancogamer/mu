package user

import "github.com/fiscaluno/mu/db"

// Migrate migration User BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "J"})

	// Read
	var user User
	db.First(&user, 1)               // find user with id 1
	db.First(&user, "name = ?", "J") // find user with name J

	// Update - update user's Name to JC
	db.Model(&user).Update("Name", "JC")

	// Delete - delete user
	// db.Delete(&user)
}
