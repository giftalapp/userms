package verification

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResendRequest struct {
	PhoneNumber string `json:"phone_number"`
	Service     string `json:"service"`
}

type ResendResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

func ResendHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := ResendResponse{
		Token: "",
		Error: "",
	}

	// Decode json request
	request := ResendRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
