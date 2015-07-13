package fakes

import (
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type FakeFinder struct {
	FindID    int
	FindVM    bltcvm.VM
	FindFound bool
	FindErr   error
}

func (f *FakeFinder) Find(id int) (bltcvm.VM, bool, error) {
	f.FindID = id
	return f.FindVM, f.FindFound, f.FindErr
}
