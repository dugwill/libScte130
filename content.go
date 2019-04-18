package scte130

type Content struct {
	Text     string    `xml:",chardata"`
	AssetRef *AssetRef `xml:"AssetRef"`
	Tracking string    `xml:"Tracking"`
}
