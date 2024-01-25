package main

import (
	"fmt"

	"github.com/Tboules/dc_go_fullstack/internal/routes"
)

func main() {
	fmt.Println("Hello World")

	router := routes.AppRouter()

	router.Logger.Fatal(router.Start(":8000"))
}
