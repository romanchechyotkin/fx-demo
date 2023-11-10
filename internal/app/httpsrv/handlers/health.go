package handlers

import (
	"fmt"
	"net/http"
)

func Health() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "health\n")
	}
}
