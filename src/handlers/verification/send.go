package verification

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/giftalapp/userms/src/middleware"
)

type SendRequest struct {
	PhoneNumber string `json:"phone_number"`
	Service     string `json:"service"`
}

type SendResponse struct {
	Token      string `json:"token"`
	Error      string `json:"error,omitempty"`
	statusCode int
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
	token := ""
	var err error = nil

	switch request.Service {
	case "sms":
		token, err = pubc.SMS.Send(request.PhoneNumber)
	case "whatsapp":
		token, err = pubc.WhatsApp.Send(request.PhoneNumber)
	default:
		err = errors.New("unsupported_service_error")
	}

	response.Token = token

	if err != nil {
		response.statusCode, response.Error = handleError(err)
	}

	// Return response
	responseBinary, _ := json.Marshal(response)

	w.WriteHeader(response.statusCode)
	fmt.Fprintln(w, string(responseBinary))
}
