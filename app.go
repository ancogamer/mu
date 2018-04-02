package main

import (
	"github.com/fiscaluno/mu/server"
	"github.com/fiscaluno/mu/user"
)

func main() {
	user.Migrate()
	server.Listen()
}
