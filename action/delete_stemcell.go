package action

type DeleteStemcell struct {
}

func NewDeleteStemcell() DeleteStemcell {
	return DeleteStemcell{}
}

func (a DeleteStemcell) Run(stemcellCID StemcellCID) (interface{}, error) {
	return nil, nil
}
