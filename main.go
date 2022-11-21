package main

import (
	f "fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := f.Fprintf(w, "Hello, World")
		if err != nil {
			f.Println(err)
		}
		f.Println(f.Sprintf("Bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
