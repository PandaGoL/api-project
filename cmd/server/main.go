package main

import (
	"fmt"
	"net/http"

	"github.com/PandaGoL/api-project/internal/database"
	intHttp "github.com/PandaGoL/api-project/internal/http"
)

func Run() error {
	fmt.Println("Running App")

	_, err := database.InitDatabase()
	if err != nil {
		return nil
	}
	handler := intHttp.NewHandler()
	handler.InitRoutes()
	if err := http.ListenAndServe(":9000", handler.Router); err != nil {
		return err
	}
	return nil
}
func main() {
	err := Run()
	if err != nil {
		fmt.Println("Error running app")
		fmt.Println(err)
	}
}
