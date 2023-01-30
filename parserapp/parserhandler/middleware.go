package parserhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"webparser/analyzerapp"
	"webparser/dbapp"
	"webparser/parserapp/parser"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
)

type KeyUser struct{}

var ctx context.Context
var scanResultDB dbapp.Result

//MiddlewareValidation will validate incoming request and authentication before scan
func (n *NewLogger) MiddlewareValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
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

			//Validate Auth
			err = AuthValidate(n.l, rw, r)
			if err != nil {
				return
			}

			//perform field validation
			err = req.Validate()
			if err != nil {
				if strings.Contains(err.Error(), "lte") {
					http.Error(rw, "Max depth equal to or greater than 3 is not allowed", http.StatusBadRequest)
				}
				if strings.Contains(err.Error(), "gt") {
					http.Error(rw, "Min depth less than 0 is not allowed", http.StatusBadRequest)
				}

				//rw.WriteHeader(http.StatusBadRequest)
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

			//Validate Auth
			err := AuthValidate(n.l, rw, r)
			if err != nil {
				return
			}

			//ctx = context.WithValue(r.Context(), KeyUser{}, req)
			ctx = context.TODO()
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)

		case "/uploadhtml":

			//Validate Auth
			err := AuthValidate(n.l, rw, r)
			if err != nil {
				return
			}

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

	//max retention count in db
	maxRetention := 2

	if len(scanResultDB) > 0 {
		countRecs := 0
		for _, res := range scanResultDB {
			if res.Person == Credential.Username {
				scanmap[countRecs] = res.ScanID
				countRecs++
			}
		}

		//logic for record exceeding max retention
		if len(results) >= maxRetention || countRecs >= maxRetention {
			// Add logic to delete all scans
			if countRecs != 0 {
				for i := 0; i < countRecs; i++ {
					dbapp.DeleteAScan(scanmap[i], l)
				}
			}

			//Add the scan results in db
			for index, result := range results {
				if index < maxRetention {
					if strings.Contains(result.Person, "guest") || strings.Contains(result.Person, "admin") {
						dbapp.AddScan(result, l)
					}
				}

			}
		}

		//logic for record within max retention
		if len(results) < maxRetention {

			// Add logic to delete additional scans
			if countRecs != 0 {
				for i := 0; i < len(results); i++ {
					dbapp.DeleteAScan(scanmap[i], l)
				}
			}

			//Add the scan results in db
			for index, result := range results {
				if index < len(results) {
					if strings.Contains(result.Person, "guest") || strings.Contains(result.Person, "admin") {
						dbapp.AddScan(result, l)
					}
				}

			}

		}
	} else {
		//Add the scan results in db
		for index, result := range results {
			if index < maxRetention {
				if strings.Contains(result.Person, "guest") || strings.Contains(result.Person, "admin") {
					dbapp.AddScan(result, l)
				}
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

//startScan will start the scan for a URL
func startScan(req *MyURLReq, l *log.Logger, base string) analyzerapp.Response {

	if req.File == nil {
		resp, err := http.Get(req.URLFromReq)
		if err != nil {
			l.Print("Error fetching URL response", err)
		}
		defer resp.Body.Close()

		logger := parser.GetMyLogger(l, req, Credential.Username)

		//results for a link
		return logger.Parse(resp.Body, base)

	} else {

		logger := parser.GetMyLogger(l, req, Credential.Username)
		return logger.Parse(req.File, base)

	}

}

//AuthValidate will validate the cookie before all requests
func AuthValidate(l *log.Logger, rw http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			l.Printf("Error : %v", err)
			l.Println("*****Exiting middleware******")
			return err
		}
		rw.WriteHeader(http.StatusBadRequest)
		l.Printf("Error : %v", err)
		l.Println("*****Exiting middleware******")
		return err
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			l.Printf("Error : %v", err)
			l.Println("*****Exiting middleware******")

			return err
		}
		rw.WriteHeader(http.StatusBadRequest)
		l.Printf("Error : %v", err)
		l.Println("*****Exiting middleware******")
		return err
	}
	if !tkn.Valid {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		l.Printf("Error : %v", err)
		l.Println("*****Exiting middleware******")
		return err
	}

	return nil

}
