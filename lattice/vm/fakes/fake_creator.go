package fakes

import (
	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type FakeCreator struct {
	CreateAgentID           string
	CreateNetworks          bltcvm.Networks
	CreateVMCloudProperties bltcvm.VMCloudProperties
	CreateEnvironment       bltcvm.Environment
	CreateVM                bltcvm.VM
	CreateErr               error
}

func (c *FakeCreator) Create(agentID string, vmCloudProperties bltcvm.VMCloudProperties, networks bltcvm.Networks, env bltcvm.Environment) (bltcvm.VM, error) {
	c.CreateAgentID = agentID
	c.CreateVMCloudProperties = vmCloudProperties
	c.CreateNetworks = networks
	c.CreateEnvironment = env
	return c.CreateVM, c.CreateErr
}
