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
	l.Println("Start processing : ARIA6")
	if stat := aria6(node); stat {
		d.Wc111Aria6 = "fail"
	} else {
		d.Wc111Aria6 = "pass"
	}

	//ARIA10
	l.Println("Start processing : ARIA10")
	if stat := aria10(node); !stat {
		d.Wc111Aria10 = "fail"
	} else {
		d.Wc111Aria10 = "pass"
	}

	//G94
	l.Println("Start processing : G94 ")
	if stat := G94(node); !stat {
		d.Wc111G94 = "fail"
	} else {
		d.Wc111G94 = "pass"
	}

	//H2
	l.Println("Start processing : H2")
	if stat := H2(node); stat {
		d.Wc111H2 = "fail"
	} else {
		d.Wc111H2 = "pass"
	}
}

//iframeRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Iframetag) iframeRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Iframe = nodeText(node)

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

//linkRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Anchortag) linkRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating div object
	d.Anchor = nodeText(node)

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

/********************************************/
// WCAG2.0 1.1.1 techniques
/********************************************/

//aria6 is for ARIA6 validation
func aria6(node *html.Node) bool {
	if attributeSearch(node.Attr, "role") {
		if attributeCheckValEmpty(node.Attr, "aria-label") {
			return true
		}
	}
	return false
}

//aria10 is for ARIA10 validation
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

//G94 is for image tag alt attribute validation
func G94(node *html.Node) bool {
	if attributeCheckValEmpty(node.Attr, "alt") {
		return true
	}
	return false
}

//H2 is for image tag alt attribute validation
func H2(node *html.Node) bool {
	if node.Parent.Data == "a" {
		if attributeCheckValEmpty(node.Attr, "alt") {
			return true
		}
	}
	return false
}
