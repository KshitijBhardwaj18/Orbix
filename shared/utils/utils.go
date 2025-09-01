package utils

import (
	"fmt"
	"strings"
)

func ParseMarketId(marketId string) (baseAsset, quoteAsset string, err error) {

	parts := strings.Split(marketId, "/")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("Invalid market id %s", marketId)
	}

	return parts[0], parts[1], nil
}
