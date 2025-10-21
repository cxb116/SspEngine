package interfaces

type IValidator interface {
	Validate(req IBidRequest) (int, error)
}
