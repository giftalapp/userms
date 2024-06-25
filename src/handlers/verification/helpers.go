package verification

import (
	"log"
	"net/http"
	"strings"
)

type UserType struct {
	Uid         string
	Username    interface{} `json:"username"`
	FirstName   interface{} `json:"first_name"`
	LastName    interface{} `json:"last_name"`
	Email       interface{} `json:"email"`
	PhoneNumber interface{} `json:"phone_number"`
	Gender      interface{} `json:"gender"`
}

func handleError(err error) (int, string) {
	if strings.HasPrefix(err.Error(), "server_") {
		log.Printf("[ERROR] %s\n", err.Error())
		return http.StatusInternalServerError, "internal_server_error"
	}

	return http.StatusBadRequest, err.Error()
}
