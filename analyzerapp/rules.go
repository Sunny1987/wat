package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
	"sync"
)

//MyAnalysisLog is the logger for analysisapp
type MyAnalysisLog struct {
	l *log.Logger
}

//Analysis will return the MyAnalysisLog logger
func Analysis(l *log.Logger) *MyAnalysisLog {
	return &MyAnalysisLog{l: l}
}

var wg sync.WaitGroup

//ApplyRules will initiate rule application
func (l *MyAnalysisLog) ApplyRules(nodeMap map[string][]*html.Node) Response {
	l.l.Println("Initiating Rules Check....")
	var ruleResults Response

	//div analysis
	wg.Add(1)
	go func() {
		myMap := divAnalysis(l.l, nodeMap)
		myList := myMap["div"]
		ruleResults.DivResults = myList
		wg.Done()
	}()

	//button analysis
	wg.Add(1)
	go func() {
		myMap := buttonAnalysis(l.l, nodeMap)
		myList := myMap["button"]
		ruleResults.ButtonResults = myList
		wg.Done()
	}()

	//input analysis
	wg.Add(1)
	go func() {
		myMap := inputAnalysis(l.l, nodeMap)
		myList := myMap["input"]
		ruleResults.InputResults = myList
		wg.Done()
	}()

	//images analysis
	wg.Add(1)
	go func() {
		myMap := imagesAnalysis(l.l, nodeMap)
		myList := myMap["image"]
		ruleResults.ImageResults = myList
		wg.Done()
	}()

	//video analysis
	wg.Add(1)
	go func() {
		myMap := videoAnalysis(l.l, nodeMap)
		myList := myMap["video"]
		ruleResults.VideoResults = myList
		wg.Done()
	}()

	//audio analysis
	wg.Add(1)
	go func() {
		myMap := audioAnalysis(l.l, nodeMap)
		myList := myMap["audio"]
		ruleResults.AudioResults = myList
		wg.Done()
	}()

	wg.Wait()
	return ruleResults
}

//divAnalysis function initiates all the div rule based analysis
func divAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Divtag {
	l.Println("Initiating div tag Analysis......")
	nodes := nodeMap["divNodes"]

	var issues = make(map[string][]Divtag)
	var list []Divtag
	for _, node := range nodes {
		var tag Divtag
		tag.divRulesWCAG111(node, l)
		l.Printf("divTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("divs: %v\n", list)
	issues["div"] = list

	return issues
}

//buttonAnalysis function initiates all the div rule based analysis
func buttonAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Buttontag {
	l.Println("Initiating button tag Analysis......")
	nodes := nodeMap["buttonNodes"]

	var issues = make(map[string][]Buttontag)
	var list []Buttontag
	for _, node := range nodes {
		var tag Buttontag
		tag.buttonRulesWCAG111(node, l)
		l.Printf("buttonTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("buttons: %v\n", list)
	issues["button"] = list

	return issues
}

//inputAnalysis function initiates all the div rule based analysis
func inputAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Inputtag {
	l.Println("Initiating input tag Analysis......")
	nodes := nodeMap["inputNodes"]

	var issues = make(map[string][]Inputtag)
	var list []Inputtag
	for _, node := range nodes {
		var tag Inputtag
		tag.inputRulesWCAG111(node, l)
		l.Printf("inputTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("inputs: %v\n", list)
	issues["input"] = list

	return issues
}

//inputAnalysis function initiates all the div rule based analysis
func imagesAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Imgtag {
	l.Println("Initiating image tag Analysis......")
	nodes := nodeMap["imgNodes"]

	var issues = make(map[string][]Imgtag)
	var list []Imgtag
	for _, node := range nodes {
		var tag Imgtag
		tag.imagesRulesWCAG111(node, l)
		l.Printf("imageTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("images: %v\n", list)
	issues["image"] = list

	return issues
}

//videoAnalysis function initiates all the div rule based analysis
func videoAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Videotag {
	l.Println("Initiating video tag Analysis......")
	nodes := nodeMap["videoNodes"]

	var issues = make(map[string][]Videotag)
	var list []Videotag
	for _, node := range nodes {
		var tag Videotag
		tag.videoRulesWCAG111(node, l)
		l.Printf("videoTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("videos: %v\n", list)
	issues["video"] = list

	return issues
}

//audioAnalysis function initiates all the div rule based analysis
func audioAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Audiotag {
	l.Println("Initiating audio tag Analysis......")
	nodes := nodeMap["audioNodes"]

	var issues = make(map[string][]Audiotag)
	var list []Audiotag
	for _, node := range nodes {
		var tag Audiotag
		tag.audioRulesWCAG111(node, l)
		l.Printf("audioTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("audios: %v\n", list)
	issues["audio"] = list

	return issues
}
