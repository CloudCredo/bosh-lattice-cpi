package action

type CreateStemcell struct {}

type CreateStemcellCloudProps struct {
	Id             int    `json:"virtual-disk-image-id"`
	Uuid           string `json:"virtual-disk-image-uuid"`
	DatacenterName string `json:"datacenter-name"`
}

func NewCreateStemcell() CreateStemcell {
	return CreateStemcell{}
}

func (a CreateStemcell) Run(imagePath string, stemcellCloudProps CreateStemcellCloudProps) (StemcellCID, error) {
	return StemcellCID(0), nil
}
