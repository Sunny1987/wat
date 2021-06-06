package analyzerapp

//Response is the response object for json
type Response struct {
	//Request parser.MyURLReq `json:"request"`
	DivResults      []Divtag      `json:"divs"`
	ButtonResults   []Buttontag   `json:"buttons"`
	InputResults    []Inputtag    `json:"inputs"`
	ImageResults    []Imgtag      `json:"images"`
	VideoResults    []Videotag    `json:"videos"`
	AudioResults    []Audiotag    `json:"audios"`
	TextareaResults []Textareatag `json:"textAreas"`
	SelectResults   []Selecttag   `json:"selects"`
	ParaResults     []Paratag     `json:"paras"`
	IframeResults   []Iframetag   `json:"iframes"`
	LinkResults     []Anchortag   `json:"links"`
}
