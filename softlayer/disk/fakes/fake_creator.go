package fakes

import (
	bslcdisk "github.com/CloudCredo/bosh-lattice-cpi/softlayer/disk"
)

type FakeCreator struct {
	CreateSize int
	CreateDisk bslcdisk.Disk
	CreateErr  error
}

func (c *FakeCreator) Create(size int, virtualGuestId int) (bslcdisk.Disk, error) {
	c.CreateSize = size
	return c.CreateDisk, c.CreateErr
}
