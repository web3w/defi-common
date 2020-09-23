// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package unpack

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
)

// unpacker is a utility interface that enables us to have
// abstraction between events and methods and also to properly
// "unpack" them; e.g. events use Inputs, methods use Outputs.
type unpacker interface {
	tupleUnpack(v interface{}, output []byte) error
	singleUnpack(v interface{}, output []byte) error
	isTupleReturn() bool
}

// reads the integer based on its kind
func readInteger(kind reflect.Kind, b []byte) interface{} {
	switch kind {
	case reflect.Uint8:
		return b[len(b)-1]
	case reflect.Uint16:
		return binary.BigEndian.Uint16(b[len(b)-2:])
	case reflect.Uint32:
		return binary.BigEndian.Uint32(b[len(b)-4:])
	case reflect.Uint64:
		return binary.BigEndian.Uint64(b[len(b)-8:])
	case reflect.Int8:
		return int8(b[len(b)-1])
	case reflect.Int16:
		return int16(binary.BigEndian.Uint16(b[len(b)-2:]))
	case reflect.Int32:
		return int32(binary.BigEndian.Uint32(b[len(b)-4:]))
	case reflect.Int64:
		return int64(binary.BigEndian.Uint64(b[len(b)-8:]))
	default:
		return new(big.Int).SetBytes(b)
	}
}

var (
	errBadBool = errors.New("abi: improperly encoded boolean value")
)

// reads a bool
func readBool(word []byte) (bool, int, error) {
	for _, b := range word[:31] {
		if b != 0 {
			return false, 0, errBadBool
		}
	}
	switch word[31] {
	case 0:
		return false, 0, nil
	case 1:
		return true, 1, nil
	default:
		return false, 0, errBadBool
	}
}

// A function type is simply the address with the function selection signature at the end.
// This enforces that standard by always presenting it as a 24-array (address + sig = 24 bytes)
func readFunctionType(t abi.Type, word []byte) (funcTy [24]byte, size int, err error) {
	if t.T != abi.FunctionTy {
		return [24]byte{}, 0, fmt.Errorf("abi: invalid type in call to make function type byte array.")
	}
	if garbage := binary.BigEndian.Uint64(word[24:32]); garbage != 0 {
		err = fmt.Errorf("abi: got improperly encoded function type, got %v", word)
	} else {
		copy(funcTy[:], word[0:24])
	}
	return funcTy, 1, nil
}

// through reflection, creates a fixed array to be read from
func readFixedBytes(t abi.Type, word []byte) (interface{}, int, error) {
	if t.T != abi.FixedBytesTy {
		return nil, 0, fmt.Errorf("abi: invalid type in call to make fixed byte array.")
	}
	// convert
	array := reflect.New(t.Type).Elem()

	reflect.Copy(array, reflect.ValueOf(word[0:t.Size]))
	return array.Interface(), 1, nil

}

// iteratively unpack elements
func forEachUnpack(t abi.Type, output []byte, start, size int) (interface{}, int, error) {
	if start+32*size > len(output) {
		return nil, 0, fmt.Errorf("abi: cannot marshal in to go array: offset %d would go over slice boundary (len=%d)", len(output), start+32*size)
	}

	// this value will become our slice or our array, depending on the type
	var refSlice reflect.Value
	slice := output[start : start+size*32]

	if t.T == abi.SliceTy {
		// declare our slice
		refSlice = reflect.MakeSlice(t.Type, size, size)
	} else if t.T == abi.ArrayTy {
		// declare our array
		refSlice = reflect.New(t.Type).Elem()
	} else {
		return nil, 0, fmt.Errorf("abi: invalid type in array/slice unpacking stage")
	}

	for i, j := start, 0; j*32 < len(slice); i, j = i+32, j+1 {
		// this corrects the arrangement so that we get all the underlying array values
		if t.Elem.T == abi.ArrayTy && j != 0 {
			i = start + t.Elem.Size*32*j
		}
		inter, _, err := ToGoType(i, *t.Elem, output)
		if err != nil {
			return nil, 0, err
		}
		// append the item to our reflect slice
		refSlice.Index(j).Set(reflect.ValueOf(inter))
	}

	// return the interface
	return refSlice.Interface(), size, nil
}

// ToGoType modifies original toGoType to return size it processed
//
// Warning: array data (corresponding to SliceTy and ArrayTy) unpacking has not been tested.
func ToGoType(index int, t abi.Type, output []byte) (interface{}, int, error) {
	if index+32 > len(output) {
		return nil, 0, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), index+32)
	}

	var (
		returnOutput []byte
		begin, end   int
		err          error
	)

	// if we require a length prefix, find the beginning word and size returned.

	if requiresLengthPrefix(t) {
		begin, end, err = lengthPrefixPointsTo(index, output)
		if err != nil {
			return nil, 0, err
		}
	} else {
		returnOutput = output[index : index+32]
	}

	switch t.T {
	case abi.SliceTy:
		return forEachUnpack(t, output, begin, end)
	case abi.ArrayTy:
		return forEachUnpack(t, output, index, t.Size)
	case abi.StringTy: // variable arrays are written at the end of the return bytes
		return string(output[begin : begin+end]), end/32 + 1, nil
	case abi.IntTy, abi.UintTy:
		return readInteger(t.Kind, returnOutput), 1, nil
	case abi.BoolTy:
		return readBool(returnOutput)
	case abi.AddressTy:
		return common.BytesToAddress(returnOutput), 1, nil
	case abi.HashTy:
		return common.BytesToHash(returnOutput), 1, nil
	case abi.BytesTy:
		return output[begin : begin+end], end/32 + 1, nil
	case abi.FixedBytesTy:
		return readFixedBytes(t, returnOutput)
	case abi.FunctionTy:
		return readFunctionType(t, returnOutput)
	default:
		return nil, 0, fmt.Errorf("abi: unknown type %v", t.T)
	}
}

// interprets a 32 byte slice as an offset and then determines which indice to look to decode the type.
func lengthPrefixPointsTo(index int, output []byte) (start int, length int, err error) {
	offset := int(binary.BigEndian.Uint64(output[index+24 : index+32]))
	if offset+32 > len(output) {
		return 0, 0, fmt.Errorf("abi: cannot marshal in to go slice: offset %d would go over slice boundary (len=%d)", len(output), offset+32)
	}
	length = int(binary.BigEndian.Uint64(output[offset+24 : offset+32]))
	if offset+32+length > len(output) {
		return 0, 0, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), offset+32+length)
	}
	start = offset + 32

	//fmt.Printf("LENGTH PREFIX INFO: \nsize: %v\noffset: %v\nstart: %v\n", length, offset, start)
	return
}

func requiresLengthPrefix(t abi.Type) bool {
	return t.T == abi.StringTy || t.T == abi.BytesTy || t.T == abi.SliceTy
}
