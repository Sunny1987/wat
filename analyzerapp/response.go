package analyzerapp

//Response is the response object for json
type Response struct {
	Request         interface{}   `json:"request"`
	Person          string        `json:"person"`
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
	AreaResults     []Areatag     `json:"areas"`
	ObjectResults   []Objecttag   `json:"objects"`
	EmbedResults    []Embedtag    `json:"embeds"`
	TrackResults    []Tracktag    `json:"tracks"`
	H1Results       []H1tag       `json:"h1s"`
	H2Results       []H2tag       `json:"h2s"`
	H3Results       []H3tag       `json:"h3s"`
	H4Results       []H4tag       `json:"h4s"`
	H5Results       []H5tag       `json:"h5s"`
	H6Results       []H6tag       `json:"h6s"`
	PreResults      []Pretag      `json:"pres"`
	CSSResults      []CSS         `json:"CSS"`
}
