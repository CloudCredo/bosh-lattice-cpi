package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type CreateVM struct {
	vmCreator         bltcvm.Creator
	vmCloudProperties bltcvm.VMCloudProperties
}

type Environment map[string]interface{}

func NewCreateVM(vmCreator bltcvm.Creator) CreateVM {
	return CreateVM{
		vmCreator:         vmCreator,
		vmCloudProperties: bltcvm.VMCloudProperties{},
	}
}

func (a CreateVM) Run(agentID string, cloudProps bltcvm.VMCloudProperties, networks Networks, env Environment) (VMCID, error) {
	a.updateCloudProperties(cloudProps)

	vmNetworks := networks.AsVMNetworks()

	vmEnv := bltcvm.Environment(env)

	vm, err := a.vmCreator.Create(agentID, cloudProps, vmNetworks, vmEnv)
	if err != nil {
		return 0, bosherr.WrapErrorf(err, "Creating VM with agent ID '%s'", agentID)
	}

	return VMCID(vm.ID()), nil
}

func (a CreateVM) updateCloudProperties(cloudProps bltcvm.VMCloudProperties) {

}
