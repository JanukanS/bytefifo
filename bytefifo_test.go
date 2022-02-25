package bytefifo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	testRR1, err := New(2)
	assert.Empty(err, "Error Initializing")
	assert.Equal(testRR1.baseRing.Value.(resData).rawData, []byte{}, "rawData not initialized properly in ResRing construction")

	_, err = New(0)
	assert.EqualError(err, LengthInitErr.Error(), "Failed to recognize bad length argument during initialization of ResRing")
}

func TestValue(t *testing.T) {
	assert := assert.New(t)

	testRR1, _ := New(2)
	testValue1, err := testRR1.Value()
	assert.Empty(err, "Error detecting expected values")
	assert.Equal(testValue1, []byte{}, "Can't Detect expected initalized value")

	testRR2 := ByteFifo{}
	_, err = testRR2.Value()
	assert.EqualError(err, BaseTypeErr.Error(), "Can't detect bad type for value extraction")
}

func TestTimeStamp(t *testing.T) {
	assert := assert.New(t)

	testRR1, _ := New(2)
	testValue1, err := testRR1.TimeStamp()
	assert.Empty(err, "Error detecting expected values")
	assert.NotEqual(testValue1, time.Time{}, "Initalization used zero values for Time value")

	testRR2 := ByteFifo{}
	_, err = testRR2.TimeStamp()
	assert.EqualError(err, BaseTypeErr.Error(), "Can't detect bad type for timestamp extraction")
}
func TestAdd(t *testing.T) {
	assert := assert.New(t)

	testRR1, err := New(2)
	assert.Empty(err, "Error Initializing Valid ResRing")

	testSlice := []byte{1, 2}
	err = testRR1.Add(testSlice)
	assert.Empty(err, "Error adding valid data")

	err = testRR1.Add(testSlice)
	assert.EqualError(err, SameDataErr.Error(), "Added Duplicate Data in Error")

	testRR2 := ByteFifo{}
	err = testRR2.Add(testSlice)
	assert.EqualError(err, BaseTypeErr.Error(), "Failed to Recognize Improperly Initialized baseRing")

}
