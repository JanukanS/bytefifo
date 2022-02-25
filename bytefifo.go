package resring

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

// ResRing is a a structure similar to ring of resData structs, New()
type ResRing struct {
	baseRing ring.Ring
}

// New creates a ResRing struct of length n
func New(length int) (ResRing, error) {
	if length < 1 {
		return ResRing{}, LengthInitErr
	}
	baseRing := ring.New(length)
	for i := 0; i <= length; i++ {
		baseRing.Value = resData{time.Now(), []byte{}}
		baseRing = baseRing.Next()
	}
	return ResRing{*baseRing}, nil
}

//Add attempts to update the ResRing with new data
func (r *ResRing) Add(newData []byte) error {
	currentData, err := r.Value()
	if err != nil {
		return err
	}
	if bytes.Equal(newData, currentData) {
		return SameDataErr
	}
	r.baseRing = *r.baseRing.Next()
	r.baseRing.Value = resData{time.Now(), newData}
	return nil
}

// Value returns the bytedata at the current position of the ring-like structure
func (r ResRing) Value() ([]byte, error) {
	resDataVal, err := r.extractValue()
	if err != nil {
		return []byte{}, err
	}
	return resDataVal.rawData, nil
}

// TimeStamp returns the time when the resource was added to the ResRing
func (r ResRing) TimeStamp() (time.Time, error) {
	resDataVal, err := r.extractValue()
	if err != nil {
		return time.Time{}, err
	}
	return resDataVal.timeStamp, nil
}

func (r ResRing) extractValue() (resData, error) {
	resDataVal, success := r.baseRing.Value.(resData)
	if success {
		return resDataVal, nil
	}
	return resData{}, BaseTypeErr
}
