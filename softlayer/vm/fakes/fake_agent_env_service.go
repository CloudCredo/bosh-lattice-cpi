package fakes

import (
	bslcvm "github.com/cloudcredo/bosh-lattice-cpi/softlayer/vm"
)

type FakeAgentEnvService struct {
	FetchCalled   bool
	FetchAgentEnv bslcvm.AgentEnv
	FetchErr      error

	UpdateAgentEnv bslcvm.AgentEnv
	UpdateErr      error

	vmId int
}

func (s *FakeAgentEnvService) Fetch() (bslcvm.AgentEnv, error) {
	s.FetchCalled = true
	return s.FetchAgentEnv, s.FetchErr
}

func (s *FakeAgentEnvService) Update(agentEnv bslcvm.AgentEnv) error {
	s.UpdateAgentEnv = agentEnv
	return s.UpdateErr
}
