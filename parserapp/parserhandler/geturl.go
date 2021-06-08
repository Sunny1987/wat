package parserhandler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
	"webparser/parserapp/parser"
)

//NewLogger is the logger for the GET function
type NewLogger struct {
	l *log.Logger
}

//GetNewLogger returns the NewLogger handler
func GetNewLogger(l *log.Logger) *NewLogger {
	return &NewLogger{l: l}
}

//GetURLResp will return the WCAG2.1 guidelines result for a URL
func (n *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
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
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	timeStart := time.Now()
	req := &parser.MyURLReq{}
	json.NewDecoder(r.Body).Decode(req)
	resp, err := http.Get(req.URLFromReq)
	if err != nil {
		n.l.Printf("Error fetching URL response", err)
	}
	defer resp.Body.Close()

	logger := parser.GetMyLogger(n.l, req)

	results := logger.Parse(resp.Body)

	n.l.Println("Initiating the response....")
	rep, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		n.l.Println(err)
	}
	fmt.Fprintln(rw, string(rep))
	n.l.Printf("Query completed in %v\n", time.Since(timeStart))

}
