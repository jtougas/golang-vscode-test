package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world from GfG")

		out, err := exec.Command("lsasdasdasd").Output()

		// if there is an error with our execution
		// handle it here
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		// as the out variable defined above is of type []byte we need to convert
		// this to a string or else we will see garbage printed out in our console
		// this is how we convert it to a string
		fmt.Fprintln(w, "Command Successfully Executed")
		output := string(out[:])
		fmt.Fprintln(w, output)
	})

	port := ":8080"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
