package order

const (
	// TypeLimit represents that the order type is a limit.
	TypeLimit = "LIMIT"
	// TypeMarket represents that the order type is a market.
	TypeMarket = "MARKET"

	// SideBuy represents that the side is a buy.
	SideBuy = "BUY"
	// SideSell represents that the side is a sell.
	SideSell = "SELL"

	// TimeInForceGTC represents that the time in force is good till canceled.
	TimeInForceGTC = "GTC"
	// TimeInForceIOC represents that the time in force is immediate or cancel.
	TimeInForceIOC = "IOC"
	// TimeInForceFOK represents that the time in force is fill or kill.
	TimeInForceFOK = "FOK"
)
