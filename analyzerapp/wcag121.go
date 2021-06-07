package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
)

const (
	Applicable = "Applicable"
	NA         = "Not Applicable"
)

//audioRulesWCAG121 will check all WCAG1.2.1 techniques
func (d *Audiotag) audioRulesWCAG121(node *html.Node, l *log.Logger) {

	// creating audio object
	if d.Audio != "" {
		d.Audio = nodeText(node)
	}

}

//videoRulesWCAG121 will check all WCAG1.2.1 techniques
func (d *Videotag) videoRulesWCAG121(node *html.Node, l *log.Logger) {

	// creating video object
	if d.Video != "" {
		d.Video = nodeText(node)
	}

}

//objectRulesWCAG121 will check all WCAG1.2.1 techniques
func (d *Objecttag) objectRulesWCAG121(node *html.Node, l *log.Logger) {

	// creating object object
	d.Object = nodeText(node)

}

//embedRulesWCAG121 will check all WCAG1.2.1 techniques
func (d *Embedtag) embedRulesWCAG121(node *html.Node, l *log.Logger) {

	// creating div object
	d.Embed = nodeText(node)

}

//trackRulesWCAG121 will check all WCAG1.2.1 techniques
func (d *Tracktag) trackRulesWCAG121(node *html.Node, l *log.Logger) {

	// creating div object
	d.Track = nodeText(node)

	//H96
	l.Println("Starting processing : H96")
	apply, stat := H96(node)
	if apply == Applicable {
		if stat {
			d.Wc121H96 = "pass"
		} else {
			d.Wc121H96 = "fail"
		}
	} else {
		d.Wc121H96 = apply
	}

}

/********************************************/
// WCAG2.0 1.2.1 techniques
/********************************************/

func H96(node *html.Node) (string, bool) {
	if node.Parent.Data == "audio" {
		if (attributeCheckVal(node.Attr, "kind", "captions") ||
			attributeCheckVal(node.Attr, "kind", "descriptions")) &&
			attributeCheckValEmpty(node.Attr, "label") {
			return "Applicable", true
		} else {
			return "Applicable", false
		}
	} else {
		return "Not Applicable", false
	}

}
