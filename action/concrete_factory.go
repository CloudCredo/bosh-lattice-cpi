package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	ltc "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type concreteFactory struct {
	availableActions map[string]Action
}

func NewConcreteFactory(latticeClient ltc.Client, options ConcreteFactoryOptions, logger boshlog.Logger) concreteFactory {
	agentEnvServiceFactory := bltcvm.NewLatticeAgentEnvServiceFactory(latticeClient, logger)

	vmCreator := bltcvm.NewLatticeCreator(
		latticeClient,
		agentEnvServiceFactory,
		options.Agent,
		logger,
	)

	vmFinder := bltcvm.NewLatticeFinder(
		latticeClient,
		agentEnvServiceFactory,
		logger,
	)

	return concreteFactory{
		availableActions: map[string]Action{
			// Stemcell management
			"create_stemcell": NewCreateStemcell(),
			"delete_stemcell": NewDeleteStemcell(),

			// VM management
			"create_vm":          NewCreateVM(vmCreator),
			"delete_vm":          NewDeleteVM(vmFinder),
			"has_vm":             NewHasVM(vmFinder),
			"reboot_vm":          NewRebootVM(vmFinder),
			"set_vm_metadata":    NewSetVMMetadata(vmFinder),
			"configure_networks": NewConfigureNetworks(vmFinder),

			// Not implemented (disk related):
			//	 create_disk
			//	 delete_disk
			//	 attach_disk
			//	 detach_disk
			//   snapshot_disk
			//   delete_snapshot
			//   get_disks

			// Not implemented (others):
			//   current_vm_id
			//   ping
		},
	}
}

func (f concreteFactory) Create(method string) (Action, error) {
	action, found := f.availableActions[method]
	if !found {
		return nil, bosherr.Errorf("Could not create action with method %s", method)
	}

	return action, nil
}
