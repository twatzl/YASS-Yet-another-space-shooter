package util

import (
	"reflect"
	"testing"
)

func Test_bitBoardImpl_Set(t *testing.T) {
	type fields struct {
		bitmask  int
		dataSize int
		xDim     int
		yDim     int
		data     []int
	}
	type args struct {
		x   int
		y   int
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &bitBoardImpl{
				bitmask:  tt.fields.bitmask,
				dataSize: tt.fields.dataSize,
				xDim:     tt.fields.xDim,
				yDim:     tt.fields.yDim,
				data:     tt.fields.data,
			}
			bb.Set(tt.args.x, tt.args.y, tt.args.val)
		})
	}
}

func Test_bitBoardImpl_Get(t *testing.T) {
	type fields struct {
		bitmask  int
		dataSize int
		xDim     int
		yDim     int
		data     []int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &bitBoardImpl{
				bitmask:  tt.fields.bitmask,
				dataSize: tt.fields.dataSize,
				xDim:     tt.fields.xDim,
				yDim:     tt.fields.yDim,
				data:     tt.fields.data,
			}
			if got := bb.Get(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("bitBoardImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitBoardImpl_GetBool(t *testing.T) {
	type fields struct {
		bitmask  int
		dataSize int
		xDim     int
		yDim     int
		data     []int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &bitBoardImpl{
				bitmask:  tt.fields.bitmask,
				dataSize: tt.fields.dataSize,
				xDim:     tt.fields.xDim,
				yDim:     tt.fields.yDim,
				data:     tt.fields.data,
			}
			if got := bb.GetBool(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("bitBoardImpl.GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitBoardImpl_calculateIndex(t *testing.T) {
	type fields struct {
		bitmask  int
		dataSize int
		xDim     int
		yDim     int
		data     []int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantSliceIndex int
		wantBitOffset  int
	}{
		{name: "TestIndex00For1Bit",
			fields: fields{bitmask: 1, dataSize: 1, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 0, y: 0}, wantSliceIndex: 0, wantBitOffset: 0},
		{name: "TestIndex13For1Bit",
			fields: fields{bitmask: 1, dataSize: 1, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 1, y: 3}, wantSliceIndex: 0, wantBitOffset: 31},
		{name: "TestIndex23For1Bit",
			fields: fields{bitmask: 1, dataSize: 1, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 2, y: 3}, wantSliceIndex: 1, wantBitOffset: 0},
		{name: "TestIndex00For4Bit",
			fields: fields{bitmask: 1, dataSize: 4, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 0, y: 0}, wantSliceIndex: 0, wantBitOffset: 0},
		{name: "TestIndex70For4Bit",
			fields: fields{bitmask: 1, dataSize: 4, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 7, y: 0}, wantSliceIndex: 0, wantBitOffset: 28},
		{name: "TestIndex80For4Bit",
			fields: fields{bitmask: 1, dataSize: 4, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 8, y: 0}, wantSliceIndex: 1, wantBitOffset: 0},
		{name: "TestIndex51For4Bit",
			fields: fields{bitmask: 1, dataSize: 4, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 5, y: 1}, wantSliceIndex: 1, wantBitOffset: 28},
		{name: "TestIndex61For4Bit",
			fields: fields{bitmask: 1, dataSize: 4, xDim: 10, yDim: 10, data: make([]int, 100, 100)},
			args:   args{x: 6, y: 1}, wantSliceIndex: 2, wantBitOffset: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &bitBoardImpl{
				bitmask:  tt.fields.bitmask,
				dataSize: tt.fields.dataSize,
				xDim:     tt.fields.xDim,
				yDim:     tt.fields.yDim,
				data:     tt.fields.data,
			}
			gotSliceIndex, gotBitOffset := bb.calculateIndex(tt.args.x, tt.args.y)
			if gotSliceIndex != tt.wantSliceIndex {
				t.Errorf("bitBoardImpl.calculateIndex() gotSliceIndex = %v, want %v", gotSliceIndex, tt.wantSliceIndex)
			}
			if gotBitOffset != tt.wantBitOffset {
				t.Errorf("bitBoardImpl.calculateIndex() gotBitOffset = %v, want %v", gotBitOffset, tt.wantBitOffset)
			}
		})
	}
}

func TestNewBitBoard(t *testing.T) {
	type args struct {
		dataSize int
		xDim     int
		yDim     int
	}
	tests := []struct {
		name string
		args args
		want BitBoard
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBitBoard(tt.args.dataSize, tt.args.xDim, tt.args.yDim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBitBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
