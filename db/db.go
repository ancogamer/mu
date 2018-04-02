package db

import (
	"github.com/fiscaluno/pandorabox"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn connect on SQL Data Base Gorm
func Conn() (db *gorm.DB) {

	typeDB := pandorabox.GetOSEnvironment("DB", "sqlite3")

	pathDB := pandorabox.GetOSEnvironment("DATABASE_URL", "test.db")

	db, err := gorm.Open(typeDB, pathDB)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
