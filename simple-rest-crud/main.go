package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akhiltn/learn-golang/go-simple-rest-crud/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is started at port :4000")
}
