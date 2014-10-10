package fakes

import (
	bslcdisk "github.com/maximilien/bosh-softlayer-cpi/softlayer/disk"
)

type FakeVM struct {
	id int

	DeleteCalled bool
	DeleteErr    error

	RebootCalled bool
	RebootErr    error

	AttachDiskDisk bslcdisk.Disk
	AttachDiskErr  error

	DetachDiskDisk bslcdisk.Disk
	DetachDiskErr  error
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

func (vm *FakeVM) AttachDisk(disk bslcdisk.Disk) error {
	vm.AttachDiskDisk = disk
	return vm.AttachDiskErr
}

func (vm *FakeVM) DetachDisk(disk bslcdisk.Disk) error {
	vm.DetachDiskDisk = disk
	return vm.DetachDiskErr
}
