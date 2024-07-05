package verification

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giftalapp/userms/src/middleware"
)

type SendRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type SendResponse struct {
	VerificationToken string `json:"verification_token,omitempty"`
	Error             string `json:"error,omitempty"`
	statusCode        int
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := SendResponse{
		statusCode: http.StatusOK,
	}

	// Reference Dependencies
	pubc := middleware.GetPub(r)

	// Decode json request
	request := SendRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Create and store token && Send OTP
	token, err := pubc.WhatsApp.Send(request.PhoneNumber, "en")

	response.VerificationToken = token

	if err != nil {
		response.VerificationToken = ""
		response.statusCode, response.Error = handleError(err)
	}

	// Return response
	responseBinary, _ := json.Marshal(response)

	w.WriteHeader(response.statusCode)
	fmt.Fprintln(w, string(responseBinary))
}
