package parserhandler

import (
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

	//track execution time for scan
	timeStart := time.Now()

	//get the request from middleware
	req := r.Context().Value(KeyUser{}).(*MyURLReq)
	resp, err := http.Get(req.URLFromReq)
	if err != nil {
		n.l.Printf("Error fetching URL response", err)
	}
	defer resp.Body.Close()

	logger := parser.GetMyLogger(n.l, req, Credential.Username)

	results := logger.Parse(resp.Body)

	//print the response
	PrintResponse(results, rw, n.l)

	n.l.Printf("Query completed in %v\n", time.Since(timeStart))

}
