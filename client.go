package scte130

type Client struct {
	Text            string           `xml:",chardata"`
	TerminalAddress *TerminalAddress `xml:"TerminalAddress"`
}

func (c *Client) Type() string {
	return c.TerminalAddress.Type
}

func (c *Client) Address() string {
	return c.TerminalAddress.Text

}
