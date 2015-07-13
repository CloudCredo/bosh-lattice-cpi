package vm

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	util "github.com/cloudcredo/bosh-lattice-cpi/util"
	ltc "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"
)

const latticeFinderLogTag = "LatticeFinder"

type LatticeFinder struct {
	latticeClient        ltc.Client
	agentEnvServiceFactory AgentEnvServiceFactory

	logger boshlog.Logger
}

func NewLatticeFinder(latticeClient ltc.Client, agentEnvServiceFactory AgentEnvServiceFactory, logger boshlog.Logger) LatticeFinder {
	return LatticeFinder{
		latticeClient:        latticeClient,
		agentEnvServiceFactory: agentEnvServiceFactory,

		logger: logger,
	}
}

func (f LatticeFinder) Find(vmID int) (VM, bool, error) {
	app, err := f.latticeClient.FindApp(vmID)
	if err != nil {
		return LatticeVM{}, false, bosherr.WrapErrorf(err, "Finding Lattice Virtual Guest with id `%d`", vmID)
	}

	vm, found := LatticeVM{}, true
	if app.Id == vmID {
		vm = NewLatticeVM(vmID, f.latticeClient, util.GetSshClient(), f.agentEnvServiceFactory.New(vmID), f.logger)
	} else {
		found = false
	}

	return vm, found, nil
}
