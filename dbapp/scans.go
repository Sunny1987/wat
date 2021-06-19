package dbapp

import (
	"encoding/json"
	"log"
	"strconv"
	"webparser/analyzerapp"
)

var (
	GetAllScansURL = "https://pgserv.herokuapp.com/scans"
	DeleteScanURL  = "https://pgserv.herokuapp.com/deletescan/"
	AddScanURL     = "https://pgserv.herokuapp.com/addscan"
)

//Result will provide the list of scans from database pgserv
type Result []struct {
	ScanID int             `json:"scan_id"`
	Person string          `json:"person"`
	URL    string          `json:"url"`
	Result json.RawMessage `json:"result"`
}

type ResultData struct {
	ScanID int    `json:"scan_id"`
	Person string `json:"person"`
	URL    string `json:"url"`
	Result string `json:"result"`
}

type JsonResp struct {
	DivResults      []analyzerapp.Divtag      `json:"divs"`
	ButtonResults   []analyzerapp.Buttontag   `json:"buttons"`
	InputResults    []analyzerapp.Inputtag    `json:"inputs"`
	ImageResults    []analyzerapp.Imgtag      `json:"images"`
	VideoResults    []analyzerapp.Videotag    `json:"videos"`
	AudioResults    []analyzerapp.Audiotag    `json:"audios"`
	TextareaResults []analyzerapp.Textareatag `json:"textAreas"`
	SelectResults   []analyzerapp.Selecttag   `json:"selects"`
	ParaResults     []analyzerapp.Paratag     `json:"paras"`
	IframeResults   []analyzerapp.Iframetag   `json:"iframes"`
	LinkResults     []analyzerapp.Anchortag   `json:"links"`
	AreaResults     []analyzerapp.Areatag     `json:"areas"`
	ObjectResults   []analyzerapp.Objecttag   `json:"objects"`
	EmbedResults    []analyzerapp.Embedtag    `json:"embeds"`
	TrackResults    []analyzerapp.Tracktag    `json:"tracks"`
}

//GetScans will return teh list of scans
func (res *Result) GetScans(l *log.Logger) {
	l.Println("**** Fetching the scans****")

	resp, err := Client.R().SetHeader("Accept", "application/json").Get(GetAllScansURL)
	if err != nil {
		l.Println(err)

	}
	err = json.Unmarshal(resp.Body(), res)
	if err != nil {
		l.Println(err)
	}

}

//DeleteAScan method will delete a scan from DB
func DeleteAScan(id int, l *log.Logger) {
	l.Println("**** Deleting a scans****")
	deleteScanURL := DeleteScanURL + strconv.Itoa(id)
	resp, err := Client.R().SetHeader("Accept", "application/json").Delete(deleteScanURL)
	if err != nil {
		l.Println(err)

	}
	l.Println(resp)

}

//AddScan method will delete a scan from DB
func AddScan(results analyzerapp.Response, l *log.Logger) {
	l.Println("**** Adding a scans****")
	resp, err := Client.R().SetHeader("Accept", "application/json").SetBody(results).Post(AddScanURL)
	if err != nil {
		l.Println(err)

	}
	l.Println(resp)
}
