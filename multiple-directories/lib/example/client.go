package example

type Client struct {}

func (c *Client) Message() string {
	return "Message from example client."
}
