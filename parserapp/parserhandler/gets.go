package parserhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"webparser/analyzerapp"
	"webparser/dbapp"
	"webparser/parserapp/parser"
	"webparser/parserapp/sitemapbuilder"
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

	//get the list of links from sitemap
	links := sitemapbuilder.SiteMap(req.URLFromReq, req.MaxDepth, n.l)

	var finalResult []analyzerapp.Response

	//scan based on depth
	for i, link := range links {
		//setup req object
		reqMod := &MyURLReq{}
		reqMod.URLFromReq = link
		reqMod.MaxDepth = req.MaxDepth
		n.l.Printf("Link# %v : %v ", i, link)

		//start scan for url
		results := startScan(reqMod, n.l, rw)
		finalResult = append(finalResult, results)

	}
	//print the response
	PrintResponse(finalResult, rw, n.l)

	n.l.Printf("Query completed in %v\n", time.Since(timeStart))

}

//startScan will start the scan for a URL
func startScan(req *MyURLReq, l *log.Logger, rw http.ResponseWriter) analyzerapp.Response {
	resp, err := http.Get(req.URLFromReq)
	if err != nil {
		l.Printf("Error fetching URL response", err)
	}
	defer resp.Body.Close()

	logger := parser.GetMyLogger(l, req, Credential.Username)

	//results for a link
	return logger.Parse(resp.Body)

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
