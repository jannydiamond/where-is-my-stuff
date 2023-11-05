package handlers

import (
	"fmt"
	"net/http"
)

func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Where is my Stuff?\n")
}
