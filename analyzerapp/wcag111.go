package analyzerapp

import "golang.org/x/net/html"

var results map[string]*html.Node

//divRulesWCAG111 will check all WCAG1.1.1 techniques
func (d *Divtag) divRulesWCAG111(div *html.Node,index int) {

	//ARIA6
	if stat,node := aria6(div,index); stat {
		d.Wc111Aria6 = node
	}

}


//aria6 is for ARIA6 div validation
func aria6(div *html.Node,index int) (bool,*html.Node) {
	if attributeSearch(div.Attr,"role") {
		if attributeCheckValEmpty(div.Attr,"aria-label") {
			return true,div
		}
	}
	return false,nil
}
