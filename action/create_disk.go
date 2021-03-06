package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bslcdisk "github.com/cloudcredo/bosh-lattice-cpi/softlayer/disk"
)

type CreateDisk struct {
	diskCreator bslcdisk.Creator
}

func NewCreateDisk(diskCreator bslcdisk.Creator) CreateDisk {
	return CreateDisk{diskCreator: diskCreator}
}

func (a CreateDisk) Run(size int, instanceId VMCID) (DiskCID, error) {
	disk, err := a.diskCreator.Create(size, instanceId.Int())
	if err != nil {
		return 0, bosherr.WrapErrorf(err, "Creating disk of size '%d'", size)
	}

	return DiskCID(disk.ID()), nil
}
