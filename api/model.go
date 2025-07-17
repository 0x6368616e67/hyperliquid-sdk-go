package api

type AssetMetadata struct {
	Name         string `json:"name"`
	SzDecimals   int    `json:"szDecimals"`
	MaxLeverage  int    `json:"maxLeverage"`
	OnlyIsolated bool   `json:"onlyIsolated,omitempty"`
	IsDelisted   bool   `json:"isDelisted,omitempty"`
}

type OrderInfo struct {
	Coin      string `json:"coin"`
	Side      string `json:"side"`
	LimitPx   string `json:"limitPx"`
	Sz        string `json:"sz"`
	Oid       int64  `json:"oid"`
	Timestamp int64  `json:"timestamp"`
	OrigSz    string `json:"origSz"`
}

type UserStat struct {
	MarginSummary struct {
		AccountValue    string `json:"accountValue"`
		TotalNtlPos     string `json:"totalNtlPos"`
		TotalRawUsd     string `json:"totalRawUsd"`
		TotalMarginUsed string `json:"totalMarginUsed"`
	} `json:"marginSummary"`
	CrossMarginSummary struct {
		AccountValue    string `json:"accountValue"`
		TotalNtlPos     string `json:"totalNtlPos"`
		TotalRawUsd     string `json:"totalRawUsd"`
		TotalMarginUsed string `json:"totalMarginUsed"`
	} `json:"crossMarginSummary"`
	CrossMaintenanceMarginUsed string `json:"crossMaintenanceMarginUsed"`
	Withdrawable               string `json:"withdrawable"`
	AssetPositions             []struct {
		Type     string `json:"type"`
		Position struct {
			Coin     string `json:"coin"`
			Szi      string `json:"szi"`
			Leverage struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"leverage"`
			EntryPx        string `json:"entryPx"`
			PositionValue  string `json:"positionValue"`
			UnrealizedPnl  string `json:"unrealizedPnl"`
			ReturnOnEquity string `json:"returnOnEquity"`
			LiquidationPx  any    `json:"liquidationPx"`
			MarginUsed     string `json:"marginUsed"`
			MaxLeverage    int    `json:"maxLeverage"`
			CumFunding     struct {
				AllTime     string `json:"allTime"`
				SinceOpen   string `json:"sinceOpen"`
				SinceChange string `json:"sinceChange"`
			} `json:"cumFunding"`
		} `json:"position"`
	} `json:"assetPositions,omitempty"`
	Time int64 `json:"time"`
}
