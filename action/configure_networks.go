package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type ConfigureNetworks struct {
	vmFinder bltcvm.Finder
}

func NewConfigureNetworks(vmFinder bltcvm.Finder) ConfigureNetworks {
	return ConfigureNetworks{
		vmFinder: vmFinder,
	}
}

func (a ConfigureNetworks) Run(vmCID VMCID, networks bltcvm.Networks) (interface{}, error) {
	vm, found, err := a.vmFinder.Find(int(vmCID))
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Finding vm '%s'", vmCID)
	}

	if found {
		err := vm.ConfigureNetworks(networks)
		if err != nil {
			return nil, bosherr.WrapErrorf(err, "Configuring networks vm '%s'", vmCID)
		}
	}

	return nil, nil
}
