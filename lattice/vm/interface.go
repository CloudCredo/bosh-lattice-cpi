package vm

type VMCloudProperties struct {}

type VMMetadata map[string]interface{}

type Creator interface {
	// Create takes an agent id and creates a VM with provided configuration
	Create(string, VMCloudProperties, Networks, Environment) (VM, error)
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
}

type Environment map[string]interface{}
