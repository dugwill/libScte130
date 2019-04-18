package scte130

type SpotScopedEvents struct {
	Text   string    `xml:",chardata"`
	Spot   *Spot     `xml:"Spot"`
	Events *[]Events `xml:"Events"`
}
