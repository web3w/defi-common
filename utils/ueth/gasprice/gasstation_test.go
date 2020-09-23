package gasprice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchGasStationReading(t *testing.T) {
	httpmock.Activate()

	httpmock.RegisterResponder("GET", ethGasStationApiUrl,
		httpmock.NewStringResponder(200,
			`{"fastWait": 0.5, "blockNum": 5887665, "average_calc": 300.0, "block_time": 15.756756756756756, "safeLow": 200.0, "average_txpool": 300.0, "average": 300.0, "safeLowWait": 12.4, "avgWait": 3.4, "safelow_txpool": 200.0, "fast": 400.0, "fastestWait": 0.5, "safelow_calc": 200.0, "speed": 0.7015095397648193, "fastest": 930.0}`))
	{
		reading, err := FetchGasStationReading()
		assert.NoError(t, err)
		assert.Equal(t, uint64(30e9), reading.Average.Uint64())
	}

	httpmock.Reset()
	httpmock.RegisterResponder("GET", ethGasStationApiUrl,
		httpmock.NewStringResponder(200,
			`{"fastWait": 0.5, "blockNum": 5887665, "average_calc": 300.0, "block_time": 15.756756756756756, "safeLow": 200.0, "average_txpool": 300.0, "average": NaN, "safeLowWait": 12.4, "avgWait": 3.4, "safelow_txpool": 200.0, "fast": 400.0, "fastestWait": 0.5, "safelow_calc": 200.0, "speed": 0.7015095397648193, "fastest": 930.0}`))
	{
		reading, err := FetchGasStationReading()
		assert.NoError(t, err)
		assert.Equal(t, uint64(40e9), reading.Average.Uint64())
	}

	httpmock.Reset()
	httpmock.RegisterResponder("GET", ethGasStationApiUrl,
		httpmock.NewStringResponder(200,
			`{"fastWait": 0.5, "blockNum": 5887665, "average_calc": 300.0, "block_time": 15.756756756756756, "safeLow": 200.0, "average_txpool": 300.0, "average": NaN, "safeLowWait": 12.4, "avgWait": 3.4, "safelow_txpool": 200.0, "fast": NaN, "fastestWait": 0.5, "safelow_calc": 200.0, "speed": 0.7015095397648193, "fastest": 930.0}`))
	{
		reading, err := FetchGasStationReading()
		assert.NoError(t, err)
		assert.Equal(t, uint64(93e9), reading.Average.Uint64())
	}

	httpmock.Reset()
	httpmock.RegisterResponder("GET", ethGasStationApiUrl,
		httpmock.NewStringResponder(200,
			`{"fastWait": 0.5, "blockNum": 5887665, "average_calc": 300.0, "block_time": 15.756756756756756, "safeLow": 200.0, "average_txpool": 300.0, "average": NaN, "safeLowWait": 12.4, "avgWait": 3.4, "safelow_txpool": 200.0, "fast": NaN, "fastestWait": 0.5, "safelow_calc": 200.0, "speed": 0.7015095397648193, "fastest": NaN}`))
	{
		_, err := FetchGasStationReading()
		assert.Error(t, err)
	}
}
