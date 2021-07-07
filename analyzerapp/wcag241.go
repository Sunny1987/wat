package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
)

//h1RulesWCAG241 will track 2.4.1 implementation
func (h1 *H1tag) h1RulesWCAG241(node *html.Node, l *log.Logger) {

	//H69
	l.Println("Starting processing : H69")
	if apply, stat := H69(node); apply == Applicable {
		if !stat {
			h1.Wcag241 = "pass"
		} else {
			h1.Wcag241 = "fail"
		}
	}

}

//h2RulesWCAG241 will track 2.4.1 implementation
func (h2 *H2tag) h2RulesWCAG241(node *html.Node, l *log.Logger) {

	//H69
	l.Println("Starting processing : H69")
	if apply, stat := H69(node); apply == Applicable {
		if stat {
			h2.Wcag241 = "pass"
		} else {
			h2.Wcag241 = "fail"
		}
	}

}

/********************************************/
// WCAG2.0 2.4.1 techniques
/********************************************/

//H69 is for header validation
func H69(node *html.Node) (string, bool) {

	if text := nodeText(node); text == "" {
		return Applicable, true
	} else {
		return Applicable, false
	}
}
