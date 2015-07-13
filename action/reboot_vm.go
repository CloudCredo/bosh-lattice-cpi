package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type RebootVM struct {
	vmFinder bltcvm.Finder
}

func NewRebootVM(vmFinder bltcvm.Finder) RebootVM {
	return RebootVM{vmFinder: vmFinder}
}

func (a RebootVM) Run(vmCID VMCID) (interface{}, error) {
	vm, found, err := a.vmFinder.Find(int(vmCID))
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Finding vm '%s'", vmCID)
	}

	if found {
		err := vm.Reboot()
		if err != nil {
			return nil, bosherr.WrapErrorf(err, "Rebooting vm '%s'", vmCID)
		}
	}

	return nil, nil
}
