package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello WebAssembly")
	log.Fatal(http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`))))
}
