package client

type latticeClient struct {

}

func NewLatticeClient() *latticeClient {
	return &latticeClient{}
}

func (client *latticeClient) CreateApp() (App, error) {
	return App{}, nil
}

func (client *latticeClient) DeleteApp(vmId int) (bool, error) {
	return true, nil
}

func (client *latticeClient) FindApp(vmId int) (App, error) {
	return App{}, nil
}

func (client *latticeClient) SetTags(vmId int, tags []string) (bool, error) {
	return true, nil
}