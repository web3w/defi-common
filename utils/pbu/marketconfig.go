package pbu

import (
	"fmt"

	pb "github.com/gisvr/defi-common/proto/dex2"
	dex2sc "github.com/gisvr/defi-common/utils/dex2"
	"github.com/gisvr/defi-common/utils/ueth"
	"github.com/golang/protobuf/proto"
)

func MarketConfigShallowCopy(cfg *pb.MarketConfig) *pb.MarketConfig {
	var cp = *cfg

	// Decouple the slice fields.
	cp.CashTokens = make([]*pb.TokenInfo, len(cfg.CashTokens))
	cp.StockTokens = make([]*pb.TokenInfo, len(cfg.StockTokens))
	cp.DisabledTokens = make([]*pb.TokenInfo, len(cfg.DisabledTokens))
	copy(cp.CashTokens, cfg.CashTokens)
	copy(cp.StockTokens, cfg.StockTokens)
	copy(cp.DisabledTokens, cfg.DisabledTokens)
	return &cp
}

func ValidateMarketConfig(cfg *pb.MarketConfig, maxFeeRateE4 uint64) error {
	if cfg.MakerFeeRateE4 < 0 || cfg.MakerFeeRateE4 > maxFeeRateE4 {
		return fmt.Errorf("invalid MakerFeeRateE4 %v", cfg.MakerFeeRateE4)
	}
	if cfg.TakerFeeRateE4 < 0 || cfg.TakerFeeRateE4 > maxFeeRateE4 {
		return fmt.Errorf("invalid TakerFeeRateE4 %v", cfg.TakerFeeRateE4)
	}
	if cfg.WithdrawFeeRateE4 < 0 || cfg.WithdrawFeeRateE4 > maxFeeRateE4 {
		return fmt.Errorf("invalid WithdrawFeeRateE4 %v", cfg.WithdrawFeeRateE4)
	}

	for _, token := range cfg.CashTokens {
		if err := ValidateTokenInfo(token); err != nil {
			return err
		}
	}
	for _, token := range cfg.StockTokens {
		if err := ValidateTokenInfo(token); err != nil {
			return err
		}
	}
	for _, token := range cfg.DisabledTokens {
		if err := ValidateTokenInfo(token); err != nil {
			return err
		}
	}
	return nil
}

func IsTokenInfoConflicting(a *pb.TokenInfo, b *pb.TokenInfo) bool {
	if a.TokenId == b.TokenId || a.TokenCode == b.TokenCode || a.TokenAddr == b.TokenAddr {
		return !proto.Equal(a, b)
	}
	return false
}

// No-op if already exists
func InsertTokenInfo(slice *([]*pb.TokenInfo), t *pb.TokenInfo) {
	for _, tk := range *slice {
		if proto.Equal(tk, t) {
			return
		}
	}
	*slice = append(*slice, t)
}

// No-op if does not exists
func RemoveTokenInfo(slice *([]*pb.TokenInfo), t *pb.TokenInfo) {
	p := 0
	for i := range *slice {
		if !proto.Equal((*slice)[i], t) {
			(*slice)[p] = (*slice)[i]
			p++
		}
	}
	*slice = (*slice)[:p]
}

func ValidateTokenInfo(t *pb.TokenInfo) error {
	if len(t.TokenId) == 0 || len(t.TokenId) > 6 {
		return fmt.Errorf("invalid TokenId %v", t.TokenId)
	}

	if t.TokenId == "ETH" {
		if t.TokenAddr != "" {
			return fmt.Errorf("TokenAddr of ETH is non-empty: %v", t.TokenAddr)
		}
	} else {
		if !ueth.IsHexAddr(t.TokenAddr) {
			return fmt.Errorf("invalid TokenAddr %v", t.TokenAddr)
		}
	}

	if t.TokenCode > 0xffff {
		return fmt.Errorf("TokenCode %v is not 16-bit", t.TokenCode)
	}
	if t.ScaleFactor == 0 {
		return fmt.Errorf("zero ScaleFactor")
	}
	if _, err := dex2sc.CheckTokenScaleFactor(t.ScaleFactor); err != nil {
		return err
	}
	return nil
}
