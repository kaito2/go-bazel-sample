package example

type Client struct{}

func (c *Client) Message() string {
	return "(Updated) Message from example client."
}
