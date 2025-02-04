package dependency

type depError struct {
	source  Dependent
	message string
}
