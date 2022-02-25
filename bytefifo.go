package bytefifo

import (
	"bytes"
	"container/ring"
	"time"
)

// resData is a struct that holds a time stamp and the raw byte data of a resource
type resData struct {
	timeStamp time.Time
	rawData   []byte
}

// ByteFifo is a a structure similar to ring of resData structs, New()
type ByteFifo struct {
	baseRing ring.Ring
}

// New creates a ByteFifo struct of length n
func New(length int) (ByteFifo, error) {
	if length < 1 {
		return ByteFifo{}, LengthInitErr
	}
	baseRing := ring.New(length)
	for i := 0; i <= length; i++ {
		baseRing.Value = resData{time.Now(), []byte{}}
		baseRing = baseRing.Next()
	}
	return ByteFifo{*baseRing}, nil
}

//Add attempts to update the ByteFifo with new data
func (b *ByteFifo) Add(newData []byte) error {
	currentData, err := b.Value()
	if err != nil {
		return err
	}
	if bytes.Equal(newData, currentData) {
		return SameDataErr
	}
	b.baseRing = *b.baseRing.Next()
	b.baseRing.Value = resData{time.Now(), newData}
	return nil
}

// Value returns the bytedata at the current position of the ring-like structure
func (b ByteFifo) Value() ([]byte, error) {
	resDataVal, err := b.extractValue()
	if err != nil {
		return []byte{}, err
	}
	return resDataVal.rawData, nil
}

// TimeStamp returns the time when the resource was added to the ByteFifo
func (b ByteFifo) TimeStamp() (time.Time, error) {
	resDataVal, err := b.extractValue()
	if err != nil {
		return time.Time{}, err
	}
	return resDataVal.timeStamp, nil
}

func (b ByteFifo) extractValue() (resData, error) {
	resDataVal, success := b.baseRing.Value.(resData)
	if success {
		return resDataVal, nil
	}
	return resData{}, BaseTypeErr
}
