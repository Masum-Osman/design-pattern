package chainofresponsibilities

type Department interface {
	execute(*Patient)
	setNext(Department)
}
