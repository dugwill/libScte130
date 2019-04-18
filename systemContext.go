package scte130

type SystemContext struct {
	Text    string `xml:",chardata"`
	Session string `xml:"Session"`
}
