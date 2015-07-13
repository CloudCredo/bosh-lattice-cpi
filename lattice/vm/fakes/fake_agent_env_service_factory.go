package fakes

import (
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type FakeAgentEnvServiceFactory struct {
	NewAgentEnvService *FakeAgentEnvService
}

func (f *FakeAgentEnvServiceFactory) New(vmId int) bltcvm.AgentEnvService {
	if f.NewAgentEnvService == nil {
		// Always return non-nil service for convenience
		return &FakeAgentEnvService{vmId: vmId}
	}

	return f.NewAgentEnvService
}
