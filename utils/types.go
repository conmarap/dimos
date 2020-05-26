package utils

import (
	"bytes"
	"encoding/binary"
	"math"
	"unsafe"
)

var Endian binary.ByteOrder

// CheckEndian checks the byte order and determines the endiannes.
func CheckEndian() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		Endian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		Endian = binary.BigEndian
	default:
		panic("Unknown native endian.")
	}
}

// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}

	return x
}

// Float64ToByte converts a float64 type number to a byte array.
func Float64ToByte(f float64) ([]byte, error) {
	if Endian == nil {
		CheckEndian()
	}

	var buf bytes.Buffer
	err := binary.Write(&buf, Endian, f)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// BytesToFloat64 converts bytes to a float64.
func BytesToFloat64(bytes []byte) float64 {
	if Endian == nil {
		CheckEndian()
	}

	bits := Endian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// Int64ToBytes converts an int64 to []bytes.
func Int64ToBytes(num int64) []byte {
	if Endian == nil {
		CheckEndian()
	}

	bytes := make([]byte, 8)
	Endian.PutUint64(bytes, uint64(num))

	return bytes
}

// BytesToInt64 converts a byte array to an int64.
func BytesToInt64(bytes []byte) int64 {
	if Endian == nil {
		CheckEndian()
	}

	return Abs(int64(Endian.Uint64(bytes)))
}