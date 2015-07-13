package fakes

import ltcclient "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"

type FakeLatticeClient struct {

}

func NewFakeLatticeClient() *FakeLatticeClient {
	return &FakeLatticeClient{}
}

func (client FakeLatticeClient) CreateApp() (ltcclient.App, error) {
	return ltcclient.App{}, nil
}

func (client FakeLatticeClient) DeleteApp(vmId int) (bool, error) {
	return true, nil
}

func (client FakeLatticeClient) FindApp(vmId int) (ltcclient.App, error) {
	return ltcclient.App{}, nil
}

func (client FakeLatticeClient) SetTags(vmId int, tags []string) (bool, error) {
	return true, nil
}