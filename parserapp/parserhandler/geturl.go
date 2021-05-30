package parserhandler

import (
	"encoding/json"
	"log"
	"net/http"
	"webparser/parserapp/parser"
)

type NewLogger struct {
	l *log.Logger
}

func GetNewLogger(l *log.Logger) *NewLogger {
	return &NewLogger{l: l}
}

func (n *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	req := &parser.MyURLReq{}
	json.NewDecoder(r.Body).Decode(req)
	resp, err := http.Get(req.URLFromReq)
	if err != nil {
		n.l.Printf("Error fetching URL response", err)
	}
	defer resp.Body.Close()

	logger := parser.GetMyLogger(n.l)

	logger.Parse(resp.Body)

}
