package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", yum)
}

func yum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yum yum yum!")
}
