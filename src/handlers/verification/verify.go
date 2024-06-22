package verification

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type VerifyRequest struct {
	PhoneNumber string `json:"phone_number"`
	Service     string `json:"service"`
}

type VerifyResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := VerifyResponse{
		Token: "",
		Error: "",
	}

	// Decode json request
	request := VerifyRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
