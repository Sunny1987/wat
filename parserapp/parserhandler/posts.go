package parserhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
	"webparser/analyzerapp"
	"webparser/parserapp/sitemapbuilder"
)

var wg sync.WaitGroup

//GetURLResp will return the WCAG2.1 guidelines result for a URL
func (n *NewLogger) GetURLResp(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	//track execution time for scan
	timeStart := time.Now()

	//get the request from middleware
	req := r.Context().Value(KeyUser{}).(*MyURLReq)

	//get the list of links from sitemap
	links, base := sitemapbuilder.SiteMap(req.URLFromReq, req.MaxDepth, n.l)
	n.l.Println("***** site map completed*****")
	var finalResult []analyzerapp.Response

	//scan based on depth
	wg.Add(len(links))
	for i, link := range links {

		go func(link string) {
			defer wg.Done()

			//setup req object
			reqMod := &MyURLReq{}
			reqMod.URLFromReq = link
			reqMod.MaxDepth = req.MaxDepth

			n.l.Printf("Link# %v : %v ", i, reqMod.URLFromReq)
			//start scan for url
			results := startScan(reqMod, n.l, base)
			//mu.Lock()
			finalResult = append(finalResult, results)
			//mu.Unlock()

		}(link)

	}

	wg.Wait()
	//print the response
	PrintResponse(finalResult, rw, n.l)

	n.l.Printf("Query completed in %v\n", time.Since(timeStart))

}

//FileScan will upload an HTML file for test
func (n *NewLogger) FileScan(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
	n.l.Println("******** Starting file upload*****")

	//track execution time for scan
	timeStart := time.Now()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(rw, "Error max file size exceeded", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("myfile")

	if err != nil {
		n.l.Println("Error Retrieving the File")
		n.l.Println(err)
		return
	}
	defer file.Close()
	n.l.Printf("Uploaded File: %+v\n", handler.Filename)
	n.l.Printf("File Size: %+v\n", handler.Size)
	n.l.Printf("MIME Header: %+v\n", handler.Header)

	reqMod := &MyURLReq{}
	reqMod.URLFromReq = ""
	reqMod.MaxDepth = 0
	reqMod.File = file
	//
	results := startScan(reqMod, n.l, "")
	//
	rep, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		n.l.Println(err)
	}
	_, err = fmt.Fprintln(rw, string(rep))
	if err != nil {
		n.l.Printf("Error : %v", err)
	}

	n.l.Printf("Query completed in %v\n", time.Since(timeStart))

}
