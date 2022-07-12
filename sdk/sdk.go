package sdk

type Sdk struct {
	client Client
}

func NewSdk() (*Sdk, error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	return &Sdk{
		client: c,
	}, nil
}
