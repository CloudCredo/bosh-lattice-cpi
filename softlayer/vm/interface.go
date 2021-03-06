package vm

import (
	bslcdisk "github.com/cloudcredo/bosh-lattice-cpi/softlayer/disk"
	bslcstem "github.com/cloudcredo/bosh-lattice-cpi/softlayer/stemcell"

	sldatatypes "github.com/maximilien/softlayer-go/data_types"
)

type VMCloudProperties struct {
	Domain                   string `json:"domain,omitempty"`
	StartCpus                int    `json:"startCpus,omitempty"`
	MaxMemory                int    `json:"maxMemory,omitempty"`
	Datacenter               sldatatypes.Datacenter
	BlockDeviceTemplateGroup sldatatypes.BlockDeviceTemplateGroup
	SshKeys                  []sldatatypes.SshKey `json:"sshKeys"`
	RootDiskSize             int                  `json:"rootDiskSize,omitempty"`
	EphemeralDiskSize        int                  `json:"ephemeralDiskSize,omitempty"`
}

type VMMetadata map[string]interface{}

type Creator interface {
	// Create takes an agent id and creates a VM with provided configuration
	Create(string, bslcstem.Stemcell, VMCloudProperties, Networks, Environment) (VM, error)
}

type Finder interface {
	Find(int) (VM, bool, error)
}

type VM interface {
	ID() int

	Delete() error
	Reboot() error

	SetMetadata(VMMetadata) error
	ConfigureNetworks(Networks) error

	AttachDisk(bslcdisk.Disk) error
	DetachDisk(bslcdisk.Disk) error
}

type Environment map[string]interface{}
