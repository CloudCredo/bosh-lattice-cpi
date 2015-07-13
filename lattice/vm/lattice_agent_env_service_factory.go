package vm

import (
	ltc "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type LatticeAgentEnvServiceFactory struct {
	client ltc.Client
	logger boshlog.Logger
}

func NewLatticeAgentEnvServiceFactory(client ltc.Client, logger boshlog.Logger) LatticeAgentEnvServiceFactory {
	return LatticeAgentEnvServiceFactory{
		client: client,
		logger: logger,
	}
}

func (f LatticeAgentEnvServiceFactory) New(vmId int) AgentEnvService {
	return NewLatticeAgentEnvService(vmId, f.client, f.logger)
}
