package resring

type resRingError struct {
	message string
}

//Error returns string for error output
func (s resRingError) Error() string {
	return s.message
}

var (
	//LengthInitErr occurs when initializing ResRing with an invalid amount of elements
	LengthInitErr = resRingError{"Length of ResRing should be greater than 0"}
	//BaseTypeErr occurs when the internal ring structure of ResRing contains elements of the wrong type
	BaseTypeErr = resRingError{"Internal Ring Value Type incorrect, ResRing instance may not be generated from resring.New()"}
	//SameDataErr occurs when attempting to add data to the ResRing that is identical to the current data on top
	SameDataErr = resRingError{"New and Current data are identical"}
)
