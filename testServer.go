package main

//required libraries
import (
	"fmt"
	"log"
	"net/http"
)

//main function to call greeting handler and actually start the server on port 8080
func main() {
	http.HandleFunc("/hello", greetingHandler)

	fmt.Printf("Hey, lets start a server on port 8080\n")
	/*if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}*/
	//updated code to include TLS
	err := http.ListenAndServeTLS(":8080", "test-server.crt", "https-server.key", nil)
	if err != nil {
		log.Fatal(err)
	}

}

//a separate function to handle the greeting "hello" with error handling
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method GET is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello and happy Friday!")
}
