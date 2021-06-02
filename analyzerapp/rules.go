package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
	"sync"
)

type MyAnalysisLog struct {
	l *log.Logger
}

func Analysis(l *log.Logger) *MyAnalysisLog{
	return &MyAnalysisLog{l: l}
}
var wg sync.WaitGroup

func (l *MyAnalysisLog) ApplyRules(nodeMap map[string][]*html.Node) Response {
	l.l.Println("Initiating Rules Check....")
	var ruleResults Response

	//div analysis
	wg.Add(1)
	go func() {
		myMap := divAnalysis(l.l,nodeMap)
		divList := myMap["div"]
		ruleResults.DivResults = divList
		wg.Done()
	}()



	wg.Wait()
	return ruleResults
}


//divAnalysis function initiates all the div rule based analysis
func divAnalysis(l *log.Logger, nodeMap map[string][]*html.Node ) map[string][]Divtag {
	l.Println("Initiating div tag Analysis......")
	divs := nodeMap["divNodes"]

	var divIssues = make(map[string][]Divtag)
	var divList []Divtag
	for i,div := range divs {
		var divTag Divtag
		divTag.divRulesWCAG111(div,i)
		divList = append(divList,divTag)
	}
	l.Println(divList)
	divIssues["div"] = divList

	return divIssues
}








