package hyperliquid

import (
	"context"
	"log"
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
	assert.Fail(t, "for debug")
}

func TestPerpetualsUserStat(t *testing.T) {
	cli, err := DialContext(context.Background(), MainnetAPIURL, "")
	assert.Equal(t, err, nil, "DialContext error")
	state, err := cli.PerpetualsUserStat(context.Background(), "0x30000Da0cb2c459EC74aE5e80bE247A4f8DAcF98", "")
	assert.Equal(t, err, nil, "Get PerpetualsUserStat error")
	assert.Greater(t, state.Time, 10000, "Get PerpetualsUserStat error")
	log.Printf("state:%+v \n", state)
}

func TestOpenOrders(t *testing.T) {
	cli, err := DialContext(context.Background(), MainnetAPIURL, "")
	assert.Equal(t, err, nil, "DialContext error")
	orders, err := cli.OpenOrders(context.Background(), "0x30000Da0cb2c459EC74aE5e80bE247A4f8DAcF98", "")
	assert.Equal(t, err, nil, "Get OpenOrders error")
	log.Printf("orders:%+v \n", orders)
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

func TestSubscribeNotification(t *testing.T) {
	cli, err := DialContext(context.Background(), MainnetAPIURL, MainnetWSURL)
	assert.Equal(t, err, nil, "DialContext error")
	err = cli.SubscribeNotification("0x30000Da0cb2c459EC74aE5e80bE247A4f8DAcF98")
	assert.Equal(t, err, nil, "Subscribe All Mids error")
	time.Sleep(10 * time.Minute)
}
