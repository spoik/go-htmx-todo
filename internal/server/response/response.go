package response

import (
	"log"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	log.Printf("Internal server error caused by: %v", err)
}
