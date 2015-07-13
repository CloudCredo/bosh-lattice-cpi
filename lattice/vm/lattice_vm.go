package vm

import (
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	util "github.com/cloudcredo/bosh-lattice-cpi/util"

	ltc "github.com/cloudcredo/bosh-lattice-cpi/lattice/client"
)

const (
	latticeVMtag = "LatticeVM"
	ROOT_USER_NAME = "root"
)

type LatticeVM struct {
	id int

	latticeClient ltc.Client

	agentEnvService AgentEnvService

	sshClient util.SshClient

	logger boshlog.Logger
}

func NewLatticeVM(id int, latticeClient ltc.Client, sshClient util.SshClient, agentEnvService AgentEnvService, logger boshlog.Logger) LatticeVM {
	return LatticeVM{
		id: id,

		latticeClient: latticeClient,

		agentEnvService: agentEnvService,

		sshClient: sshClient,

		logger: logger,
	}
}

func (vm LatticeVM) ID() int { return vm.id }

func (vm LatticeVM) Delete() error {
	deleted, err := vm.latticeClient.DeleteApp(vm.ID())
	if err != nil {
		return bosherr.WrapError(err, "Deleting Lattice VirtualGuest from client")
	}

	if !deleted {
		return bosherr.WrapError(nil, "Did not delete Lattice VirtualGuest from client")
	}

	return nil
}

func (vm LatticeVM) Reboot() error {

	return nil
}

func (vm LatticeVM) SetMetadata(vmMetadata VMMetadata) error {
	tags, err := vm.extractTagsFromVMMetadata(vmMetadata)
	if err != nil {
		return err
	}

	if len(tags) == 0 {
		return nil
	}

	//Check below needed since Golang strings.Split return [""] on strings.Split("", ",")
	if len(tags) == 1 && tags[0] == "" {
		return nil
	}

	success, err := vm.latticeClient.SetTags(vm.ID(), tags)
	if !success {
		return bosherr.WrapErrorf(err, "Settings tags on Lattice VirtualGuest `%d`", vm.ID())
	}

	if err != nil {
		return bosherr.WrapErrorf(err, "Settings tags on Lattice VirtualGuest `%d`", vm.ID())
	}

	return nil
}

func (vm LatticeVM) ConfigureNetworks(networks Networks) error {
	return NotSupportedError{}
}

// Private methods
func (vm LatticeVM) extractTagsFromVMMetadata(vmMetadata VMMetadata) ([]string, error) {
	tags := []string{}
	for key, value := range vmMetadata {
		if key == "tags" {
			stringValue, ok := value.(string)
			if !ok {
				return []string{}, bosherr.Errorf("Could not convert tags metadata value `%v` to string", value)
			}

			tags = vm.parseTags(stringValue)
		}
	}

	return tags, nil
}

func (vm LatticeVM) parseTags(value string) []string {
	return strings.Split(value, ",")
}
