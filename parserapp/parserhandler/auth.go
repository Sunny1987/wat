package parserhandler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	"webparser/dbapp"
)

//authentication variables
var (
	jwtKey    = []byte("secret_key")
	users     = map[string]string{}
	usersList dbapp.UserList
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

var Credential Credentials

//Login method will verify the credentials for usage
func (n *NewLogger) Login(w http.ResponseWriter, r *http.Request) {
	n.l.Println("Starting login analysis...")

	//get the users list
	if usersList == nil {
		usersList.GetUsers(n.l)
	}

	//validate update before proceed
	if usersList == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&Credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n.l.Printf("userlist: %v", usersList)
	for _, u := range usersList {
		n.l.Printf("user: %v", u)
		users[u.Username] = u.Password

	}

	//n.l.Printf("username: %v", users["admin"])
	n.l.Printf("username: %v", Credential.Username)
	n.l.Printf("password: %v", Credential.Password)

	expectedPass, ok := users[Credential.Username]
	if !ok || expectedPass != Credential.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return

	}
	expirationTimeAdmin := time.Now().Add(time.Minute * 60)
	expirationTimeGuest := time.Now().Add(time.Minute * 10)

	var claims *Claims
	if Credential.Username == "admin" {
		claims = &Claims{
			Username: Credential.Username,
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
	if Credential.Username == "guest" {
		claims = &Claims{
			Username: Credential.Username,
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
