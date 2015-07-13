package vm

import (
	"errors"

	ltc "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

const (
	softLayerAgentEnvServiceLogTag = "softLayerAgentEnvService"

	softLayerAgentEnvServiceSettingsFileName  = "softlayer-cpi-agent-env.json"
	softLayerAgentEnvServiceTmpSettingsPath   = "/tmp/" + softLayerAgentEnvServiceSettingsFileName
	softLayerAgentEnvServiceFinalSettingsPath = "/var/vcap/bosh/" + softLayerAgentEnvServiceSettingsFileName
)

type LatticeAgentEnvService struct {
	vmId            int
	latticeClient   ltc.Client
	logger          boshlog.Logger
}

func NewLatticeAgentEnvService(vmId int, latticeClient ltc.Client, logger boshlog.Logger) LatticeAgentEnvService {
	return LatticeAgentEnvService{
		vmId:            vmId,
		latticeClient:   latticeClient,
		logger:          logger,
	}
}

func (s LatticeAgentEnvService) Fetch() (AgentEnv, error) {
	return AgentEnv{}, errors.New("Implement me!")
}

func (s LatticeAgentEnvService) Update(agentEnv AgentEnv) error {
	return errors.New("Implement me!")
}
