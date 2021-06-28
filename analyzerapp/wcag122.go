package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
)

//trackRulesWCAG121 will check all WCAG1.2.1 techniques
func (d *Tracktag) trackRulesWCAG122(node *html.Node, l *log.Logger) {

	// creating div object
	d.Track = nodeText(node)

	//H95
	l.Println("Starting processing : H95")
	apply, stat := H95(node)
	if apply == Applicable {
		if stat {
			d.Wcag122G87 = "pass"
		} else {
			d.Wcag122G87 = "fail"
		}
	} else {
		d.Wcag122G87 = apply
	}

}

/********************************************/
// WCAG2.0 1.2.2 techniques
/********************************************/

//H95 for track implementation
func H95(node *html.Node) (string, bool) {
	if node.Parent.Data == "video" {
		if attributeCheckVal(node.Attr, "kind", "captions") {
			return "Applicable", true
		} else {
			return "Applicable", false
		}
	} else {
		return "Not Applicable", false
	}

}
