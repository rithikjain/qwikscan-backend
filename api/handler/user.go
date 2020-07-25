package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/rithikjain/quickscan-backend/api/middleware"
	"github.com/rithikjain/quickscan-backend/api/view"
	"github.com/rithikjain/quickscan-backend/pkg/entities"
	"github.com/rithikjain/quickscan-backend/pkg/user"
	"net/http"
	"os"
)

func register(svc user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		var user entities.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			view.Wrap(err, w)
			return
		}

		u, err := svc.Register(&user)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		// Handling JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": u.Email,
			"role":  "user",
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
		if err != nil {
			view.Wrap(err, w)
			return
		}
		u.Password = ""
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Account Created",
			"token":   tokenString,
			"user":    u,
		})
	})
}

func login(svc user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}
		var user entities.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			view.Wrap(err, w)
			return
		}

		u, err := svc.Login(user.Email, user.Password)
		if err != nil {
			view.Wrap(err, w)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": u.Email,
			"role":  "user",
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
		if err != nil {
			view.Wrap(err, w)
			return
		}
		u.Password = ""
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Login Successful",
			"token":   tokenString,
			"user":    u,
		})
	})
}

// Protected Request
func userDetails(svc user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			view.Wrap(view.ErrMethodNotAllowed, w)
			return
		}

		claims, err := middleware.ValidateAndGetClaims(r.Context(), "user")
		if err != nil {
			view.Wrap(err, w)
			return
		}
		u, err := svc.GetUserByEmail(claims["email"].(string))
		if err != nil {
			view.Wrap(err, w)
			return
		}
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User Found",
			"user":    u,
		})
	})
}

// Handlers
func MakeUserHandler(r *http.ServeMux, svc user.Service) {
	r.Handle("/api/user/register", register(svc))
	r.Handle("/api/user/login", login(svc))
	r.Handle("/api/user/details", middleware.Validate(userDetails(svc)))
}
