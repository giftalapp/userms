package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	Verified bool  `json:"verified"`
	Error    error `json:"error"`
}

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request, fbAuth *sns.Client) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Decode json request
	request := VerifyTokenRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Prepare result map
	response := VerifyTokenResponse{
		Verified: true,
		Error:    nil,
	}

	// // Check token validity
	// if _, err := fbAuth.VerifyIDTokenAndCheckRevoked(context.Background(), request.Token); err != nil {
	// 	response.Verified = false
	// 	response.Error = err
	// 	w.WriteHeader(400)
	// }

	// Return success
	responseBinary, _ := json.Marshal(response)
	fmt.Fprintln(w, string(responseBinary))
}
