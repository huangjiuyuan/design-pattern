package delegate

type Client struct {
	Delegate *BusinessDelegate
}

func NewClient(delegate *BusinessDelegate) *Client {
	return &Client{
		Delegate: delegate,
	}
}

func (c *Client) DoTask() {
	c.Delegate.DoTask()
}
