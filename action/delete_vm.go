package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type DeleteVM struct {
	vmFinder bltcvm.Finder
}

func NewDeleteVM(vmFinder bltcvm.Finder) DeleteVM {
	return DeleteVM{vmFinder: vmFinder}
}

func (a DeleteVM) Run(vmCID VMCID) (interface{}, error) {
	vm, found, err := a.vmFinder.Find(int(vmCID))
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Finding vm '%s'", vmCID)
	}

	if found {
		err := vm.Delete()
		if err != nil {
			return nil, bosherr.WrapErrorf(err, "Deleting vm '%s'", vmCID)
		}
	}

	return nil, nil
}
