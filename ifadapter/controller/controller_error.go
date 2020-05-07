package controller

type unknownReferenceTypeError struct{}

type unexpectedReferenceError struct{}

func (e *unknownReferenceTypeError) Error() string {
	return "unknown reference type"
}

func (e *unexpectedReferenceError) Error() string {
	return "reference is unexpected syntax"
}
