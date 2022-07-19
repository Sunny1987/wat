package analyzerapp

import (
	"golang.org/x/net/html"
	"log"
	"strings"
)

//divRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Divtag) divRulesWCAG111(node *html.Node, l *log.Logger) {

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
	//ARIA15
	l.Println("Start processing : ARIA15")
	if apply, stat := ARIA15(node); apply == Applicable {
		if stat {
			d.Wc111Aria15 = "fail"
		} else {
			d.Wc111Aria15 = "pass"
		}
	} else {
		d.Wc111Aria15 = apply
	}

}

//iframeRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Iframetag) iframeRulesWCAG111(node *html.Node, l *log.Logger) {

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
	l.Println("Starting processing : H35")
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

//objectRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Objecttag) objectRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating link object
	d.Object = nodeText(node)

	//H53 implementation
	l.Println("Starting processing : H53")
	if apply, stat := H53(node); apply == Applicable {
		if stat {
			d.Wc111H53 = "pass"
		} else {
			d.Wc111H53 = "fail"
		}
	} else {
		d.Wc111H53 = apply
	}

}

//embedRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Embedtag) embedRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating link object
	d.Embed = nodeText(node)

	//H53 implementation
	l.Println("Starting processing : H53")
	if apply, stat := H53(node); apply == Applicable {
		if stat {
			d.Wc111H53 = "pass"
		} else {
			d.Wc111H53 = "fail"
		}
	} else {
		d.Wc111H53 = apply
	}

}

//preRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Pretag) preRulesWCAG111(node *html.Node, l *log.Logger) {

	// creating link object
	d.Pre = nodeText(node)

	//H86 implementation
	l.Println("Starting processing : H86")
	if apply, stat := H86(node); apply == Applicable {
		if stat {
			d.Wc111H86 = "pass"
		} else {
			d.Wc111H86 = "fail"
		}
	} else {
		d.Wc111H86 = apply
	}

}

func (c *CSS) cssAnalysisWCAG111(css string, l *log.Logger, nodes []*html.Node) {

	//H86 implementation
	l.Println("Starting processing : H86")
	for _, node := range nodes {
		if apply, stat := H86CSS(css, node); apply == Applicable {
			if stat {
				c.WC111H86 = "pass"
			} else {
				c.WC111H86 = "fail"
			}
		} else {
			c.WC111H86 = apply
		}
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

//H35 for applet based implementation
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

//H53 for object based implementation
func H53(node *html.Node) (string, bool) {
	if hasOneChild(node) {
		if node.FirstChild.Data != "p" {
			return Applicable, false
		}
		if node.FirstChild.Data == "img" {
			if attributeCheckValEmpty(node.FirstChild.Attr, "alt") {
				return Applicable, true
			} else {
				return Applicable, false
			}
		}

	}
	return NA, false
}

//ARIA15 for Image based implementation
func ARIA15(node *html.Node) (string, bool) {
	if attributeSearch(node.Attr, "aria-describedby") {
		if attributeCheckValEmpty(node.Attr, "aria-describedby") {
			return Applicable, true
		} else {
			return Applicable, false
		}
	} else {
		return NA, false
	}
}

//H86 for pre based implementation
func H86(node *html.Node) (string, bool) {
	if attributeSearch(node.Attr, "alt") {
		return Applicable, true
	} else {
		return Applicable, false
	}
}

//H86CSS update
func H86CSS(css string, node *html.Node) (string, bool) {
	if strings.Contains(css, "white-space: pre") {
		name := strings.Split(css, " ")
		className := strings.Trim(name[0], "@")
		if attributeCheckVal(node.Attr, "class", className) {
			if attributeSearch(node.Attr, "alt") {
				return Applicable, true
			} else {
				return Applicable, false
			}
		} else {
			return Applicable, false
		}
	}
	return NA, false
}
