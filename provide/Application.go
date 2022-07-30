package provide

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloWord(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Path", r.URL.Path)
	name := r.Form.Get("name")
	result := "hello" + name
	fmt.Fprintf(w, result)
}

func main() {
	http.HandleFunc("/hello", sayHelloWord)
	err := http.ListenAndServe(":9098", nil)
	if err != nil {
		log.Fatalf("ListenAndServe.error:", err)
	}
}
