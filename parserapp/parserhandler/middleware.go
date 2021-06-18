package parserhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"log"
	"net/http"
	"strings"
	"webparser/analyzerapp"
)

type KeyUser struct{}

var ctx context.Context

//MiddlewareValidation will validate incoming request and authentication before scan
func (n *NewLogger) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		n.l.Println("*****Entering middleware******")

		if r.URL.Path != "/login" {
			//var req parser.MyURLReq
			req := &MyURLReq{}
			err := json.NewDecoder(r.Body).Decode(req)
			if err != nil {
				n.l.Printf("Error : %v", err)
			}
			n.l.Printf("req: %v", req)
			cookie, err := r.Cookie("token")
			if err != nil {
				if err == http.ErrNoCookie {
					rw.WriteHeader(http.StatusUnauthorized)
					n.l.Printf("Error : %v", err)
					n.l.Println("*****Exiting middleware******")
					return
				}
				rw.WriteHeader(http.StatusBadRequest)
				n.l.Printf("Error : %v", err)
				n.l.Println("*****Exiting middleware******")
				return
			}
			tokenStr := cookie.Value
			claims := &Claims{}
			tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					rw.WriteHeader(http.StatusUnauthorized)
					n.l.Printf("Error : %v", err)
					n.l.Println("*****Exiting middleware******")

					return
				}
				rw.WriteHeader(http.StatusBadRequest)
				n.l.Printf("Error : %v", err)
				n.l.Println("*****Exiting middleware******")
				return
			}
			if !tkn.Valid {
				rw.WriteHeader(http.StatusUnauthorized)
				n.l.Printf("Error : %v", err)
				n.l.Println("*****Exiting middleware******")
				return
			}

			//perform field validation
			err = req.Validate()
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				n.l.Printf("Error : %v", err)
				n.l.Println("*****Exiting middleware******")
				return
			}

			ctx = context.WithValue(r.Context(), KeyUser{}, req)
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		} else {
			next.ServeHTTP(rw, r)
		}

		n.l.Println("No Error")
		n.l.Println("*****Exiting middleware******")
	})
}

//PrintResponse will print the final response for scan
func PrintResponse(results analyzerapp.Response, rw http.ResponseWriter, l *log.Logger) {
	l.Println("Initiating the response....")
	rep, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		l.Println(err)
	}
	_, err = fmt.Fprintln(rw, string(rep))
	if err != nil {
		l.Printf("Error : %v", err)
	}
}

//Validate the request object
func (req *MyURLReq) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("url", ValidateReqUrl)
	if err != nil {
		return err
	}
	return validate.Struct(req)
}

//ValidateReqUrl validates Request object URL parameter
func ValidateReqUrl(fl validator.FieldLevel) bool {

	//check if url is empty
	if fl.Field().String() == "" {
		return false
	}

	//check if url has http/https
	if !strings.Contains(fl.Field().String(), "https://") &&
		!strings.Contains(fl.Field().String(), "http://") {
		return false
	}

	return true

}
