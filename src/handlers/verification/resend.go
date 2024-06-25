package verification

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/giftalapp/userms/src/middleware"
)

type ResendRequest struct {
	VerificationToken string `json:"verification_token"`
	Service           string `json:"service"`
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

	// Create and store token && Send OTP
	var err error = nil

	switch request.Service {
	case "sms":
		err = pubc.SMS.Resend(request.VerificationToken)
	case "whatsapp":
		err = pubc.WhatsApp.Resend(request.VerificationToken)
	default:
		err = errors.New("unsupported_service_error")
	}

	if err != nil {
		response.statusCode, response.Error = handleError(err)
	}

	fmt.Printf("%s (TODO: DATABASE OP)\n", request.VerificationToken)

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
