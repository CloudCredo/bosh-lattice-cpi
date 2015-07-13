package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bltcvm "github.com/cloudcredo/bosh-lattice-cpi/lattice/vm"
)

type ConcreteFactoryOptions struct {
	StemcellsDir string

	Agent bltcvm.AgentOptions
}

func (o ConcreteFactoryOptions) Validate() error {
	if o.StemcellsDir == "" {
		return bosherr.Error("Must provide non-empty StemcellsDir")
	}

	err := o.Agent.Validate()
	if err != nil {
		return bosherr.WrapError(err, "Validating Agent configuration")
	}

	return nil
}
