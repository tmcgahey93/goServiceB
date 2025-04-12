package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", serviceBHandler)

	fmt.Println("Server Running on port 8081....")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func serviceBHandler(w http.ResponseWriter, r *http.Request) {
	returnMessage := "Hello from Service B"
	fmt.Fprintln(w, returnMessage)
}
