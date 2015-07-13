package fakes

import (
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type FakeVM struct {
	id int

	DeleteCalled bool
	DeleteErr    error

	RebootCalled bool
	RebootErr    error

	SetMetadataCalled bool
	SetMetadataErr    error
	VMMetadata        bltcvm.VMMetadata

	ConfigureNetworksCalled bool
	ConfigureNetworksErr    error
	Networks                bltcvm.Networks
}

func NewFakeVM(id int) *FakeVM {
	return &FakeVM{id: id}
}

func (vm FakeVM) ID() int { return vm.id }

func (vm *FakeVM) Delete() error {
	vm.DeleteCalled = true
	return vm.DeleteErr
}

func (vm *FakeVM) Reboot() error {
	vm.RebootCalled = true
	return vm.RebootErr
}

func (vm *FakeVM) SetMetadata(metadata bltcvm.VMMetadata) error {
	vm.SetMetadataCalled = true
	vm.VMMetadata = metadata
	return vm.SetMetadataErr
}

func (vm *FakeVM) ConfigureNetworks(networks bltcvm.Networks) error {
	vm.ConfigureNetworksCalled = true
	vm.Networks = networks
	return vm.ConfigureNetworksErr
}