package parserhandler

import (
	"encoding/json"
	"fmt"
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
