package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
)

//divRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Divtag) divRulesWCAG111(div *html.Node, l *log.Logger) {

	// creating div object
	d.Div = nodeText(div)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(div); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}


	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(div); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}


//aria6 is for ARIA6 div validation
func aria6(div *html.Node) bool {
	if attributeSearch(div.Attr,"role") {
		if attributeCheckValEmpty(div.Attr,"aria-label") {
			return true
		}
	}
	return false
}

//aria10 is for ARIA10 div validation
func aria10(div *html.Node) bool {

	if hasChildren(div){
		for c := div.FirstChild; c!= nil ; c= c.NextSibling {
			if attributeSearch(c.Attr,"src") {
				if attributeSearch(div.Attr,"aria-labelledby") {
					return false
				}
			}
		}
	}

	return true

}


