package unuls

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestVerifyMsg(t *testing.T) {
	msg := "444578323a20547261646572203c744e554c536542614d6642714c6257527876767565417a674a6d637a69424a536a416843545a3e20576974686472617773203c4e554c533e203c3130303030303030303e2061742054696d657374616d702031353734383431303735"
	signed := "692103657bd5a70162bd15cb605896acd0d0985edd4eed6886fd02ef22b145c61a18f846304402202d8765279c2d8d3065e947c3cab729fb18c4d9de318c84c89e4cb9725110d00a0220675d0edeb6020e58e31ab172f3a6cd12bd8ab4fcc8d83f89ce0fd4ccab37f944"
	address := "tNULSeBaMfBqLbWRxvvueAzgJmcziBJSjAhCTZ"
	result := VerifyRpcSingedMsg(msg, signed, address)
	fmt.Println(result)
}

func TestReadVarInt(t *testing.T) {
	lengthByteArr, _ := hex.DecodeString("fe07000100")
	length, size := ReadVarInt(lengthByteArr, 0)
	fmt.Println(length, size)
}

func TestBytesToInt32(t *testing.T) {
	buf := "0b6465706f7369744e554c53940200010321e"
	fmt.Println(buf[0])
	result := 0xFF & buf[0]
	ss := uint64(result)
	fmt.Println(ss)
}