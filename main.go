package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%v", port)
	http.ListenAndServe(addr, nil)
}
