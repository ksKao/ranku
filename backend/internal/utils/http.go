package utils

import "net/http"

func WriteGenericInternalServerError(w http.ResponseWriter) {
	http.Error(w, "Something went wrong. Please try again.", http.StatusInternalServerError)
}
