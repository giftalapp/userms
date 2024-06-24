package verification

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giftalapp/authsrv/src/middleware"
)

type VerifyRequest struct {
	VerificationToken string `json:"verification_token"`
	Passcode          string `json:"passcode"`
}

type VerifyResponse struct {
	Error      string `json:"error,omitempty"`
	statusCode int
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := VerifyResponse{
		statusCode: http.StatusCreated,
	}

	// Reference Dependencies
	pubc := middleware.GetPub(r)

	// Decode json request
	request := VerifyRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Verify token
	phoneNumber, err := pubc.Verify(request.Passcode, request.VerificationToken)

	if err != nil {
		response.statusCode, response.Error = handleError(err)
	}

	fmt.Printf("%s (TODO: DATABASE OP)\n", phoneNumber)

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
