package controller

type unKnownReferenceTypeError struct{}

type unExpectedReferenceError struct{}

func (e *unKnownReferenceTypeError) Error() string {
	return "unknown reference type"
}

func (e *unExpectedReferenceError) Error() string {
	return "reference is unexpected syntax"
}
