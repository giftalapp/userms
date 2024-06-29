package verification

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giftalapp/userms/src/middleware"
)

type ResendRequest struct {
	VerificationToken string `json:"verification_token"`
}

type ResendResponse struct {
	Error      string `json:"error,omitempty"`
	statusCode int
}

func ResendHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := ResendResponse{
		statusCode: http.StatusOK,
	}

	// Reference Dependencies
	pubc := middleware.GetPub(r)

	// Decode json request
	request := ResendRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Send OTP
	err := pubc.WhatsApp.Resend(request.VerificationToken, "en")

	if err != nil {
		response.statusCode, response.Error = handleError(err)
	}

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
