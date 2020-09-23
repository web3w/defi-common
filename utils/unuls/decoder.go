package unuls

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Decoder struct {
	input []byte
}

// NewDecoder returns a new Decoder instance.
func NewDecoder(input []byte) *Decoder {
	return &Decoder{input}
}


func (sd *Decoder) ParseUint8() (uint8, error) {
	if len(sd.input) < 1 {
		return 0, fmt.Errorf("parse int32 fail: invalid len %v", sd.input)
	}
	result := sd.input[:1]
	sd.input = sd.input[1:]
	return result[0], nil
}

func (sd *Decoder) ParseInt32() (int32, error) {
	if len(sd.input) < 4 {
		return 0, fmt.Errorf("parse int32 fail: invalid len %v", sd.input)
	}
	result := BytesToInt32(sd.input[:4])
	sd.input = sd.input[4:]
	return result, nil
}

func (sd *Decoder) ParseUint64() (uint64, error) {
	if len(sd.input) < 16 {
		return 0, fmt.Errorf("parse int32 fail: invalid len %v", sd.input)
	}

	decodeStr, _ := hex.DecodeString(string(sd.input[:16]))
	result := binary.LittleEndian.Uint64(decodeStr)
	sd.input = sd.input[16:]
	return result, nil
}

func (sd *Decoder) ParseByteByLength() ([]byte, uint64, uint, error) {
	length, originallyEncodedSize := ReadVarInt(sd.input, 0)
	var actualLength uint64
	actualLength = length

	if len(sd.input) < int(actualLength) {
		return nil, 0, 0, fmt.Errorf("bytes length too large: %v > %v", length, len(sd.input))
	}
	result := sd.input[originallyEncodedSize:actualLength+uint64(originallyEncodedSize)]
	sd.input = sd.input[actualLength+uint64(originallyEncodedSize):]
	return result, actualLength, originallyEncodedSize, nil
}

func (sd *Decoder) ParseBytesWithLength(length int) ([]byte, error) {
	if len(sd.input) < length {
		return nil, fmt.Errorf("bytes length too large: %v > %v", length, len(sd.input))
	}
	result := sd.input[:length]
	sd.input = sd.input[length:]
	return result, nil
}

func BytesToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(b))
}

func BytesToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func BytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}


func ReadVarInt(buf []byte , offset int) (length uint64, originallyEncodedSize uint){

	var first  = uint64(0xFF & buf[offset])
	if first < 253 {
		// 1 entity byte (8 bits)
		length = first
		originallyEncodedSize = 1
	} else if (first == 253) {
		// 1 marker + 2 entity bytes (16 bits)
		length = (uint64((0xFF & buf[offset + 1])) | (uint64((0xFF & buf[offset + 2])) << 8));
		originallyEncodedSize = 3
	} else if (first == 254) {
		// 1 marker + 4 entity bytes (32 bits)
		length = uint64(BytesToUint32(buf[offset + 1:5]));
		originallyEncodedSize = 5
	} else {
		// 1 marker + 8 entity bytes (64 bits)
		length = uint64(BytesToUint64(buf[offset + 1:9]));
		originallyEncodedSize = 9
	}

	return length, originallyEncodedSize
}


