package verification

import (
	"log"
	"net/http"
	"strings"
)

func handleError(err error) (int, string) {
	if strings.HasPrefix(err.Error(), "server_") {
		log.Printf("[ERROR] %s\n", err.Error())
		return http.StatusInternalServerError, "internal_server_error"
	}

	return http.StatusBadRequest, err.Error()
}
