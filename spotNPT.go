package scte130

type SpotNPT struct {
	Text  string `xml:",chardata"`
	Scale string `xml:"scale,attr"`
}
