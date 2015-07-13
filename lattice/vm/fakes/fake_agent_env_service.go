package fakes

import (
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type FakeAgentEnvService struct {
	FetchCalled   bool
	FetchAgentEnv bltcvm.AgentEnv
	FetchErr      error

	UpdateAgentEnv bltcvm.AgentEnv
	UpdateErr      error

	vmId int
}

func (s *FakeAgentEnvService) Fetch() (bltcvm.AgentEnv, error) {
	s.FetchCalled = true
	return s.FetchAgentEnv, s.FetchErr
}

func (s *FakeAgentEnvService) Update(agentEnv bltcvm.AgentEnv) error {
	s.UpdateAgentEnv = agentEnv
	return s.UpdateErr
}
