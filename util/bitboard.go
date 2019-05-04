package util

const bitsPerInt = 32

// A bit board which can be used to efficiently store 2 dimensional
// data that is only a few bits long.
type BitBoard interface {
	// Set sets the bits at the position x,y.
	// It does nothing if x or y is out of bounds.
	Set(x, y, val int)

	// Get returns the stored value at the position x,y.
	Get(x, y int) int

	// GetBool returns true if the value at x,y is not 0.
	GetBool(x, y int) bool
}

type bitBoardImpl struct {
	bitmask  int
	dataSize int
	xDim     int
	yDim     int
	data     []int
}

func (bb *bitBoardImpl) Set(x, y, val int) {
	sliceIdx, bitOffset := bb.calculateIndex(x, y)
	// mask out only the number of bits allowed
	val = val & bb.bitmask
	shamt := uint(bitsPerInt - bb.dataSize - bitOffset)
	// delete bitmask
	dbm := bb.bitmask << shamt
	dbm = ^dbm
	// shift the data to the right position
	bm := val << shamt

	data := bb.data[sliceIdx]
	// delete old data
	data = data & dbm
	// write new data
	bb.data[sliceIdx] = data | bm
}

func (bb *bitBoardImpl) Get(x, y int) int {
	sliceIdx, bitOffset := bb.calculateIndex(x, y)
	shamt := uint(bitsPerInt - bb.dataSize - bitOffset)
	// calculate bitmask for data retrieval
	bm := bb.bitmask << shamt

	data := bb.data[sliceIdx]
	data = data & bm
	// shift data back to LSB
	data = data >> shamt
	return data
}

func (bb *bitBoardImpl) GetBool(x, y int) bool {
	return bb.Get(x, y) != 0
}

func (bb *bitBoardImpl) calculateIndex(x, y int) (sliceIndex, bitOffset int) {
	bitIndex := (y*bb.xDim + x) * bb.dataSize
	sliceIndex = bitIndex / bitsPerInt
	bitOffset = bitIndex - sliceIndex*bitsPerInt
	return sliceIndex, bitOffset
}

//NewBitBoard creates a new bit board with the given dimensions and
// data size.
// dataSize specifies how many bits each value will hold.
// For implementation reasons valid values for dataSize are 1, 2, 4, 8, 16, 32.
func NewBitBoard(dataSize, xDim, yDim int) BitBoard {
	if dataSize != 1 &&
		dataSize != 2 &&
		dataSize != 4 &&
		dataSize != 8 &&
		dataSize != 16 &&
		dataSize != 32 {
		return nil
	}

	bb := &bitBoardImpl{
		dataSize: dataSize,
		xDim:     xDim,
		yDim:     yDim,
	}

	sliceLen := dataSize * xDim * yDim / bitsPerInt;
	bb.data = make([]int, sliceLen, sliceLen)
	// some nice bit shift magic ƪ(ړײ)‎ƪ​​ to initialize a base bitmask
	bb.bitmask = ^0>>bitsPerInt - dataSize

	return bb
}
