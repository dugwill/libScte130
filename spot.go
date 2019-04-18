package scte130

type Spot struct {
	Text    string   `xml:",chardata"`
	Content *Content `xml:"Content"`
}
