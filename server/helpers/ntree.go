package helpers

type Ntree interface {
}
type ntree struct {
	cardinality int
}

func NewNTree(cardinality int) Ntree {
	return ntree{cardinality: cardinality}
}
