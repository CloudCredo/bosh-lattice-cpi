package fakes

import (
	bslcvm "github.com/cloudcredo/bosh-lattice-cpi/softlayer/vm"
)

type FakeAgentEnvServiceFactory struct {
	NewAgentEnvService *FakeAgentEnvService
}

func (f *FakeAgentEnvServiceFactory) New(vmId int) bslcvm.AgentEnvService {
	if f.NewAgentEnvService == nil {
		// Always return non-nil service for convenience
		return &FakeAgentEnvService{vmId: vmId}
	}

	return f.NewAgentEnvService
}
