package scte130

type AssetRef struct {
	Text       string `xml:",chardata"`
	AssetID    string `xml:"assetID,attr"`
	ProviderID string `xml:"providerID,attr"`
}
