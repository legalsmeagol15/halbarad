package graph

type depError struct {
	msg string
	src *Dep
}

func NewError(msg string, src *Dep) depError {
	return depError{
		msg: msg,
		src: src,
	}
}

func (e *depError) Equals(other *depError) bool {
	return e.src == other.src && e.msg == other.msg
}
