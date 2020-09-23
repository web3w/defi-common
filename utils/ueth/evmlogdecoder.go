package ueth

import (
	"fmt"
	"github.com/gisvr/defi-common/utils/ueth/unpack"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO: add unittests

// EvmLogDecoder can decodes EVM logs emitted from a set of contracts to events.
type EvmLogDecoder struct {
	eventAbis map[common.Hash]abi.Event // keyed by event id (event's signature)
}

// Returns (nil, error) if any ABI string cannot be parsed or if there are duplicate event ABIs.
func NewEvmLogDecoder(contractAbiStrs ...string) (*EvmLogDecoder, error) {
	eventAbis := make(map[common.Hash]abi.Event)
	for _, abiStr := range contractAbiStrs {
		contractAbi, err := abi.JSON(strings.NewReader(abiStr))
		if err != nil {
			return nil, err
		}

		for _, event := range contractAbi.Events {
			id := event.ID()
			if _, ok := eventAbis[id]; ok {
				return nil, fmt.Errorf("duplication event ABI: %v", event.Name)
			}
			eventAbis[id] = event
		}
	}
	return &EvmLogDecoder{eventAbis: eventAbis}, nil
}

// Decodes an EVM log. Returns the name and the argument map of the event. Returns an error if `log`
// is nil, or has a unknown event ABI, or cannot be decoded.
func (dec *EvmLogDecoder) Decode(log *types.Log) (string, map[string]interface{}, error) {
	if log == nil {
		return "", nil, fmt.Errorf("cannot decode nil log")
	}

	eventAbi, ok := dec.eventAbis[log.Topics[0]]
	if !ok {
		return "", nil, fmt.Errorf("unknown event ABI %v", log.Topics[0])
	}

	args := make(map[string]interface{})
	topicIndex := 1
	dataIndex := 0
	for _, input := range eventAbi.Inputs {
		if input.Indexed == true {
			val, _, err := unpack.ToGoType(0, input.Type, log.Topics[topicIndex].Bytes())
			if err != nil {
				return "", nil, err
			}
			args[input.Name] = val
			topicIndex++
		} else {
			val, size, err := unpack.ToGoType(dataIndex, input.Type, log.Data)
			if err != nil {
				return "", nil, err
			}
			args[input.Name] = val
			dataIndex += size * 32
		}
	}
	return eventAbi.Name, args, nil
}
