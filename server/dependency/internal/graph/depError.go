package graph

type DepError struct {
	msg string
	src *Dep
}

func NewError(msg string, src *Dep) DepError {
	return DepError{
		msg: msg,
		src: src,
	}
}

func (e *DepError) Equals(other *DepError) bool {
	return e.src == other.src && e.msg == other.msg
}
