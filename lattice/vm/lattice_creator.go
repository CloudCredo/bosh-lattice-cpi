package vm

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	ltc "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"

	util "github.com/cloudcredo/bosh-lattice-cpi/util"
)

const latticeCreatorLogTag = "LatticeCreator"

type LatticeCreator struct {
	latticeClient ltc.Client
	agentEnvServiceFactory AgentEnvServiceFactory

	agentOptions AgentOptions
	logger       boshlog.Logger
}

func NewLatticeCreator(latticeClient ltc.Client, agentEnvServiceFactory AgentEnvServiceFactory, agentOptions AgentOptions, logger boshlog.Logger) LatticeCreator {
	return LatticeCreator{
		latticeClient:        latticeClient,
		agentEnvServiceFactory: agentEnvServiceFactory,
		agentOptions:           agentOptions,
		logger:                 logger,
	}
}

func (c LatticeCreator) Create(agentID string, cloudProps VMCloudProperties, networks Networks, env Environment) (VM, error) {
	app, err := c.latticeClient.CreateApp()
	if err != nil {
		return LatticeVM{}, bosherr.WrapError(err, "Creating application from Lattice client")
	}

	agentEnvService := c.agentEnvServiceFactory.New(app.Id)

	vm := NewLatticeVM(app.Id, c.latticeClient, util.GetSshClient(), agentEnvService, c.logger)

	return vm, nil
}

func (c LatticeCreator) resolveNetworkIP(networks Networks) (string, error) {
	var network Network

	switch len(networks) {
	case 0:
		return "", bosherr.Error("Expected exactly one network; received zero")
	case 1:
		network = networks.First()
	default:
		return "", bosherr.Error("Expected exactly one network; received multiple")
	}

	if network.IsDynamic() {
		return "", nil
	}

	return network.IP, nil
}
