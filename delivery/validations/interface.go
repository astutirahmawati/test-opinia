package validations

type Validation interface {
	Validation(request interface{}) error
}
