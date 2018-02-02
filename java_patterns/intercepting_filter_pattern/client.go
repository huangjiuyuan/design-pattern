package filter

type Client struct {
	fm *FilterManager
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetFilterManager(fm *FilterManager) {
	c.fm = fm
}

func (c *Client) SendRequest(req string) {
	c.fm.FilterRequest(req)
}
