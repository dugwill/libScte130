package scte130

type PlacementStatusEvent struct {
	Text       string   `xml:",chardata"`
	Time       string   `xml:"time,attr"`
	Type       string   `xml:"type,attr"`
	MessageRef string   `xml:"messageRef,attr"`
	SpotNPT    *SpotNPT `xml:"SpotNPT"`
}
