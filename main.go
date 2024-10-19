package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/manishSharma1-dev/goPractice/routes"
)

func main() {
	fmt.Println("Creating a small go api using Mongodb")
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listen at 4000")
}
