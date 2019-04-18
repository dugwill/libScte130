package scte130

type TerminalAddress struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}


