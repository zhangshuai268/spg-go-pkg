package sdk

type sdk struct {
	Client clientService
}

func NewSdk() (*sdk, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}
	return &sdk{
		Client: c,
	}, nil
}
