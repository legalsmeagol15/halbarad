package helpers

type Ntree interface {
}
type ntree struct {
}

type Npoint interface {
}

type npoint struct {
}

func NewPoint(...float64) {

}

func NewNTree(dimension int) Ntree {
	return ntree{}
}
