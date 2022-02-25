package bytefifo

type byteFifoError struct {
	message string
}

//Error returns string for error output
func (s byteFifoError) Error() string {
	return s.message
}

var (
	//LengthInitErr occurs when initializing ByteFifo with an invalid amount of elements
	LengthInitErr = byteFifoError{"Length of ResRing should be greater than 0"}
	//BaseTypeErr occurs when the internal ring structure of ByteFifo contains elements of the wrong type
	BaseTypeErr = byteFifoError{"Internal Ring Value Type incorrect, ResRing instance may not be generated from resring.New()"}
	//SameDataErr occurs when attempting to add data to the ByteFifo that is identical to the current data on top
	SameDataErr = byteFifoError{"New and Current data are identical"}
)
