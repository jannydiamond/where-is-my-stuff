package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Where is my Stuff?\n")
}
