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
	"webparser/dbapp"
)

type KeyUser struct{}

var ctx context.Context
var scanResultDB dbapp.Result

//MiddlewareValidation will validate incoming request and authentication before scan
func (n *NewLogger) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		n.l.Println("*****Entering middleware******")

		switch r.URL.Path {
		case "/scan":
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

		case "/login":
			next.ServeHTTP(rw, r)

		case "/getscans":
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

			//ctx = context.WithValue(r.Context(), KeyUser{}, req)
			ctx = context.TODO()
			r = r.WithContext(ctx)

			next.ServeHTTP(rw, r)

		}

		n.l.Println("No Error")
		n.l.Println("*****Exiting middleware******")
	})
}

//PrintResponse will print the final response for scan
func PrintResponse(results []analyzerapp.Response, rw http.ResponseWriter, l *log.Logger) {
	l.Println("Initiating the response....")

	//Get all Scans list
	scanResultDB.GetScans(l)

	//Add logic to check scan count in db
	scanmap := make(map[int]int)

	if len(scanResultDB) > 0 {
		countRecs := 0
		for _, res := range scanResultDB {
			if res.Person == Credential.Username {
				scanmap[countRecs] = res.ScanID
				countRecs++
			}
		}
		//maximum number of record to be retained
		maxRetention := 20

		if countRecs > maxRetention {

			// Add logic to delete additional scans
			for i := 0; i < countRecs-maxRetention; i++ {
				dbapp.DeleteAScan(scanmap[i], l)
			}

			//Add the scan results in db
			for _, result := range results {
				if result.Person != "" {
					dbapp.AddScan(result, l)
				}
			}

		} else {
			//Add the scan results in db
			for _, result := range results {
				if result.Person != "" {
					dbapp.AddScan(result, l)
				}
			}

		}
	} else {
		//Add the scan results in db
		for _, result := range results {
			if result.Person != "" {
				dbapp.AddScan(result, l)
			}
		}

	}

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
