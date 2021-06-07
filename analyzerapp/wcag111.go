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
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//buttonRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Buttontag) buttonRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating button object
	d.Button = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//inputRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Inputtag) inputRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating input object
	d.Input = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//videoRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Videotag) videoRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating video object
	d.Video = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//audioRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Audiotag) audioRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating audio object
	d.Audio = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//selectRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Selecttag) selectRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating select object
	d.Select = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//textAreaRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Textareatag) textAreaRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating textarea object
	d.Textarea = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//imagesRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Imgtag) imagesRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating image object
	d.Img = nodeText(node)

	//ARIA6
	apply, stat := aria6(node)
	if apply == Applicable {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Start processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

	//G94
	l.Println("Start processing : G94 ")
	if apply, stat := G94(node); apply == Applicable {
		if !stat {
			d.Wc111G94 = "fail"
		} else {
			d.Wc111G94 = "pass"
		}
	} else {
		d.Wc111G94 = apply
	}

	//H2
	l.Println("Start processing : H2")
	if apply, stat := H2(node); apply == Applicable {
		if stat {
			d.Wc111H2 = "fail"
		} else {
			d.Wc111H2 = "pass"
		}
	} else {
		d.Wc111H2 = apply
	}
}

//iframeRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Iframetag) iframeRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating iframe object
	d.Iframe = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")
	apply, stat := aria6(node)
	if apply == "A" {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//linkRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Anchortag) linkRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating link object
	d.Anchor = nodeText(node)

	//ARIA6
	l.Println("Starting processing : ARIA6")

	apply, stat := aria6(node)

	if apply == "A" {
		if stat {
			d.Wc111Aria6 = "fail"
		} else {
			d.Wc111Aria6 = "pass"
		}
	} else {
		d.Wc111Aria6 = apply
	}

	//ARIA10
	l.Println("Starting processing : ARIA10")
	if apply, stat := aria10(node); apply == Applicable {
		if stat {
			d.Wc111Aria10 = "fail"
		} else {
			d.Wc111Aria10 = "pass"
		}
	} else {
		d.Wc111Aria10 = apply
	}

}

//areaRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Areatag) areaRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating link object
	d.Area = nodeText(node)

	//G94
	l.Println("Start processing : G94 ")
	if apply, stat := G94(node); apply == Applicable {
		if !stat {
			d.Wc111G94 = "fail"
		} else {
			d.Wc111G94 = "pass"
		}
	} else {
		d.Wc111G94 = apply
	}

	//H2
	l.Println("Start processing : H2")
	if apply, stat := H2(node); apply == Applicable {
		if stat {
			d.Wc111H2 = "fail"
		} else {
			d.Wc111H2 = "pass"
		}
	} else {
		d.Wc111H2 = apply
	}

}

//appletRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Applettag) appletRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating link object
	d.Applet = nodeText(node)

	//H35 implementation
	if apply, stat := H35(node); apply == Applicable {
		if stat {
			d.Wc111H35 = "pass"
		} else {
			d.Wc111H35 = "fail"
		}
	} else {
		d.Wc111H35 = apply
	}

}

/********************************************/
// WCAG2.0 1.1.1 techniques
/********************************************/

//aria6 is for ARIA6 validation
func aria6(node *html.Node) (string, bool) {
	if attributeSearch(node.Attr, "role") {
		if attributeCheckValEmpty(node.Attr, "aria-label") {
			return Applicable, true
		} else {
			return Applicable, false
		}
	}
	return NA, false
}

//aria10 is for ARIA10 validation
func aria10(node *html.Node) (string, bool) {

	if hasChildren(node) {
		if attributeCheckValEmpty(node.Attr, "aria-labelledby") {
			for c := node.FirstChild; c != nil; c = c.NextSibling {
				if attributeSearch(c.Attr, "src") {
					if attributeCheckValEmpty(node.Attr, "alt") {
						return Applicable, true
					} else {
						return Applicable, false
					}
				}
			}
		}

	}

	return NA, true

}

//G94 is for image tag alt attribute validation
func G94(node *html.Node) (string, bool) {
	if attributeCheckValEmpty(node.Attr, "alt") {
		return Applicable, true
	}
	if attributeCheckVal(node.Attr, "role", "img") {
		if attributeCheckValEmpty(node.Attr, "aria-labelledby") {
			return Applicable, true
		} else {
			return Applicable, false
		}
	}
	if attributeCheckVal(node.Attr, "role", "img") {
		if attributeCheckValEmpty(node.Attr, "aria-label") {
			return Applicable, true
		} else {
			return NA, false
		}
	}
	if attributeCheckValEmpty(node.Attr, "title") {
		return Applicable, true
	}
	if !attributeCheckValEmpty(node.Attr, "alt") && attributeCheckVal(node.Attr, "role", "presentation") {
		return Applicable, true
	}

	return Applicable, false
}

//H2 is for image tag alt attribute validation
func H2(node *html.Node) (string, bool) {
	if node.Parent.Data == "a" {
		if attributeCheckValEmpty(node.Attr, "alt") {
			return Applicable, true
		} else {
			return Applicable, false
		}
	}
	return NA, false
}

func H35(node *html.Node) (string, bool) {
	if hasOneChild(node) {
		if isTextNode(node.FirstChild) {
			if attributeCheckValEmpty(node.Attr, "alt") {
				return Applicable, true
			} else {
				return Applicable, false
			}
		}
	}
	return NA, false

}
