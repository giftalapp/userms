package verification

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/giftalapp/userms/config"
	"github.com/giftalapp/userms/src/middleware"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type VerifyRequest struct {
	VerificationToken string `json:"verification_token"`
	Passcode          string `json:"passcode"`
}

type VerifyResponse struct {
	UserToken  string `json:"user_token,omitempty"`
	Error      string `json:"error,omitempty"`
	statusCode int
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize headers
	w.Header().Add("Content-Type", "application/json")

	// Prepare result map
	response := VerifyResponse{
		statusCode: http.StatusOK,
	}

	// Reference Dependencies
	pubc := middleware.GetPub(r)
	db := middleware.GetDB(r)

	// Decode json request
	request := VerifyRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	// Verify token
	phoneNumber, err := pubc.Verify(request.Passcode, request.VerificationToken)

	if err != nil {
		response.statusCode, response.Error = handleError(err)

		responseBinary, _ := json.Marshal(response)
		w.WriteHeader(response.statusCode)
		fmt.Fprintln(w, string(responseBinary))
		return
	}

	// Register user if he doesn't exist
	userRow := db.QueryRow(
		"SELECT * FROM user WHERE phone = ?",
		phoneNumber,
	)

	user := UserType{}

	err = userRow.Scan(
		&user.Uid,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.Gender,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			user.Uid = uuid.New().String()
			user.PhoneNumber = phoneNumber
			response.statusCode = http.StatusCreated

			_, err = db.Exec(
				"INSERT INTO user (uid, phone) VALUES (?, ?)",
				user.Uid,
				user.PhoneNumber,
			)

			if err != nil {
				response.statusCode, response.Error = handleError(fmt.Errorf("server_sql_error %s", err.Error()))
			}
		} else {
			response.statusCode, response.Error = handleError(fmt.Errorf("server_sql_error %s", err.Error()))
		}

		responseBinary, _ := json.Marshal(response)

		w.WriteHeader(response.statusCode)
		fmt.Fprintln(w, string(responseBinary))
	}

	// Calculate user_token jwt
	userToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": config.Env.AppName,
		"iat": time.Now(),
		"uid": user.Uid,
	})

	if response.UserToken, err = userToken.SignedString([]byte(config.Env.JWTSecret)); err != nil {
		response.statusCode, response.Error = handleError(fmt.Errorf("server_jwt_error %s", err.Error()))
	}

	// Return response
	responseBinary, _ := json.Marshal(response)

	w.WriteHeader(response.statusCode)
	fmt.Fprintln(w, string(responseBinary))
}
