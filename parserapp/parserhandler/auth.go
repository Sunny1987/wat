package parserhandler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//authentication variables
var (
	jwtKey = []byte("secret_key")
	users  = map[string]string{
		"admin": "admin",
		"guest": "guest",
	}
)

//Credentials will receive the required credentials for usage
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Claims is used to set the expiredtime for a username
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Login method will verify the credentials for usage
func (n *NewLogger) Login(w http.ResponseWriter, r *http.Request) {
	n.l.Println("Starting login analysis...")
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	n.l.Printf("username: %v", credentials.Username)
	n.l.Printf("password: %v", credentials.Password)

	expectedPass, ok := users[credentials.Username]
	if !ok || expectedPass != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	expirationTimeAdmin := time.Now().Add(time.Minute * 60)
	expirationTimeGuest := time.Now().Add(time.Minute * 10)

	var claims *Claims
	if credentials.Username == "admin" {
		claims = &Claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTimeAdmin.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTimeAdmin,
		})
	}
	if credentials.Username == "guest" {
		claims = &Claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTimeGuest.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTimeGuest,
		})
	}

}
