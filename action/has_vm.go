package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type HasVM struct {
	vmFinder bltcvm.Finder
}

func NewHasVM(vmFinder bltcvm.Finder) HasVM {
	return HasVM{vmFinder: vmFinder}
}

func (a HasVM) Run(vmCID VMCID) (bool, error) {
	_, found, err := a.vmFinder.Find(int(vmCID))
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding VM '%s'", vmCID)
	}

	return found, nil
}
