package user

import "github.com/fiscaluno/mu/db"

// Migrate migration User BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{FacebookID: "1234"})

	// Read
	var user User
	db.First(&user, 1)                         // find user with id 1
	db.First(&user, "facebook_id = ?", "1234") // find user with FacebookID 1234

	// Update - update user's FacebookID to 12345
	db.Model(&user).Update("FacebookID", "12345")

	// Delete - delete user
	// db.Delete(&user)
}
