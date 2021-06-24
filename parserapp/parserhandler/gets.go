package parserhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"webparser/dbapp"
)

//NewLogger is the logger for the GET function
type NewLogger struct {
	l *log.Logger
}

//GetNewLogger returns the NewLogger handler
func GetNewLogger(l *log.Logger) *NewLogger {
	return &NewLogger{l: l}
}

//GetUserBasedScans will return the list of saved scans for the user
func (n *NewLogger) GetUserBasedScans(rw http.ResponseWriter, r *http.Request) {

	n.l.Println("******* Fetching the scans******")
	//track execution time for scan
	timeStart := time.Now()

	scanResultDB.GetScans(n.l)

	if len(scanResultDB) > 0 {
		for _, res := range scanResultDB {
			if res.Person == Credential.Username {
				var re dbapp.JsonResp
				err := json.Unmarshal(res.Result, &re)

				data := struct {
					Id         int            `json:"id"`
					RequestUrl string         `json:"request_url"`
					ResData    dbapp.JsonResp `json:"scanRes"`
				}{Id: res.ScanID, RequestUrl: res.URL, ResData: re}

				rep, err := json.MarshalIndent(&data, "", " ")
				if err != nil {
					n.l.Println(err)
				}
				_, err = fmt.Fprintln(rw, string(rep))
				if err != nil {
					n.l.Printf("Error : %v", err)
				}

			}
		}

	}

	n.l.Printf("Query completed in %v\n", time.Since(timeStart))
}
