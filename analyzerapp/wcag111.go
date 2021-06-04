package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
)

//divRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Divtag) divRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Div = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//buttonRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Buttontag) buttonRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Button = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//inputRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Inputtag) inputRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Input = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//videoRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Videotag) videoRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Video = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//audioRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Audiotag) audioRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Audio = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//selectRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Selecttag) selectRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Select = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//textAreaRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Textareatag) textAreaRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Textarea = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//imagesRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Imgtag) imagesRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Img = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

}

//aria6 is for ARIA6 div validation
func aria6(node *html.Node) bool {
	if attributeSearch(node.Attr, "role") {
		if attributeCheckValEmpty(node.Attr, "aria-label") {
			return true
		}
	}
	return false
}

//aria10 is for ARIA10 div validation
func aria10(node *html.Node) bool {

	if hasChildren(node) {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if attributeSearch(c.Attr, "src") {
				if attributeSearch(node.Attr, "aria-labelledby") {
					return false
				}
			}
		}
	}

	return true

}
