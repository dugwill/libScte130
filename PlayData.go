package scte130

type PlayData struct {
	Text             string            `xml:",chardata"`
	IdentityADS      string            `xml:"identityADS,attr"`
	SystemContext    *SystemContext    `xml:"SystemContext"`
	Client           *Client           `xml:"Client"`
	SpotScopedEvents *SpotScopedEvents `xml:"SpotScopedEvents"`
}