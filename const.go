package hyperliquid

import "time"

const (
	MainnetAPIURL = "https://api.hyperliquid.xyz"
	TestnetAPIURL = "https://api.hyperliquid-testnet.xyz"
	LocalAPIURL   = "http://localhost:3001"
	MainnetWSURL  = "wss://api.hyperliquid.xyz/ws"
	TestnetWSURL  = "wss://api.hyperliquid-testnet.xyz/ws"
)

const (
	WSPongWait          = 60 * time.Second
	WSMaxMessageSize    = 1024 * 100
	WSHeartbeatInterval = 30 * time.Second
)

type CandleIntervalType string

const (
	Interval1m  = CandleIntervalType("1m")
	Interval3m  = CandleIntervalType("3m")
	Interval5m  = CandleIntervalType("5m")
	Interval15m = CandleIntervalType("15m")
	Interval30m = CandleIntervalType("30m")

	Interval1h  = CandleIntervalType("1h")
	Interval2h  = CandleIntervalType("2h")
	Interval4h  = CandleIntervalType("4h")
	Interval8h  = CandleIntervalType("8h")
	Interval12h = CandleIntervalType("12h")

	Interval1d = CandleIntervalType("1d")
	Interval3d = CandleIntervalType("3d")

	Interval1w = CandleIntervalType("1w")
	Interval1M = CandleIntervalType("1M")
)
