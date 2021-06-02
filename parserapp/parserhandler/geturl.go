package parserhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"webparser/parserapp/parser"
)

type NewLogger struct {
	l *log.Logger
}

func GetNewLogger(l *log.Logger) *NewLogger {
	return &NewLogger{l: l}
}

func (n *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	req := &parser.MyURLReq{}
	json.NewDecoder(r.Body).Decode(req)
	resp, err := http.Get(req.URLFromReq)
	if err != nil {
		n.l.Printf("Error fetching URL response", err)
	}
	defer resp.Body.Close()

	logger := parser.GetMyLogger(n.l)

	results := logger.Parse(resp.Body)
	n.l.Println("Initiating the response....")
	rep,err := json.MarshalIndent(results,""," ")
	if err != nil {
		n.l.Println(err)
	}
	fmt.Fprintln(rw,string(rep))
	n.l.Printf("Query completed in %v\n",time.Since(timeStart))

}
