package main

type Client struct {
	address     string
	port        string
	forwardPort string
}

func NewClient(address, port, forwardPort string) *Client {
	client := new(Client)
	client.address = address
	client.port = port
	client.forwardPort = forwardPort
	return client
}

func (c *Client) Accept() {

	//c.clientBackendDialer()

	for {

	}
}

/*func (c *Client) clientBackendDialer() (func(net.Dialer, error), error)  {
	toCon, err := net.Dial("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", ))
}*/
