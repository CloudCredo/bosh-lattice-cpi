package fakes

import (
	bslcdisk "github.com/CloudCredo/bosh-lattice-cpi/softlayer/disk"
)

type FakeFinder struct {
	FindID    int
	FindDisk  bslcdisk.Disk
	FindFound bool
	FindErr   error
}

func (f *FakeFinder) Find(id int) (bslcdisk.Disk, bool, error) {
	f.FindID = id
	return f.FindDisk, f.FindFound, f.FindErr
}
