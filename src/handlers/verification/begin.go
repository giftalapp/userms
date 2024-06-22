package verification

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giftalapp/authsrv/src/middleware"
)

type BeginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Service     string `json:"service"`
}

type BeginResponse struct {
	Token string `json:"token"`
	Error string `json:"error,omitempty"`
}

func BeginHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := BeginResponse{
		Token: "",
		Error: "",
	}

	// Reference Dependencies
	pubc := middleware.GetPub(r)

	// Decode json request
	request := BeginRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Create and store token
	token := ""
	var err error = nil

	switch request.Service {
	case "sms":
		token, err = pubc.SMS.Send(request.PhoneNumber)
	case "whatsapp":
		token, err = pubc.WhatsApp.Send(request.PhoneNumber)
	default:
		err = fmt.Errorf("the service %s is not supported", request.Service)
	}

	if err != nil {
		w.WriteHeader(400)
		response.Token = token
		response.Error = err.Error()
	}

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
