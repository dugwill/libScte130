package scte130

import "time"

type Events struct {
	Text                 string                `xml:",chardata"`
	PlacementStatusEvent *PlacementStatusEvent `xml:"PlacementStatusEvent"`
}

func (e *Events) PlayDuration() (dur time.Duration, err error) {
	dur, err = time.ParseDuration(e.PlacementStatusEvent.SpotNPT.Text + "s")
	if err != nil {
		return 0, err
	}
	return dur, nil
}

func (e *Events) EventType() string {
	return e.PlacementStatusEvent.Type
}
