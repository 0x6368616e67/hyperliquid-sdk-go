package hyperliquid

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	cli, err := DialContext(context.Background(), MainnetAPIURL, "")
	assert.Equal(t, err, nil, "DialContext error")
	coins, err := cli.Metadata(context.Background(), "")
	assert.Equal(t, err, nil, "Get Metadata error")
	assert.Greater(t, len(coins), 1, "Metadata is empty")
}

func TestAllMids(t *testing.T) {
	cli, err := DialContext(context.Background(), MainnetAPIURL, "")
	assert.Equal(t, err, nil, "DialContext error")
	mids, err := cli.AllMids(context.Background(), "")
	assert.Equal(t, err, nil, "Get All Mids error")
	assert.NotEqual(t, nil, mids["BTC"], "No BTC mids")
}

func TestSubscribeAllMids(t *testing.T) {
	cli, err := DialContext(context.Background(), MainnetAPIURL, MainnetWSURL)
	assert.Equal(t, err, nil, "DialContext error")
	err = cli.SubscribeAllMids()
	assert.Equal(t, err, nil, "Subscribe All Mids error")
	time.Sleep(1 * time.Minute)
}
