package main

import (
	"fmt"

	db "github.com/luisupbeat/go-gorm/controllers"
)

func main() {
	db.DBConnection()
	fmt.Println("Hello men")
}