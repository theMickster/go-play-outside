package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{ \n\t\"response\": \"Hello World\" \n}")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Web server started and listening on localhost port 16926")
	log.Fatal(http.ListenAndServe(":16926", nil))
}
