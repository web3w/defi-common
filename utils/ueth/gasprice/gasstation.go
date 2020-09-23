package gasprice

import (
	"encoding/json"
	"fmt"
	"github.com/gisvr/deif-common/utils/utime"
	"io/ioutil"
	"math/big"
	"net/http"

	"bytes"
)

// Example response:
// {
// 	"safelow_calc": 20,
// 	"average_calc": 40,
// 	"average": 40,
// 	"safelow_txpool": 20,
// 	"fastest": 200,
// 	"fast": 80,
// 	"average_txpool": 40,
// 	"safeLow": 20
// }
const ethGasStationApiUrl = "https://ethgasstation.info/json/ethgasAPI.json"

// Prices in Wei
type gasStationReading struct {
	SafeLow *big.Int
	Average *big.Int
	Fast    *big.Int
	Fastest *big.Int
}

func FetchGasStationReading() (gasStationReading, error) {
	req, err := http.NewRequest(http.MethodGet, ethGasStationApiUrl, nil)
	if err != nil {
		return gasStationReading{}, err
	}

	var client http.Client
	req = req.WithContext(utime.CtxWithTimeoutMs(15e3))
	resp, err := client.Do(req)
	if err != nil {
		return gasStationReading{}, err
	}
	defer resp.Body.Close()

	bin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gasStationReading{}, err
	}

	if len(bin) == 0 {
		return gasStationReading{}, fmt.Errorf("empty response from %v", ethGasStationApiUrl)
	}

	raw := struct {
		Fast    float64 `json:"fast"`    // in 0.1 Gwei
		Fastest float64 `json:"fastest"` // in 0.1 Gwei
		Average float64 `json:"average"` // in 0.1 Gwei
		SafeLow float64 `json:"safeLow"` // in 0.1 Gwei
	}{}
	// replace NaN with -1
	bin = bytes.Replace(bin, []byte("NaN"), []byte("-1"), -1)
	err = json.Unmarshal(bin, &raw)
	if err != nil {
		// API may have changed. In that case, this code has to be updated.
		return gasStationReading{}, fmt.Errorf("Failed to parse body %v: %v", string(bin), err)
	}
	// handle NaN
	if raw.Fast <= 0 {
		raw.Fast = raw.Fastest
	}
	if raw.Average <= 0 {
		raw.Average = raw.Fast
	}
	if raw.SafeLow <= 0 {
		raw.SafeLow = raw.Average
	}
	if raw.Fast <= 0 || raw.Fastest <= 0 || raw.Average <= 0 || raw.SafeLow <= 0 {
		return gasStationReading{}, fmt.Errorf("Invalid response %+v", raw)
	}

	reading := gasStationReading{
		Fast:    big.NewInt(int64(raw.Fast*1e8 + 0.5)),
		Fastest: big.NewInt(int64(raw.Fastest*1e8 + 0.5)),
		Average: big.NewInt(int64(raw.Average*1e8 + 0.5)),
		SafeLow: big.NewInt(int64(raw.SafeLow*1e8 + 0.5)),
	}
	return reading, nil
}
