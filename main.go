package main

import (
	"fmt"

	"github.com/itokun99/bang-jago/app/database"
)

func main() {
	fmt.Println("Bang Jago: Boilerplate Go-lang for Build REST API (>_<)")
	database.InitDb()
}
