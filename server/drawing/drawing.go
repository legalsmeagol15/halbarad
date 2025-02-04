package drawing

var (
	Updates = make(chan Drawing, 512)
)

type Drawing interface {
	GetPoints() [][]float64
}
