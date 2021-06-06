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

	//textarea analysis
	wg.Add(1)
	go func() {
		myMap := textareaAnalysis(l.l, nodeMap)
		myList := myMap["textarea"]
		ruleResults.TextareaResults = myList
		wg.Done()
	}()

	//select analysis
	wg.Add(1)
	go func() {
		myMap := selectAnalysis(l.l, nodeMap)
		myList := myMap["select"]
		ruleResults.SelectResults = myList
		wg.Done()
	}()

	//iframe analysis
	wg.Add(1)
	go func() {
		myMap := iframeAnalysis(l.l, nodeMap)
		myList := myMap["iframe"]
		ruleResults.IframeResults = myList
		wg.Done()
	}()

	//link analysis
	wg.Add(1)
	go func() {
		myMap := linkAnalysis(l.l, nodeMap)
		myList := myMap["anchor"]
		ruleResults.LinkResults = myList
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

//buttonAnalysis function initiates all the button rule based analysis
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

//inputAnalysis function initiates all the input rule based analysis
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

//imagesAnalysis function initiates all the images rule based analysis
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

//videoAnalysis function initiates all the video rule based analysis
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

//audioAnalysis function initiates all the audio rule based analysis
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

//textareaAnalysis function initiates all the textarea rule based analysis
func textareaAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Textareatag {
	l.Println("Initiating textarea tag Analysis......")
	nodes := nodeMap["textareaNodes"]

	var issues = make(map[string][]Textareatag)
	var list []Textareatag
	for _, node := range nodes {
		var tag Textareatag
		tag.textAreaRulesWCAG111(node, l)
		l.Printf("textareaTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("textarea: %v\n", list)
	issues["textarea"] = list

	return issues
}

//selectAnalysis function initiates all the select rule based analysis
func selectAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Selecttag {
	l.Println("Initiating select tag Analysis......")
	nodes := nodeMap["selectNodes"]

	var issues = make(map[string][]Selecttag)
	var list []Selecttag
	for _, node := range nodes {
		var tag Selecttag
		tag.selectRulesWCAG111(node, l)
		l.Printf("selectTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("select: %v\n", list)
	issues["select"] = list

	return issues
}

//iframeAnalysis function initiates all the iframe rule based analysis
func iframeAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Iframetag {
	l.Println("Initiating iFrame tag Analysis......")
	nodes := nodeMap["iframeNodes"]

	var issues = make(map[string][]Iframetag)
	var list []Iframetag
	for _, node := range nodes {
		var tag Iframetag
		tag.iframeRulesWCAG111(node, l)
		l.Printf("iframeTag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("iframe: %v\n", list)
	issues["iframe"] = list

	return issues
}

//linkAnalysis function initiates all the link rule based analysis
func linkAnalysis(l *log.Logger, nodeMap map[string][]*html.Node) map[string][]Anchortag {
	l.Println("Initiating Anchor tag Analysis......")
	nodes := nodeMap["linkNodes"]

	var issues = make(map[string][]Anchortag)
	var list []Anchortag
	for _, node := range nodes {
		var tag Anchortag
		tag.linkRulesWCAG111(node, l)
		l.Printf("Anchortag : %v", tag)
		list = append(list, tag)
	}
	l.Printf("anchor: %v\n", list)
	issues["anchor"] = list

	return issues
}
