// Webアプリ導入とて、LocalhostでHelloworld表示


package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Used Golang")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is listening on :3000...")
	http.ListenAndServe(":3000", nil)
}