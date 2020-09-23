package dex2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeployMarketEventTopic(t *testing.T) {
	for _, event := range Dex2Abi.Events {
		id := event.ID()
		if id == DeployMarketEventTopic {
			return
		}
	}
	require.Fail(t, "DeployMarketEventTopic is missing from the DEx2 ABI")
}
