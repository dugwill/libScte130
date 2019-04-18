package scte130

import (
	"encoding/xml"
)

//PSN hols the data from a SCTE130 PlacementStatusNotification message
type PSN struct {
	XMLName        xml.Name  `xml:"PlacementStatusNotification"`
	Text           string    `xml:",chardata"`
	Identity       string    `xml:"identity,attr"`
	MessageID      string    `xml:"messageId,attr"`
	Version        string    `xml:"version,attr"`
	System         string    `xml:"system,attr"`
	Adm            string    `xml:"adm,attr"`
	Core           string    `xml:"core,attr"`
	SchemaLocation string    `xml:"schemaLocation,attr"`
	Xsi            string    `xml:"xsi,attr"`
	Cmcst          string    `xml:"cmcst,attr"`
	PlayData       *PlayData `xml:"PlayData"`
}

// NewPSN returns a pointer to a PSN
func NewPSN() (p *PSN) {
	p = new(PSN)
	return
}

//UnMarshalPSN populates the PSN struct with data from b
func (p *PSN) UnMarshalPSN(b []byte) (err error) {
	err = xml.Unmarshal(b, p)
	if err != nil {
		return err
	}
	return nil
}

//Client returns a pointer to the Client
func (p *PSN) Client() *Client {
	return p.PlayData.Client
}

//Tracking returns the tracking value
func (p *PSN) Tracking() string {
	return p.PlayData.SpotScopedEvents.Spot.Content.Tracking
}

//AssetID returns the assetID
func (p *PSN) AssetID() string {
	return p.PlayData.SpotScopedEvents.Spot.Content.AssetRef.AssetID
}

//NumEvents returns the number of Events from the spotScopedEvents
func (p *PSN) NumEvents() int {
	return len(*p.PlayData.SpotScopedEvents.Events)
}

//GetEvents returns a slice of pointers to the events
func (p *PSN) GetEvents() []Events {
	return *p.PlayData.SpotScopedEvents.Events
}
